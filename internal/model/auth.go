package model

import "github.com/golang-jwt/jwt/v4"

type RequestAuth struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type JWTPayload struct {
	UserID   string `json:"userId"`
	UserRole uint8  `json:"userRole"`
	//CompanyIDs UserCompanies `json:"company_ids,omitempty" db:"company_id"`
}

type Tokens struct {
	JWTToken     string
	RefreshToken string
}

type AuthClaims struct {
	UserId   string `json:"userId"`
	UserRole uint8  `json:"userRole"`
	jwt.RegisteredClaims
}
