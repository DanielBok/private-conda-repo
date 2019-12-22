package application

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"private-conda-repo/store/models"
)

func TestCreateUser(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	ts := newTestServer(CreateUser)
	defer ts.Close()

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
	assert := assert.New(t)

	ts := newTestServer(ListUsers)
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	assert.NoError(err)
	assert.EqualValues(resp.StatusCode, 200)

	var users []*models.User
	err = json.NewDecoder(resp.Body).Decode(&users)
	assert.NoError(err)
	defer func() { _ = resp.Body.Close() }()

	assert.Len(users, 2) // hard-coded from the mock interface
}

func TestRemoveUser(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	ts := newTestServer(RemoveUser)
	defer ts.Close()

	payload := strings.NewReader(`
	{
		"name": "daniel",
		"password": "Password123"
	}`)

	client := &http.Client{}
	req, err := http.NewRequest("DELETE", ts.URL, payload)
	assert.NoError(err)

	resp, err := client.Do(req)
	assert.NoError(err)

	assert.EqualValues(resp.StatusCode, 200)
}
