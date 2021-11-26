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

// func BearerVerify(rsaPublicKey *rsa.PublicKey, mustAuthorized bool) echo.MiddlewareFunc {
// 	return func(next echo.HandlerFunc) echo.HandlerFunc {
// 		return func(c echo.Context) error {
// 			if os.Getenv("NO_TOKEN") == "1" {
// 				return next(c)
// 			}

// 			auth := c.Request().Header.Get("Authorization")

// 			if len(auth) <= 0 {
// 				return helper.NewResponse(http.StatusUnauthorized, "authorization is empty", nil).WriteResponse(c, 0, nil)
// 			}

// 			splitToken := strings.Split(auth, " ")
// 			if len(splitToken) < 2 {
// 				return helper.NewResponse(http.StatusUnauthorized, "authorization is empty", nil).WriteResponse(c, 0, nil)
// 			}

// 			if splitToken[0] != "Bearer" {
// 				return helper.NewResponse(http.StatusUnauthorized, "authorization is invalid", nil).WriteResponse(c, 0, nil)
// 			}

// 			tokenStr := splitToken[1]
// 			token, err := jwt.ParseWithClaims(tokenStr, &BearerClaims{}, func(token *jwt.Token) (interface{}, error) {
// 				if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
// 					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 				}
// 				return rsaPublicKey, nil
// 			})

// 			refreshToken := c.QueryParam("refresh_token")
// 			if err != nil && refreshToken == "" && !strings.Contains(c.Path(), "token/refresh") {
// 				if strings.Contains(err.Error(), "token") && strings.Contains(err.Error(), "expired") {
// 					err = errors.New("Your session has expired. Please login again")
// 				}
// 				return helper.NewResponse(http.StatusUnauthorized, err.Error(), nil).WriteResponse(c, 0, nil)
// 			}

// 			if claims, ok := token.Claims.(*BearerClaims); err == nil && token.Valid && ok {
// 				if mustAuthorized {
// 					if claims.UserAuthorized {
// 						c.Set("token", token)
// 						return next(c)
// 					}
// 					fmt.Printf("%+v", claims)
// 					return helper.NewResponse(http.StatusUnauthorized, "Resource need an authorised user", nil).WriteResponse(c, 0, nil)
// 				}
// 				c.Set("token", token)
// 				return next(c)
// 			} else if ve, ok := err.(*jwt.ValidationError); ok {
// 				if refreshToken != "" && strings.Contains(c.Path(), "token/refresh") {
// 					c.Set("token", token)
// 					return next(c)
// 				}
// 				var errorStr string
// 				if ve.Errors&jwt.ValidationErrorMalformed != 0 {
// 					errorStr = fmt.Sprintf("Invalid token format: %s", tokenStr)
// 				} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
// 					errorStr = "Token has been expired"
// 				} else {
// 					errorStr = fmt.Sprintf("Token Parsing Error: %s", err.Error())
// 				}
// 				return helper.NewResponse(http.StatusUnauthorized, errorStr, nil).WriteResponse(c, 0, nil)
// 			} else {
// 				return helper.NewResponse(http.StatusUnauthorized, "Unknown token error", nil).WriteResponse(c, 0, nil)
// 			}
// 		}
// 	}
// }

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
