package net

import (
	"io/ioutil"
	"net/http"
)

// RequestURL makes HTTP GET request to the URL
func RequestURL(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}

		return string(data), nil
	}

	return "", nil
}
