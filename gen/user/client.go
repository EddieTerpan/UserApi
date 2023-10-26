// Code generated by goa v3.13.2, DO NOT EDIT.
//
// user client
//
// Command:
// $ goa gen UserApi

package user

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "user" service client.
type Client struct {
	CreateEndpoint goa.Endpoint
	ReadEndpoint   goa.Endpoint
	UpdateEndpoint goa.Endpoint
	DeleteEndpoint goa.Endpoint
	TokenEndpoint  goa.Endpoint
}

// NewClient initializes a "user" service client given the endpoints.
func NewClient(create, read, update, delete_, token goa.Endpoint) *Client {
	return &Client{
		CreateEndpoint: create,
		ReadEndpoint:   read,
		UpdateEndpoint: update,
		DeleteEndpoint: delete_,
		TokenEndpoint:  token,
	}
}

// Create calls the "create" endpoint of the "user" service.
func (c *Client) Create(ctx context.Context, p *CreatePayload) (res string, err error) {
	var ires any
	ires, err = c.CreateEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(string), nil
}

// Read calls the "read" endpoint of the "user" service.
func (c *Client) Read(ctx context.Context, p *ReadPayload) (res string, err error) {
	var ires any
	ires, err = c.ReadEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(string), nil
}

// Update calls the "update" endpoint of the "user" service.
func (c *Client) Update(ctx context.Context, p *UpdatePayload) (res string, err error) {
	var ires any
	ires, err = c.UpdateEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(string), nil
}

// Delete calls the "delete" endpoint of the "user" service.
func (c *Client) Delete(ctx context.Context, p *DeletePayload) (res string, err error) {
	var ires any
	ires, err = c.DeleteEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(string), nil
}

// Token calls the "token" endpoint of the "user" service.
func (c *Client) Token(ctx context.Context, p *TokenPayload) (res string, err error) {
	var ires any
	ires, err = c.TokenEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(string), nil
}