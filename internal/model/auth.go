package model

type RequestAuth struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type JWTPayload struct {
	UserId   string `json:"userId"`
	UserRole uint8  `json:"userRole"`
	//CompanyIDs UserCompanies `json:"company_ids,omitempty" db:"company_id"`
}

type Tokens struct {
	JWTToken     string
	RefreshToken string
}
