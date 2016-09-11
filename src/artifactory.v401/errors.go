package artifactory

import (
	"fmt"
)

type ErrorsJson struct {
	Errors []ErrorJson `json:"errors"`
}

type ErrorJson struct {
	Status  int `json:"status"`
	Message string `json:"message"`
}

type ArtifactoryHTTPError struct {
	HttpStatus int
	Body []byte
	Errors ErrorsJson
}

func (e *ArtifactoryHTTPError) Error() string {
	return fmt.Sprintf("Artifactory HTTP Error %d: %s", e.HttpStatus, string(e.Body))
}
