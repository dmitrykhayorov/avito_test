# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FlatCreatePost**](AuthOnlyApi.md#FlatCreatePost) | **Post** /flat/create | 
[**HouseIdGet**](AuthOnlyApi.md#HouseIdGet) | **Get** /house/{id} | 
[**HouseIdSubscribePost**](AuthOnlyApi.md#HouseIdSubscribePost) | **Post** /house/{id}/subscribe | 

# **FlatCreatePost**
> Flat FlatCreatePost(ctx, optional)


Создание квартиры. Квартира создается в статусе created

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AuthOnlyApiFlatCreatePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthOnlyApiFlatCreatePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of FlatCreateBody**](FlatCreateBody.md)|  | 

### Return type

[**Flat**](Flat.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HouseIdGet**
> InlineResponse2002 HouseIdGet(ctx, id)


Получение квартир в выбранном доме. Для обычных пользователей возвращаются только квартиры в статусе approved, для модераторов - в любом статусе

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**int32**](.md)|  | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HouseIdSubscribePost**
> HouseIdSubscribePost(ctx, id, optional)


Дополнительное задание. Подписаться на уведомления о новых квартирах в доме.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**int32**](.md)|  | 
 **optional** | ***AuthOnlyApiHouseIdSubscribePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthOnlyApiHouseIdSubscribePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of IdSubscribeBody**](IdSubscribeBody.md)|  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

