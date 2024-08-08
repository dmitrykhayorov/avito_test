# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FlatUpdatePost**](ModerationsOnlyApi.md#FlatUpdatePost) | **Post** /flat/update | 
[**HouseCreatePost**](ModerationsOnlyApi.md#HouseCreatePost) | **Post** /house/create | 

# **FlatUpdatePost**
> Flat FlatUpdatePost(ctx, optional)


Обновление квартиры.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ModerationsOnlyApiFlatUpdatePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ModerationsOnlyApiFlatUpdatePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of FlatUpdateBody**](FlatUpdateBody.md)|  | 

### Return type

[**Flat**](Flat.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HouseCreatePost**
> House HouseCreatePost(ctx, optional)


Создание нового дома.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ModerationsOnlyApiHouseCreatePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ModerationsOnlyApiHouseCreatePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of HouseCreateBody**](HouseCreateBody.md)|  | 

### Return type

[**House**](House.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

