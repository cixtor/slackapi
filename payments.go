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

type PaymentsBillingAddressesFields struct {
	CompanyName     string `json:"company_name"`
	Street1         string `json:"street1"`
	Street2         string `json:"street2"`
	City            string `json:"city"`
	State           string `json:"state"`
	Zip             string `json:"zip"`
	Country         string `json:"country"`
	CountryCode     string `json:"country_code"`
	VatID           string `json:"vat_id"`
	AbnID           string `json:"abn_id"`
	TaxID           string `json:"regional_tax_id"`
	GeoCode         string `json:"geocode"`
	IsBusiness      bool   `json:"is_business"`
	IsCheckoutV2    bool   `json:"is_checkout_v2"`
	IsVatRegistered bool   `json:"is_vat_registered"`
	WaitingForVat   bool   `json:"waiting_for_vat"`
	Notes           string `json:"notes"`
}

type PaymentsBillingAddressesValidateAndSetInput struct {
	CheckoutStep string `json:"checkout_step"`
	PaymentsBillingAddressesFields
}

type PaymentsBillingAddressesValidateAndSetResponse struct {
	Response
	TeamID     int                            `json:"team_id"`
	DateCreate int                            `json:"date_create"`
	DateUpdate int                            `json:"date_update"`
	DateDelete int                            `json:"date_delete"`
	Fields     PaymentsBillingAddressesFields `json:"fields"`
}

// PaymentsBillingAddressesValidateAndSet is https://api.slack.com/methods/payments.billing.addresses.validateAndSet
func (s *SlackAPI) PaymentsBillingAddressesValidateAndSet(input PaymentsBillingAddressesValidateAndSetInput) PaymentsBillingAddressesValidateAndSetResponse {
	in := url.Values{}
	in.Add("checkout_step", input.CheckoutStep)
	in.Add("company_name", input.CompanyName)
	in.Add("street1", input.Street1)
	in.Add("street2", input.Street2)
	in.Add("city", input.City)
	in.Add("state", input.State)
	in.Add("zip", input.Zip)
	in.Add("country", input.Country)
	in.Add("country_code", input.CountryCode)
	in.Add("vat_id", input.VatID)
	in.Add("abn_id", input.AbnID)
	in.Add("regional_tax_id", input.TaxID)
	in.Add("geocode", input.GeoCode)
	if input.IsBusiness {
		in.Add("is_business", "true")
	} else {
		in.Add("is_business", "false")
	}
	if input.IsCheckoutV2 {
		in.Add("is_checkout_v2", "true")
	} else {
		in.Add("is_checkout_v2", "false")
	}
	if input.IsVatRegistered {
		in.Add("is_vat_registered", "true")
	} else {
		in.Add("is_vat_registered", "false")
	}
	if input.WaitingForVat {
		in.Add("waiting_for_vat", "true")
	} else {
		in.Add("waiting_for_vat", "false")
	}
	in.Add("notes", input.Notes)
	var out PaymentsBillingAddressesValidateAndSetResponse
	if err := s.baseFormPOST("/api/payments.billing.addresses.validateAndSet", in, &out); err != nil {
		return PaymentsBillingAddressesValidateAndSetResponse{Response: Response{Error: err.Error()}}
	}
	return out
}
