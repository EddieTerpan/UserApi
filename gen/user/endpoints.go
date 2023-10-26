// Code generated by goa v3.13.2, DO NOT EDIT.
//
// user endpoints
//
// Command:
// $ goa gen UserApi

package user

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "user" service endpoints.
type Endpoints struct {
	Create goa.Endpoint
	Read   goa.Endpoint
	Update goa.Endpoint
	Delete goa.Endpoint
	Token  goa.Endpoint
}

// NewEndpoints wraps the methods of the "user" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		Create: NewCreateEndpoint(s),
		Read:   NewReadEndpoint(s, a.JWTAuth),
		Update: NewUpdateEndpoint(s, a.JWTAuth),
		Delete: NewDeleteEndpoint(s, a.JWTAuth),
		Token:  NewTokenEndpoint(s),
	}
}

// Use applies the given middleware to all the "user" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Create = m(e.Create)
	e.Read = m(e.Read)
	e.Update = m(e.Update)
	e.Delete = m(e.Delete)
	e.Token = m(e.Token)
}

// NewCreateEndpoint returns an endpoint function that calls the method
// "create" of service "user".
func NewCreateEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*CreatePayload)
		return s.Create(ctx, p)
	}
}

// NewReadEndpoint returns an endpoint function that calls the method "read" of
// service "user".
func NewReadEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*ReadPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"api:access"},
			RequiredScopes: []string{"api:access"},
		}
		var token string
		if p.JWT != nil {
			token = *p.JWT
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		return s.Read(ctx, p)
	}
}

// NewUpdateEndpoint returns an endpoint function that calls the method
// "update" of service "user".
func NewUpdateEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*UpdatePayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"api:access"},
			RequiredScopes: []string{"api:access"},
		}
		var token string
		if p.JWT != nil {
			token = *p.JWT
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		return s.Update(ctx, p)
	}
}

// NewDeleteEndpoint returns an endpoint function that calls the method
// "delete" of service "user".
func NewDeleteEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*DeletePayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"api:access"},
			RequiredScopes: []string{"api:access"},
		}
		var token string
		if p.JWT != nil {
			token = *p.JWT
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		return s.Delete(ctx, p)
	}
}

// NewTokenEndpoint returns an endpoint function that calls the method "token"
// of service "user".
func NewTokenEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*TokenPayload)
		return s.Token(ctx, p)
	}
}
