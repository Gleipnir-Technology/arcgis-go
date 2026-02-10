package arcgis

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

type apiErrorType string

const (
	APIErrorUnrecognized   apiErrorType = ""
	APIErrorInvalidRequest apiErrorType = "invalid_request"
	APIErrorNotPermitted apiErrorType = "not_permitted"
)

type apiError struct {
	Code        int
	Description string
	Details     []string
	ErrorType   apiErrorType
	Message     string
}

func (ae apiError) Error() string {
	return fmt.Sprintf("API error %d: %s (%s)", ae.Code, ae.Message, ae.Description)
}
func (ae apiError) Is(target error) bool {
	return ae.Error() == target.Error()
}

var (
	ErrorInvalidatedRefreshToken *apiError = &apiError{
		Code:        498,
		Description: "invalidated refresh_token",
		Details:     []string{},
		ErrorType:   APIErrorInvalidRequest,
		Message:     "invalidated refresh_token",
	}
	ErrorNotPermitted *apiError = &apiError{
		Code: 403,
		Description: "User does not have permissions to access this service",
		Details: []string{},
		ErrorType:   APIErrorNotPermitted,
		Message: "not permitted",
	}

)

func errorTypeFromString(s string) apiErrorType {
	switch s {
	case "invalid_request":
		return APIErrorInvalidRequest
	default:
		log.Warn().Str("s", s).Msg("Did not recognize API error type")
		return APIErrorUnrecognized
	}
}

func hasString(strs []string, to_find string) bool {
	for _, str := range strs {
		if str == to_find {
			return true
		}
	}
	return false
}
func newAPIError(e ErrorResponse) apiError {
	log.Debug().Int("code", e.Error.Code).Strs("details", e.Error.Details).Str("error", e.Error.Error).Str("description", e.Error.Description).Str("message", e.Error.Message).Msg("got API error")
	if /*e.Error.Error == "" &&*/ e.Error.Code == 403 /*&& hasString(e.Error.Details, "User does not have permissions to access this service ()")*/ {
		log.Debug().Msg("Recognized error as 'not premitted'")
		return *ErrorNotPermitted
	}
	return apiError{
		Code:        e.Error.Code,
		Description: e.Error.Description,
		Details:     e.Error.Details,
		ErrorType:   errorTypeFromString(e.Error.Error),
		Message:     e.Error.Message,
	}
}
