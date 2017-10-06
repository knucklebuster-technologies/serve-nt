package handlers

import (
	"encoding/json"
	"net/http"
)

func restGet(url string) []byte {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil
	}
	c := &http.Client{}
	res, err := c.Do(req)
	if err != nil {
		return nil
	}

	defer res.Body.Close()
	b, err := json.Marshal(res.Body)
	if err != nil {
		return nil
	}
	return b
}
