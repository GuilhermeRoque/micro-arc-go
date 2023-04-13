package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type JsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// readJSON tries to read the body of a request and converts it into JSON
func ReadJSON(w http.ResponseWriter, r *http.Request) (map[string]any, error) {
	maxBytes := 1048576 // one megabyte
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))
	return GetJsonPayload(r.Body)

}

func GetJsonPayload(body io.ReadCloser) (map[string]any, error) {
	data := map[string]any{}

	dec := json.NewDecoder(body)
	err := dec.Decode(&data)
	if err != nil {
		return nil, err
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return nil, errors.New("body must have only a single JSON value")
	}

	return data, nil
}

// writeJSON takes a response status code and arbitrary data and writes a json response to the client
func WriteJSON(r http.ResponseWriter, status int, data any) error {
	err := AddJsonPayloadResponse(data, r)
	SetStatusCode(r, status)
	return err
}

func CopyHeaders(sourceHeader http.Header, destHeader http.Header) {
	for key, value := range sourceHeader {
		destHeader[key] = value
	}
}

func SetStatusCode(r http.ResponseWriter, status int) {
	r.WriteHeader(status)
}

func CopyPayload(responseSource *http.Response, responseDest http.ResponseWriter) error {
	defer responseSource.Body.Close()
	dataService, err := ioutil.ReadAll(responseSource.Body)
	if err != nil {
		log.Println("Error reading response body from responseSource")
		return err
	}
	// log.Printf("%d bytes read", len(dataService))

	_, err = responseDest.Write(dataService)
	if err != nil {
		log.Println("Error writing data to responseDest")
		return err
	}
	// log.Printf("%d bytes written\n", lenBytes)
	return nil
}

func AddJsonPayloadResponse(data any, r http.ResponseWriter) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	r.Header().Set("Content-Type", "application/json")
	_, err = r.Write(out)
	if err != nil {
		return err
	}
	return nil
}

func NewPostJsonRequest(data any, path string) (*http.Request, error) {
	jsonData, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", path, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	return request, nil
}

// errorJSON takes an error, and optionally a response status code, and generates and sends
// a json error response
func ErrorJSON(w http.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload JsonResponse
	payload.Error = true
	payload.Message = err.Error()

	return WriteJSON(w, statusCode, payload)
}

func CopyRequestNewPath(request *http.Request, responseWriter http.ResponseWriter, path string) (*http.Request, error) {
	defer request.Body.Close()

	data, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Println("Error reading body from request")
		return nil, err
	}
	// log.Printf("%d bytes read", len(data))

	requestNew, err := http.NewRequest(request.Method, path, bytes.NewBuffer(data))
	if err != nil {
		log.Println("Error creating new request with same data")
		return nil, err
	}
	CopyHeaders(request.Header, requestNew.Header)
	return requestNew, nil
}
