# \NetworksApi

All URIs are relative to *https://api.demo.netmaker.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetwork**](NetworksApi.md#CreateNetwork) | **Post** /api/networks | Create a network.
[**DeleteNetwork**](NetworksApi.md#DeleteNetwork) | **Delete** /api/networks/{networkname} | Delete a network.  Will not delete if there are any nodes that belong to the network.
[**GetNetwork**](NetworksApi.md#GetNetwork) | **Get** /api/networks/{networkname} | Get a network.
[**GetNetworkACL**](NetworksApi.md#GetNetworkACL) | **Get** /api/networks/{networkname}/acls | Get a network ACL (Access Control List).
[**GetNetworks**](NetworksApi.md#GetNetworks) | **Get** /api/networks | Lists all networks.
[**UpdateNetwork**](NetworksApi.md#UpdateNetwork) | **Put** /api/networks/{networkname} | Update pro settings for a network.
[**UpdateNetworkACL**](NetworksApi.md#UpdateNetworkACL) | **Put** /api/networks/{networkname}/acls | Update a network ACL (Access Control List).
[**UpdateNetworkACL_0**](NetworksApi.md#UpdateNetworkACL_0) | **Put** /api/networks/{networkname}/acls/v2 | Update a network ACL (Access Control List).


# **CreateNetwork**
> Network CreateNetwork(ctx, optional)
Create a network.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NetworksApiCreateNetworkOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NetworksApiCreateNetworkOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **network** | [**optional.Interface of Network**](Network.md)| Network | 

### Return type

[**Network**](Network.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNetwork**
> SuccessResponse DeleteNetwork(ctx, networkname)
Delete a network.  Will not delete if there are any nodes that belong to the network.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **networkname** | **string**| name: network name | 

### Return type

[**SuccessResponse**](SuccessResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNetwork**
> Network GetNetwork(ctx, networkname)
Get a network.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **networkname** | **string**| name: network name | 

### Return type

[**Network**](Network.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNetworkACL**
> AclContainer GetNetworkACL(ctx, networkname)
Get a network ACL (Access Control List).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **networkname** | **string**| name: network name | 

### Return type

[**AclContainer**](ACLContainer.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNetworks**
> []Network GetNetworks(ctx, )
Lists all networks.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Network**](Network.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNetwork**
> Network UpdateNetwork(ctx, networkname, optional)
Update pro settings for a network.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **networkname** | **string**| name: network name | 
 **optional** | ***NetworksApiUpdateNetworkOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NetworksApiUpdateNetworkOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **network** | [**optional.Interface of Network**](Network.md)| Network | 

### Return type

[**Network**](Network.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNetworkACL**
> AclContainer UpdateNetworkACL(ctx, networkname, optional)
Update a network ACL (Access Control List).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **networkname** | **string**| name: network name | 
 **optional** | ***NetworksApiUpdateNetworkACLOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NetworksApiUpdateNetworkACLOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aclContainer** | [**optional.Interface of AclContainer**](AclContainer.md)| ACL Container | 

### Return type

[**AclContainer**](ACLContainer.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNetworkACL_0**
> AclContainer UpdateNetworkACL_0(ctx, networkname, optional)
Update a network ACL (Access Control List).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **networkname** | **string**| name: network name | 
 **optional** | ***NetworksApiUpdateNetworkACL_1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NetworksApiUpdateNetworkACL_1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aclContainer** | [**optional.Interface of AclContainer**](AclContainer.md)| ACL Container | 

### Return type

[**AclContainer**](ACLContainer.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

