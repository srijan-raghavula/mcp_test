package main

import (
	"fmt"
	"io"
	"net/http"
)

func (app *application) reqAuthentication() error {
	req, err := app.newGet("/authentication", map[string]string{"accept": "application/json"})
	if err != nil {
		return err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(body))

	return nil
}
