package web

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"canteen-app/internal/adapter/http/common"
	domUser "canteen-app/internal/domain/user"
	"canteen-app/internal/usecase"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	auth       common.AuthUseCase
	accessTTL  time.Duration
	refreshTTL time.Duration
	tokenSvc   usecase.TokenService
	validator  common.Validator
}

func NewAuthHandler(
	router *gin.Engine,
	auth common.AuthUseCase,
	accessTTL time.Duration,
	refreshTTL time.Duration,
	tokenSvc usecase.TokenService,
	validator common.Validator,
) {
	handler := &AuthHandler{
		auth:       auth,
		accessTTL:  accessTTL,
		refreshTTL: refreshTTL,
		tokenSvc:   tokenSvc,
		validator:  validator,
	}

	router.LoadHTMLGlob("internal/adapter/http/web/templates/*")

	router.GET("/register", handler.RegisterGET)
	router.POST("/register", CSRFMiddleware(), handler.RegisterPOST)

	router.GET("/login", handler.LoginGET)
	router.POST("/login", CSRFMiddleware(), handler.LoginPOST)

	router.GET("/home", AuthMiddleware(handler.tokenSvc), handler.HomeGET)
}

func (ah *AuthHandler) RegisterGET(c *gin.Context) {
	reason := getFlash(c, "flash_auth")
	csrfToken := setCsrfCookie(c)

	c.HTML(http.StatusOK, "register.html", gin.H{
		"reason":    reason,
		"csrfToken": csrfToken,
	})
}

type RegisterFormData struct {
	Login    string `validate:"required,min=2,max=50"`
	Password string `validate:"required,min=8,max=100"`
	Name     string `validate:"required,max=100,alpha"`
	Surname  string `validate:"required,max=100,alpha"`
	Role     string `validate:"required"`
}

func (ah *AuthHandler) RegisterPOST(c *gin.Context) {
	formData := RegisterFormData{}
	formData.Login = c.PostForm("login")
	formData.Name = c.PostForm("name")
	formData.Surname = c.PostForm("surname")
	formData.Password = c.PostForm("password")
	formData.Role = c.PostForm("role")

	if err := ah.validator.Struct(formData); err != nil {
		fmt.Println(err.Error())
		_, msg := common.ErrorToHTTP(common.ErrValidationError)
		redirectToAuthPage(c, "/register", msg)
		return
	}

	tokens, err := ah.auth.Register(formData.Login, formData.Password, formData.Name, formData.Surname, formData.Role)
	if err != nil {
		_, msg := common.ErrorToHTTP(err)
		redirectToAuthPage(c, "/register", msg)
		return
	}

	c.SetCookieData(&http.Cookie{
		Name:     "access_token",
		Value:    tokens.AccessToken,
		Path:     "/",
		Domain:   "",
		Expires:  time.Now().Add(ah.accessTTL),
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	c.Redirect(http.StatusSeeOther, "/home")
}

func (ah *AuthHandler) LoginGET(c *gin.Context) {
	reason := getFlash(c, "flash_auth")
	csrfToken := setCsrfCookie(c)
	c.HTML(http.StatusOK, "login.html", gin.H{
		"reason":    reason,
		"csrfToken": csrfToken,
	})
}

type LoginFormData struct {
	Login    string `validate:"required,max=50"`
	Password string `validate:"required,max=100"`
}

func (ah *AuthHandler) LoginPOST(c *gin.Context) {
	formData := LoginFormData{}
	formData.Login = c.PostForm("login")
	formData.Password = c.PostForm("password")

	if err := ah.validator.Struct(formData); err != nil {
		_, msg := common.ErrorToHTTP(common.ErrValidationError)
		redirectToAuthPage(c, "/login", msg)
		return
	}

	tokens, err := ah.auth.Login(formData.Login, formData.Password)
	if err != nil {
		_, msg := common.ErrorToHTTP(err)
		redirectToAuthPage(c, "/login", msg)
		return
	}

	c.SetCookieData(&http.Cookie{
		Name:     "access_token",
		Value:    tokens.AccessToken,
		Path:     "/",
		Domain:   "",
		Expires:  time.Now().Add(ah.accessTTL),
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	c.Redirect(http.StatusSeeOther, "/home")
}

func (ah *AuthHandler) HomeGET(c *gin.Context) {
	userID, ok := c.Get("userID")
	if !ok {
		_, msg := common.ErrorToHTTP(errors.New("internal server error"))
		redirectToAuthPage(c, "/login", msg)
		return
	}
	user, err := ah.auth.GetUserByID(userID.(domUser.UserID))
	if err != nil {
		_, msg := common.ErrorToHTTP(err)
		redirectToAuthPage(c, "/login", msg)
		return
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"name":    user.Name,
		"surname": user.Surname,
		"role":    user.Role,
	})
}
