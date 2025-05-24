# \ExtClientApi

All URIs are relative to *https://api.demo.netmaker.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateExtClient**](ExtClientApi.md#CreateExtClient) | **Post** /api/extclients/{network}/{nodeid} | Create an individual extclient.  Must have valid key and be unique.
[**DeleteExtClient**](ExtClientApi.md#DeleteExtClient) | **Delete** /api/extclients/{network}/{clientid} | Delete an individual extclient.
[**GetAllExtClients**](ExtClientApi.md#GetAllExtClients) | **Get** /api/extclients | A separate function to get all extclients, not just extclients for a particular network.
[**GetExtClient**](ExtClientApi.md#GetExtClient) | **Get** /api/extclients/{network}/{clientid} | Get an individual extclient.
[**GetExtClientConf**](ExtClientApi.md#GetExtClientConf) | **Get** /api/extclients/{network}/{clientid}/{type} | Get an individual extclient.
[**GetNetworkExtClients**](ExtClientApi.md#GetNetworkExtClients) | **Get** /api/extclients/{network} | Get all extclients associated with network.
[**UpdateExtClient**](ExtClientApi.md#UpdateExtClient) | **Put** /api/extclients/{network}/{clientid} | Update an individual extclient.


# **CreateExtClient**
> CreateExtClient(ctx, network, nodeid, optional)
Create an individual extclient.  Must have valid key and be unique.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **network** | **string**| Network | 
  **nodeid** | **string**| Node ID | 
 **optional** | ***ExtClientApiCreateExtClientOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExtClientApiCreateExtClientOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **customExtClient** | [**optional.Interface of CustomExtClient**](CustomExtClient.md)| Custom ExtClient | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteExtClient**
> SuccessResponse DeleteExtClient(ctx, clientid, network)
Delete an individual extclient.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientid** | **string**| Client ID | 
  **network** | **string**| Network | 

### Return type

[**SuccessResponse**](SuccessResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllExtClients**
> []ExtClient GetAllExtClients(ctx, optional)
A separate function to get all extclients, not just extclients for a particular network.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExtClientApiGetAllExtClientsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExtClientApiGetAllExtClientsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **networks** | **optional.[]string**| Networks | 

### Return type

[**[]ExtClient**](ExtClient.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExtClient**
> ExtClient GetExtClient(ctx, clientid, network)
Get an individual extclient.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientid** | **string**| Client ID | 
  **network** | **string**| Network | 

### Return type

[**ExtClient**](ExtClient.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExtClientConf**
> ExtClient GetExtClientConf(ctx, type_, clientid, network)
Get an individual extclient.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**| Type | 
  **clientid** | **string**| Client ID | 
  **network** | **string**| Network | 

### Return type

[**ExtClient**](ExtClient.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNetworkExtClients**
> []ExtClient GetNetworkExtClients(ctx, network)
Get all extclients associated with network.

Gets all extclients associated with network, including pending extclients.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **network** | **string**| Network | 

### Return type

[**[]ExtClient**](ExtClient.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateExtClient**
> ExtClient UpdateExtClient(ctx, clientid, network, optional)
Update an individual extclient.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientid** | **string**| Client ID | 
  **network** | **string**| Network | 
 **optional** | ***ExtClientApiUpdateExtClientOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExtClientApiUpdateExtClientOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **extClient** | [**optional.Interface of ExtClient**](ExtClient.md)| ExtClient | 

### Return type

[**ExtClient**](ExtClient.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

