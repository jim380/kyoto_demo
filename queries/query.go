package queries

import (
	"io/ioutil"
	"net/http"
)

var Addr string = "https://rpc.cosmos.network"

func RESTQuery(route string) ([]byte, error) {
	req, err := http.NewRequest("GET", Addr+route, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return body, err
}
