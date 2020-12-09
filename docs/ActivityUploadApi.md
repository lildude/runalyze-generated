# {{classname}}

All URIs are relative to *https://runalyze.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetActivityByID**](ActivityUploadApi.md#GetActivityByID) | **Get** /api/v1/activities/uploads/{id} | 
[**UploadActivity**](ActivityUploadApi.md#UploadActivity) | **Post** /api/v1/activities/uploads | 

# **GetActivityByID**
> GetActivityByID(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadActivity**
> UploadActivity(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ActivityUploadApiUploadActivityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ActivityUploadApiUploadActivityOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **title** | **optional.**|  | 
 **note** | **optional.**|  | 
 **route** | **optional.**|  | 
 **file** | [**optional.Interface of []*os.File**](*os.File.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

