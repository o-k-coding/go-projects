package gapi

import (
	"context"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

type Metadata struct {
	UserAgent string
	ClientIP  string
}

const (
	grpcGatewayUserAgentHeader = "grpcgateway-user-agent"
	xForwardedForHeader        = "x-forwarded-for"
	userAgentHeader            = "user-agent" // This is for straight grpc requests
)

func (server *Server) extractMetadata(ctx context.Context) *Metadata {
	meta := &Metadata{}

	if metaCtx, ok := metadata.FromIncomingContext(ctx); ok {
		if userAgents := metaCtx.Get(grpcGatewayUserAgentHeader); len(userAgents) > 0 {
			meta.UserAgent = userAgents[0]
		}

		// This one is for grpc request that does not go through the gateway
		if userAgents := metaCtx.Get(userAgentHeader); len(userAgents) > 0 && meta.UserAgent == "" {
			meta.UserAgent = userAgents[0]
		}

		if clientIPs := metaCtx.Get(xForwardedForHeader); len(clientIPs) > 0 {
			meta.ClientIP = clientIPs[0]
		}
	}

	// This one is for grpc request that does not go through the gateway
	if p, ok := peer.FromContext(ctx); ok && meta.ClientIP == "" {
		meta.ClientIP = p.Addr.String()
	}

	return meta
}
