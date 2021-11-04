package tables

import (
	bmcapisdk "gitlab.com/phoenixnap/bare-metal-cloud/go-sdk.git/bmcapi"
	models "phoenixnap.com/pnap-cli/common/models"
)

type Quota struct {
	Id                           string   `header:"ID"`
	Name                         string   `header:"Name"`
	Description                  string   `header:"Description"`
	Status                       string   `header:"Status"`
	Limit                        int32    `header:"Limit"`
	Unit                         string   `header:"Unit"`
	Used                         int32    `header:"Used"`
	QuotaEditLimitRequestDetails []string `header:"Quota Edit Limit Request Details"`
}

func ToQuotaTable(quota bmcapisdk.Quota) Quota {
	return Quota{
		Id:                           quota.Id,
		Name:                         quota.Name,
		Description:                  quota.Description,
		Status:                       quota.Status,
		Limit:                        quota.Limit,
		Unit:                         quota.Unit,
		Used:                         quota.Used,
		QuotaEditLimitRequestDetails: models.QuotaEditLimitRequestDetailsToTableString(quota.QuotaEditLimitRequestDetails),
	}
}
