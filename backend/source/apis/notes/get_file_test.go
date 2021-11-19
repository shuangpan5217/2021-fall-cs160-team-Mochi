package notes_test

import (
	"2021-fall-cs160-team-Mochi/backend/source/apis/testUtils"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestGetFileAPI(t *testing.T) {

	db, testServer := testUtils.SetupTestServer()
	testUtils.PrepareUsers(db)
	note := testUtils.PrePareNoteAndFile(db)

	type Input struct {
		Path string
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
			name: "getFile-notfound",
			input: Input{
				Path: "test.pdf",
			},
			expected: Expected{
				StatusCode:   http.StatusNotFound,
				ErrorMessage: "record not found",
			},
		},
		{
			name: "getFile-success",
			input: Input{
				Path: note.NoteReference,
			},
			expected: Expected{
				StatusCode: http.StatusOK,
			},
		},
	}

	for _, collection := range collections {
		t.Run(collection.name, func(t *testing.T) {
			apiurl := testServer.URL + "/v1/notes/file/" + collection.input.Path

			// cretae request
			req, err := http.NewRequest("GET", apiurl, nil)
			if err != nil {
				t.Errorf("Encounter error: %s", err.Error())
			}

			// send request
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

	testUtils.CleanNoteAndFile(db, note)
	testUtils.CleanUsers(db)
	testUtils.ShutDownTestServer(db, testServer)
}
