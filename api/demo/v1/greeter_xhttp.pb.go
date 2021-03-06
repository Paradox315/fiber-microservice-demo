// Code generated by protoc-gen-go-xhttp. DO NOT EDIT.
// versions:
// protoc-gen-go-xhttp v1.0.0

package v1

import (
	context "context"
	middleware "github.com/go-kratos/kratos/v2/middleware"
	xhttp "github.com/go-kratos/kratos/v2/transport/xhttp"
	"github.com/go-kratos/kratos/v2/transport/xhttp/apistate"
	binding "github.com/go-kratos/kratos/v2/transport/xhttp/binding"
	"github.com/gofiber/fiber/v2"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.BindBody

const _ = xhttp.SupportPackageIsVersion1
const _ = middleware.SupportPackageIsVersion1

// The greeting service definition.
type GreeterXHTTPServer interface {
	Add(context.Context, *HelloRequest) (*HelloReply, error)
	Del(context.Context, *HelloRequest) (*HelloReply, error)
	Edit(context.Context, *HelloRequest) (*HelloReply, error)
	Get(context.Context, *HelloRequest) (*HelloReply, error)
	List(context.Context, *PageRequest) (*HelloReply, error)
}

func RegisterGreeterXHTTPServer(s *xhttp.Server, srv GreeterXHTTPServer) {
	s.Route(func(r fiber.Router) {
		api := r.Group("api/demo")
		// Register all service annotation
		{
		}
		api.Get(
			"/list", _Greeter_List0_XHTTP_Handler(srv))
		api.Get(
			"/get/:name", _Greeter_Get0_XHTTP_Handler(srv))
		api.Post(
			"/", _Greeter_Add0_XHTTP_Handler(srv))
		api.Put(
			"/", _Greeter_Edit0_XHTTP_Handler(srv))
		api.Delete(
			"/", _Greeter_Del0_XHTTP_Handler(srv))
	})
}

// Gets a greeting for a user.
func _Greeter_List0_XHTTP_Handler(srv GreeterXHTTPServer) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var in PageRequest
		if err := binding.BindQuery(ctx, &in); err != nil {
			return err
		}
		if err := in.ValidateAll(); err != nil {
			return apistate.InvalidError().WithError(err).Send(ctx)
		}
		reply, err := srv.List(ctx.Context(), &in)
		if err != nil {
			return apistate.Error().WithError(err).Send(ctx)
		}
		return apistate.Success().WithData(reply).Send(ctx)
	}
}

// Gets a greeting for a user.
func _Greeter_Get0_XHTTP_Handler(srv GreeterXHTTPServer) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var in HelloRequest
		if err := binding.BindParams(ctx, &in); err != nil {
			return err
		}
		reply, err := srv.Get(ctx.Context(), &in)
		if err != nil {
			return apistate.Error().WithError(err).Send(ctx)
		}
		return apistate.Success().WithData(reply).Send(ctx)
	}
}

// Adds a greeting for a user.
func _Greeter_Add0_XHTTP_Handler(srv GreeterXHTTPServer) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var in HelloRequest
		if err := binding.BindBody(ctx, &in); err != nil {
			return err
		}
		reply, err := srv.Add(ctx.Context(), &in)
		if err != nil {
			return apistate.Error().WithError(err).Send(ctx)
		}
		return apistate.Success().WithData(reply).Send(ctx)
	}
}

// Adds a greeting for a user.
func _Greeter_Edit0_XHTTP_Handler(srv GreeterXHTTPServer) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var in HelloRequest
		if err := binding.BindBody(ctx, &in); err != nil {
			return err
		}
		reply, err := srv.Edit(ctx.Context(), &in)
		if err != nil {
			return apistate.Error().WithError(err).Send(ctx)
		}
		return apistate.Success().WithData(reply).Send(ctx)
	}
}

// Adds a greeting for a user.
func _Greeter_Del0_XHTTP_Handler(srv GreeterXHTTPServer) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var in HelloRequest
		if err := binding.BindQuery(ctx, &in); err != nil {
			return err
		}
		reply, err := srv.Del(ctx.Context(), &in)
		if err != nil {
			return apistate.Error().WithError(err).Send(ctx)
		}
		return apistate.Success().WithData(reply).Send(ctx)
	}
}
