# \PaymentsApi

All URIs are relative to *https://api.facestore.local/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPaymentById**](PaymentsApi.md#GetPaymentById) | **Get** /payments/{id}/ | 
[**GetPayments**](PaymentsApi.md#GetPayments) | **Get** /payments | 


# **GetPaymentById**
> InlineResponse2005 GetPaymentById(ctx, id)


Returns all payments from the system that the user has access to 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **int64**| ID of payment | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPayments**
> InlineResponse2004 GetPayments(ctx, optional)


Returns all payments from the system that the user has access to 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includes** | [**[]string**](string.md)| Include associated objects within response | 
 **limit** | **int32**| max records to return | 
 **orderBy** | [**[]string**](string.md)| Specify the field to be sorted, examples:  - &#x60;?order_by&#x3D;id|desc&#x60; - &#x60;?order_by&#x3D;updated_at|desc,position|asc&#x60;  | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

