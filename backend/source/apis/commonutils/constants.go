package commonutils

import (
	"2021-fall-cs160-team-Mochi/backend/source/generated/models"
	"net/http"
	"os"
)

// http error
const (
	InternalServerError = "Internal Server Error"
)

// jwt
const (
	SecretTokenKey = "secret_token_key"
)

func GetFileDir(username string) (path string, errResp *models.ErrResponse) {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		errResp = GenerateErrResp(http.StatusInternalServerError, err.Error())
		return
	}
	return userHomeDir + "/mochiNote/" + username, nil
}

func GetMochiNoteFilesDir() (path string, errResp *models.ErrResponse) {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		errResp = GenerateErrResp(http.StatusInternalServerError, err.Error())
		return
	}
	return userHomeDir + "/mochiNote", nil
}

func GetFileReferende(fileName, username string) string {
	return "/mochiNote/" + username + "/" + fileName
}
