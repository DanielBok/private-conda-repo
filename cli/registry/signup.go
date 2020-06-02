package registry

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"cli/config"
	"cli/request"
)

func init() {
	registerCmd.Flags().StringP("channel", "c", "", "Channel name")
	registerCmd.Flags().StringP("password", "p", "", "Password")
	registerCmd.Flags().StringP("email", "e", "", "Email address")
}

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Registers a new channel",
	Long:  `Creates a new account and logs the cli tool with the channel's credentials.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, _ []string) {
		handler := RegisterHandler{cmd: cmd}

		payload, err := handler.getAllInputs()
		if err != nil {
			cmd.PrintErr(err)
			return
		}

		err = handler.registerChannel(payload)
		if err != nil {
			cmd.PrintErr(err)
			return
		}

		conf := config.New()
		conf.Channel.Channel = payload.Channel
		conf.Channel.Password = payload.Password

		conf.Save()
		cmd.Printf("Created and logged into channel '%s'", payload.Channel)
	},
}

type RegisterHandler struct {
	cmd *cobra.Command
}

type SignUpPayload struct {
	Channel  string `json:"channel"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (h *RegisterHandler) getAllInputs() (*SignUpPayload, error) {
	var errs []string

	channel, err := h.getChannel()
	if err != nil {
		errs = append(errs, err.Error())
	}

	password, err := h.getPassword()
	if err != nil {
		errs = append(errs, err.Error())
	}

	email, err := h.getEmail()
	if err != nil {
		errs = append(errs, err.Error())
	}

	if len(errs) != 0 {
		return nil, errors.Errorf("Invalid values: \n%s", strings.Join(errs, "\n"))
	}

	return &SignUpPayload{
		Channel:  channel,
		Password: password,
		Email:    email,
	}, nil
}

func (h *RegisterHandler) getChannel() (string, error) {
	return h.getValue("channel", 0, func(input string) error {
		input = strings.TrimSpace(input)
		if input == "" {
			return errors.New("Channel cannot be empty")
		}
		if len(input) < 2 {
			return errors.New("Channel length must be >= 2 characters")
		}
		return nil
	})
}

func (h *RegisterHandler) getPassword() (string, error) {
	return h.getValue("password", '*', func(input string) error {
		if input == "" {
			return errors.New("Password cannot be empty")
		}
		if len(input) < 4 {
			return errors.New("Password length must be >= 4 characters")
		}
		return nil
	})
}

func (h *RegisterHandler) getEmail() (string, error) {
	re, err := regexp.Compile(`^[^@\s]+@[^@\s]+$`)
	if err != nil {
		return "", err
	}

	return h.getValue("email", 0, func(input string) error {
		input = strings.TrimSpace(input)

		if !re.MatchString(input) {
			return errors.New("invalid email address")
		}
		return nil
	})
}

func (h *RegisterHandler) getValue(flag string, mask rune, validateFunc promptui.ValidateFunc) (string, error) {
	value, err := h.cmd.Flags().GetString(flag)
	if err != nil {
		return "", errors.Wrapf(err, "could not get %s flag value", flag)
	}

	// checks and returns a value if flags are specified
	if value != "" {
		err = validateFunc(value)
		if err != nil {
			return "", err
		}
		return value, nil
	}

	// value is unspecified from flags, prompt user instead
	label := strings.Title(flag)
	prompt := promptui.Prompt{
		Label:    label,
		Validate: validateFunc,
		Mask:     mask,
	}

	value, err = prompt.Run()
	if err != nil {
		return "", err
	}
	return value, nil
}

func (h *RegisterHandler) registerChannel(payload *SignUpPayload) error {
	conf := config.New()
	if !conf.HasRegistry() {
		return nil
	}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(payload)
	if err != nil {
		return err
	}

	resp, err := request.Post(conf.Registry+"/channel", "application/json", &buf)
	if err != nil {
		return errors.Wrap(err, "could not create account")
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return errors.Wrap(err, "error encountered and could not decipher error message from server")
		}
		return errors.Errorf("could not create account: %s", string(body))
	}

	return nil
}
