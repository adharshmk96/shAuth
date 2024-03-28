package apihandler

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/adharshmk96/shAuth/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// Prototype: change oauthState to a random string
// store it in a cookie
var oauthState string = "state"
var googleOauthConfig *oauth2.Config

func getGoogleOAuthConfig() *oauth2.Config {
	if googleOauthConfig == nil {
		googleOauthConfig = &oauth2.Config{
			RedirectURL:  viper.GetString("oauth.google.redirect_url"),
			ClientID:     viper.GetString("oauth.google.client_id"),
			ClientSecret: viper.GetString("oauth.google.client_secret"),
			Scopes:       viper.GetStringSlice("oauth.google.scopes"),
			Endpoint:     google.Endpoint,
		}
	}
	return googleOauthConfig
}

type GoogleUserInfo struct {
	Email         string `json:"email"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	EmailVerified bool   `json:"email_verified"`
}

func (h *authHandler) GoogleLogin(c *fiber.Ctx) error {
	googleOauthConfig := getGoogleOAuthConfig()
	url := googleOauthConfig.AuthCodeURL(oauthState, oauth2.AccessTypeOffline)
	return c.Redirect(url, http.StatusTemporaryRedirect)
}

func (h *authHandler) GoogleLoginCallback(c *fiber.Ctx) error {
	state := c.Query("state")
	if state != oauthState {
		h.logger.Error(
			"google login callback: invalid state",
			slog.String("state", state),
		)
		return c.Status(http.StatusBadRequest).SendString("Login Failed")
	}

	googleOauthConfig := getGoogleOAuthConfig()
	code := c.Query("code")
	bgCtx := context.Background()
	token, err := googleOauthConfig.Exchange(bgCtx, code)
	if err != nil {
		h.logger.Error(
			"google login callback: code exchange failed",
			slog.String("error", err.Error()),
		)
		return c.Status(http.StatusUnauthorized).SendString("Login Failed")
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return c.Status(http.StatusUnauthorized).SendString("Failed getting user info: " + err.Error())
	}

	defer response.Body.Close()

	// decode response to save in storage
	// it registers the user if not already registered
	userInfo := &GoogleUserInfo{}
	err = json.NewDecoder(response.Body).Decode(userInfo)
	if err != nil {
		h.logger.Error(
			"google login callback: failed to decode user info",
			slog.String("error", err.Error()),
		)
		return c.Status(http.StatusInternalServerError).SendString("login failed")
	}

	claims := utils.NewClaims()
	(*claims)["email"] = userInfo.Email

	signedJWT, err := utils.EncodeJWT(claims)
	if err != nil {
		h.logger.Error(
			"google login callback: failed to encode JWT",
			slog.String("error", err.Error()),
		)
		return c.Status(http.StatusInternalServerError).SendString("login failed")
	}

	c.Cookie(&fiber.Cookie{
		Name:     viper.GetString("auth.cookie_name"),
		Value:    signedJWT,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
		SameSite: "Lax",
	})

	return c.Redirect(viper.GetString("auth.redirect_url"), http.StatusTemporaryRedirect)
}
