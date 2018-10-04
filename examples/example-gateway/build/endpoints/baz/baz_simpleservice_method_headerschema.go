// Code generated by zanzibar
// @generated

// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package bazendpoint

import (
	"context"
	"encoding/json"
	"fmt"
	"runtime/debug"

	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	zanzibar "github.com/uber/zanzibar/runtime"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	workflow "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/baz/workflow"
	endpointsBazBaz "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/endpoints/baz/baz"

	module "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/baz/module"
)

// SimpleServiceHeaderSchemaHandler is the handler for "/baz/header-schema"
type SimpleServiceHeaderSchemaHandler struct {
	Dependencies *module.Dependencies
	endpoint     *zanzibar.RouterEndpoint
}

// NewSimpleServiceHeaderSchemaHandler creates a handler
func NewSimpleServiceHeaderSchemaHandler(deps *module.Dependencies) *SimpleServiceHeaderSchemaHandler {
	handler := &SimpleServiceHeaderSchemaHandler{
		Dependencies: deps,
	}
	handler.endpoint = zanzibar.NewRouterEndpoint(
		deps.Default.ContextExtractor, deps.Default.Logger, deps.Default.Scope, deps.Default.Tracer,
		"baz", "headerSchema",
		handler.HandleRequest,
	)

	return handler
}

// Register adds the http handler to the gateway's http router
func (h *SimpleServiceHeaderSchemaHandler) Register(g *zanzibar.Gateway) error {
	return g.HTTPRouter.Register(
		"POST", "/baz/header-schema",
		h.endpoint,
	)
}

// HandleRequest handles "/baz/header-schema".
func (h *SimpleServiceHeaderSchemaHandler) HandleRequest(
	ctx context.Context,
	req *zanzibar.ServerHTTPRequest,
	res *zanzibar.ServerHTTPResponse,
) {
	defer func() {
		if r := recover(); r != nil {
			stacktrace := string(debug.Stack())
			e := errors.Errorf("enpoint panic: %v, stacktrace: %v", r, stacktrace)
			h.Dependencies.Default.ContextLogger.Error(
				ctx,
				"endpoint panic",
				zap.Error(e),
				zap.String("stacktrace", stacktrace),
				zap.String("endpoint", h.endpoint.EndpointName))

			h.endpoint.EndpointMetrics.Panic.Inc(1)
			res.SendError(502, "Unexpected workflow panic, recovered at endpoint.", nil)
		}
	}()

	if !req.CheckHeaders([]string{"Auth", "Content-Type"}) {
		return
	}
	var requestBody endpointsBazBaz.SimpleService_HeaderSchema_Args
	if ok := req.ReadAndUnmarshalBody(&requestBody); !ok {
		return
	}

	// log endpoint request to downstream services
	if ce := h.Dependencies.Default.ContextLogger.Check(zapcore.DebugLevel, "stub"); ce != nil {
		zfields := []zapcore.Field{
			zap.String("endpoint", h.endpoint.EndpointName),
		}
		zfields = append(zfields, zap.String("body", fmt.Sprintf("%s", req.GetRawBody())))
		for _, k := range req.Header.Keys() {
			if val, ok := req.Header.Get(k); ok {
				zfields = append(zfields, zap.String(k, val))
			}
		}
		h.Dependencies.Default.ContextLogger.Debug(ctx, "endpoint request to downstream", zfields...)
	}

	w := workflow.NewSimpleServiceHeaderSchemaWorkflow(h.Dependencies)
	if span := req.GetSpan(); span != nil {
		ctx = opentracing.ContextWithSpan(ctx, span)
	}

	response, cliRespHeaders, err := w.Handle(ctx, req.Header, &requestBody)

	// log downstream response to endpoint
	if ce := h.Dependencies.Default.ContextLogger.Check(zapcore.DebugLevel, "stub"); ce != nil {
		zfields := []zapcore.Field{
			zap.String("endpoint", h.endpoint.EndpointName),
		}
		if body, err := json.Marshal(response); err == nil {
			zfields = append(zfields, zap.String("body", fmt.Sprintf("%s", body)))
		}
		for _, k := range cliRespHeaders.Keys() {
			if val, ok := cliRespHeaders.Get(k); ok {
				zfields = append(zfields, zap.String(k, val))
			}
		}
		if traceKey, ok := req.Header.Get("x-trace-id"); ok {
			zfields = append(zfields, zap.String("x-trace-id", traceKey))
		}
		h.Dependencies.Default.ContextLogger.Debug(ctx, "downstream service response", zfields...)
	}

	if err != nil {
		switch errValue := err.(type) {

		case *endpointsBazBaz.AuthErr:
			res.WriteJSON(
				401, cliRespHeaders, errValue,
			)
			return

		case *endpointsBazBaz.OtherAuthErr:
			res.WriteJSON(
				403, cliRespHeaders, errValue,
			)
			return

		default:
			res.SendError(500, "Unexpected server error", err)
			return
		}

	}

	res.WriteJSON(200, cliRespHeaders, response)
}
