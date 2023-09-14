package model

type CompanyRequest struct {
	Company
	PlatesUse `json:"platesUsed,omitempty"`
}

type PlatesUse []uint8
