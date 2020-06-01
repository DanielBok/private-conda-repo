package registry

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"cli/config"
	"cli/request"
)

func init() {
	loginCmd.Flags().StringP("channel", "c", "", "Channel name")
	loginCmd.Flags().StringP("password", "p", "", "Registry password")
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log into the registry",
	Long: `Logs the cli tool with the channel's credentials. The channel's credentials will be verified
against the server. `,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, _ []string) {
		handler := loginHandler{cmd: cmd}
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

		err = handler.validateChannelCredentials(channel, password)
		if err != nil {
			cmd.PrintErr(err)
			return
		}

		conf := config.New()
		conf.Channel.Channel = channel
		conf.Channel.Password = password

		conf.Save()
		cmd.Printf("Logged into '%s'", channel)
	},
}

type loginHandler struct {
	cmd *cobra.Command
}

func (h *loginHandler) getValue(flag string, mask rune) (string, error) {
	value, err := h.getFlag(flag)
	if err != nil {
		return "", err
	} else if value != "" {
		return value, nil
	}

	return h.promptValue(strings.Title(flag), mask)
}

func (h *loginHandler) getFlag(flag string) (string, error) {
	value, err := h.cmd.Flags().GetString(flag)
	if err != nil {
		return "", errors.Wrapf(err, "could not get '%s' flag value", flag)
	}

	return strings.TrimSpace(value), nil
}

func (h *loginHandler) promptValue(label string, mask rune) (string, error) {
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

func (h *loginHandler) validateChannelCredentials(channel, password string) error {
	conf := config.New()
	if !conf.HasRegistry() {
		return nil
	}

	payload := strings.NewReader(fmt.Sprintf(`{
		"channel": "%s",
		"password": "%s"
	}`, channel, password))

	resp, err := request.Post(conf.Registry+"/channel/check", "application/json", payload)
	if err != nil {
		return errors.Wrap(err, "could not check user information against server")
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return errors.Errorf("channel credentials for '%s' is incorrect. Did you create the account created or have you forgotten the password?", channel)
	}

	return nil
}
