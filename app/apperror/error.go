package apperror

import "errors"

// Validation Error
// Code start with 1xxxxx
var (
	ErrorInvalidUsername = errors.New("invalid Username, please try again")

	ErrorInvalidPassword = errors.New("invalid Password, please try again")

	ErrorInvalidFullname = errors.New("invalid Fullname, please try again")

	ErrorEmptyField = errors.New("empty field")

	ErrorSpaceDetected = errors.New("there should be no space on Username / Password field")

	ErrorWrongDateFormat = errors.New("wrong given format for date")

	ErrorPasswordIncorrect = errors.New("password is incorrect")
)

// Client errors.
// Code start with 2xxxx
var (
	// create a record but it has already existed
	ErrorEntryExists = errors.New("record already exists")

	// wrong input.
	ErrorInputInvalid = errors.New("input invalid")

	// cookie is not found.
	ErrorCookieNotFound = errors.New("cookie not found")

	// users response to a blocked IP address.
	ErrorBlockedIP = errors.New("blocked IP")

	// authentication failed for user
	ErrorNotAuthenticated = errors.New("authentication failed")

	// user is not authorized to access this resource
	ErrorNotAuthorized = errors.New("not authorized for this action ")

	// cookie outdated
	ErrorCookieOutdated = errors.New("cookie outdated")

	// cannot delete user, like admin
	ErrorCannotDeleteThisEntity = errors.New("cannot delete this user  because it will violate the system")

	// action cannot be performed because it is not a suitable user
	ErrorRelationViolation = errors.New("cannot add this relationship the current user, please check that you select the right users and/or objects of this action")

	ErrorUsernameAlreadyExist = errors.New("username already exists, please choose another username")

	ErrorNameAlreadyExist = errors.New("the name of the object has already existed. Please choose another name")
)

// Server errors.
var (
	// error related to server
	ErrorInternal = errors.New("error internal")

	// query some entry that doesnt exist as some relationship may be wrong
	ErrorEntryNotExist = errors.New("entry not exist")

	// Any confict that should not have been occurred.
	ErrorConflict = errors.New("conflict, report immediately to the admin with record of behaviour")

	// binding json failed.
	ErrorBindJSON = errors.New("binding input failed")

	// user fails to login
	ErrorContextLostValue = errors.New("authentication failed during processing, please login again")

	// request is sent but server does not receive it
	ErrorNoParamReceived = errors.New("no params have been received by the server")
)

// Sql error, CRUD operation
var (
	ErrorDeleteFailed = errors.New("deleting failed, please try again")

	ErrorUpdateFailed = errors.New("updaing failed, please try again")

	ErrorCreateFailed = errors.New("createing failed, please try again")

	ErrorQueryFailed = errors.New("querying failed, please try again")
)
