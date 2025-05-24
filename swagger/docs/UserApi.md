# \UserApi

All URIs are relative to *https://api.demo.netmaker.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApprovePendingUser**](UserApi.md#ApprovePendingUser) | **Post** /api/users_pending/user/{username} | approve pending user.
[**AttachUserToRemoteAccessGateway**](UserApi.md#AttachUserToRemoteAccessGateway) | **Post** /api/users/{username}/remote_access_gw | Attach User to a remote access gateway.
[**CreateAdmin**](UserApi.md#CreateAdmin) | **Post** /api/users/adm/createsuperadmin | Make a user an admin.
[**CreateUser**](UserApi.md#CreateUser) | **Post** /api/users/{username} | Create a user.
[**DeleteAllPendingUsers**](UserApi.md#DeleteAllPendingUsers) | **Delete** /api/users_pending/{username}/pending | delete all pending users.
[**DeletePendingUser**](UserApi.md#DeletePendingUser) | **Delete** /api/users_pending/user/{username} | delete pending user.
[**DeleteUser**](UserApi.md#DeleteUser) | **Delete** /api/users/{username} | Delete a user.
[**GetPendingUsers**](UserApi.md#GetPendingUsers) | **Get** /api/users_pending | Get all pending users.
[**GetUser**](UserApi.md#GetUser) | **Get** /api/users/{username} | Get an individual user.
[**GetUsers**](UserApi.md#GetUsers) | **Get** /api/users | Get all users.
[**HasSuperAdmin**](UserApi.md#HasSuperAdmin) | **Get** /api/users/adm/hassuperadmin | Checks whether the server has an admin.
[**RemoveUserFromRemoteAccessGW**](UserApi.md#RemoveUserFromRemoteAccessGW) | **Delete** /api/users/{username}/remote_access_gw | Delete User from a remote access gateway.
[**TransferSuperAdmin**](UserApi.md#TransferSuperAdmin) | **Post** /api/users/adm/transfersuperadmin | Transfers superadmin role to an admin user.
[**UpdateUser**](UserApi.md#UpdateUser) | **Put** /api/users/{username} | Update a user.


# **ApprovePendingUser**
> User ApprovePendingUser(ctx, )
approve pending user.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AttachUserToRemoteAccessGateway**
> User AttachUserToRemoteAccessGateway(ctx, username)
Attach User to a remote access gateway.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**|  | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateAdmin**
> User CreateAdmin(ctx, optional)
Make a user an admin.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserApiCreateAdminOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiCreateAdminOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | [**optional.Interface of User**](User.md)| User | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateUser**
> User CreateUser(ctx, username, optional)
Create a user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| Username | 
 **optional** | ***UserApiCreateUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiCreateUserOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **user** | [**optional.Interface of User**](User.md)| User | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAllPendingUsers**
> User DeleteAllPendingUsers(ctx, )
delete all pending users.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePendingUser**
> User DeletePendingUser(ctx, )
delete pending user.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUser**
> User DeleteUser(ctx, username)
Delete a user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| Username | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPendingUsers**
> User GetPendingUsers(ctx, )
Get all pending users.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUser**
> User GetUser(ctx, username)
Get an individual user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| Username | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsers**
> User GetUsers(ctx, )
Get all users.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HasSuperAdmin**
> HasSuperAdmin(ctx, )
Checks whether the server has an admin.

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

# **RemoveUserFromRemoteAccessGW**
> User RemoveUserFromRemoteAccessGW(ctx, username)
Delete User from a remote access gateway.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**|  | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TransferSuperAdmin**
> User TransferSuperAdmin(ctx, )
Transfers superadmin role to an admin user.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUser**
> User UpdateUser(ctx, username, optional)
Update a user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| Username | 
 **optional** | ***UserApiUpdateUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiUpdateUserOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **user** | [**optional.Interface of User**](User.md)| User | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

