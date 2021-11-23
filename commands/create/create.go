package create

import (
	"os"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/commands/create/cluster"
	"phoenixnap.com/pnap-cli/commands/create/server"
	"phoenixnap.com/pnap-cli/commands/create/server/privatenetwork"
	"phoenixnap.com/pnap-cli/commands/create/sshkey"
	"phoenixnap.com/pnap-cli/commands/create/tag"
)

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a resource.",
	Long:  `Create a resource.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

func init() {
	CreateCmd.AddCommand(server.CreateServerCmd)
	CreateCmd.AddCommand(cluster.CreateClusterCmd)
	CreateCmd.AddCommand(tag.CreateTagCmd)
	CreateCmd.AddCommand(privatenetwork.CreateServerPrivateNetworkCmd)
	CreateCmd.AddCommand(sshkey.CreateSshKeyCmd)
}
