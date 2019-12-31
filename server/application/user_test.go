package application

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"private-conda-repo/store/models"
)

func TestCreateUser(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	ts := newTestServer(CreateUser)
	defer ts.Close()

	payload := strings.NewReader(`
	{
		"channel": "daniel",
		"password": "Password123",
		"email": "daniel@gmail.com"
	}`)

	resp, err := http.Post(ts.URL, ApplicationJson, payload)
	assert.NoError(err)
	assert.EqualValues(resp.StatusCode, 200)

	var u *models.User
	err = json.NewDecoder(resp.Body).Decode(&u)
	assert.NoError(err)
	defer func() { _ = resp.Body.Close() }()

	assert.EqualValues(u.Channel, "daniel")
	assert.Empty(u.Password)
}

func TestGetUserInfo(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	channel := "test-get-user-info"
	email := "daniel@gmail.com"

	_, err := db.AddUser(channel, "password", email)
	assert.NoError(err)

	ts := newTestServerWithRouteContext("GET", "/{user}", GetUserInfo)
	defer ts.Close()

	resp, err := http.Get(fmt.Sprintf("%s/%s", ts.URL, channel))
	assert.NoError(err)
	assert.EqualValues(resp.StatusCode, http.StatusOK)

	var output models.User
	err = json.NewDecoder(resp.Body).Decode(&output)
	assert.NoError(err)
	defer func() { _ = resp.Body.Close() }()

	assert.EqualValues(output.Password, "")
	assert.EqualValues(output.Channel, channel)
	assert.EqualValues(output.Email, email)
}

func TestListUsers(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	ts := newTestServer(ListUsers)
	defer ts.Close()

	_, err := db.AddUser("Pikachu", "pika-pi!!", "daniel@gmail.com")
	assert.NoError(err)

	resp, err := http.Get(ts.URL)
	assert.NoError(err)
	assert.EqualValues(resp.StatusCode, 200)

	var users []*models.User
	err = json.NewDecoder(resp.Body).Decode(&users)
	assert.NoError(err)
	defer func() { _ = resp.Body.Close() }()

	allUsers, err := db.GetAllUsers()
	assert.NoError(err)
	assert.Len(users, len(allUsers))
}

func TestRemoveUser(t *testing.T) {
	type TestRow struct {
		Payload    string
		StatusCode int
	}

	assert := require.New(t)

	ts := newTestServer(RemoveUser)
	defer ts.Close()

	channel := "daniel-remove-user"
	password := "Password123"
	err := createChannelAndAddPackages(channel)
	assert.NoError(err)

	_, err = db.AddUser(channel, password, "daniel@gmail.com")
	assert.NoError(err)

	tests := []TestRow{
		{
			Payload: fmt.Sprintf(`{
			"channel": "%s",
			"password": "%s"
			}`, channel, password),
			StatusCode: http.StatusOK,
		},
		{
			Payload: fmt.Sprintf(`{
			"channel": "BadChannel",
			"password": "%s"
			}`, password),
			StatusCode: http.StatusBadRequest,
		},
		{
			Payload: fmt.Sprintf(`{
			"channel": "%s",
			"password": "BadPassword"
			}`, channel),
			StatusCode: http.StatusBadRequest,
		},
	}

	runTest := func(test TestRow) {
		client := &http.Client{}
		req, err := http.NewRequest("DELETE", ts.URL, strings.NewReader(test.Payload))
		assert.NoError(err)

		resp, err := client.Do(req)
		assert.NoError(err)
		assert.EqualValues(test.StatusCode, resp.StatusCode)
	}

	for _, test := range tests {
		runTest(test)
	}
}
