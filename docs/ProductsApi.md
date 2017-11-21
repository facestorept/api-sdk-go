# \ProductsApi

All URIs are relative to *https://api.facestore.local/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddProduct**](ProductsApi.md#AddProduct) | **Post** /products | 
[**DeleteProductById**](ProductsApi.md#DeleteProductById) | **Delete** /products/{id}/ | 
[**GetProductById**](ProductsApi.md#GetProductById) | **Get** /products/{id}/ | 
[**GetProducts**](ProductsApi.md#GetProducts) | **Get** /products | 
[**UpdateProductById**](ProductsApi.md#UpdateProductById) | **Put** /products/{id}/ | 
[**UpdateProductById_0**](ProductsApi.md#UpdateProductById_0) | **Patch** /products/{id}/ | 


# **AddProduct**
> InlineResponse2014 AddProduct(ctx, product)


Creates a new product in the store.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **product** | [**Product**](Product.md)| Product to add to the store | 

### Return type

[**InlineResponse2014**](inline_response_201_4.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProductById**
> DeleteProductById(ctx, id)


deletes a single product based on the ID supplied

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **int64**| ID of product to delete | 

### Return type

 (empty response body)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProductById**
> InlineResponse2014 GetProductById(ctx, id, optional)


Returns a product based on a single ID  ### Includes You can give the following values on includes parameter: `brands, categories, routes, stocks` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **int64**| ID of product to fetch | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| ID of product to fetch | 
 **includes** | [**[]string**](string.md)| Include associated objects within response | 
 **limit** | **int32**| max records to return | 

### Return type

[**InlineResponse2014**](inline_response_201_4.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProducts**
> InlineResponse2006 GetProducts(ctx, optional)


Returns all products from the system that the user has access to  ### Includes You can give the following values on includes parameter: `brands, categories, routes, stocks` 

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

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateProductById**
> UpdateProductById(ctx, id, tax)


update a single product based on the ID supplied

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **int64**| ID of product to update | 
  **tax** | [**Product**](Product.md)| Product to add to the store | 

### Return type

 (empty response body)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateProductById_0**
> UpdateProductById_0(ctx, id, tax)


update a single product based on the ID supplied

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **int64**| ID of product to update | 
  **tax** | [**Product**](Product.md)| Product to add to the store | 

### Return type

 (empty response body)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

