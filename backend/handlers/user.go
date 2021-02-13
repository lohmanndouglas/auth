package handlers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

// from commit 1780e6e31487655a46d27a8997061d07c60481ff
// 1 - get provider origin url param
// 2 - get access token req
// 3 - get user from req
// 4 - check if access token is valid token
// 5 - check if the user already exists
//   -- if exists - go to 6
//   -- if not exists - create a new user
// 6 - generate a access token and a user id then return them to frontend (here would be great to use as crfs http ready only token)
func (h *HTTPHandler) GetUser(context echo.Context) error {
	log.Println("GetUser handler")
	tokenString := context.Request().Header.Get("Authorization")
	log.Println("received token: ", tokenString)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("secret"), nil
	})
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	claims, claimsErr := token.Claims.(jwt.MapClaims)
	if !claimsErr || !token.Valid  {
		return context.JSON(http.StatusInternalServerError, "invalid token")
	}
	userId := claims["id"]
	log.Println("claim user id:", userId)
	user, err := h.repository.GetUserView(fmt.Sprintf("%v", userId))
	if err != nil {
		return context.JSON(http.StatusInternalServerError, fmt.Errorf("could not get user"))
	}
	return context.JSON(http.StatusOK, user)
}