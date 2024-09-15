package service

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/SevvyP/todo_web_v1/internal/middleware"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

type TokenService interface {
	GenerateToken() (string, error)
}

type tokenService struct {
	Config *middleware.AuthConfig
}

func NewTokenService(config *middleware.AuthConfig) *tokenService {
	return &tokenService{Config: config}
}

func (t *tokenService) GenerateToken() (string, error) {
	url := "https://" + t.Config.Domain + "/oauth/token"

	payload := strings.NewReader("grant_type=client_credentials&client_id=" + t.Config.ClientID + "&client_secret=" + t.Config.Secret + "&audience=" + t.Config.Audience)

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return "", err
	}

	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	token := TokenResponse{}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	if res.StatusCode != 200 {
		return "", errors.New("Failed to get token from server status code: " + res.Status + " body: " + string(body))
	}
	err = json.Unmarshal(body, &token)
	if err != nil {
		return "", err
	}
	return token.AccessToken, nil

}
