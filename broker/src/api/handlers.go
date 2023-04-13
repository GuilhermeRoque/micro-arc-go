package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func Echo(w http.ResponseWriter, r *http.Request) {
	data, err := ReadJSON(w, r)
	if err != nil {
		log.Println("Error first read json")
		ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
	messageResponse := fmt.Sprintf("Request received %+v\n", data)
	log.Println(messageResponse)
	payload := JsonResponse{
		Error:   false,
		Message: messageResponse,
		Data:    data,
	}
	WriteJSON(w, http.StatusOK, payload)
}

func ProxyAuthService(w http.ResponseWriter, r *http.Request) {

	// call the service
	authServiceURL := os.Getenv("AUTH_SERVICE_URL")
	loginPath := fmt.Sprintf("%s/login", authServiceURL)
	log.Printf("Proxing to %s", loginPath)

	// create some json we'll send to the auth microservice
	data, err := ReadJSON(w, r)
	if err != nil {
		log.Println("Error Reading Json from Request")
		ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	request, err := NewPostJsonRequest(data, loginPath)
	if err != nil {
		log.Println("Error creating new request with same data")
		ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println("Error client do request")
		ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	dataService, err := GetJsonPayload(response.Body)
	if err != nil {
		log.Println("Error GetJsonPayload from response")
		ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	defer response.Body.Close()
	WriteJSON(w, response.StatusCode, dataService)

}

func ProxyRequestHost(host string) func(http.ResponseWriter, *http.Request) {

	return func(responseWriter http.ResponseWriter, request *http.Request) {
		// call the service
		path := fmt.Sprintf("%s%s", host, request.URL.Path)
		log.Printf("Proxing to %s", path)

		request, err := CopyRequestNewPath(request, responseWriter, path)
		if err != nil {
			ErrorJSON(responseWriter, err, http.StatusInternalServerError)
			return
		}

		client := &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			log.Println("Error client do request")
			ErrorJSON(responseWriter, err, http.StatusInternalServerError)
			return
		}
		SetStatusCode(responseWriter, response.StatusCode)
		err = CopyPayload(response, responseWriter)
		if err != nil {
			ErrorJSON(responseWriter, err, http.StatusInternalServerError)
			return
		}

		CopyHeaders(response.Header, responseWriter.Header())
	}
}
