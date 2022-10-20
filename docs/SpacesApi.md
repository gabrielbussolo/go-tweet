# {{classname}}

All URIs are relative to *https://api.twitter.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindSpaceById**](SpacesApi.md#FindSpaceById) | **Get** /2/spaces/{id} | Space lookup by Space ID
[**FindSpacesByCreatorIds**](SpacesApi.md#FindSpacesByCreatorIds) | **Get** /2/spaces/by/creator_ids | Space lookup by their creators
[**FindSpacesByIds**](SpacesApi.md#FindSpacesByIds) | **Get** /2/spaces | Space lookup up Space IDs
[**SearchSpaces**](SpacesApi.md#SearchSpaces) | **Get** /2/spaces/search | Search for Spaces
[**SpaceBuyers**](SpacesApi.md#SpaceBuyers) | **Get** /2/spaces/{id}/buyers | Retrieve the list of Users who purchased a ticket to the given space
[**SpaceTweets**](SpacesApi.md#SpaceTweets) | **Get** /2/spaces/{id}/tweets | Retrieve Tweets from a Space.

# **FindSpaceById**
> Get2SpacesIdResponse FindSpaceById(ctx, id, optional)
Space lookup by Space ID

Returns a variety of information about the Space specified by the requested ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the Space to be retrieved. | 
 **optional** | ***SpacesApiFindSpaceByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpacesApiFindSpaceByIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **spaceFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Space fields to display. | 
 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 
 **topicFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Topic fields to display. | 

### Return type

[**Get2SpacesIdResponse**](Get2SpacesIdResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken), [OAuth2UserToken](../README.md#OAuth2UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindSpacesByCreatorIds**
> Get2SpacesByCreatorIdsResponse FindSpacesByCreatorIds(ctx, userIds, optional)
Space lookup by their creators

Returns a variety of information about the Spaces created by the provided User IDs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userIds** | [**[]string**](string.md)| The IDs of Users to search through. | 
 **optional** | ***SpacesApiFindSpacesByCreatorIdsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpacesApiFindSpacesByCreatorIdsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **spaceFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Space fields to display. | 
 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 
 **topicFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Topic fields to display. | 

### Return type

[**Get2SpacesByCreatorIdsResponse**](Get2SpacesByCreatorIdsResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken), [OAuth2UserToken](../README.md#OAuth2UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindSpacesByIds**
> Get2SpacesResponse FindSpacesByIds(ctx, ids, optional)
Space lookup up Space IDs

Returns a variety of information about the Spaces specified by the requested IDs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ids** | [**[]string**](string.md)| The list of Space IDs to return. | 
 **optional** | ***SpacesApiFindSpacesByIdsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpacesApiFindSpacesByIdsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **spaceFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Space fields to display. | 
 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 
 **topicFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Topic fields to display. | 

### Return type

[**Get2SpacesResponse**](Get2SpacesResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken), [OAuth2UserToken](../README.md#OAuth2UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchSpaces**
> Get2SpacesSearchResponse SearchSpaces(ctx, query, optional)
Search for Spaces

Returns Spaces that match the provided query.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **query** | **string**| The search query. | 
 **optional** | ***SpacesApiSearchSpacesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpacesApiSearchSpacesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **state** | **optional.String**| The state of Spaces to search for. | [default to all]
 **maxResults** | **optional.Int32**| The number of results to return. | [default to 100]
 **spaceFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Space fields to display. | 
 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 
 **topicFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Topic fields to display. | 

### Return type

[**Get2SpacesSearchResponse**](Get2SpacesSearchResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken), [OAuth2UserToken](../README.md#OAuth2UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SpaceBuyers**
> Get2SpacesIdBuyersResponse SpaceBuyers(ctx, id, optional)
Retrieve the list of Users who purchased a ticket to the given space

Retrieves the list of Users who purchased a ticket to the given space

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the Space to be retrieved. | 
 **optional** | ***SpacesApiSpaceBuyersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpacesApiSpaceBuyersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paginationToken** | [**optional.Interface of string**](.md)| This parameter is used to get a specified &#x27;page&#x27; of results. | 
 **maxResults** | **optional.Int32**| The maximum number of results. | [default to 100]
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 
 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **tweetFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Tweet fields to display. | 

### Return type

[**Get2SpacesIdBuyersResponse**](Get2SpacesIdBuyersResponse.md)

### Authorization

[OAuth2UserToken](../README.md#OAuth2UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SpaceTweets**
> Get2SpacesIdTweetsResponse SpaceTweets(ctx, id, optional)
Retrieve Tweets from a Space.

Retrieves Tweets shared in the specified Space.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the Space to be retrieved. | 
 **optional** | ***SpacesApiSpaceTweetsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpacesApiSpaceTweetsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maxResults** | **optional.Int32**| The number of Tweets to fetch from the provided space. If not provided, the value will default to the maximum of 100. | [default to 100]
 **tweetFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Tweet fields to display. | 
 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **mediaFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Media fields to display. | 
 **pollFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Poll fields to display. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 
 **placeFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Place fields to display. | 

### Return type

[**Get2SpacesIdTweetsResponse**](Get2SpacesIdTweetsResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken), [OAuth2UserToken](../README.md#OAuth2UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

