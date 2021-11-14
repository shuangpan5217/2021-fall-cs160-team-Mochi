package usermgmt_test

import (
	"2021-fall-cs160-team-Mochi/backend/source/apis/dbpackages"
	"2021-fall-cs160-team-Mochi/backend/source/apis/testUtils"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestUpdatePasswordAPI(t *testing.T) {

	type Input struct {
		Password string
	}

	type Expected struct {
		StatusCode      int32
		UpdatedPassword string
		ErrorMessage    string
	}
	type TestCollection struct {
		name     string
		input    Input
		expected Expected
	}

	collections := []TestCollection{
		{
			name: "allSpacesStringPasword-badrequest",
			input: Input{
				Password: "     ",
			},
			expected: Expected{
				StatusCode:   http.StatusBadRequest,
				ErrorMessage: "empty password is not allowed",
			},
		},
		{
			name: "updatepassword-success",
			input: Input{
				Password: "newpassword",
			},
			expected: Expected{
				StatusCode:      http.StatusOK,
				UpdatedPassword: "newpassword",
			},
		},
	}

	db, testServer := testUtils.SetupTestServer()
	testUtils.PrepareUsers(db)

	for _, collection := range collections {
		t.Run(collection.name, func(t *testing.T) {
			url := testServer.URL + "/v1/password/" + collection.input.Password

			req, err := http.NewRequest("PATCH", url, nil)
			req.Header.Set("Authorization", "moke token")
			if err != nil {
				t.Errorf("%s. Encountered error while creating request", err.Error())
			}
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Errorf("%s. Encountered error while sending request", err.Error())
			}

			body, err := ioutil.ReadAll(resp.Body)
			defer resp.Body.Close()
			if err != nil {
				t.Errorf("Error reading resp. %s", err.Error())
			}

			m := make(map[string]interface{})
			json.Unmarshal(body, &m)

			if resp.StatusCode != int(collection.expected.StatusCode) {
				t.Errorf("Excepted status code - %d Got - %d", collection.expected.StatusCode, resp.StatusCode)
			} else {
				user := dbpackages.User{}
				db.Table("users").Where("username = ?", "admin").First(&user)
				if resp.StatusCode == int(http.StatusOK) {
					if collection.expected.UpdatedPassword != user.Password {
						t.Errorf("Expected password=%s, Got password=%s", collection.expected.UpdatedPassword, user.Password)
					}
				} else {
					if !strings.Contains(m["errMessage"].(string), collection.expected.ErrorMessage) {
						t.Errorf("Expected error message=%s, Got error message=%s", collection.expected.ErrorMessage, m["errMessage"].(string))
					}
				}
			}
		})
	}

	testUtils.CleanUsers(db)
	testUtils.ShutDownTestServer(db, testServer)
}
