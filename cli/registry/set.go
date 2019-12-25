package registry

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"cli/config"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets the package registry",
	Long: `Verifies the package registry and sets it if successful. The registry needs to be specified
for the cli to work correctly`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		handler := setHandler{cmd: cmd}
		handler.fetchMeta(args[0])

		conf := config.New()
		conf.Registry = handler.Registry
		conf.PackageRepository = handler.Repository

		conf.Save()
		log.Printf(strings.TrimSpace(fmt.Sprintf(`Set registry target to:
	Registry:   %s
	Repository: %s
`, conf.Registry, conf.PackageRepository)))
	},
}

type setHandler struct {
	cmd *cobra.Command
	registryMeta
}

type registryMeta struct {
	Registry   string `json:"registry"`
	Repository string `json:"repository"`
}

func (h *setHandler) fetchMeta(host string) {
	resp, err := http.Get(host + "/meta")
	if err != nil {
		log.Fatal(errors.Wrap(err, "could not fetch meta information from registry. Is this a valid Private Conda Repo?"))
	}
	defer func() { _ = resp.Body.Close() }()

	var meta registryMeta
	if err := json.NewDecoder(resp.Body).Decode(&meta); err != nil {
		log.Fatal(errors.Wrap(err, "could not parse meta information from registry"))
	}

	h.Registry = meta.Registry
	h.Repository = meta.Repository
}
