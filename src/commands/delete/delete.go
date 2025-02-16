package delete

import (
	"os"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/commands/delete/cluster"
	"phoenixnap.com/pnapctl/commands/delete/server"
	"phoenixnap.com/pnapctl/commands/delete/sshkey"
	"phoenixnap.com/pnapctl/commands/delete/tag"

	"phoenixnap.com/pnapctl/commands/delete/privatenetwork"
	serverprivatenetwork "phoenixnap.com/pnapctl/commands/delete/server/privatenetwork"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a resource.",
	Long:  `Delete a resource.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

func init() {
	DeleteCmd.AddCommand(server.DeleteServerCmd)
	DeleteCmd.AddCommand(serverprivatenetwork.DeleteServerPrivateNetworkCmd)
	DeleteCmd.AddCommand(cluster.DeleteClusterCmd)
	DeleteCmd.AddCommand(tag.DeleteTagCmd)
	DeleteCmd.AddCommand(sshkey.DeleteSshKeyCmd)
	DeleteCmd.AddCommand(cluster.DeleteClusterCmd)
	DeleteCmd.AddCommand(privatenetwork.DeletePrivateNetworkCmd)
}
