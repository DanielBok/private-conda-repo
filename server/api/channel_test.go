package api_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	. "private-conda-repo/api"
	"private-conda-repo/domain/entity"
)

const (
	email    = "daniel.bok@outlook.com"
	password = "TestPassword"
)

func NewChannelHandler() *ChannelHandler {
	return &ChannelHandler{
		DB:      NewMockDb(),
		FileSys: NewMockFileSys(),
	}
}

func TestChannelHandler_CreateChannel(t *testing.T) {
	handler := NewChannelHandler()
	assert := require.New(t)

	channelName := "create-channel"
	w := httptest.NewRecorder()

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(&entity.Channel{
		Channel:  channelName,
		Password: password,
		Email:    email,
	})
	assert.NoError(err)
	r := NewTestRequest("POST", "/", &buf, nil)

	handler.CreateChannel()(w, r)
	assert.Equal(http.StatusOK, w.Code)

	var result *entity.Channel
	err = json.NewDecoder(w.Body).Decode(&result)
	assert.NoError(err)

	assert.Equal(result.Channel, channelName)
	assert.Empty(result.Password)
	assert.Equal(result.Email, email)
}

func TestChannelHandler_GetChannelInfo(t *testing.T) {
	handler := NewChannelHandler()
	assert := require.New(t)

	channelName := "get-channel-info"
	_, err := handler.DB.CreateChannel(channelName, password, email)
	assert.NoError(err)

	for _, test := range []struct {
		Channel      string
		ExpectedCode int
	}{
		{channelName, http.StatusOK},
		{"bad-channel-name-that-does-not-exist", http.StatusBadRequest},
	} {
		w := httptest.NewRecorder()
		r := NewTestRequest("GET", fmt.Sprintf("/%s", test.Channel), nil, map[string]string{
			"channel": test.Channel,
		})
		handler.GetChannelInfo()(w, r)

		assert.Equal(test.ExpectedCode, w.Code)
		if test.ExpectedCode == http.StatusOK {
			var result *entity.Channel
			err = json.NewDecoder(w.Body).Decode(&result)
			assert.NoError(err)

			assert.Equal(result.Channel, channelName)
			assert.Empty(result.Password)
			assert.Equal(result.Email, email)
		}
	}

}

func TestChannelHandler_ListChannels(t *testing.T) {
	handler := NewChannelHandler()
	assert := require.New(t)
	n := 10

	for i := 0; i < n; i++ {
		_, err := handler.DB.CreateChannel(fmt.Sprintf("channel-%d", i), password, email)
		assert.NoError(err)
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)

	handler.ListChannels()(w, r)

	assert.Equal(http.StatusOK, w.Code)
	var result []*entity.Channel
	err := json.NewDecoder(w.Body).Decode(&result)
	assert.NoError(err)

	assert.Len(result, n)
	for _, r := range result {
		assert.NotNil(r)
		assert.Empty(r.Password)
	}
}

func TestChannelHandler_RemoveChannel(t *testing.T) {
	handler := NewChannelHandler()
	assert := require.New(t)

	channelName := "remove-channel"
	_, err := handler.DB.CreateChannel(channelName, password, email)
	assert.NoError(err)

	for _, test := range []struct {
		Channel      string
		Password     string
		ExpectedCode int
	}{
		{"does-not-exist", password, http.StatusBadRequest},
		{channelName, "wrong" + password, http.StatusBadRequest},
		{channelName, password, http.StatusBadRequest},
	} {
		w := httptest.NewRecorder()

		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(&entity.Channel{
			Channel:  test.Channel,
			Password: test.Password,
			Email:    email,
		})
		assert.NoError(err)
		r := NewTestRequest("DELETE", "/", &buf, nil)

		handler.RemoveChannel()(w, r)
		assert.Equal(test.ExpectedCode, w.Code)
	}

}
