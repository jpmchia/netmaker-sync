# \DnsApi

All URIs are relative to *https://api.demo.netmaker.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDNS**](DnsApi.md#CreateDNS) | **Post** /api/dns/{network} | Create a DNS entry.
[**DeleteDNS**](DnsApi.md#DeleteDNS) | **Delete** /api/dns/{network}/{domain} | Delete a DNS entry.
[**GetAllDNS**](DnsApi.md#GetAllDNS) | **Get** /api/dns | Gets all DNS entries.
[**GetCustomDNS**](DnsApi.md#GetCustomDNS) | **Get** /api/dns/adm/{network}/custom | Gets custom DNS entries associated with a network.
[**GetDNS**](DnsApi.md#GetDNS) | **Get** /api/dns/adm/{network} | Gets all DNS entries associated with the network.
[**GetNodeDNS**](DnsApi.md#GetNodeDNS) | **Get** /api/dns/adm/{network}/nodes | Gets node DNS entries associated with a network.
[**PushDNS**](DnsApi.md#PushDNS) | **Post** /api/dns/adm/pushdns | Push DNS entries to nameserver.


# **CreateDNS**
> []DnsEntry CreateDNS(ctx, network, optional)
Create a DNS entry.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **network** | **string**| Network | 
 **optional** | ***DnsApiCreateDNSOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DnsApiCreateDNSOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of []DnsEntry**](DNSEntry.md)| DNS Entry | 

### Return type

[**[]DnsEntry**](DNSEntry.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDNS**
> DeleteDNS(ctx, network, domain)
Delete a DNS entry.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **network** | **string**| Network | 
  **domain** | **string**| Domain | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllDNS**
> []DnsEntry GetAllDNS(ctx, )
Gets all DNS entries.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]DnsEntry**](DNSEntry.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomDNS**
> []DnsEntry GetCustomDNS(ctx, network)
Gets custom DNS entries associated with a network.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **network** | **string**| Network | 

### Return type

[**[]DnsEntry**](DNSEntry.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDNS**
> []DnsEntry GetDNS(ctx, network)
Gets all DNS entries associated with the network.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **network** | **string**| Network | 

### Return type

[**[]DnsEntry**](DNSEntry.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNodeDNS**
> []DnsEntry GetNodeDNS(ctx, network)
Gets node DNS entries associated with a network.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **network** | **string**| Network | 

### Return type

[**[]DnsEntry**](DNSEntry.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PushDNS**
> []DnsEntry PushDNS(ctx, )
Push DNS entries to nameserver.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]DnsEntry**](DNSEntry.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

