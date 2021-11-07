package notes

import (
	"2021-fall-cs160-team-Mochi/backend/source/apis/commonutils"
	"2021-fall-cs160-team-Mochi/backend/source/apis/dbpackages"
	"2021-fall-cs160-team-Mochi/backend/source/generated/models"
	"2021-fall-cs160-team-Mochi/backend/source/generated/restapi/operations/notes_v1"
	"net/http"
	"os"

	"github.com/go-openapi/runtime/middleware"
	"github.com/jinzhu/gorm"
)

func DeleteNoteByIdV1Handler(db *gorm.DB) notes_v1.DeleteNoteV1HandlerFunc {
	return func(params notes_v1.DeleteNoteV1Params) middleware.Responder {
		deleteNoteByIdResp, errResp := processdeleteNoteByIdRequest(db, params)
		if errResp != nil {
			switch errResp.StatusCode {
			case http.StatusBadRequest:
				return notes_v1.NewDeleteNoteV1BadRequest().WithPayload(errResp)
			case http.StatusUnauthorized:
				return notes_v1.NewDeleteNoteV1Unauthorized().WithPayload(errResp)
			case http.StatusForbidden:
				return notes_v1.NewDeleteNoteV1Forbidden().WithPayload(errResp)
			case http.StatusNotFound:
				return notes_v1.NewDeleteNoteV1NotFound().WithPayload(errResp)
			case http.StatusConflict:
				return notes_v1.NewDeleteNoteV1Conflict().WithPayload(errResp)
			default:
				return notes_v1.NewDeleteNoteV1InternalServerError().WithPayload(errResp)
			}
		}
		resp := notes_v1.NewDeleteNoteV1OK()
		resp.SetPayload(deleteNoteByIdResp)
		return resp
	}
}

func processdeleteNoteByIdRequest(db *gorm.DB, params notes_v1.DeleteNoteV1Params) (resp *models.NoteResponse, errResp *models.ErrResponse) {
	payload, errResp := commonutils.ExtractJWT(params.HTTPRequest)
	if errResp != nil {
		return
	}

	username := payload.Username

	// check if note exists
	note, err := checkIfNoteExists(db, params.ID)
	if gorm.IsRecordNotFoundError(err) {
		errResp = commonutils.GenerateErrResp(http.StatusNotFound, err.Error())
		return
	} else if err != nil {
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
		return
	}

	// delete note by username and id
	tx := db.Begin()
	errResp = deleteNoteByID(tx, params.ID, username)
	if errResp != nil {
		tx.Rollback()
		return
	}

	noteRef := note.NoteReference
	// check file owner
	file, err := checkFileOwner(tx, noteRef)
	if gorm.IsRecordNotFoundError(err) {

	} else if err != nil {
		tx.Rollback()
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
		return
	} else if file.FileOwner == username {
		// remove file
		var path string
		path, errResp = commonutils.GetMochiNoteFilesDir()
		if errResp != nil {
			tx.Rollback()
			return
		}
		err = os.Remove(path + "/" + noteRef)
		if os.IsNotExist(err) {

		} else if err != nil {
			errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
			tx.Rollback()
			return
		}
	}

	if file.FileName != "" {
		// remove filename and fileowner from files table
		errResp = removeFileFormFileTable(tx, noteRef)
		if errResp != nil {
			tx.Rollback()
			return
		}
	}

	// commit
	if err = tx.Commit().Error; err != nil {
		commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
		tx.Rollback()
		return
	}

	resp = &models.NoteResponse{
		NoteID: params.ID,
	}
	return
}

func checkFileOwner(db *gorm.DB, fileName string) (file dbpackages.File, err error) {
	err = db.Table(dbpackages.FileTable).Where("file_name = ?", fileName).First(&file).Error
	return file, err
}

func removeFileFormFileTable(db *gorm.DB, fileName string) (errResp *models.ErrResponse) {
	result := db.Table(dbpackages.FileTable).Where("file_name = ?", fileName).Delete(&dbpackages.File{})
	if result.Error != nil {
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, result.Error.Error())
	} else if result.RowsAffected < 1 {
		errResp = commonutils.GenerateErrResp(http.StatusNotFound, "record not found")
	}
	return
}

func deleteNoteByID(db *gorm.DB, noteId, username string) (errResp *models.ErrResponse) {
	result := db.Table(dbpackages.NoteTable).Where("note_owner = ? AND note_id = ?", username, noteId).Delete(&dbpackages.Note{})
	if result.Error != nil {
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, result.Error.Error())
	} else if result.RowsAffected < 1 {
		errResp = commonutils.GenerateErrResp(http.StatusNotFound, "record not found")
	}
	return
}
