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
	ID    uint32 `json:"id"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

// GenerateTokensAndSetCookies generates jwt token and saves it to the http-only cookie.
// func GenerateTokensAndSetCookies(user *models.User, c echo.Context) error {
// 	accessToken, exp, err := generateAccessToken(user)
// 	if err != nil {
// 		return err
// 	}

// 	setTokenCookie(accessTokenCookieName, accessToken, exp, c)
// 	setUserCookie(user, exp, c)

// 	return nil
// }

func GenerateAccessToken(user *models.User) (string, time.Time, error) {
	// Declare the expiration time of the token (1h).
	expirationTime := time.Now().Add(1 * time.Hour)

	return generateToken(user, expirationTime, []byte(configs.Auth.Secret))
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

// // Here we are creating a new cookie, which will store the valid JWT token.
// func setTokenCookie(name, token string, expiration time.Time, c echo.Context) {
// 	cookie := new(http.Cookie)
// 	cookie.Name = name
// 	cookie.Value = token
// 	cookie.Expires = expiration
// 	cookie.Path = "/"
// 	// Http-only helps mitigate the risk of client side script accessing the protected cookie.
// 	cookie.HttpOnly = true

// 	c.SetCookie(cookie)
// }

// // Purpose of this cookie is to store the user's name.
// func setUserCookie(user *models.User, expiration time.Time, c echo.Context) {
// 	cookie := new(http.Cookie)
// 	cookie.Name = "user"
// 	cookie.Value = user.Name
// 	cookie.Expires = expiration
// 	cookie.Path = "/"
// 	c.SetCookie(cookie)
// }

// // JWTErrorChecker will be executed when user try to access a protected path.
// func JWTErrorChecker(err error, c echo.Context) error {
// 	// Redirects to the signIn form.
// 	return c.Redirect(http.StatusMovedPermanently, c.Echo().Reverse("userSignInForm"))
// }

// After finishing the main JWT token functionality, let’s add the SignIn controllers, which will handle user authentication. First, we need to add the new routers inside main() function:

//     e.GET("/user/signin", controllers.SignInForm()).Name = "userSignInForm"
// 	e.POST("/user/signin", controllers.SignIn())
// In the code below I created a user package with the user structure and a function that loads a dummy user from imaginary database. We gonna need it in our controllers to process and validate user data.

// package user

// import "golang.org/x/crypto/bcrypt"

// type User struct {
// 	Password string `json:"password" form:"password"`
// 	Name string `json:"name" form:"name"`
// }

// func LoadTestUser() *User {
//     // Just for demonstration purpose, we create a user with the encrypted "test" password.
//     // In real-world applications, you might load the user from the database by specific parameters (email, username, etc.)
// 	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("test"), 8)
// 	return &User{Password: string(hashedPassword), Name: "Test user"}
// }
// After this, we will create a controllers package, where we add SignInForm() and SignIn() functions:

// package controllers

// import (
// 	"github.com/alexsergivan/blog-examples/authentication/auth"
// 	"github.com/alexsergivan/blog-examples/authentication/user"
// 	"github.com/labstack/echo/v4"
// 	"golang.org/x/crypto/bcrypt"
// 	"html/template"
// 	"net/http"
// 	"path"
// )

// // SignInForm responsible for signIn Form rendering.
// func SignInForm() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		fp := path.Join("templates", "signIn.html")
// 		tmpl, err := template.ParseFiles(fp)
// 		if err != nil {
// 			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
// 		}
// 		if err := tmpl.Execute(c.Response().Writer, nil); err != nil {
// 			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
// 		}

// 		return nil
// 	}
// }

// // SignIn will be executed after SignInForm submission.
// func SignIn() echo.HandlerFunc {
// 	return func(c echo.Context) error {
//         // Load our "test" user.
// 		storedUser := user.LoadTestUser()
//         // Initiate a new User struct.
//         u := new(user.User)
//         // Parse the submitted data and fill the User struct with the data from the SignIn form.
// 		if err := c.Bind(u); err != nil {
// 			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
// 		}
// 		// Compare the stored hashed password, with the hashed version of the password that was received.
// 		if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(u.Password)); err != nil {
// 			// If the two passwords don't match, return a 401 status.
// 			return echo.NewHTTPError(http.StatusUnauthorized, "Password is incorrect")
//         }
//         // If password is correct, generate tokens and set cookies.
// 		err := auth.GenerateTokensAndSetCookies(storedUser, c)

// 		if err != nil {
// 			return echo.NewHTTPError(http.StatusUnauthorized, "Token is incorrect")
// 		}

// 		return c.Redirect(http.StatusMovedPermanently, "/admin")
// 	}
// }
// Now, in the /templates folder we need to create a signIn.html template with the simple SignIn form:

// <!DOCTYPE html>
// <html lang="en">
//   <form class="w-full" method="post" action="/user/signin">
//     <label for="password">Password:</label>
//     <input type="password" id="password" name="password">
//     <button type="submit">Sign In</button>
//   </form>
// </html>
// Let’s also modify the admin controller in controllers/admin.go:

// package controllers

// import (
// 	"fmt"
// 	"github.com/labstack/echo/v4"
// 	"net/http"
// )

// func Admin() echo.HandlerFunc {
// 	return func(c echo.Context) error {
//         // Gets user cookie.
// 		userCookie, _ := c.Cookie("user")
// 		return c.String(http.StatusOK, fmt.Sprintf("Hi, %s! You have been authenticated!", userCookie.Value))
// 	}
// }
