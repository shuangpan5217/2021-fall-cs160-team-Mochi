package usermgmt_test

import (
	"2021-fall-cs160-team-Mochi/backend/source/apis/dbpackages"
	"2021-fall-cs160-team-Mochi/backend/source/apis/testUtils"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestUpdateUserInfoAPI(t *testing.T) {

	type Input struct {
		Description string `json:"description"`
		Email       string `json:"email"`
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"`
		MiddleName  string `json:"middle_name"`
	}

	type Expected struct {
		StatusCode   int32
		Description  string
		Email        string
		FirstName    string
		LastName     string
		MiddleName   string
		ErrorMessage string
	}
	type TestCollection struct {
		name     string
		input    Input
		expected Expected
	}

	collections := []TestCollection{
		{
			name: "UpdateDescription-success",
			input: Input{
				Description: "new description",
			},
			expected: Expected{
				Description: "new description",
				StatusCode:  http.StatusOK,
			},
		},
		{
			name: "UpdateEmail-success",
			input: Input{
				Email: "new email",
			},
			expected: Expected{
				Email:      "new email",
				StatusCode: http.StatusOK,
			},
		},
		{
			name: "UpdateFirstName-success",
			input: Input{
				FirstName: "new first name",
			},
			expected: Expected{
				FirstName:  "new first name",
				StatusCode: http.StatusOK,
			},
		},
		{
			name: "UpdateLastName-success",
			input: Input{
				LastName: "new last name",
			},
			expected: Expected{
				LastName:   "new last name",
				StatusCode: http.StatusOK,
			},
		},
		{
			name: "UpdateDescription-success",
			input: Input{
				MiddleName: "new middle name",
			},
			expected: Expected{
				MiddleName: "new middle name",
				StatusCode: http.StatusOK,
			},
		},
		{
			name: "UpdateAll-success",
			input: Input{
				Description: "update all description",
				Email:       "update all email",
				FirstName:   "update all first name",
				MiddleName:  "update all middle name",
				LastName:    "update all last name",
			},
			expected: Expected{
				Description: "update all description",
				Email:       "update all email",
				FirstName:   "update all first name",
				MiddleName:  "update all middle name",
				LastName:    "update all last name",
				StatusCode:  http.StatusOK,
			},
		},
	}

	db, testServer := testUtils.SetupTestServer()
	testUtils.PrepareUsers(db)

	for _, collection := range collections {
		t.Run(collection.name, func(t *testing.T) {
			url := testServer.URL + "/v1/user"

			bytesBody, err := json.Marshal(collection.input)
			if err != nil {
				t.Errorf(err.Error())
			}

			req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(bytesBody))
			req.Header.Set("Authorization", "moke token")
			req.Header.Add("content-type", "application/json")
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

			// get relfection values of Input struct
			inputValues := reflect.ValueOf(collection.input)
			numOfInputFields := inputValues.NumField()

			// get relfection vaules of Expected struct
			expectedValues := reflect.ValueOf(collection.expected)
			numOfExpectedFields := expectedValues.NumField()

			if resp.StatusCode != int(collection.expected.StatusCode) {
				t.Errorf("Excepted status code - %d Got - %d", collection.expected.StatusCode, resp.StatusCode)
			} else {
				user := dbpackages.User{}
				db.Table("users").Where("username = ?", "admin").First(&user)
				if resp.StatusCode == int(http.StatusOK) {
					for i := 0; i < numOfInputFields; i++ {
						// if input field is not empty string, check result
						if inputValues.Field(i).String() != "" {
							// get field name of input struct
							fieldName := inputValues.Type().Field(i).Name
							tagName := inputValues.Type().Field(i).Tag.Get("json")

							for j := 0; j < numOfExpectedFields; j++ {
								if expectedValues.Type().Field(j).Name == fieldName {
									if m[tagName].(string) != expectedValues.Field(j).String() {
										t.Errorf("Expected value=%s, got value=%s", expectedValues.Field(j).String(), m[tagName].(string))
									}
									break
								}
							}
						}
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
