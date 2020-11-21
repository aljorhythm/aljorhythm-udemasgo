package utility

import (
	"io/ioutil"
	"net/http"
)

// Theysaidso what they said so in bytes
func Theysaidso() (*[]byte, error) {
	resp, err := http.Get("http://api.theysaidso.com/qod.json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return &contents, nil
}
