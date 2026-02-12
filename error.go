package arcgis

import (
	"context"
	"fmt"

	"github.com/rs/zerolog"
)

type apiErrorType string

const (
	APIErrorUnrecognized     apiErrorType = ""
	APIErrorInvalidAuthToken apiErrorType = "invalid_auth_token"
	APIErrorInvalidRequest   apiErrorType = "invalid_request"
	APIErrorNotPermitted     apiErrorType = "not_permitted"
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
func (e ErrorFromAPI) HasDetail(s string) bool {
	return hasString(e.Details, s)
}

var (
	ErrorInvalidAuthToken *apiError = &apiError{
		Code:        400,
		Description: "invalid auth token",
		Details:     []string{},
		ErrorType:   APIErrorInvalidAuthToken,
		Message:     "invalid auth token",
	}
	ErrorInvalidRefreshToken *apiError = &apiError{
		Code:        498,
		Description: "invalidated refresh_token",
		Details:     []string{},
		ErrorType:   APIErrorInvalidRequest,
		Message:     "invalidated refresh_token",
	}
	ErrorNotPermitted *apiError = &apiError{
		Code:        403,
		Description: "User does not have permissions to access this service",
		Details:     []string{},
		ErrorType:   APIErrorNotPermitted,
		Message:     "not permitted",
	}
)

func errorTypeFromString(ctx context.Context, s string) apiErrorType {
	logger := zerolog.Ctx(ctx)
	switch s {
	case "invalid_request":
		return APIErrorInvalidRequest
	default:
		logger.Warn().Str("s", s).Msg("Did not recognize API error type")
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
func newAPIError(ctx context.Context, e ErrorResponse) apiError {
	logger := zerolog.Ctx(ctx)
	logger.Debug().
		Int("code", e.Error.Code).
		Strs("details", e.Error.Details).
		Str("error", e.Error.Error).
		Str("description", e.Error.Description).
		Str("message", e.Error.Message).
		Msg("got API error")
	if e.Error.Code == 400 && e.Error.Message == "Invalid token" {
		logger.Debug().Msg("Recognized error as 'invalid token'")
		return *ErrorInvalidAuthToken
	} else if e.Error.Error == "" && e.Error.Code == 403 /*&& hasString(e.Error.Details, "User does not have permissions to access this service ()")*/ {
		logger.Debug().Msg("Recognized error as 'not premitted'")
		return *ErrorNotPermitted
	} else if e.Error.Code == 498 && e.Error.HasDetail("Invalid token.") {
		// error= code=498 description= details=["Invalid token."]
		logger.Debug().Msg("Recognized error as 'invalid token'")
		return *ErrorInvalidAuthToken
	} else if e.Error.Code == 498 && e.Error.Error == "invalid_request" {
		// 9:26PM DBG got API error error= code=498 description= details=["Invalid token."]
		logger.Debug().Msg("Recognized error as 'invalidated refresh token'")
		return *ErrorInvalidRefreshToken
	}

	return apiError{
		Code:        e.Error.Code,
		Description: e.Error.Description,
		Details:     e.Error.Details,
		ErrorType:   errorTypeFromString(ctx, e.Error.Error),
		Message:     e.Error.Message,
	}
}
