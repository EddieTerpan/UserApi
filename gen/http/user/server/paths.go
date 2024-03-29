// Code generated by goa v3.13.2, DO NOT EDIT.
//
// HTTP request path constructors for the user service.
//
// Command:
// $ goa gen UserApi

package server

// CreateUserPath returns the URL path to the user service create HTTP endpoint.
func CreateUserPath() string {
	return "/user/create"
}

// ReadUserPath returns the URL path to the user service read HTTP endpoint.
func ReadUserPath() string {
	return "/user"
}

// UpdateUserPath returns the URL path to the user service update HTTP endpoint.
func UpdateUserPath() string {
	return "/user/update"
}

// DeleteUserPath returns the URL path to the user service delete HTTP endpoint.
func DeleteUserPath() string {
	return "/user/delete"
}

// TokenUserPath returns the URL path to the user service token HTTP endpoint.
func TokenUserPath() string {
	return "/auth/token"
}
