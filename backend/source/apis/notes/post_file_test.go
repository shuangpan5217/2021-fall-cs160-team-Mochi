package notes_test

import (
	"2021-fall-cs160-team-Mochi/backend/source/apis/commonutils"
	"2021-fall-cs160-team-Mochi/backend/source/apis/dbpackages"
	"2021-fall-cs160-team-Mochi/backend/source/apis/testUtils"
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/jinzhu/gorm"
)

const (
	expected_count = 2
)

var (
	count int32
)

func TestPostFileAPI(t *testing.T) {

	type Input struct {
		NoteFile string
	}

	type Expected struct {
		StatusCode   int32
		ErrorMessage string
	}
	type TestCollection struct {
		name     string
		input    Input
		expected Expected
	}

	collections := []TestCollection{
		{
			name: "postfile-success",
			input: Input{
				NoteFile: "test.pdf",
			},
			expected: Expected{
				StatusCode: http.StatusOK,
			},
		},
		{
			name: "postfile-success-2",
			input: Input{
				NoteFile: "test.pdf",
			},
			expected: Expected{
				StatusCode: http.StatusOK,
			},
		},
		{
			name: "postnonpdffile-badrequest",
			input: Input{
				NoteFile: "test.txt",
			},
			expected: Expected{
				StatusCode:   http.StatusBadRequest,
				ErrorMessage: "Only pdf file is allowed.",
			},
		},
	}
	db, testServer := testUtils.SetupTestServer()

	for _, collection := range collections {
		t.Run(collection.name, func(t *testing.T) {
			apiurl := testServer.URL + "/v1/notes/file"

			// get working directory
			workingDir, err := os.Getwd()
			if err != nil {
				t.Errorf("Encounter error: %s", err.Error())
			}
			// get file absolute path xx
			filePath := path.Join(workingDir, "test_data", collection.input.NoteFile)

			// open file
			file, err := os.Open(filePath)
			if err != nil {
				t.Errorf("Encounter error: %s", err.Error())
			}
			defer file.Close()

			body := &bytes.Buffer{}
			// create a writer
			writer := multipart.NewWriter(body)
			// create form data
			part, err := writer.CreateFormFile("noteFile", filePath)
			if err != nil {
				t.Errorf("Encounter error: %s", err.Error())
			}
			_, err = io.Copy(part, file)
			if err != nil {
				t.Errorf("Encounter error: %s", err.Error())
			}
			writer.Close()

			// cretae request
			req, err := http.NewRequest("POST", apiurl, body)
			if err != nil {
				t.Errorf("Encounter error: %s", err.Error())
			}

			// send request
			req.Header.Add("Content-Type", writer.FormDataContentType())
			req.Header.Set("Authorization", "moke token")
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Errorf("Encounter error: %s", err.Error())
			}
			if resp.StatusCode != int(collection.expected.StatusCode) {
				t.Errorf("Excepted status code - %d Got - %d", collection.expected.StatusCode, resp.StatusCode)
			}

			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("Encounter error: %s", err.Error())
			}
			defer resp.Body.Close()

			m := make(map[string]interface{})
			err = json.Unmarshal(bodyBytes, &m)
			if err != nil {
				t.Errorf("Encounter error: %s", err.Error())
			}
			if collection.expected.ErrorMessage != "" {
				if !strings.Contains(m["errMessage"].(string), collection.expected.ErrorMessage) {
					t.Errorf("Expected error message=%s, Got error message=%s", collection.expected.ErrorMessage, m["errMessage"].(string))
				}
			}
		})
	}
	// check file counts
	GetCountFileFromDirectory(t)
	if count != expected_count {
		t.Errorf("Expected value=%d, got value=%d", expected_count, count)
	}
	// remove files
	CleanCreatedFiles(t, db)

	//shutdown
	testUtils.ShutDownTestServer(db, testServer)
}

func GetCountFileFromDirectory(t *testing.T) {
	funcName := "GetCountFileFromDirectory: "
	path, errResp := commonutils.GetMochiNoteFilesDir()
	if errResp != nil {
		t.Errorf(funcName, errResp.ErrMessage)
	}
	files, err := os.ReadDir(path)
	if err != nil {
		t.Errorf(funcName, err.Error())
	}

	for _, file := range files {
		if strings.HasPrefix(file.Name(), "admin") {
			count += 1
		}
	}
}

func CleanCreatedFiles(t *testing.T, db *gorm.DB) {
	funcName := "CleanCreatedFiles: "
	path, errResp := commonutils.GetMochiNoteFilesDir()
	if errResp != nil {
		t.Errorf(funcName, errResp.ErrMessage)
	}
	files, err := os.ReadDir(path)
	if err != nil {
		t.Errorf(funcName, err.Error())
	}

	for _, file := range files {
		if strings.HasPrefix(file.Name(), "admin") {
			err = os.Remove(path + "/" + file.Name())
			if err != nil {
				t.Errorf(funcName, err.Error())
			}
		}
	}

	db.Table(dbpackages.FileTable).Where("file_owner = ?", "admin").Delete(&dbpackages.File{})
}
