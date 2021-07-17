package http

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type randomInnerObject struct {
	Status string
	Value  int64
}

type randomObject struct {
	randomInnerObject
	Date time.Time
}

func TestWriterResponse(t *testing.T) {
	t.Run("WHEN error", func(t *testing.T) {
		w := httptest.NewRecorder()
		msg := "Something went wrong"
		WriteResponse(w, msg, 500, nil, errors.New(msg))

		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)
		assert.Equal(t, 500, resp.StatusCode)
		assert.Equal(t, "application/json", resp.Header.Get("content-type"))
		assert.Contains(t, string(body), `"message":"Something went wrong"`)
	})

	t.Run("WHEN success", func(t *testing.T) {
		w := httptest.NewRecorder()
		msg := "default.success"
		currTime := time.Now()
		data := randomObject{
			randomInnerObject{
				Status: "abcd",
				Value:  1234,
			},
			currTime,
		}
		WriteResponse(w, msg, 200, data, nil)

		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)
		assert.Equal(t, 200, resp.StatusCode)
		assert.Equal(t, "application/json", resp.Header.Get("content-type"))

		dataWithBaseResponse := BaseResponse{
			200,
			msg,
			nil,
			data,
			currTime.Unix(),
		}
		expectedVal, _ := json.Marshal(dataWithBaseResponse)
		assert.Equal(t, string(expectedVal), string(body))
	})
}
