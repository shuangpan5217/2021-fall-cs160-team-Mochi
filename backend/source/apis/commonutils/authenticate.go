package commonutils

import (
	"2021-fall-cs160-team-Mochi/backend/source/generated/models"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Payload struct {
	Username string
}

func GenerateJWT(username string) (tokenString string, errResp *models.ErrResponse) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Minute * 1440).Unix()
	/*
		In real world, we need to move "secret_token_key" into something like "secret.json" file of Kubernetes etc
	*/
	tokenString, err := token.SignedString([]byte(SecretTokenKey))
	if err != nil {
		errResp = GenerateErrResp(http.StatusInternalServerError, err.Error())
		return "", errResp
	}
	return tokenString, nil
}

func ExtractJWT(req *http.Request) (payload *Payload, errResp *models.ErrResponse) {
	if checkTestingEnv() {
		return &Payload{
			Username: "admin",
		}, nil
	}
	var tokenString = ""
	tokenHeader := req.Header.Get("Authorization")
	if strings.TrimSpace(tokenHeader) == "" {
		errResp = GenerateErrResp(http.StatusBadRequest, "Bearer token is not presented")
		return
	}
	if strings.HasPrefix(strings.ToLower(tokenHeader), "bearer") {
		tokenString = (strings.Split(tokenHeader, " "))[1]
	}
	if len(tokenString) == 0 {
		errResp = GenerateErrResp(http.StatusBadRequest, "Bearer token is not presented")
		return
	}

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("error decoding token")
		}
		return []byte(SecretTokenKey), nil
	})
	if err != nil {
		errResp = GenerateErrResp(http.StatusBadRequest, err.Error())
		return
	}
	if !token.Valid {
		errResp = GenerateErrResp(http.StatusBadRequest, "invalid tokeen")
		return
	}
	payload = &Payload{
		Username: claims["username"].(string),
	}
	return payload, nil
}

func checkTestingEnv() bool {
	return os.Getenv("TESTING") == "true"
}
