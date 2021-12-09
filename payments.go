package slackapi

import (
	"net/url"
)

type PaymentsBillingAddressesGetResponse struct {
	Response
	TeamID          string `json:"team_id"`
	CompanyName     string `json:"company_name"`
	Street1         string `json:"street1"`
	Street2         string `json:"street2"`
	City            string `json:"city"`
	State           string `json:"state"`
	Zip             string `json:"zip"`
	Plus4           string `json:"plus4"`
	Country         string `json:"country"`
	VatID           string `json:"vat_id"`
	AbnID           string `json:"abn_id"`
	TaxID           string `json:"tax_id"`
	IsBusiness      bool   `json:"is_business"`
	IsVatRegistered bool   `json:"is_vat_registered"`
	WaitingForVat   bool   `json:"waiting_for_vat"`
	Notes           string `json:"notes"`
	RegionalTaxID   string `json:"regional_tax_id"`
}

// PaymentsBillingAddressesGet is https://api.slack.com/methods/payments.billing.addresses.get
func (s *SlackAPI) PaymentsBillingAddressesGet() PaymentsBillingAddressesGetResponse {
	in := url.Values{}
	var out PaymentsBillingAddressesGetResponse
	if err := s.baseGET("/api/payments.billing.addresses.get", in, &out); err != nil {
		return PaymentsBillingAddressesGetResponse{Response: Response{Error: err.Error()}}
	}
	return out
}
