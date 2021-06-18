// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// VirtualWansClient contains the methods for the VirtualWans group.
// Don't use this type directly, use NewVirtualWansClient() instead.
type VirtualWansClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewVirtualWansClient creates a new instance of VirtualWansClient with the specified values.
func NewVirtualWansClient(con *armcore.Connection, subscriptionID string) *VirtualWansClient {
	return &VirtualWansClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates a VirtualWAN resource if it doesn't exist else updates the existing VirtualWAN.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualWansClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, virtualWANName string, wanParameters VirtualWAN, options *VirtualWansBeginCreateOrUpdateOptions) (VirtualWANPollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, virtualWANName, wanParameters, options)
	if err != nil {
		return VirtualWANPollerResponse{}, err
	}
	result := VirtualWANPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("VirtualWansClient.CreateOrUpdate", "azure-async-operation", resp, client.con.Pipeline(), client.createOrUpdateHandleError)
	if err != nil {
		return VirtualWANPollerResponse{}, err
	}
	poller := &virtualWANPoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (VirtualWANResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new VirtualWANPoller from the specified resume token.
// token - The value must come from a previous call to VirtualWANPoller.ResumeToken().
func (client *VirtualWansClient) ResumeCreateOrUpdate(ctx context.Context, token string) (VirtualWANPollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("VirtualWansClient.CreateOrUpdate", token, client.con.Pipeline(), client.createOrUpdateHandleError)
	if err != nil {
		return VirtualWANPollerResponse{}, err
	}
	poller := &virtualWANPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return VirtualWANPollerResponse{}, err
	}
	result := VirtualWANPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (VirtualWANResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// CreateOrUpdate - Creates a VirtualWAN resource if it doesn't exist else updates the existing VirtualWAN.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualWansClient) createOrUpdate(ctx context.Context, resourceGroupName string, virtualWANName string, wanParameters VirtualWAN, options *VirtualWansBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, virtualWANName, wanParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *VirtualWansClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, virtualWANName string, wanParameters VirtualWAN, options *VirtualWansBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans/{VirtualWANName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualWANName == "" {
		return nil, errors.New("parameter virtualWANName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{VirtualWANName}", url.PathEscape(virtualWANName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(wanParameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *VirtualWansClient) createOrUpdateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// BeginDelete - Deletes a VirtualWAN.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualWansClient) BeginDelete(ctx context.Context, resourceGroupName string, virtualWANName string, options *VirtualWansBeginDeleteOptions) (HTTPPollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, virtualWANName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("VirtualWansClient.Delete", "location", resp, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDelete creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client *VirtualWansClient) ResumeDelete(ctx context.Context, token string) (HTTPPollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("VirtualWansClient.Delete", token, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Delete - Deletes a VirtualWAN.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualWansClient) deleteOperation(ctx context.Context, resourceGroupName string, virtualWANName string, options *VirtualWansBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, virtualWANName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *VirtualWansClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, virtualWANName string, options *VirtualWansBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans/{VirtualWANName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualWANName == "" {
		return nil, errors.New("parameter virtualWANName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{VirtualWANName}", url.PathEscape(virtualWANName))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *VirtualWansClient) deleteHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Get - Retrieves the details of a VirtualWAN.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualWansClient) Get(ctx context.Context, resourceGroupName string, virtualWANName string, options *VirtualWansGetOptions) (VirtualWANResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, virtualWANName, options)
	if err != nil {
		return VirtualWANResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return VirtualWANResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return VirtualWANResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VirtualWansClient) getCreateRequest(ctx context.Context, resourceGroupName string, virtualWANName string, options *VirtualWansGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans/{VirtualWANName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualWANName == "" {
		return nil, errors.New("parameter virtualWANName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{VirtualWANName}", url.PathEscape(virtualWANName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VirtualWansClient) getHandleResponse(resp *azcore.Response) (VirtualWANResponse, error) {
	var val *VirtualWAN
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return VirtualWANResponse{}, err
	}
	return VirtualWANResponse{RawResponse: resp.Response, VirtualWAN: val}, nil
}

// getHandleError handles the Get error response.
func (client *VirtualWansClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// List - Lists all the VirtualWANs in a subscription.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualWansClient) List(options *VirtualWansListOptions) ListVirtualWANsResultPager {
	return &listVirtualWANsResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp ListVirtualWANsResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ListVirtualWANsResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client *VirtualWansClient) listCreateRequest(ctx context.Context, options *VirtualWansListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/virtualWans"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VirtualWansClient) listHandleResponse(resp *azcore.Response) (ListVirtualWANsResultResponse, error) {
	var val *ListVirtualWANsResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ListVirtualWANsResultResponse{}, err
	}
	return ListVirtualWANsResultResponse{RawResponse: resp.Response, ListVirtualWANsResult: val}, nil
}

// listHandleError handles the List error response.
func (client *VirtualWansClient) listHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// ListByResourceGroup - Lists all the VirtualWANs in a resource group.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualWansClient) ListByResourceGroup(resourceGroupName string, options *VirtualWansListByResourceGroupOptions) ListVirtualWANsResultPager {
	return &listVirtualWANsResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.listByResourceGroupHandleResponse,
		errorer:   client.listByResourceGroupHandleError,
		advancer: func(ctx context.Context, resp ListVirtualWANsResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ListVirtualWANsResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *VirtualWansClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *VirtualWansListByResourceGroupOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *VirtualWansClient) listByResourceGroupHandleResponse(resp *azcore.Response) (ListVirtualWANsResultResponse, error) {
	var val *ListVirtualWANsResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ListVirtualWANsResultResponse{}, err
	}
	return ListVirtualWANsResultResponse{RawResponse: resp.Response, ListVirtualWANsResult: val}, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *VirtualWansClient) listByResourceGroupHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// UpdateTags - Updates a VirtualWAN tags.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualWansClient) UpdateTags(ctx context.Context, resourceGroupName string, virtualWANName string, wanParameters TagsObject, options *VirtualWansUpdateTagsOptions) (VirtualWANResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, virtualWANName, wanParameters, options)
	if err != nil {
		return VirtualWANResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return VirtualWANResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return VirtualWANResponse{}, client.updateTagsHandleError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *VirtualWansClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, virtualWANName string, wanParameters TagsObject, options *VirtualWansUpdateTagsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans/{VirtualWANName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualWANName == "" {
		return nil, errors.New("parameter virtualWANName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{VirtualWANName}", url.PathEscape(virtualWANName))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(wanParameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *VirtualWansClient) updateTagsHandleResponse(resp *azcore.Response) (VirtualWANResponse, error) {
	var val *VirtualWAN
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return VirtualWANResponse{}, err
	}
	return VirtualWANResponse{RawResponse: resp.Response, VirtualWAN: val}, nil
}

// updateTagsHandleError handles the UpdateTags error response.
func (client *VirtualWansClient) updateTagsHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
