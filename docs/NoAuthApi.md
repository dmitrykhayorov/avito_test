# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DummyLoginGet**](NoAuthApi.md#DummyLoginGet) | **Get** /dummyLogin | 
[**LoginPost**](NoAuthApi.md#LoginPost) | **Post** /login | 
[**RegisterPost**](NoAuthApi.md#RegisterPost) | **Post** /register | 

# **DummyLoginGet**
> InlineResponse200 DummyLoginGet(ctx, userType)


Упрощенный процесс получения токена для дальнейшего прохождения авторизации

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userType** | [**UserType**](.md)|  | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoginPost**
> InlineResponse200 LoginPost(ctx, optional)


Дополнительное задание. Процесс аутентификации путем передачи идентификатор+пароля пользователя и получения токена для дальнейшего прохождения авторизации

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NoAuthApiLoginPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NoAuthApiLoginPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of LoginBody**](LoginBody.md)|  | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegisterPost**
> InlineResponse2001 RegisterPost(ctx, optional)


Дополнительное задание. Регистрация нового пользователя

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NoAuthApiRegisterPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NoAuthApiRegisterPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of RegisterBody**](RegisterBody.md)|  | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

