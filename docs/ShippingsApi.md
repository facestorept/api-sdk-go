# \ShippingsApi

All URIs are relative to *https://api.facestore.local/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddShipping**](ShippingsApi.md#AddShipping) | **Post** /shippings | 
[**DeleteShippingById**](ShippingsApi.md#DeleteShippingById) | **Delete** /shippings/{id}/ | 
[**GetShippingById**](ShippingsApi.md#GetShippingById) | **Get** /shippings/{id}/ | 
[**GetShippings**](ShippingsApi.md#GetShippings) | **Get** /shippings | 
[**UpdateShippingById**](ShippingsApi.md#UpdateShippingById) | **Put** /shippings/{id}/ | 
[**UpdateShippingById_0**](ShippingsApi.md#UpdateShippingById_0) | **Patch** /shippings/{id}/ | 


# **AddShipping**
> InlineResponse2013 AddShipping(ctx, shipping)


Creates a new shipping in the store.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **shipping** | [**Shipping**](Shipping.md)| Shipping to add to the store | 

### Return type

[**InlineResponse2013**](inline_response_201_3.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteShippingById**
> DeleteShippingById(ctx, id)


deletes a single shipping based on the ID supplied

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **int64**| ID of shipping to delete | 

### Return type

 (empty response body)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetShippingById**
> InlineResponse2013 GetShippingById(ctx, id, optional)


Returns a shipping based on a single ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **int64**| ID of shipping to fetch | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| ID of shipping to fetch | 
 **limit** | **int32**| max records to return | 

### Return type

[**InlineResponse2013**](inline_response_201_3.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetShippings**
> InlineResponse2003 GetShippings(ctx, optional)


Returns all shippings from the system that the user has access to

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

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateShippingById**
> InlineResponse2013 UpdateShippingById(ctx, id, tax)


update a single shipping based on the ID supplied

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **int64**| ID of shipping to update | 
  **tax** | [**Shipping**](Shipping.md)| Shipping to update in store | 

### Return type

[**InlineResponse2013**](inline_response_201_3.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateShippingById_0**
> InlineResponse2013 UpdateShippingById_0(ctx, id, tax)


update a single shipping based on the ID supplied

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **int64**| ID of shipping to update | 
  **tax** | [**Shipping**](Shipping.md)| Shipping to update in store | 

### Return type

[**InlineResponse2013**](inline_response_201_3.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

