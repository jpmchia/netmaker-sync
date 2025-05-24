# \EnrollmentKeysApi

All URIs are relative to *https://api.demo.netmaker.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEnrollmentKey**](EnrollmentKeysApi.md#CreateEnrollmentKey) | **Post** /api/v1/enrollment-keys | Creates an EnrollmentKey for hosts to use on Netmaker server.
[**DeleteEnrollmentKey**](EnrollmentKeysApi.md#DeleteEnrollmentKey) | **Delete** /api/v1/enrollment-keys/{keyid} | Deletes an EnrollmentKey from Netmaker server.
[**GetEnrollmentKeys**](EnrollmentKeysApi.md#GetEnrollmentKeys) | **Get** /api/v1/enrollment-keys | Lists all EnrollmentKeys for admins.
[**HandleHostRegister**](EnrollmentKeysApi.md#HandleHostRegister) | **Post** /api/v1/enrollment-keys/{token} | Handles a Netclient registration with server and add nodes accordingly.
[**UpdateEnrollmentKey**](EnrollmentKeysApi.md#UpdateEnrollmentKey) | **Put** /api/v1/enrollment-keys/{keyid} | Updates an EnrollmentKey for hosts to use on Netmaker server. Updates only the relay to use.


# **CreateEnrollmentKey**
> EnrollmentKey CreateEnrollmentKey(ctx, optional)
Creates an EnrollmentKey for hosts to use on Netmaker server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EnrollmentKeysApiCreateEnrollmentKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EnrollmentKeysApiCreateEnrollmentKeyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ApiEnrollmentKey**](ApiEnrollmentKey.md)| APIEnrollmentKey | 

### Return type

[**EnrollmentKey**](EnrollmentKey.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEnrollmentKey**
> DeleteEnrollmentKey(ctx, keyid)
Deletes an EnrollmentKey from Netmaker server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **keyid** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEnrollmentKeys**
> []EnrollmentKey GetEnrollmentKeys(ctx, )
Lists all EnrollmentKeys for admins.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]EnrollmentKey**](EnrollmentKey.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HandleHostRegister**
> RegisterResponse HandleHostRegister(ctx, token, optional)
Handles a Netclient registration with server and add nodes accordingly.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**|  | 
 **optional** | ***EnrollmentKeysApiHandleHostRegisterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EnrollmentKeysApiHandleHostRegisterOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **host** | [**optional.Interface of Host**](Host.md)|  | 

### Return type

[**RegisterResponse**](RegisterResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEnrollmentKey**
> EnrollmentKey UpdateEnrollmentKey(ctx, keyid, optional)
Updates an EnrollmentKey for hosts to use on Netmaker server. Updates only the relay to use.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **keyid** | **string**| KeyID | 
 **optional** | ***EnrollmentKeysApiUpdateEnrollmentKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EnrollmentKeysApiUpdateEnrollmentKeyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ApiEnrollmentKey**](ApiEnrollmentKey.md)| APIEnrollmentKey | 

### Return type

[**EnrollmentKey**](EnrollmentKey.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

