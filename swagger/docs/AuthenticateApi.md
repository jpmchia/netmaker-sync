# \AuthenticateApi

All URIs are relative to *https://api.demo.netmaker.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Authenticate**](AuthenticateApi.md#Authenticate) | **Post** /api/nodes/adm/{network}/authenticate | Authenticate to make further API calls related to a network.
[**AuthenticateHost**](AuthenticateApi.md#AuthenticateHost) | **Post** /api/hosts/adm/authenticate | Host based authentication for making further API calls.
[**AuthenticateUser**](AuthenticateApi.md#AuthenticateUser) | **Post** /api/users/adm/authenticate | User authenticates using its password and retrieves a JWT for authorization.


# **Authenticate**
> SuccessResponse Authenticate(ctx, network, optional)
Authenticate to make further API calls related to a network.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **network** | **string**| network | 
 **optional** | ***AuthenticateApiAuthenticateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthenticateApiAuthenticateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authParams** | [**optional.Interface of AuthParams**](AuthParams.md)| AuthParams | 

### Return type

[**SuccessResponse**](SuccessResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthenticateHost**
> SuccessResponse AuthenticateHost(ctx, )
Host based authentication for making further API calls.

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

# **AuthenticateUser**
> SuccessResponse AuthenticateUser(ctx, optional)
User authenticates using its password and retrieves a JWT for authorization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AuthenticateApiAuthenticateUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthenticateApiAuthenticateUserOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userAuthParams** | [**optional.Interface of UserAuthParams**](UserAuthParams.md)| User Auth Params | 

### Return type

[**SuccessResponse**](SuccessResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

