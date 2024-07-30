package cli

import (
	"log"

	"clitemplate/cmd/web"

	"github.com/spf13/cobra"
)

var ServerCmd = &cobra.Command{
	Use:              "server",
	Short:            "API server",
	PersistentPreRun: state.ConnectDatabase,
	Run: func(cmd *cobra.Command, args []string) {
		log.Fatal(
			web.
				CreateServer().
				Listen(state.Flags.ServerAddr),
		)
	},
}

func init() {
	ServerCmd.Flags().StringVar(&state.Flags.ServerAddr, "address", state.Flags.ServerAddr, "")
}
