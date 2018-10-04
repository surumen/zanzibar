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

package barendpoint

import (
	"context"
	"encoding/json"
	"fmt"
	"runtime/debug"

	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	zanzibar "github.com/uber/zanzibar/runtime"
	"go.uber.org/thriftrw/ptr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	workflow "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/bar/workflow"
	endpointsBarBar "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/endpoints/bar/bar"

	module "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/bar/module"
)

// BarArgWithManyQueryParamsHandler is the handler for "/bar/argWithManyQueryParams"
type BarArgWithManyQueryParamsHandler struct {
	Dependencies *module.Dependencies
	endpoint     *zanzibar.RouterEndpoint
}

// NewBarArgWithManyQueryParamsHandler creates a handler
func NewBarArgWithManyQueryParamsHandler(deps *module.Dependencies) *BarArgWithManyQueryParamsHandler {
	handler := &BarArgWithManyQueryParamsHandler{
		Dependencies: deps,
	}
	handler.endpoint = zanzibar.NewRouterEndpoint(
		deps.Default.ContextExtractor, deps.Default.Logger, deps.Default.Scope, deps.Default.Tracer,
		"bar", "argWithManyQueryParams",
		handler.HandleRequest,
	)

	return handler
}

// Register adds the http handler to the gateway's http router
func (h *BarArgWithManyQueryParamsHandler) Register(g *zanzibar.Gateway) error {
	return g.HTTPRouter.Register(
		"GET", "/bar/argWithManyQueryParams",
		h.endpoint,
	)
}

