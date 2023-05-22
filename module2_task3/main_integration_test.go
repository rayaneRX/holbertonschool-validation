package main

import (
  "io/ioutil"
  "net/http"
  "net/http/httptest"
  "testing"
)

func Test_server(t *testing.T) {
  if testing.Short() {
    t.Skip("Flag `-short` provided: skipping Integration Tests.")
  }

  tests := []struct {
    name         string
    URI          string
    responseCode int
    body         string
  }{
    {
      name:         "Home page",
      URI:          "",
      responseCode: 404,
      body:         "404 page not found\n",
    },
    {
      name:         "Health page",
      URI:          "/hello?name=Holberton",
      responseCode: 200,
      body:         "Hello Holberton!",
    },
  }
    reqNoName, _ := http.NewRequest("GET", "/hello", nil)
    reqChris, _ := http.NewRequest("GET", "/hello?name=Chris", nil)
    reqEmptyName, _ := http.NewRequest("GET", "/hello?name=", nil)

    response := httptest.NewRecorder()
    HelloHandler(response, reqNoName)
    if response.Body.String() != "Hello there!" {
        t.Errorf("unexpected body: got %v want %v", response.Body.String(), "Hello there!")
    }

    response = httptest.NewRecorder()
    HelloHandler(response, reqChris)
    if response.Body.String() != "Hello Chris!" {
        t.Errorf("unexpected body: got %v want %v", response.Body.String(), "Hello Chris!")
    }

    response = httptest.NewRecorder()
    HelloHandler(response, reqEmptyName)
    if status := response.Code; status != http.StatusBadRequest {
        t.Errorf("wrong status code: got %v want %v", status, 400)
    }



  ts := httptest.NewServer(http.HandlerFunc(HealthCheckHandler))
    defer ts.Close()

    res, err := http.Get(ts.URL)

    body, err := ioutil.ReadAll(res.Body)
    res.Body.Close()
    if err != nil {
        t.Fatal(err)
    }

    if string(body) != "ALIVE" {
        t.Errorf("expected body: 'ALIVE'; response body: %q", string(body))
    }

  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      ts := httptest.NewServer(setupRouter())
      defer ts.Close()

      res, err := http.Get(ts.URL + tt.URI)
      if err != nil {
        t.Fatal(err)
      }

      // Check that the status code is what you expect.
      expectedCode := tt.responseCode
      gotCode := res.StatusCode
      if gotCode != expectedCode {
        t.Errorf("handler returned wrong status code: got %q want %q", gotCode, expectedCode)
      }

      // Check that the response body is what you expect.
      expectedBody := tt.body
      bodyBytes, err := ioutil.ReadAll(res.Body)
      res.Body.Close()
      if err != nil {
        t.Fatal(err)
      }
      gotBody := string(bodyBytes)
      if gotBody != expectedBody {
        t.Errorf("handler returned unexpected body: got %q want %q", gotBody, expectedBody)
      }
    })
  }
}