package network

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"

	mgmtnetwork "github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-07-01/network"
	"github.com/Azure/go-autorest/autorest"
)

// PrivateLinkServicesClient is a minimal interface for azure PrivateLinkServicesClient
type PrivateLinkServicesClient interface {
	DeletePrivateEndpointConnection(ctx context.Context, resourceGroupName string, serviceName string, peConnectionName string) (result mgmtnetwork.PrivateLinkServicesDeletePrivateEndpointConnectionFuture, err error)
	PrivateLinkServicesClientAddons
}

type privateLinkServicesClient struct {
	mgmtnetwork.PrivateLinkServicesClient
}

var _ PrivateLinkServicesClient = &privateLinkServicesClient{}

// NewPrivateLinkServicesClient creates a new PrivateLinkServicesClient
func NewPrivateLinkServicesClient(subscriptionID string, authorizer autorest.Authorizer) PrivateLinkServicesClient {
	client := mgmtnetwork.NewPrivateLinkServicesClient(subscriptionID)
	client.Authorizer = authorizer

	return &privateLinkServicesClient{
		PrivateLinkServicesClient: client,
	}
}
