package arcgis

import (
	"fmt"
)

type APIErrorType string

const (
	APIErrorUnrecognized   APIErrorType = ""
	APIErrorInvalidRequest APIErrorType = "invalid_request"
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
func (ae ArcGISAPIError) Is(target error) bool {
	return ae.Error() == target.Error()
}

var (
	InvalidatedRefreshTokenError *ArcGISAPIError = &ArcGISAPIError{
		Code:        498,
		Description: "invalidated refresh_token",
		Details:     []string{},
		ErrorType:   APIErrorInvalidRequest,
		Message:     "invalidated refresh_token",
	}
)

func errorTypeFromString(s string) APIErrorType {
	switch s {
	case "invalid_request":
		return APIErrorInvalidRequest
	default:
		return APIErrorUnrecognized
	}
}
