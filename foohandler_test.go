package foohandler

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

//we need to prefix it with test so, go will automatically run on : go test -v ./...
//-v => verbose : will put out Println() calls

/*
To test, we can use 2 things that Go provides in the standard testing library:
- the httpTest.server
- the httpResponseRecorder
*/

func TestHandleGetFoo(t *testing.T){
	server := httptest.NewServer(http.HandlerFunc(handleGetFoo));

	response, err := http.Get(server.URL);
	if err != nil {
		t.Error(err);
	}

	defer response.Body.Close();

	//we can test: headers, api tokens, refresh tokens; everything in the response;

	if response.StatusCode != http.StatusOK {
		t.Errorf("expected 200 but got %d", response.StatusCode)
	}

	expected := "FOO";

	bytes, err := io.ReadAll(response.Body);

	if err != nil {
		t.Error(err);
	}

	if string(bytes) != expected {
		t.Errorf("expected %s but we got %s", expected, string(bytes));
	}

}

//Response recorder will capture everything written with the responseWriter
//we will actually know who has set, which header, ourselves or a third-party api, or an auth flow api, or a reverse proxy, etc...

func TestHandleGetFooResponseRecorder(t *testing.T){
	responseRecorder := httptest.NewRecorder();

	//Method, urlSearchParams, Body => no body -> nil
	request, err := http.NewRequest(http.MethodGet, "", nil);

	if err != nil {
		t.Error(err);
	}

	handleGetFoo(responseRecorder, request);

	if responseRecorder.Result().StatusCode != http.StatusOK {
		t.Errorf("expected 200 but got %d", responseRecorder.Result().StatusCode)
	}

	expected := "FOO";

	bytes, err := io.ReadAll(responseRecorder.Result().Body);

	if err != nil {
		t.Error(err);
	}

	if string(bytes) != expected {
		t.Errorf("expected %s but we got %s", expected, string(bytes));
	}

}