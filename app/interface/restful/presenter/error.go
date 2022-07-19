package presenter

import "github.com/thvinhtruong/legoha/app/apperror"

type ErrorResponse struct {
	ErrorCode  int    `json:"error_code"`
	ErrorMsg   string `json:"error_msg"`
	ErrorField string `json:"error_field"`
}

func getErrorField(err error) (field string) {
	switch err {
	case nil:
		return ""
	// The following are error field, not elaboration.
	case apperror.ErrorInvalidUsername:
		return "Invalid Username"
	case apperror.ErrorInvalidPassword:
		return "Invalid Password"
	case apperror.ErrorInvalidFullname:
		return "Invalid Fullname"
	case apperror.ErrorSpaceDetected:
		return "Space in Username and Password"
	case apperror.ErrorCookieNotFound:
		return "Cookie Not Found"
	case apperror.ErrorBlockedIP:
		return "IP is Blocked"
	case apperror.ErrorWrongDateFormat:
		return "Invalid Date Format"

	// The following are elaboration, not error field.
	case apperror.ErrorEntryExists:
		return "Entry has already existed"
	case apperror.ErrorCookieOutdated:
		return "Session expired. Please login again"
	case apperror.ErrorInputInvalid:
		return "Input invalid. Please try again"
	case apperror.ErrorNotAuthorized:
		return "You are not authorized to do this action"
	case apperror.ErrorNotAuthenticated:
		return "You are not authenticated, please login"
	case apperror.ErrorInternal:
		return "There is an internal error, please contact admin or try again later"
	case apperror.ErrorEntryNotExist:
		return "There is an internal error, please contact admin or try again later"
	case apperror.ErrorBindJSON:
		return "Connection is unstable, please try again later"
	case apperror.ErrorContextLostValue:
		return "The system authentication of your account has been deleted, please login again"
	case apperror.ErrorNameAlreadyExist:
		return "choose another name for your object (new todo name, new username, etc)"
	default:
		return "Unidentified Error"
	}
}
