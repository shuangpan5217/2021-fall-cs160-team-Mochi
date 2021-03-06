package comments

import (
	"2021-fall-cs160-team-Mochi/backend/source/apis/commonutils"
	"2021-fall-cs160-team-Mochi/backend/source/apis/dbpackages"
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
	payload, errResp := commonutils.ExtractJWT(params.HTTPRequest)
	if errResp != nil {
		return
	}
	username := payload.Username
	commentID := params.CommentID

	db = db.Table(dbpackages.CommentTable).Where("comment_id = ? and username = ?", commentID, username).Delete(&dbpackages.Comment{})
	if db.Error != nil {
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, db.Error.Error())
		return
	} else if db.RowsAffected < 1 {
		errResp = commonutils.GenerateErrResp(http.StatusNotFound, "record not found")
		return
	}
	resp = &models.CommentResponse{
		CommentID: commentID,
	}

	return
}
