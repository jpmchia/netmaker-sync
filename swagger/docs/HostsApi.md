# \HostsApi

All URIs are relative to *https://api.demo.netmaker.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddHostToNetwork**](HostsApi.md#AddHostToNetwork) | **Post** /api/hosts/{hostid}/networks/{network} | Given a network, a host is added to the network.
[**DelEmqxHosts**](HostsApi.md#DelEmqxHosts) | **Delete** /api/emqx/hosts | Lists all hosts.
[**DeleteHost**](HostsApi.md#DeleteHost) | **Delete** /api/hosts/{hostid} | Deletes a Netclient host from Netmaker server.
[**DeleteHostFromNetwork**](HostsApi.md#DeleteHostFromNetwork) | **Delete** /api/hosts/{hostid}/networks/{network} | Given a network, a host is removed from the network.
[**GetHosts**](HostsApi.md#GetHosts) | **Get** /api/hosts | Lists all hosts.
[**HostUpdateFallback**](HostsApi.md#HostUpdateFallback) | **Put** /api/v1/fallback/host/{hostid} | Updates a Netclient host on Netmaker server.
[**PullHost**](HostsApi.md#PullHost) | **Get** /api/v1/host | 
[**SignalPeer**](HostsApi.md#SignalPeer) | **Post** /api/hosts/{hostid}/signalpeer | send signal to peer.
[**Synchost**](HostsApi.md#Synchost) | **Post** /api/hosts/{hostid}/sync | Requests a host to pull.
[**UpdateAllKeys**](HostsApi.md#UpdateAllKeys) | **Post** /api/hosts/keys | Update keys for a network.
[**UpdateHost**](HostsApi.md#UpdateHost) | **Put** /api/hosts/{hostid} | Updates a Netclient host on Netmaker server.
[**UpdateKeys**](HostsApi.md#UpdateKeys) | **Post** /api/hosts/{hostid}keys | Update keys for a network.


# **AddHostToNetwork**
> AddHostToNetwork(ctx, hostid, network)
Given a network, a host is added to the network.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hostid** | **string**| hostid to add or delete from network | 
  **network** | **string**| network | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DelEmqxHosts**
> ApiHost DelEmqxHosts(ctx, )
Lists all hosts.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiHost**](ApiHost.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteHost**
> ApiHost DeleteHost(ctx, hostid)
Deletes a Netclient host from Netmaker server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hostid** | **string**| HostID | 

### Return type

[**ApiHost**](ApiHost.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteHostFromNetwork**
> DeleteHostFromNetwork(ctx, hostid, network)
Given a network, a host is removed from the network.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hostid** | **string**| hostid to add or delete from network | 
  **network** | **string**| network | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHosts**
> []ApiHost GetHosts(ctx, )
Lists all hosts.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ApiHost**](ApiHost.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HostUpdateFallback**
> ApiHost HostUpdateFallback(ctx, )
Updates a Netclient host on Netmaker server.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiHost**](ApiHost.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PullHost**
> HostPull PullHost(ctx, )


Used by clients for \"pull\" command

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

# **SignalPeer**
> Signal SignalPeer(ctx, hostid)
send signal to peer.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hostid** | **string**| HostID | 

### Return type

[**Signal**](Signal.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Synchost**
> Network Synchost(ctx, hostid)
Requests a host to pull.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hostid** | **string**| HostID | 

### Return type

[**Network**](Network.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAllKeys**
> Network UpdateAllKeys(ctx, )
Update keys for a network.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Network**](Network.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateHost**
> ApiHost UpdateHost(ctx, hostid)
Updates a Netclient host on Netmaker server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hostid** | **string**| HostID | 

### Return type

[**ApiHost**](ApiHost.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateKeys**
> Network UpdateKeys(ctx, hostid)
Update keys for a network.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hostid** | **string**| HostID | 

### Return type

[**Network**](Network.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

