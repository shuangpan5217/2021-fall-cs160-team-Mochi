package comments

import (
	"2021-fall-cs160-team-Mochi/backend/source/apis/commonutils"
	"2021-fall-cs160-team-Mochi/backend/source/apis/dbpackages"
	"2021-fall-cs160-team-Mochi/backend/source/generated/models"
	"2021-fall-cs160-team-Mochi/backend/source/generated/restapi/operations/comments_v1"
	"net/http"

	"github.com/google/uuid"

	"github.com/go-openapi/runtime/middleware"
	"github.com/jinzhu/gorm"
)

func PostCommentsV1Handler(db *gorm.DB) comments_v1.PostCommentsV1HandlerFunc {
	return func(params comments_v1.PostCommentsV1Params) (responder middleware.Responder) {
		commentResp, errResp := processPostCommentRequest(db, params)
		if errResp != nil {
			switch errResp.StatusCode {
			case http.StatusBadRequest:
				return comments_v1.NewPostCommentsV1BadRequest().WithPayload(errResp)
			case http.StatusUnauthorized:
				return comments_v1.NewPostCommentsV1Unauthorized().WithPayload(errResp)
			case http.StatusForbidden:
				return comments_v1.NewPostCommentsV1Forbidden().WithPayload(errResp)
			case http.StatusNotFound:
				return comments_v1.NewPostCommentsV1NotFound().WithPayload(errResp)
			case http.StatusConflict:
				return comments_v1.NewPostCommentsV1Conflict().WithPayload(errResp)
			default:
				return comments_v1.NewPostCommentsV1InternalServerError().WithPayload(errResp)
			}
		}
		resp := comments_v1.NewPostCommentsV1OK()
		resp.SetPayload(commentResp)
		return resp
	}
}

func processPostCommentRequest(db *gorm.DB, params comments_v1.PostCommentsV1Params) (resp *models.CommentResponse, errResp *models.ErrResponse) {
	payload, errResp := commonutils.ExtractJWT(params.HTTPRequest)
	if errResp != nil {
		return
	}
	username := payload.Username
	body := *params.Body
	noteID := *body.NoteID
	convertedNoteID, err := uuid.Parse(noteID)
	if err != nil {
		return
	}
	content := *body.Content
	commentID := uuid.New()

	var comment dbpackages.Comment
	comment = dbpackages.Comment{
		CommentID: commentID,
		NoteID:    convertedNoteID,
		Username:  username,
		Content:   content,
	}
	tx := db.Begin()
	err = tx.Save(&comment).Error
	if err != nil {
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
		return
	}
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
		return
	}
	resp = &models.CommentResponse{
		NoteID: noteID,
	}
	return
}
