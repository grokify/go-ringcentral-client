package lite

import (
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

func GetUsers(serverURL, authOrApiToken string) ([]byte, error) {
	bytes := []byte("")
	headers, apiURL, err := APIInfo(
		serverURL, APIUsersURLPath, authOrApiToken)
	if err != nil {
		return bytes, err
	}

	req, err := http.NewRequest(http.MethodGet, apiURL, nil)
	if err != nil {
		return bytes, err
	}

	req.Header = headers
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return bytes, err
	} else if resp.StatusCode >= 300 {
		return bytes, errors.New("E_RESPONSE_STATUS_CODE_NOT_2XX")
	}
	return ioutil.ReadAll(resp.Body)
}
