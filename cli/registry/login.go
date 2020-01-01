package registry

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"cli/config"
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
		channel := handler.getValue("channel", 0)
		password := handler.getValue("password", '*')

		handler.validateChannelCredentials(channel, password)
		conf := config.New()
		conf.Channel.Channel = channel
		conf.Channel.Password = password

		conf.Save()
		log.Printf("Logged into '%s'", channel)
	},
}

type loginHandler struct {
	cmd *cobra.Command
}

func (h *loginHandler) getValue(flag string, mask rune) string {
	value := h.getFlag(flag)
	if value != "" {
		return value
	}

	return h.promptValue(strings.Title(flag), mask)
}

func (h *loginHandler) getFlag(flag string) string {
	value, err := h.cmd.Flags().GetString(flag)
	if err != nil {
		log.Fatalln(err)
	}

	return strings.TrimSpace(value)
}

func (h *loginHandler) promptValue(label string, mask rune) string {
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
		log.Fatalln(err)
	}
	return value
}

func (h *loginHandler) validateChannelCredentials(channel, password string) {
	conf := config.New()
	if !conf.HasRegistry() {
		return
	}

	payload := strings.NewReader(fmt.Sprintf(`{
		"channel": "%s",
		"password": "%s"
	}`, channel, password))

	resp, err := http.Post(conf.Registry+"/user/check", "application/json", payload)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "could not check user information against server"))
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("channel credentials for '%s' is incorrect. Did you create the account created or have you forgotten the password?", channel)
	}
}
