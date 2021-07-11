package http

import "net/http"

type ResponseWriterWrap struct {
	http.ResponseWriter
	statusCode int
	error      error
	body       []byte
}

func GetResponseWriterWrap(w http.ResponseWriter) *ResponseWriterWrap {
	if _w, ok := w.(*ResponseWriterWrap); ok {
		return _w
	}
	return &ResponseWriterWrap{w, http.StatusOK, nil, nil}
}

func (w *ResponseWriterWrap) Write(data []byte) (int, error) {
	return w.ResponseWriter.Write(data)
}

func (w *ResponseWriterWrap) SetStatusCode(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

func (w *ResponseWriterWrap) WriteError(err error) {
	w.error = err
}

func (w *ResponseWriterWrap) Status() int {
	return w.statusCode
}

func (w *ResponseWriterWrap) Error() error {
	return w.error
}

func (w *ResponseWriterWrap) Body() []byte {
	return w.body
}
