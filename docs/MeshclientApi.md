# \MeshclientApi

All URIs are relative to *https://api.demo.netmaker.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFile**](MeshclientApi.md#GetFile) | **Get** /meshclient/files/{filename} | Retrieve a file from the file server.


# **GetFile**
> *os.File GetFile(ctx, filename)
Retrieve a file from the file server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **filename** | **string**| Filename | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

