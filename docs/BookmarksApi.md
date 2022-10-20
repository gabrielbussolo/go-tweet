# {{classname}}

All URIs are relative to *https://api.twitter.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUsersIdBookmarks**](BookmarksApi.md#GetUsersIdBookmarks) | **Get** /2/users/{id}/bookmarks | Bookmarks by User
[**PostUsersIdBookmarks**](BookmarksApi.md#PostUsersIdBookmarks) | **Post** /2/users/{id}/bookmarks | Add Tweet to Bookmarks
[**UsersIdBookmarksDelete**](BookmarksApi.md#UsersIdBookmarksDelete) | **Delete** /2/users/{id}/bookmarks/{tweet_id} | Remove a bookmarked Tweet

# **GetUsersIdBookmarks**
> Get2UsersIdBookmarksResponse GetUsersIdBookmarks(ctx, id, optional)
Bookmarks by User

Returns Tweet objects that have been bookmarked by the requesting User

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the authenticated source User for whom to return results. | 
 **optional** | ***BookmarksApiGetUsersIdBookmarksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BookmarksApiGetUsersIdBookmarksOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maxResults** | **optional.Int32**| The maximum number of results. | 
 **paginationToken** | [**optional.Interface of string**](.md)| This parameter is used to get the next &#x27;page&#x27; of results. | 
 **tweetFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Tweet fields to display. | 
 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **mediaFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Media fields to display. | 
 **pollFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Poll fields to display. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 
 **placeFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Place fields to display. | 

### Return type

[**Get2UsersIdBookmarksResponse**](Get2UsersIdBookmarksResponse.md)

### Authorization

[OAuth2UserToken](../README.md#OAuth2UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUsersIdBookmarks**
> BookmarkMutationResponse PostUsersIdBookmarks(ctx, body, id)
Add Tweet to Bookmarks

Adds a Tweet (ID in the body) to the requesting User's (in the path) bookmarks

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BookmarkAddRequest**](BookmarkAddRequest.md)|  | 
  **id** | [**string**](.md)| The ID of the authenticated source User for whom to add bookmarks. | 

### Return type

[**BookmarkMutationResponse**](BookmarkMutationResponse.md)

### Authorization

[OAuth2UserToken](../README.md#OAuth2UserToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersIdBookmarksDelete**
> BookmarkMutationResponse UsersIdBookmarksDelete(ctx, id, tweetId)
Remove a bookmarked Tweet

Removes a Tweet from the requesting User's bookmarked Tweets.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the authenticated source User whose bookmark is to be removed. | 
  **tweetId** | [**string**](.md)| The ID of the Tweet that the source User is removing from bookmarks. | 

### Return type

[**BookmarkMutationResponse**](BookmarkMutationResponse.md)

### Authorization

[OAuth2UserToken](../README.md#OAuth2UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

