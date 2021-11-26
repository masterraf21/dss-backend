package middleware

import (
	// "github.com/geziyor/geziyor/middleware"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/masterraf21/dss-backend/auth"
	"github.com/masterraf21/dss-backend/configs"
)

func CreateBearerVerify() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(configs.Auth.Secret),
		Claims:     &auth.Claims{},
	})
}

// func TokenRefresherMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		// If the user is not authenticated (no user token data in the context), don't do anything.
// 		if c.Get("user") == nil {
// 			return next(c)
// 		}
// 		// Gets user token from the context.
// 		u := c.Get("user").(*jwt.Token)

// 		claims := u.Claims.(*Claims)

// 		// We ensure that a new token is not issued until enough time has elapsed.
// 		// In this case, a new token will only be issued if the old token is within
// 		// 15 mins of expiry.
// 		if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) < 15*time.Minute {
// 			// Gets the refresh token from the cookie.
// 			rc, err := c.Cookie(refreshTokenCookieName)
// 			if err == nil && rc != nil {
// 				// Parses token and checks if it valid.
// 				tkn, err := jwt.ParseWithClaims(rc.Value, claims, func(token *jwt.Token) (interface{}, error) {
// 					return []byte(GetRefreshJWTSecret()), nil
// 				})
// 				if err != nil {
// 					if err == jwt.ErrSignatureInvalid {
// 						c.Response().Writer.WriteHeader(http.StatusUnauthorized)
// 					}
// 				}

// 				if tkn != nil && tkn.Valid {
// 					// If everything is good, update tokens.
// 					_ = GenerateTokensAndSetCookies(&user.User{
// 						Name: claims.Name,
// 					}, c)
// 				}
// 			}
// 		}

// 		return next(c)
// 	}
// }
