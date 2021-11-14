package testUtils

import (
	"2021-fall-cs160-team-Mochi/backend/source/apis/commonutils"
	"2021-fall-cs160-team-Mochi/backend/source/apis/dbpackages"
	"2021-fall-cs160-team-Mochi/backend/source/generated/models"
	"2021-fall-cs160-team-Mochi/backend/source/generated/restapi"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"strings"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

func SetupTestServer() (db *gorm.DB, testServer *httptest.Server) {
	os.Setenv("TESTING", "true")
	apiHandler, err := restapi.GetAPIHandler()
	if err != nil {
		log.Fatal(err.Error())
	}
	testServer = httptest.NewServer(apiHandler)
	db = restapi.CreateDBConnection()
	return
}

func ShutDownTestServer(db *gorm.DB, testServer *httptest.Server) {
	db.Close()
	testServer.Close()
}

func PrepareUsers(db *gorm.DB) {
	user := &dbpackages.User{
		Username:   "admin",
		Password:   "password",
		Email:      "email",
		MiddleName: "middle",
		FirstName:  "first",
		LastName:   "last",
	}
	db.Save(&user)
}

func CleanUsers(db *gorm.DB) {
	db.Exec(`DELETE from users where username = ?`, "admin")
}

func PrePareNoteAndFile(db *gorm.DB) (note dbpackages.Note) {
	// get working directory
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err.Error())
	}
	// get file absolute path xx
	filePath := path.Join(workingDir, "..", "notes", "test_data", "test.pdf")
	errResp, noteRef := writeFile("admin", filePath, db)
	if errResp != nil {
		log.Fatal(errResp.ErrMessage)
	}

	note = dbpackages.Note{
		NoteID:        uuid.NewString(),
		NoteOwner:     "admin",
		Type:          "shared",
		Tag:           "math",
		Style:         "outline",
		NoteReference: noteRef,
	}
	err = db.Save(&note).Error
	if err != nil {
		log.Fatal(err.Error())
	}
	return
}

func CleanNoteAndFile(db *gorm.DB, note dbpackages.Note) {
	removeFileFormFileTable(db, note.NoteReference, note.NoteOwner)
	deleteNoteByID(db, note.NoteID, note.NoteOwner)
	removeFile(note.NoteReference)
}

func removeFile(noteRef string) {
	// remove file
	var path string
	path, _ = commonutils.GetMochiNoteFilesDir()
	_ = os.Remove(path + "/" + noteRef)
}

func removeFileFormFileTable(db *gorm.DB, fileName, username string) (errResp *models.ErrResponse) {
	result := db.Table(dbpackages.FileTable).Where("file_name = ? AND file_owner = ?", fileName, username).Delete(&dbpackages.File{})
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

func writeFile(username, fileAbsolutePath string, db *gorm.DB) (errResp *models.ErrResponse, fileName string) {
	// read file to bytes
	fileBytes, _ := ioutil.ReadFile(fileAbsolutePath)
	// check file type
	fileType := http.DetectContentType(fileBytes)
	if !strings.Contains(fileType, "application/pdf") {
		errResp = commonutils.GenerateErrResp(http.StatusBadRequest, "Only pdf file is allowed.")
		return
	}

	// file dir
	fileDir, errResp := commonutils.GetMochiNoteFilesDir()
	if errResp != nil {
		return
	}
	fileName = username + uuid.New().String() + ".pdf"
	exist, err := exists(fileName)
	if exist || err != nil {
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, commonutils.InternalServerError)
		return
	}
	// write dir and files
	err = os.MkdirAll(fileDir, 0777)
	if err != nil {
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
		return
	}
	err = os.WriteFile(fileDir+"/"+fileName, fileBytes, 0666)
	if err != nil {
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
		return
	}

	errResp = addFile(db, username, fileName)
	return
}

func addFile(db *gorm.DB, username, fileName string) (errResp *models.ErrResponse) {
	file := dbpackages.File{
		FileName:  fileName,
		FileOwner: username,
	}
	tx := db.Begin()
	err := db.Save(&file).Error
	if err != nil {
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
		return
	}
	if err = tx.Commit().Error; err != nil {
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
		return
	}
	return
}

func exists(name string) (bool, error) {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
