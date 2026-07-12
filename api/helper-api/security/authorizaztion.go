package security

import (
	"fmt"
	"time"

	dtoM "taskManager/db/model"
	"taskManager/envconf"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// Claims defines the field(s) for claims jwt.
type Claims struct {
	Role uuid.UUID        `json:"role"`
	Iss  string           `json:"iss"`
	Sub  uuid.UUID        `json:"sub"`
	Aud  string           `json:"aud"`
	Iat  *jwt.NumericDate `json:"iat"`
	Exp  *jwt.NumericDate `json:"exp"`
	jwt.RegisteredClaims
}

// CreatingToken for Creating a token jwt.
func CreatingToken(user dtoM.User) (string, error) {
	currentTime := time.Now()
	env, err := envconf.EnvSetting()
	if err != nil {
		return "", fmt.Errorf("something error")
	}

	// Note: (Akbar): For right now we using just firs assign acount
	role := uuid.Nil
	if len(user.UserRoles) > 0 {
		role = user.UserRoles[0].Role.ID
	}

	claims := &Claims{
		Iss:  "http://api",
		Sub:  user.ID,
		Aud:  "api",
		Iat:  jwt.NewNumericDate(currentTime),
		Exp:  jwt.NewNumericDate(currentTime.Add(time.Hour)),
		Role: role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenResult, err := token.SignedString([]byte(env.SecretKeyJwt))
	if err != nil {
		return "", fmt.Errorf("something error")
	}

	return tokenResult, nil
}

// ValidateToken for validation jwt.
func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	env, err := envconf.EnvSetting()
	if err != nil {
		return nil, fmt.Errorf("something error")
	}

	secretKey := []byte(env.SecretKeyJwt)

	tokenResult, errToken := jwt.ParseWithClaims(tokenString, &Claims{}, func(_ *jwt.Token) (any, error) {
		return secretKey, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	if errToken != nil {
		return nil, fmt.Errorf("token not valid")
	}

	claims, ok := tokenResult.Claims.(*Claims)
	if !ok || !tokenResult.Valid {
		return nil, fmt.Errorf("invalid token claims")
	}

	result := jwt.MapClaims{
		"sub":  claims.Sub,
		"role": claims.Role,
	}

	return result, nil
}
