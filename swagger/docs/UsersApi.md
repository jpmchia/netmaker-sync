# \UsersApi

All URIs are relative to *https://api.demo.netmaker.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IngressGatewayUsers**](UsersApi.md#IngressGatewayUsers) | **Get** /api/nodes/{network}/{nodeid}/ingress/users | Lists all the users attached to an ingress gateway.


# **IngressGatewayUsers**
> LegacyNode IngressGatewayUsers(ctx, network, nodeid)
Lists all the users attached to an ingress gateway.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **network** | **string**|  | 
  **nodeid** | **string**|  | 

### Return type

[**LegacyNode**](LegacyNode.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

