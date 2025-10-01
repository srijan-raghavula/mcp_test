package main

import (
	"fmt"
	"net/http"
)

func (app *application) newGet(endpoint string, headers map[string]string) (*http.Request, error) {
	url := fmt.Sprintf("%s%s", app.baseURL, endpoint)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", app.token))
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	return req, nil
}
