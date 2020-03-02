package registry

import (
	"fmt"
	"io/ioutil"
	"net/http"
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
	Short: "Registers a new channel account",
	Long:  `Creates a new account and logs the cli tool with the channel's credentials.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, _ []string) {
		handler := registerHandler{cmd: cmd}
		channel, err := handler.getValue("channel", 0)
		if err != nil {
			cmd.PrintErr(err)
			return
		}

		password, err := handler.getValue("password", '*')
		if err != nil {
			cmd.PrintErr(err)
			return
		}

		email, err := handler.getValue("email", 0)
		if err != nil {
			cmd.PrintErr(err)
			return
		}

		err = handler.registerChannel(channel, password, email)
		if err != nil {
			cmd.PrintErr(err)
			return
		}

		conf := config.New()
		conf.Channel.Channel = channel
		conf.Channel.Password = password

		conf.Save()
		cmd.Printf("Created and logged in as channel '%s'", channel)
	},
}

type registerHandler struct {
	cmd *cobra.Command
}

func (h *registerHandler) getValue(flag string, mask rune) (string, error) {
	value, err := h.getFlag(flag)
	if err != nil {
		return "", err
	} else if value != "" {
		return value, nil
	}

	return h.promptValue(strings.Title(flag), mask)
}

func (h *registerHandler) getFlag(flag string) (string, error) {
	value, err := h.cmd.Flags().GetString(flag)
	if err != nil {
		return "", errors.Wrapf(err, "could not get %s flag value", flag)
	}

	return strings.TrimSpace(value), nil
}

func (h *registerHandler) promptValue(label string, mask rune) (string, error) {
	prompt := promptui.Prompt{
		Label: label,
		Validate: func(input string) error {
			input = strings.TrimSpace(input)
			if input == "" {
				return errors.Errorf("%s cannot be empty", label)
			}
			if len(input) < 4 {
				return errors.Errorf("%s length must be >= 4 characters", label)
			}
			return nil
		},
		Mask: mask,
	}

	value, err := prompt.Run()
	if err != nil {
		return "", err
	}
	return value, nil
}

func (h *registerHandler) registerChannel(channel, password, email string) error {
	conf := config.New()
	if !conf.HasRegistry() {
		return nil
	}

	payload := strings.NewReader(fmt.Sprintf(`{
		"channel": "%s",
		"password": "%s",
		"email": "%s"
	}`, channel, password, email))

	resp, err := request.Post(conf.Registry+"/user", "application/json", payload)
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
