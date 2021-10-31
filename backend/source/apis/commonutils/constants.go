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

func GetMochiNoteFilesDir() (path string, errResp *models.ErrResponse) {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		errResp = GenerateErrResp(http.StatusInternalServerError, err.Error())
		return
	}
	return userHomeDir + "/mochiNote", nil
}

func GetUserImagesDir() (path string, errResp *models.ErrResponse) {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		errResp = GenerateErrResp(http.StatusInternalServerError, err.Error())
		return
	}
	return userHomeDir + "/mochiNote/userImages", nil
}
