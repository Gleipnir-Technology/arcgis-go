package arcgis

import (
	"fmt"
)

type APIErrorType int

const (
	APIErrorUnrecognized APIErrorType = iota
	APIErrorInvalidRequest
)

type ArcGISAPIError struct {
	Code        int
	Description string
	Details     []string
	ErrorType   APIErrorType
	Message     string
}

func (ae ArcGISAPIError) Error() string {
	return fmt.Sprintf("API error %d: %s (%s)", ae.Code, ae.Message, ae.Description)
}

type InvalidatedRefreshTokenError struct {
	ArcGISAPIError
}

func errorTypeFromString(s string) APIErrorType {
	switch s {
	case "invalid_request":
		return APIErrorInvalidRequest
	default:
		return APIErrorUnrecognized
	}
}
