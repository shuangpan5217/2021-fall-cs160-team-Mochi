package comments

import (
	"2021-fall-cs160-team-Mochi/backend/source/apis/commonutils"
	"2021-fall-cs160-team-Mochi/backend/source/generated/models"
	"2021-fall-cs160-team-Mochi/backend/source/generated/restapi/operations/comments_v1"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/jinzhu/gorm"
)

func RemoveComnentV1Handler(db *gorm.DB) comments_v1.RemoveComnentV1HandlerFunc {
	return func(params comments_v1.RemoveComnentV1Params) (responder middleware.Responder) {
		commentResp, errResp := processRemoveCommentRequest(db, params)
		if errResp != nil {
			switch errResp.StatusCode {
			case http.StatusBadRequest:
				return comments_v1.NewRemoveComnentV1BadRequest().WithPayload(errResp)
			case http.StatusUnauthorized:
				return comments_v1.NewRemoveComnentV1Unauthorized().WithPayload(errResp)
			case http.StatusForbidden:
				return comments_v1.NewRemoveComnentV1Forbidden().WithPayload(errResp)
			case http.StatusNotFound:
				return comments_v1.NewRemoveComnentV1NotFound().WithPayload(errResp)
			case http.StatusConflict:
				return comments_v1.NewRemoveComnentV1Conflict().WithPayload(errResp)
			default:
				return comments_v1.NewRemoveComnentV1InternalServerError().WithPayload(errResp)
			}
		}
		resp := comments_v1.NewRemoveComnentV1OK()
		resp.SetPayload(commentResp)
		return resp
	}
}

func processRemoveCommentRequest(db *gorm.DB, params comments_v1.RemoveComnentV1Params) (resp *models.CommentResponse, errResp *models.ErrResponse) {
	_, errResp = commonutils.ExtractJWT(params.HTTPRequest)
	if errResp != nil {
		return
	}
	commentID := params.CommentID
	row := db.Raw("DELETE from comments where comment_id = ?", commentID).Row()
	err := row.Scan(row)
	if err != nil {
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
		return
	}
	resp = &models.CommentResponse{
		CommentID: commentID,
	}

	return
}
