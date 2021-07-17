package http

import (
	"errors"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var data = "some data"

func TestResponseWriterWrap_Write(t *testing.T) {
	w := httptest.NewRecorder()
	wrap := GetResponseWriterWrap(w)

	_, err := wrap.Write([]byte(data))
	assert.NoError(t, err)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, data, string(body))
	assert.Equal(t, data, string(wrap.body))
}

func TestResponseWriterWrap_SetStatusCode(t *testing.T) {
	w := httptest.NewRecorder()
	wrap := GetResponseWriterWrap(w)

	wrap.SetStatusCode(400)
	_, err := wrap.Write([]byte("some data"))
	assert.NoError(t, err)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, 400, resp.StatusCode)
	assert.Equal(t, data, string(body))
	assert.Equal(t, data, string(wrap.body))
}

func TestResponseWriterWrap_WriteError(t *testing.T) {
	w := httptest.NewRecorder()
	wrap := GetResponseWriterWrap(w)

	wrap.WriteError(errors.New("some error"))
	assert.Equal(t, "some error", wrap.error.Error())
}

func TestResponseWriterWrap_Getter(t *testing.T) {
	w := httptest.NewRecorder()
	wrap := ResponseWriterWrap{
		w,
		405,
		errors.New("error"),
		[]byte(data),
	}

	assert.Equal(t, 405, wrap.Status())
	assert.Equal(t, "error", wrap.Error().Error())
	assert.Equal(t, data, string(wrap.Body()))
}