// HandleRequest handles "/bar/argWithManyQueryParams".
func (h *BarArgWithManyQueryParamsHandler) HandleRequest(
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

	var requestBody endpointsBarBar.Bar_ArgWithManyQueryParams_Args

	aStrOk := req.CheckQueryValue("aStr")
	if !aStrOk {
		return
	}
	aStrQuery, ok := req.GetQueryValue("aStr")
	if !ok {
		return
	}
	requestBody.AStr = aStrQuery

	anOptStrOk := req.HasQueryValue("anOptStr")
	if anOptStrOk {
		anOptStrQuery, ok := req.GetQueryValue("anOptStr")
		if !ok {
			return
		}
		requestBody.AnOptStr = ptr.String(anOptStrQuery)
	}

	aBoolOk := req.CheckQueryValue("aBool")
	if !aBoolOk {
		return
	}
	aBoolQuery, ok := req.GetQueryBool("aBool")
	if !ok {
		return
	}
	requestBody.ABool = aBoolQuery

	anOptBoolOk := req.HasQueryValue("anOptBool")
	if anOptBoolOk {
		anOptBoolQuery, ok := req.GetQueryBool("anOptBool")
		if !ok {
			return
		}
		requestBody.AnOptBool = ptr.Bool(anOptBoolQuery)
	}

	aInt8Ok := req.CheckQueryValue("aInt8")
	if !aInt8Ok {
		return
	}
	aInt8Query, ok := req.GetQueryInt8("aInt8")
	if !ok {
		return
	}
	requestBody.AInt8 = aInt8Query

	anOptInt8Ok := req.HasQueryValue("anOptInt8")
	if anOptInt8Ok {
		anOptInt8Query, ok := req.GetQueryInt8("anOptInt8")
		if !ok {
			return
		}
		requestBody.AnOptInt8 = ptr.Int8(anOptInt8Query)
	}

	aInt16Ok := req.CheckQueryValue("aInt16")
	if !aInt16Ok {
		return
	}
	aInt16Query, ok := req.GetQueryInt16("aInt16")
	if !ok {
		return
	}
	requestBody.AInt16 = aInt16Query

	anOptInt16Ok := req.HasQueryValue("anOptInt16")
	if anOptInt16Ok {
		anOptInt16Query, ok := req.GetQueryInt16("anOptInt16")
		if !ok {
			return
		}
		requestBody.AnOptInt16 = ptr.Int16(anOptInt16Query)
	}

	aInt32Ok := req.CheckQueryValue("aInt32")
	if !aInt32Ok {
		return
	}
	aInt32Query, ok := req.GetQueryInt32("aInt32")
	if !ok {
		return
	}
	requestBody.AInt32 = aInt32Query

	anOptInt32Ok := req.HasQueryValue("anOptInt32")
	if anOptInt32Ok {
		anOptInt32Query, ok := req.GetQueryInt32("anOptInt32")
		if !ok {
			return
		}
		requestBody.AnOptInt32 = ptr.Int32(anOptInt32Query)
	}

	aInt64Ok := req.CheckQueryValue("aInt64")
	if !aInt64Ok {
		return
	}
	aInt64Query, ok := req.GetQueryInt64("aInt64")
	if !ok {
		return
	}
	requestBody.AInt64 = aInt64Query

	anOptInt64Ok := req.HasQueryValue("anOptInt64")
	if anOptInt64Ok {
		anOptInt64Query, ok := req.GetQueryInt64("anOptInt64")
		if !ok {
			return
		}
		requestBody.AnOptInt64 = ptr.Int64(anOptInt64Query)
	}

	aFloat64Ok := req.CheckQueryValue("aFloat64")
	if !aFloat64Ok {
		return
	}
	aFloat64Query, ok := req.GetQueryFloat64("aFloat64")
	if !ok {
		return
	}
	requestBody.AFloat64 = aFloat64Query

	anOptFloat64Ok := req.HasQueryValue("anOptFloat64")
	if anOptFloat64Ok {
		anOptFloat64Query, ok := req.GetQueryFloat64("anOptFloat64")
		if !ok {
			return
		}
		requestBody.AnOptFloat64 = ptr.Float64(anOptFloat64Query)
	}

	aUUIDOk := req.CheckQueryValue("aUUID")
	if !aUUIDOk {
		return
	}
	aUUIDQuery, ok := req.GetQueryValue("aUUID")
	if !ok {
		return
	}
	requestBody.AUUID = endpointsBarBar.UUID(aUUIDQuery)

	anOptUUIDOk := req.HasQueryValue("anOptUUID")
	if anOptUUIDOk {
		anOptUUIDQuery, ok := req.GetQueryValue("anOptUUID")
		if !ok {
			return
		}
		requestBody.AnOptUUID = (*endpointsBarBar.UUID)(ptr.String(anOptUUIDQuery))
	}

	aListUUIDOk := req.CheckQueryValue("aListUUID[]")
	if !aListUUIDOk {
		return
	}
	aListUUIDQuery, ok := req.GetQueryValues("aListUUID[]")
	if !ok {
		return
	}
	aListUUIDQueryFinal := make([]endpointsBarBar.UUID, len(aListUUIDQuery))
	for i, v := range aListUUIDQuery {
		aListUUIDQueryFinal[i] = endpointsBarBar.UUID(v)
	}
	requestBody.AListUUID = aListUUIDQueryFinal

	anOptListUUIDOk := req.HasQueryValue("anOptListUUID[]")
	if anOptListUUIDOk {
		anOptListUUIDQuery, ok := req.GetQueryValues("anOptListUUID[]")
		if !ok {
			return
		}
		anOptListUUIDQueryFinal := make([]endpointsBarBar.UUID, len(anOptListUUIDQuery))
		for i, v := range anOptListUUIDQuery {
			anOptListUUIDQueryFinal[i] = endpointsBarBar.UUID(v)
		}
		requestBody.AnOptListUUID = anOptListUUIDQueryFinal
	}

	aStringListOk := req.CheckQueryValue("aStringList[]")
	if !aStringListOk {
		return
	}
	aStringListQuery, ok := req.GetQueryValues("aStringList[]")
	if !ok {
		return
	}
	requestBody.AStringList = endpointsBarBar.StringList(aStringListQuery)

	anOptStringListOk := req.HasQueryValue("anOptStringList[]")
	if anOptStringListOk {
		anOptStringListQuery, ok := req.GetQueryValues("anOptStringList[]")
		if !ok {
			return
		}
		requestBody.AnOptStringList = endpointsBarBar.StringList(anOptStringListQuery)
	}

	aUUIDListOk := req.CheckQueryValue("aUUIDList[]")
	if !aUUIDListOk {
		return
	}
	aUUIDListQuery, ok := req.GetQueryValues("aUUIDList[]")
	if !ok {
		return
	}
	aUUIDListQueryFinal := make([]endpointsBarBar.UUID, len(aUUIDListQuery))
	for i, v := range aUUIDListQuery {
		aUUIDListQueryFinal[i] = endpointsBarBar.UUID(v)
	}
	requestBody.AUUIDList = endpointsBarBar.UUIDList(aUUIDListQueryFinal)

	anOptUUIDListOk := req.HasQueryValue("anOptUUIDList[]")
	if anOptUUIDListOk {
		anOptUUIDListQuery, ok := req.GetQueryValues("anOptUUIDList[]")
		if !ok {
			return
		}
		anOptUUIDListQueryFinal := make([]endpointsBarBar.UUID, len(anOptUUIDListQuery))
		for i, v := range anOptUUIDListQuery {
			anOptUUIDListQueryFinal[i] = endpointsBarBar.UUID(v)
		}
		requestBody.AnOptUUIDList = endpointsBarBar.UUIDList(anOptUUIDListQueryFinal)
	}

	aTsOk := req.CheckQueryValue("aTs")
	if !aTsOk {
		return
	}
	aTsQuery, ok := req.GetQueryInt64("aTs")
	if !ok {
		return
	}
	requestBody.ATs = endpointsBarBar.Timestamp(aTsQuery)

	anOptTsOk := req.HasQueryValue("anOptTs")
	if anOptTsOk {
		anOptTsQuery, ok := req.GetQueryInt64("anOptTs")
		if !ok {
			return
		}
		requestBody.AnOptTs = (*endpointsBarBar.Timestamp)(ptr.Int64(anOptTsQuery))
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

	w := workflow.NewBarArgWithManyQueryParamsWorkflow(h.Dependencies)
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
		res.SendError(500, "Unexpected server error", err)
		return

	}

	res.WriteJSON(200, cliRespHeaders, response)
}
