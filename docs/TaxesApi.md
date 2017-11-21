# \TaxesApi

All URIs are relative to *https://api.facestore.local/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTaxes**](TaxesApi.md#AddTaxes) | **Post** /taxes | 
[**DeleteTaxById**](TaxesApi.md#DeleteTaxById) | **Delete** /taxes/{id}/ | 
[**GetTaxById**](TaxesApi.md#GetTaxById) | **Get** /taxes/{id}/ | 
[**GetTaxes**](TaxesApi.md#GetTaxes) | **Get** /taxes | 
[**UpdateTaxById**](TaxesApi.md#UpdateTaxById) | **Put** /taxes/{id}/ | 
[**UpdateTaxById_0**](TaxesApi.md#UpdateTaxById_0) | **Patch** /taxes/{id}/ | 


# **AddTaxes**
> InlineResponse2012 AddTaxes(ctx, tax)


Creates a new tax in the store.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **tax** | [**Tax**](Tax.md)| Tax to add to the store | 

### Return type

[**InlineResponse2012**](inline_response_201_2.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTaxById**
> DeleteTaxById(ctx, id)


deletes a single tax based on the ID supplied

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **int64**| ID of tax to delete | 

### Return type

 (empty response body)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTaxById**
> InlineResponse2012 GetTaxById(ctx, id, optional)


Returns a tax based on a single ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **int64**| ID of tax to fetch | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| ID of tax to fetch | 
 **limit** | **int32**| max records to return | 

### Return type

[**InlineResponse2012**](inline_response_201_2.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTaxes**
> InlineResponse2002 GetTaxes(ctx, optional)


Returns all taxes from the system that the user has access to

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

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTaxById**
> InlineResponse2012 UpdateTaxById(ctx, id, tax)


update a single tax based on the ID supplied

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **int64**| ID of tax to update | 
  **tax** | [**Tax**](Tax.md)| Tax to add to the store | 

### Return type

[**InlineResponse2012**](inline_response_201_2.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTaxById_0**
> InlineResponse2012 UpdateTaxById_0(ctx, id, tax)


update a single tax based on the ID supplied

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **int64**| ID of tax to update | 
  **tax** | [**Tax**](Tax.md)| Tax to add to the store | 

### Return type

[**InlineResponse2012**](inline_response_201_2.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

