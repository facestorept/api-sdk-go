# \CustomersApi

All URIs are relative to *https://api.facestore.local/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCustomerById**](CustomersApi.md#GetCustomerById) | **Get** /customers/{id}/ | 
[**GetCustomers**](CustomersApi.md#GetCustomers) | **Get** /customers | 


# **GetCustomerById**
> []Customer GetCustomerById(ctx, id, optional)


Returns all customers from the system that the user has access to  ### Includes You can give the following values on includes parameter: `orders, groups` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **int64**| ID of customer | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| ID of customer | 
 **includes** | [**[]string**](string.md)| Include associated objects within response | 

### Return type

[**[]Customer**](Customer.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomers**
> []Customer GetCustomers(ctx, optional)


Returns all customers from the system that the user has access to  ### Includes You can give the following values on includes parameter: `orders, groups` 

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

[**[]Customer**](Customer.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

