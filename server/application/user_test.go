package application

import (
	"encoding/json"
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

	_, err := db.AddUser("daniel", "Password123")
	assert.NoError(err)

	payload := strings.NewReader(`
	{
		"name": "daniel",
		"password": "Password123"
	}`)

	resp, err := http.Post(ts.URL, ApplicationJson, payload)
	assert.NoError(err)
	assert.EqualValues(resp.StatusCode, 200)

	var u *models.User
	err = json.NewDecoder(resp.Body).Decode(&u)
	assert.NoError(err)
	defer func() { _ = resp.Body.Close() }()

	assert.EqualValues(u.Name, "daniel")
	assert.Empty(u.Password)
}

func TestListUsers(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	ts := newTestServer(ListUsers)
	defer ts.Close()

	_, err := db.AddUser("Pikachu", "pika-pi!!")
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
	t.Parallel()

	type TestRow struct {
		Payload    string
		StatusCode int
	}

	assert := require.New(t)

	ts := newTestServer(RemoveUser)
	defer ts.Close()
	_, err := db.AddUser("daniel", "Password123")
	assert.NoError(err)

	tests := []TestRow{
		{
			Payload: `{
			"name": "daniel",
			"password": "Password123"
			}`,
			StatusCode: http.StatusOK,
		},
		{
			Payload: `{
			"name": "daniel123",
			"password": "Password123"
			}`,
			StatusCode: http.StatusBadRequest,
		},
		{
			Payload: `{
			"name": "daniel",
			"password": "Password"
			}`,
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
