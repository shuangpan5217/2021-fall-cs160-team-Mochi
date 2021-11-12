package notes_test

import (
	"2021-fall-cs160-team-Mochi/backend/source/apis/commonutils"
	"2021-fall-cs160-team-Mochi/backend/source/apis/testUtils"
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"
)

var filedir string

const expected_count = 1

var count int32

func TestPostFileAPI(t *testing.T) {

	type Input struct {
		NoteFile string
	}

	type Expected struct {
		StatusCode int32

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
				NoteFile: "Homework_RIP_OSPF_GNS3.pdf",
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
				ErrorMessage: "only allow pdf file",
			},
		},
	}
	_, testServer := testUtils.SetupTestServer()

	for _, collection := range collections {
		t.Run(collection.name, func(t *testing.T) {
			apiurl := testServer.URL + "/v1/notes/file"

			fileDir, _ := os.UserHomeDir()
			fileName := collection.input.NoteFile
			filePath := path.Join(fileDir, fileName)

			file, _ := os.Open(filePath)
			defer file.Close()

			body := &bytes.Buffer{}
			writer := multipart.NewWriter(body)
			filedir = filepath.Base(file.Name())
			part, _ := writer.CreateFormFile("noteFile", fileDir)

			io.Copy(part, file)
			writer.Close()

			req, _ := http.NewRequest("POST", apiurl, body)
			req.Header.Add("Content-Type", writer.FormDataContentType())
			req.Header.Set("Authorization", "moke token")

			http.DefaultClient.Do(req)

		})
	}
	getCountFileFromDirectory()

	if count != expected_count {
		t.Errorf("Expected value=%d, got value=%d", expected_count, count)
	}
	if count > expected_count {
		t.Errorf("Expected value=%d, got value=%d", expected_count, count)
	}

}

func getCountFileFromDirectory() {
	path, _ := commonutils.GetMochiNoteFilesDir()
	files, _ := os.ReadDir(path)

	for _, file := range files {
		if strings.HasPrefix(file.Name(), "admin") {
			count += 1
		}
	}
}
