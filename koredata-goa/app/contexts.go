// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "koredata": Application Contexts
//
// Command:
// $ goagen
// --design=github.com/hegemone/kore-poc/koredata-goa/design
// --out=$(GOPATH)/src/github.com/hegemone/kore-poc/koredata-goa
// --version=v1.3.1

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
	"unicode/utf8"
)

// CreateQuoteContext provides the quote create action context.
type CreateQuoteContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *CreateQuotePayload
}

// NewCreateQuoteContext parses the incoming request URL and body, performs validations and creates the
// context used by the quote controller create action.
func NewCreateQuoteContext(ctx context.Context, r *http.Request, service *goa.Service) (*CreateQuoteContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := CreateQuoteContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// createQuotePayload is the quote create action payload.
type createQuotePayload struct {
	Name  *string `form:"Name,omitempty" json:"Name,omitempty" xml:"Name,omitempty"`
	Quote *string `form:"Quote,omitempty" json:"Quote,omitempty" xml:"Quote,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *createQuotePayload) Validate() (err error) {
	if payload.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "Name"))
	}
	if payload.Quote == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "Quote"))
	}
	if payload.Name != nil {
		if utf8.RuneCountInString(*payload.Name) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.Name`, *payload.Name, utf8.RuneCountInString(*payload.Name), 2, true))
		}
	}
	if payload.Quote != nil {
		if utf8.RuneCountInString(*payload.Quote) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.Quote`, *payload.Quote, utf8.RuneCountInString(*payload.Quote), 2, true))
		}
	}
	return
}

// Publicize creates CreateQuotePayload from createQuotePayload
func (payload *createQuotePayload) Publicize() *CreateQuotePayload {
	var pub CreateQuotePayload
	if payload.Name != nil {
		pub.Name = *payload.Name
	}
	if payload.Quote != nil {
		pub.Quote = *payload.Quote
	}
	return &pub
}

// CreateQuotePayload is the quote create action payload.
type CreateQuotePayload struct {
	Name  string `form:"Name" json:"Name" xml:"Name"`
	Quote string `form:"Quote" json:"Quote" xml:"Quote"`
}

// Validate runs the validation rules defined in the design.
func (payload *CreateQuotePayload) Validate() (err error) {
	if payload.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "Name"))
	}
	if payload.Quote == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "Quote"))
	}
	if utf8.RuneCountInString(payload.Name) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.Name`, payload.Name, utf8.RuneCountInString(payload.Name), 2, true))
	}
	if utf8.RuneCountInString(payload.Quote) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.Quote`, payload.Quote, utf8.RuneCountInString(payload.Quote), 2, true))
	}
	return
}

// OK sends a HTTP response with status code 200.
func (ctx *CreateQuoteContext) OK(r *JSON) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *CreateQuoteContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// ListQuoteContext provides the quote list action context.
type ListQuoteContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewListQuoteContext parses the incoming request URL and body, performs validations and creates the
// context used by the quote controller list action.
func NewListQuoteContext(ctx context.Context, r *http.Request, service *goa.Service) (*ListQuoteContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ListQuoteContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListQuoteContext) OK(r *JSON) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *ListQuoteContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// ListByIDQuoteContext provides the quote list by ID action context.
type ListByIDQuoteContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	UserID string
}

// NewListByIDQuoteContext parses the incoming request URL and body, performs validations and creates the
// context used by the quote controller list by ID action.
func NewListByIDQuoteContext(ctx context.Context, r *http.Request, service *goa.Service) (*ListByIDQuoteContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ListByIDQuoteContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramUserID := req.Params["userId"]
	if len(paramUserID) > 0 {
		rawUserID := paramUserID[0]
		rctx.UserID = rawUserID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListByIDQuoteContext) OK(r *JSON) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ListByIDQuoteContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// LoginQuoteContext provides the quote login action context.
type LoginQuoteContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewLoginQuoteContext parses the incoming request URL and body, performs validations and creates the
// context used by the quote controller login action.
func NewLoginQuoteContext(ctx context.Context, r *http.Request, service *goa.Service) (*LoginQuoteContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := LoginQuoteContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *LoginQuoteContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *LoginQuoteContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}