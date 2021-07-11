package http

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type BaseResponse struct {
	Code       int         `json:"code"`
	Message    string      `json:"message,omitempty"`
	Errors     []string    `json:"errors,omitempty"`
	Data       interface{} `json:"data"`
	ServerTime int64       `json:"serverTime"`
}

func sendResponse(w http.ResponseWriter, body BaseResponse, err error) {
	w.Header().Set("content-type", "application/json")
	responseWriterWrap := GetResponseWriterWrap(w)
	responseWriterWrap.SetStatusCode(body.Code)
	responseWriterWrap.WriteError(err)

	jsonData, _ := json.Marshal(body)
	_, errWrite := responseWriterWrap.Write(jsonData)
	if errWrite != nil {
		logrus.Errorf("Error when writing response: %s", errWrite.Error())
	}
}

func WriteResponse(w http.ResponseWriter, msg string, code int, data interface{}, err error) {
	body := BaseResponse{
		Code:       code,
		Message:    msg,
		Data:       data,
		ServerTime: time.Now().Unix(),
	}

	if err != nil {
		body.Errors = []string{err.Error()}
	}
	sendResponse(w, body, err)
}
