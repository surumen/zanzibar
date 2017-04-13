// Code generated by zanzibar
// @generated

package bar

import (
	"context"

	"github.com/uber/zanzibar/examples/example-gateway/build/clients"
	zanzibar "github.com/uber/zanzibar/runtime"
	"go.uber.org/zap"

	clientsBarBar "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/bar/bar"
	endpointsBarBar "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/endpoints/bar/bar"
)

// HandleNoRequestRequest handles "/bar/no-request-path".
func HandleNoRequestRequest(
	ctx context.Context,
	req *zanzibar.ServerHTTPRequest,
	res *zanzibar.ServerHTTPResponse,
	clients *clients.Clients,
) {

	headers := map[string]string{}

	workflow := NoRequestEndpoint{
		Clients: clients,
		Logger:  req.Logger,
		Request: req,
	}

	response, _, err := workflow.Handle(ctx, headers)
	if err != nil {
		req.Logger.Warn("Workflow for endpoint returned error",
			zap.String("error", err.Error()),
		)
		res.SendErrorString(500, "Unexpected server error")
		return
	}

	res.WriteJSON(200, response)
}

// NoRequestEndpoint calls thrift client Bar.NoRequest
type NoRequestEndpoint struct {
	Clients *clients.Clients
	Logger  *zap.Logger
	Request *zanzibar.ServerHTTPRequest
}

// Handle calls thrift client.
func (w NoRequestEndpoint) Handle(
	ctx context.Context,
	headers map[string]string,
) (*endpointsBarBar.BarResponse, map[string]string, error) {

	clientRespBody, _, err := w.Clients.Bar.NoRequest(
		ctx, nil,
	)
	if err != nil {
		w.Logger.Warn("Could not make client request",
			zap.String("error", err.Error()),
		)
		return nil, nil, err
	}

	response := convertNoRequestClientResponse(clientRespBody)
	return response, nil, nil
}

func convertNoRequestClientResponse(body *clientsBarBar.BarResponse) *endpointsBarBar.BarResponse {
	// TODO: Add response fields mapping here.
	downstreamResponse := (*endpointsBarBar.BarResponse)(body)
	return downstreamResponse
}
