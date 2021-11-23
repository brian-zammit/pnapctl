package quotas

import (
	"fmt"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/common/client/bmcapi"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
	"phoenixnap.com/pnap-cli/common/models/bmcapimodels"
)

// Filename is the filename from which to retrieve a complex object
var Filename string

var commandName = "request-edit quota"

// RequestEditQuotaCmd is the command for requesting a quota modification.
var RequestEditQuotaCmd = &cobra.Command{
	Use:          "quota QUOTA_ID",
	Short:        "Submit a quota modification request.",
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	Long: `Submit a quota modification request.

Requires a file (yaml or json) containing the information needed to submit the request.`,
	Example: `# modify an existing quota as per quotaModificationRequest.yaml
pnapctl request-edit quota  --filename ~/quotaEditRequest.yaml

# quotaEditRequest.yaml
limit: 75
reason: "My current limit is not enough."`,
	RunE: func(cmd *cobra.Command, args []string) error {
		quotaEditRequest, err := bmcapimodels.CreateQuotaEditRequestFromFile(Filename, commandName)
		if err != nil {
			return err
		}

		httpResponse, err := bmcapi.Client.QuotaEditById(args[0], *quotaEditRequest)

		if err != nil {
			return ctlerrors.GenericFailedRequestError(err, commandName, ctlerrors.ErrorSendingRequest)
		} else if httpResponse.StatusCode != 202 {
			return ctlerrors.HandleBMCError(httpResponse, commandName)
		}

		fmt.Println("Quota Edit Limit Request Accepted.")
		return nil
	},
}

func init() {
	RequestEditQuotaCmd.Flags().StringVarP(&Filename, "filename", "f", "", "File containing required information for creation")
	RequestEditQuotaCmd.MarkFlagRequired("filename")
}
