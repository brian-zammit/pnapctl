package poweroff

import (
	"bytes"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/pnapctl/client"
	"phoenixnap.com/pnap-cli/pnapctl/ctlerrors"
)

var P_OffCmd = &cobra.Command{
	Use:          "power-off",
	Short:        "Powers off a specific server.",
	Long:         "Powers off a specific server.",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		// If more than one argument is passed, report error and panic.
		if len(args) != 1 {
			return ctlerrors.InvalidNumberOfArgs(1, len(args), "power-off")
		}

		var resource = "servers/" + args[0] + "/actions/power-off"
		var response, err = client.MainClient.PerformPost(resource, bytes.NewBuffer([]byte{}))

		if err != nil {
			// Generic error with PerformPost
			return ctlerrors.GenericFailedRequestError("power-off")
		}

		return ctlerrors.Result("power-off").
			IfOk("Powered off successfully.").
			IfNotFound("Error: Server with ID " + args[0] + " not found").
			UseResponse(response)
	},
}

func init() {
}
