package commonutils

import "2021-fall-cs160-team-Mochi/backend/source/generated/models"

func GenerateErrResp(httpCode int32, errMessage string) (errResp *models.ErrResponse) {
	return &models.ErrResponse{
		StatusCode: httpCode,
		ErrMessage: errMessage,
	}
}
