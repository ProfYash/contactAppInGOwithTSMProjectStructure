package web

import (
	"contactApp/errors"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// UnmarshalJSON parses data from request and return otherwise error return.
func UnmarshalJSON(request *http.Request, out interface{}) error {
	if request.Body == nil {
		return errors.NewHTTPError(errors.ErrorCodeEmptyRequestBody, http.StatusBadRequest)
	}

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return errors.NewHTTPError(err.Error(), http.StatusBadRequest)
	}

	if len(body) == 0 {
		return errors.NewHTTPError(errors.ErrorCodeEmptyRequestBody, http.StatusBadRequest)
	}

	err = json.Unmarshal(body, out)
	if err != nil {
		return errors.NewHTTPError(err.Error(), http.StatusBadRequest)
	}
	return nil
}