package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/masterraf21/dss-backend/configs"
	"github.com/masterraf21/dss-backend/models"
)

// jwtCustomClaims are custom claims extending default ones.
// See https://github.com/golang-jwt/jwt for more examples
type Claims struct {
	ID uint32 `json:"id"`
	// Admin bool   `json:"admin"`
	jwt.StandardClaims
}

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	// AccessUUID   string
	// RefreshUUID  string
	AtExpires int64
	RtExpires int64
}

func generateAccessToken(user *models.User) (string, time.Time, error) {
	// Declare the expiration time of the token (1h).
	expirationTime := time.Now().Add(24 * time.Hour)

	return generateToken(user, expirationTime, []byte(configs.Auth.Secret))
}

func GenerateToken(user *models.User) (token *TokenDetails, err error) {
	accessToken, expAt, err := generateAccessToken(user)
	if err != nil {
		return
	}

	refreshToken, expRt, err := generateRefreshToken(user)
	if err != nil {
		return
	}

	token = &TokenDetails{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		AtExpires:    expAt.Unix(),
		RtExpires:    expRt.Unix(),
	}

	return
}

// Pay attention to this function. It holds the main JWT token generation logic.
func generateToken(user *models.User, expirationTime time.Time, secret []byte) (string, time.Time, error) {
	// Create the JWT claims, which includes the username and expiry time.
	claims := &Claims{
		ID: user.ID,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds.
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the HS256 algorithm used for signing, and the claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string.
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", time.Now(), err
	}

	return tokenString, expirationTime, nil
}

func generateRefreshToken(user *models.User) (string, time.Time, error) {
	// Declare the expiration time of the token - 24 hours.
	expirationTime := time.Now().Add(24 * time.Hour)

	return generateToken(user, expirationTime, []byte(configs.Auth.RefreshSecret))
}
