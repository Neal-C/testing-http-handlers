package foohandler

import "net/http"

func handleGetFoo(responseWriter http.ResponseWriter, request *http.Request){
	if request.Method != http.MethodGet{
		responseWriter.WriteHeader(http.StatusMethodNotAllowed);
		return;
	}

	responseWriter.WriteHeader(http.StatusOK);

	responseWriter.Write([]byte("FOO"));
}