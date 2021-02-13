package handlers

import (
	"github.com/auth/backend/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/fatih/structs"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"time"
)

//// jwtCustomClaims are custom claims extending default ones.
//type jwtCustomClaims struct {
//	Version uint64 `json:"version"`
//	Id uint64 `json:"id"`
//	jwt.StandardClaims
//}

func (h *HTTPHandler) GenerateToken(user *model.User) (string, error) {
		// Create token
		token := jwt.New(jwt.SigningMethodHS256)

		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["id"] = user.UserID
		claims["name"] = user.FirstName
		//claims["admin"] = user.Admin
		claims["exp"] = h.clock.Now().Add(time.Minute * time.Duration(60)).Unix()

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return "", err
		}
		return t, nil
	}


type ResponseProviders struct {
	// The error message string
	Message string `json:"message"`
	Providers []string `json:"providers"`
}

// return a list of available providers
func (h *HTTPHandler) Providers(context echo.Context) error {
	pro := &ResponseProviders{
		Message:   "Success",
		Providers: h.providers.ProviderIndex.Providers,
	}
	return context.JSON(http.StatusOK, pro)
}

// 1 - get provider origin url param - ok
// 2 - get access token req -
// 3 - get user from req - ok
// 4 - check if access token is valid token
// 5 - check if the user already exists - ok
//   -- if exists - go to 6
//   -- if not exists - create a new user
// 6 - generate a access token and a user id then return them to frontend (here would be great to use as crfs http ready only token) - ok
func (h *HTTPHandler) Login(context echo.Context) error {
	log.Println("Login")
	//log.Println(context.)
	provider := context.Request().Header.Get("provider")
	receivedUser := new(model.GithubUser)
	if provider == "github" {
		if err := context.Bind(&receivedUser); err != nil {
			log.Println("111")
			return context.JSON(http.StatusInternalServerError, "could not bind user")
		}
	 } else {
		log.Println("222")
		return context.JSON(http.StatusInternalServerError, "unsupported provider")
	}

	if !h.repository.ExistsUser(receivedUser.ID) {
		log.Println("login with new user")
		h.repository.AddUser(model.User{
			RawData:           structs.Map(receivedUser),
			Provider:          "github",
			Email:             receivedUser.Email,
			Name:              receivedUser.Name,
			FirstName:         "",
			LastName:          "",
			NickName:          "",
			Description:       "",
			UserID:            receivedUser.ID,
			AvatarURL:         receivedUser.Image,
			Location:          "",
			AccessToken:       "",
			AccessTokenSecret: "",
			RefreshToken:      "",
			ExpiresAt:         time.Time{},
			IDToken:           "",
		})
	} else {
		log.Println("User alredy exists")
	}

	// get the user from our database
	user, err := h.repository.GetUser(receivedUser.ID)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, "could not get user from database")
	}
	log.Println(user)

	// generate a token
	t, err := h.GenerateToken(user)
	log.Println("generated token:", t)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, "could not generate token")
	}

	return context.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}
