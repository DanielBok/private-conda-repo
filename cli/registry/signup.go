package registry

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"cli/config"
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
		channel := handler.getValue("channel", 0)
		password := handler.getValue("password", '*')
		email := handler.getValue("email", 0)

		handler.registerChannel(channel, password, email)
		conf := config.New()
		conf.Channel.Channel = channel
		conf.Channel.Password = password

		conf.Save()
		log.Printf("Created and logged in as channel '%s'", channel)
	},
}

type registerHandler struct {
	cmd *cobra.Command
}

func (h *registerHandler) getValue(flag string, mask rune) string {
	value := h.getFlag(flag)
	if value != "" {
		return value
	}

	return h.promptValue(strings.Title(flag), mask)
}

func (h *registerHandler) getFlag(flag string) string {
	value, err := h.cmd.Flags().GetString(flag)
	if err != nil {
		log.Fatalln(err)
	}

	return strings.TrimSpace(value)
}

func (h *registerHandler) promptValue(label string, mask rune) string {
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

func (h *registerHandler) registerChannel(channel, password, email string) {
	conf := config.New()
	if !conf.HasRegistry() {
		return
	}

	payload := strings.NewReader(fmt.Sprintf(`{
		"channel": "%s",
		"password": "%s",
		"email": "%s"
	}`, channel, password, email))

	resp, err := http.Post(conf.Registry+"/user", "application/json", payload)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "could not create account"))
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln("error encountered and could not decipher error message from server", err)
		}
		log.Fatalf("could not create account: %s", string(body))
	}
}
