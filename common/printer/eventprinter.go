package printer

import (
	auditapisdk "gitlab.com/phoenixnap/bare-metal-cloud/go-sdk.git/auditapi"
	"phoenixnap.com/pnap-cli/common/models/auditmodels"
	"phoenixnap.com/pnap-cli/common/models/tables"
)

func PrintEventListResponse(events []auditapisdk.Event, commandName string) error {
	eventListToPrint := PrepareEventListForPrinting(events)
	return MainPrinter.PrintOutput(eventListToPrint, commandName)
}

func PrepareEventForPrinting(event auditapisdk.Event) interface{} {
	table := OutputIsTable()

	switch {
	case table:
		return tables.ToEventTable(event)
	default:
		return auditmodels.EventFromSdk(&event)
	}
}

func PrepareEventListForPrinting(events []auditapisdk.Event) []interface{} {
	var eventList []interface{}

	for _, event := range events {
		eventList = append(eventList, PrepareEventForPrinting(event))
	}

	return eventList
}
