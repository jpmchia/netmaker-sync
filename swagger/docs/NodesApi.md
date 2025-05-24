# \NodesApi

All URIs are relative to *https://api.demo.netmaker.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEgressGateway**](NodesApi.md#CreateEgressGateway) | **Post** /api/nodes/{network}/{nodeid}/creategateway | Create an egress gateway.
[**CreateIngressGateway**](NodesApi.md#CreateIngressGateway) | **Post** /api/nodes/{network}/{nodeid}/createingress | Create an ingress gateway.
[**CreateInternetGw**](NodesApi.md#CreateInternetGw) | **Post** /api/nodes/{network}/{nodeid}/inet_gw | Create an inet node.
[**CreateRelay**](NodesApi.md#CreateRelay) | **Post** /api/nodes/{network}/{nodeid}/createrelay | Create a relay.
[**DeleteEgressGateway**](NodesApi.md#DeleteEgressGateway) | **Delete** /api/nodes/{network}/{nodeid}/deletegateway | Delete an egress gateway.
[**DeleteIngressGateway**](NodesApi.md#DeleteIngressGateway) | **Delete** /api/nodes/{network}/{nodeid}/deleteingress | Delete an ingress gateway.
[**DeleteInternetGw**](NodesApi.md#DeleteInternetGw) | **Delete** /api/nodes/{network}/{nodeid}/inet_gw | Delete an internet gw.
[**DeleteNode**](NodesApi.md#DeleteNode) | **Delete** /api/nodes/{network}/{nodeid} | Delete an individual node.
[**DeleteRelay**](NodesApi.md#DeleteRelay) | **Delete** /api/nodes/{network}/{nodeid}/deleterelay | Remove a relay.
[**GetAllNodes**](NodesApi.md#GetAllNodes) | **Get** /api/nodes | Get all nodes across all networks.
[**GetNetworkNodes**](NodesApi.md#GetNetworkNodes) | **Get** /api/nodes/{network} | Gets all nodes associated with network including pending nodes.
[**GetNode**](NodesApi.md#GetNode) | **Get** /api/nodes/{network}/{nodeid} | Get an individual node.
[**HandleAuthLogin**](NodesApi.md#HandleAuthLogin) | **Get** /api/oauth/login | Handles OAuth login.
[**MigrateData**](NodesApi.md#MigrateData) | **Put** /api/v1/nodes/migrate | Used to migrate a legacy node.
[**UpdateInternetGw**](NodesApi.md#UpdateInternetGw) | **Put** /api/nodes/{network}/{nodeid}/inet_gw | update an inet node.
[**UpdateNode**](NodesApi.md#UpdateNode) | **Put** /api/nodes/{network}/{nodeid} | Update an individual node.
[**WipeLegacyNodes**](NodesApi.md#WipeLegacyNodes) | **Delete** /api/v1/legacy/nodes | Delete all legacy nodes from DB.


# **CreateEgressGateway**
> LegacyNode CreateEgressGateway(ctx, network, nodeid, optional)
Create an egress gateway.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **network** | **string**|  | 
  **nodeid** | **string**|  | 
 **optional** | ***NodesApiCreateEgressGatewayOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiCreateEgressGatewayOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **egressGatewayRequest** | [**optional.Interface of EgressGatewayRequest**](EgressGatewayRequest.md)| Egress Gateway Request | 

### Return type

[**LegacyNode**](LegacyNode.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateIngressGateway**
> LegacyNode CreateIngressGateway(ctx, network, nodeid)
Create an ingress gateway.

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

# **CreateInternetGw**
> LegacyNode CreateInternetGw(ctx, )
Create an inet node.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**LegacyNode**](LegacyNode.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRelay**
> LegacyNode CreateRelay(ctx, network, nodeid, optional)
Create a relay.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **network** | **string**|  | 
  **nodeid** | **string**|  | 
 **optional** | ***NodesApiCreateRelayOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiCreateRelayOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **relayRequest** | [**optional.Interface of RelayRequest**](RelayRequest.md)| Relay Request | 

### Return type

[**LegacyNode**](LegacyNode.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEgressGateway**
> LegacyNode DeleteEgressGateway(ctx, network, nodeid)
Delete an egress gateway.

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

# **DeleteIngressGateway**
> LegacyNode DeleteIngressGateway(ctx, network, nodeid)
Delete an ingress gateway.

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

# **DeleteInternetGw**
> LegacyNode DeleteInternetGw(ctx, )
Delete an internet gw.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**LegacyNode**](LegacyNode.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNode**
> LegacyNode DeleteNode(ctx, network, nodeid, optional)
Delete an individual node.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **network** | **string**|  | 
  **nodeid** | **string**|  | 
 **optional** | ***NodesApiDeleteNodeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiDeleteNodeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **node** | [**optional.Interface of LegacyNode**](LegacyNode.md)| Node | 

### Return type

[**LegacyNode**](LegacyNode.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRelay**
> LegacyNode DeleteRelay(ctx, network, nodeid)
Remove a relay.

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

# **GetAllNodes**
> []ApiNode GetAllNodes(ctx, )
Get all nodes across all networks.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ApiNode**](ApiNode.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNetworkNodes**
> []ApiNode GetNetworkNodes(ctx, network)
Gets all nodes associated with network including pending nodes.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **network** | **string**| Network | 

### Return type

[**[]ApiNode**](ApiNode.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNode**
> LegacyNode GetNode(ctx, network, nodeid)
Get an individual node.

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

# **HandleAuthLogin**
> HandleAuthLogin(ctx, )
Handles OAuth login.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MigrateData**
> HostPull MigrateData(ctx, )
Used to migrate a legacy node.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**HostPull**](HostPull.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateInternetGw**
> LegacyNode UpdateInternetGw(ctx, )
update an inet node.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**LegacyNode**](LegacyNode.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNode**
> LegacyNode UpdateNode(ctx, network, nodeid, optional)
Update an individual node.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **network** | **string**|  | 
  **nodeid** | **string**|  | 
 **optional** | ***NodesApiUpdateNodeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiUpdateNodeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **node** | [**optional.Interface of LegacyNode**](LegacyNode.md)| Node | 

### Return type

[**LegacyNode**](LegacyNode.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WipeLegacyNodes**
> SuccessResponse WipeLegacyNodes(ctx, )
Delete all legacy nodes from DB.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SuccessResponse**](SuccessResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

