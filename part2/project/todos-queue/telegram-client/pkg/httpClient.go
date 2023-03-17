package pkg 

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func HttpGet(url string) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	if !(res.StatusCode == http.StatusOK || res.StatusCode == http.StatusAccepted) {
		resBody, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return errors.New(fmt.Sprintf("Unsuccessful %d http response. Cannot read response body.\n", res.StatusCode))
		}

		return errors.New(fmt.Sprintf("Unsuccessful %d http response.\n%s", res.StatusCode, resBody))
	}

	return nil
}
