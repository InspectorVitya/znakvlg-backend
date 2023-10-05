package manager

import (
	"errors"
	"fmt"
	"github.com/InspectorVitya/znakvlg-backend/internal/model"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type JWTServices struct {
	signingKey string
}

func JWTAuthService(signingKey string) (*JWTServices, error) {
	if signingKey == "" {
		return nil, errors.New("empty signing key")
	}
	return &JWTServices{signingKey: signingKey}, nil
}

type authCustomClaimsAdmin struct {
	UserId   string `json:"userId"`
	UserRole uint8  `json:"userRole"`
	jwt.RegisteredClaims
}

func (service *JWTServices) GenerateToken(userID string, userRole uint8) (string, error) {

	claims := &model.AuthClaims{
		userID,
		userRole,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			Issuer:    "znakvlg",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(service.signingKey))
	if err != nil {
		return "", model.ErrInternalService
	}
	return t, err

}

func (service *JWTServices) ValidateToken(encodedToken string) (*model.AuthClaims, error) {
	token, err := jwt.ParseWithClaims(encodedToken, &model.AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(service.signingKey), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*model.AuthClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("invalid token")
	}
}
