package azurermext

import (
	"github.com/Azure/go-autorest/autorest"
	"net/http"
)

func resourceNotFound(resp autorest.Response) bool {
	return responseWasStatusCode(resp, http.StatusNotFound)
}

func responseWasStatusCode(resp autorest.Response, statusCode int) bool {
	if r := resp.Response; r != nil {
		if r.StatusCode == statusCode {
			return true
		}
	}

	return false
}