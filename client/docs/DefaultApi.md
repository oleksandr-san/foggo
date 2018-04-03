# \DefaultApi

All URIs are relative to *http://127.0.0.1:3001*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HelloPost**](DefaultApi.md#HelloPost) | **Post** /hello | 
[**ListGet**](DefaultApi.md#ListGet) | **Get** /list | 


# **HelloPost**
> bool HelloPost(ctx, id, temperature)


Interface establishment 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**| Student Id | 
  **temperature** | **float32**| Temperature | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListGet**
> []Data ListGet(ctx, )


5 min list 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Data**](Data.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

