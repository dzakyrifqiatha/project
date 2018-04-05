package models

type Selling struct {
	Item        string `json:"Item"`
	Total       string `json:"Total"`
	Paid        string `json:"Paid"`
	OfficerCode string `json:"OfficerCode"`
}
