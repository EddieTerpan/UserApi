// Code generated by goa v3.13.2, DO NOT EDIT.
//
// user HTTP server types
//
// Command:
// $ goa gen UserApi

package server

import (
	user "UserApi/gen/user"

	goa "goa.design/goa/v3/pkg"
)

// CreateRequestBody is the type of the "user" service "create" endpoint HTTP
// request body.
type CreateRequestBody struct {
	// User's email address
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// User's password
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
}

// UpdateRequestBody is the type of the "user" service "update" endpoint HTTP
// request body.
type UpdateRequestBody struct {
	// User name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// User surname
	Surname *string `form:"surname,omitempty" json:"surname,omitempty" xml:"surname,omitempty"`
	// User address
	Address *string `form:"address,omitempty" json:"address,omitempty" xml:"address,omitempty"`
	// phone
	Phone *string `form:"phone,omitempty" json:"phone,omitempty" xml:"phone,omitempty"`
}

// DeleteRequestBody is the type of the "user" service "delete" endpoint HTTP
// request body.
type DeleteRequestBody struct {
	// Delete id
	ID *uint `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
}

// TokenRequestBody is the type of the "user" service "token" endpoint HTTP
// request body.
type TokenRequestBody struct {
	// User's email address
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// User's password
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
}

// NewCreatePayload builds a user service create endpoint payload.
func NewCreatePayload(body *CreateRequestBody) *user.CreatePayload {
	v := &user.CreatePayload{
		Email:    body.Email,
		Password: body.Password,
	}

	return v
}

// NewReadPayload builds a user service read endpoint payload.
func NewReadPayload(jwt *string) *user.ReadPayload {
	v := &user.ReadPayload{}
	v.JWT = jwt

	return v
}

// NewUpdatePayload builds a user service update endpoint payload.
func NewUpdatePayload(body *UpdateRequestBody, jwt *string) *user.UpdatePayload {
	v := &user.UpdatePayload{
		Name:    body.Name,
		Surname: body.Surname,
		Address: body.Address,
		Phone:   body.Phone,
	}
	v.JWT = jwt

	return v
}

// NewDeletePayload builds a user service delete endpoint payload.
func NewDeletePayload(body *DeleteRequestBody, jwt *string) *user.DeletePayload {
	v := &user.DeletePayload{
		ID: body.ID,
	}
	v.JWT = jwt

	return v
}

// NewTokenPayload builds a user service token endpoint payload.
func NewTokenPayload(body *TokenRequestBody) *user.TokenPayload {
	v := &user.TokenPayload{
		Email:    body.Email,
		Password: body.Password,
	}

	return v
}

// ValidateCreateRequestBody runs the validations defined on CreateRequestBody
func ValidateCreateRequestBody(body *CreateRequestBody) (err error) {
	if body.Email != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.email", *body.Email, "^[a-z0-9._%+\\-]+@[a-z0-9.\\-]+\\.[a-z]{2,4}$"))
	}
	return
}

// ValidateTokenRequestBody runs the validations defined on TokenRequestBody
func ValidateTokenRequestBody(body *TokenRequestBody) (err error) {
	if body.Email != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.email", *body.Email, "^[a-z0-9._%+\\-]+@[a-z0-9.\\-]+\\.[a-z]{2,4}$"))
	}
	return
}
