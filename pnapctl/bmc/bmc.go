package bmc

import (
	"os"

	"phoenixnap.com/pnap-cli/pnapctl/bmc/delete"
	"phoenixnap.com/pnap-cli/pnapctl/bmc/deploy"
	"phoenixnap.com/pnap-cli/pnapctl/bmc/reboot"
	"phoenixnap.com/pnap-cli/pnapctl/bmc/reset"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/pnapctl/bmc/get"
	"phoenixnap.com/pnap-cli/pnapctl/bmc/poweroff"
	"phoenixnap.com/pnap-cli/pnapctl/bmc/poweron"
	"phoenixnap.com/pnap-cli/pnapctl/bmc/shutdown"
)

var BmcCmd = &cobra.Command{
	Use:   "bmc",
	Short: "Bare Metal Cloud - Short",
	Long:  "Bare Metal Cloud - Long",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

func init() {
	BmcCmd.AddCommand(poweroff.P_OffCmd)
	BmcCmd.AddCommand(get.GetCmd)
	BmcCmd.AddCommand(poweron.P_OnCmd)
	BmcCmd.AddCommand(shutdown.ShutdownCmd)
	BmcCmd.AddCommand(reset.ResetCmd)
	BmcCmd.AddCommand(deploy.DeployCmd)
	BmcCmd.AddCommand(delete.DeleteCmd)
	BmcCmd.AddCommand(reboot.RebootCmd)
}
