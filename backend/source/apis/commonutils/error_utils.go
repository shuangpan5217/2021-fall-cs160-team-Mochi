package commonutils

import (
	"2021-fall-cs160-team-Mochi/backend/source/generated/models"
)

func GenerateErrResp(httpCode int32, errMessage string) (errResp *models.ErrResponse) {
	return &models.ErrResponse{
		StatusCode: httpCode,
		ErrMessage: errMessage,
	}
}

func GenerateSuccessResp(httpCode int32, successMessage string) (successResp *models.SuccessResponse) {
	return &models.SuccessResponse{
		Message:    successMessage,
		StatusCode: httpCode,
	}
}
