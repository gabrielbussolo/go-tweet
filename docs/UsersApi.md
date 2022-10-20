# {{classname}}

All URIs are relative to *https://api.twitter.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindMyUser**](UsersApi.md#FindMyUser) | **Get** /2/users/me | User lookup me
[**FindUserById**](UsersApi.md#FindUserById) | **Get** /2/users/{id} | User lookup by ID
[**FindUserByUsername**](UsersApi.md#FindUserByUsername) | **Get** /2/users/by/username/{username} | User lookup by username
[**FindUsersById**](UsersApi.md#FindUsersById) | **Get** /2/users | User lookup by IDs
[**FindUsersByUsername**](UsersApi.md#FindUsersByUsername) | **Get** /2/users/by | User lookup by usernames
[**ListGetFollowers**](UsersApi.md#ListGetFollowers) | **Get** /2/lists/{id}/followers | Returns User objects that follow a List by the provided List ID
[**ListGetMembers**](UsersApi.md#ListGetMembers) | **Get** /2/lists/{id}/members | Returns User objects that are members of a List by the provided List ID.
[**TweetsIdLikingUsers**](UsersApi.md#TweetsIdLikingUsers) | **Get** /2/tweets/{id}/liking_users | Returns User objects that have liked the provided Tweet ID
[**TweetsIdRetweetingUsers**](UsersApi.md#TweetsIdRetweetingUsers) | **Get** /2/tweets/{id}/retweeted_by | Returns User objects that have retweeted the provided Tweet ID
[**UsersIdBlock**](UsersApi.md#UsersIdBlock) | **Post** /2/users/{id}/blocking | Block User by User ID
[**UsersIdBlocking**](UsersApi.md#UsersIdBlocking) | **Get** /2/users/{id}/blocking | Returns User objects that are blocked by provided User ID
[**UsersIdFollow**](UsersApi.md#UsersIdFollow) | **Post** /2/users/{id}/following | Follow User
[**UsersIdFollowers**](UsersApi.md#UsersIdFollowers) | **Get** /2/users/{id}/followers | Followers by User ID
[**UsersIdFollowing**](UsersApi.md#UsersIdFollowing) | **Get** /2/users/{id}/following | Following by User ID
[**UsersIdMute**](UsersApi.md#UsersIdMute) | **Post** /2/users/{id}/muting | Mute User by User ID.
[**UsersIdMuting**](UsersApi.md#UsersIdMuting) | **Get** /2/users/{id}/muting | Returns User objects that are muted by the provided User ID
[**UsersIdUnblock**](UsersApi.md#UsersIdUnblock) | **Delete** /2/users/{source_user_id}/blocking/{target_user_id} | Unblock User by User ID
[**UsersIdUnfollow**](UsersApi.md#UsersIdUnfollow) | **Delete** /2/users/{source_user_id}/following/{target_user_id} | Unfollow User
[**UsersIdUnmute**](UsersApi.md#UsersIdUnmute) | **Delete** /2/users/{source_user_id}/muting/{target_user_id} | Unmute User by User ID

# **FindMyUser**
> Get2UsersMeResponse FindMyUser(ctx, optional)
User lookup me

This endpoint returns information about the requesting User.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersApiFindMyUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiFindMyUserOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 
 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **tweetFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Tweet fields to display. | 

### Return type

[**Get2UsersMeResponse**](Get2UsersMeResponse.md)

### Authorization

[OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindUserById**
> Get2UsersIdResponse FindUserById(ctx, id, optional)
User lookup by ID

This endpoint returns information about a User. Specify User by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the User to lookup. | 
 **optional** | ***UsersApiFindUserByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiFindUserByIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 
 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **tweetFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Tweet fields to display. | 

### Return type

[**Get2UsersIdResponse**](Get2UsersIdResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken), [OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindUserByUsername**
> Get2UsersByUsernameUsernameResponse FindUserByUsername(ctx, username, optional)
User lookup by username

This endpoint returns information about a User. Specify User by username.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| A username. | 
 **optional** | ***UsersApiFindUserByUsernameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiFindUserByUsernameOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 
 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **tweetFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Tweet fields to display. | 

### Return type

[**Get2UsersByUsernameUsernameResponse**](Get2UsersByUsernameUsernameResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken), [OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindUsersById**
> Get2UsersResponse FindUsersById(ctx, ids, optional)
User lookup by IDs

This endpoint returns information about Users. Specify Users by their ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ids** | [**[]string**](string.md)| A list of User IDs, comma-separated. You can specify up to 100 IDs. | 
 **optional** | ***UsersApiFindUsersByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiFindUsersByIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 
 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **tweetFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Tweet fields to display. | 

### Return type

[**Get2UsersResponse**](Get2UsersResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken), [OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindUsersByUsername**
> Get2UsersByResponse FindUsersByUsername(ctx, usernames, optional)
User lookup by usernames

This endpoint returns information about Users. Specify Users by their username.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **usernames** | [**[]string**](string.md)| A list of usernames, comma-separated. | 
 **optional** | ***UsersApiFindUsersByUsernameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiFindUsersByUsernameOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 
 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **tweetFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Tweet fields to display. | 

### Return type

[**Get2UsersByResponse**](Get2UsersByResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken), [OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListGetFollowers**
> Get2ListsIdFollowersResponse ListGetFollowers(ctx, id, optional)
Returns User objects that follow a List by the provided List ID

Returns a list of Users that follow a List by the provided List ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the List. | 
 **optional** | ***UsersApiListGetFollowersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiListGetFollowersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maxResults** | **optional.Int32**| The maximum number of results. | [default to 100]
 **paginationToken** | [**optional.Interface of string**](.md)| This parameter is used to get a specified &#x27;page&#x27; of results. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 
 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **tweetFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Tweet fields to display. | 

### Return type

[**Get2ListsIdFollowersResponse**](Get2ListsIdFollowersResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken), [OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListGetMembers**
> Get2ListsIdMembersResponse ListGetMembers(ctx, id, optional)
Returns User objects that are members of a List by the provided List ID.

Returns a list of Users that are members of a List by the provided List ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the List. | 
 **optional** | ***UsersApiListGetMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiListGetMembersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maxResults** | **optional.Int32**| The maximum number of results. | [default to 100]
 **paginationToken** | [**optional.Interface of string**](.md)| This parameter is used to get a specified &#x27;page&#x27; of results. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 
 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **tweetFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Tweet fields to display. | 

### Return type

[**Get2ListsIdMembersResponse**](Get2ListsIdMembersResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken), [OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TweetsIdLikingUsers**
> Get2TweetsIdLikingUsersResponse TweetsIdLikingUsers(ctx, id, optional)
Returns User objects that have liked the provided Tweet ID

Returns a list of Users that have liked the provided Tweet ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A single Tweet ID. | 
 **optional** | ***UsersApiTweetsIdLikingUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiTweetsIdLikingUsersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maxResults** | **optional.Int32**| The maximum number of results. | [default to 100]
 **paginationToken** | [**optional.Interface of string**](.md)| This parameter is used to get the next &#x27;page&#x27; of results. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 
 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **tweetFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Tweet fields to display. | 

### Return type

[**Get2TweetsIdLikingUsersResponse**](Get2TweetsIdLikingUsersResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken), [OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TweetsIdRetweetingUsers**
> Get2TweetsIdRetweetedByResponse TweetsIdRetweetingUsers(ctx, id, optional)
Returns User objects that have retweeted the provided Tweet ID

Returns a list of Users that have retweeted the provided Tweet ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A single Tweet ID. | 
 **optional** | ***UsersApiTweetsIdRetweetingUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiTweetsIdRetweetingUsersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maxResults** | **optional.Int32**| The maximum number of results. | [default to 100]
 **paginationToken** | [**optional.Interface of string**](.md)| This parameter is used to get the next &#x27;page&#x27; of results. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 
 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **tweetFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Tweet fields to display. | 

### Return type

[**Get2TweetsIdRetweetedByResponse**](Get2TweetsIdRetweetedByResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken), [OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersIdBlock**
> BlockUserMutationResponse UsersIdBlock(ctx, body, id)
Block User by User ID

Causes the User (in the path) to block the target User. The User (in the path) must match the User context authorizing the request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BlockUserRequest**](BlockUserRequest.md)|  | 
  **id** | [**string**](.md)| The ID of the authenticated source User that is requesting to block the target User. | 

### Return type

[**BlockUserMutationResponse**](BlockUserMutationResponse.md)

### Authorization

[OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersIdBlocking**
> Get2UsersIdBlockingResponse UsersIdBlocking(ctx, id, optional)
Returns User objects that are blocked by provided User ID

Returns a list of Users that are blocked by the provided User ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the authenticated source User for whom to return results. | 
 **optional** | ***UsersApiUsersIdBlockingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersIdBlockingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maxResults** | **optional.Int32**| The maximum number of results. | 
 **paginationToken** | [**optional.Interface of string**](.md)| This parameter is used to get a specified &#x27;page&#x27; of results. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 
 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **tweetFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Tweet fields to display. | 

### Return type

[**Get2UsersIdBlockingResponse**](Get2UsersIdBlockingResponse.md)

### Authorization

[OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersIdFollow**
> UsersFollowingCreateResponse UsersIdFollow(ctx, id, optional)
Follow User

Causes the User(in the path) to follow, or “request to follow” for protected Users, the target User. The User(in the path) must match the User context authorizing the request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the authenticated source User that is requesting to follow the target User. | 
 **optional** | ***UsersApiUsersIdFollowOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersIdFollowOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UsersFollowingCreateRequest**](UsersFollowingCreateRequest.md)|  | 

### Return type

[**UsersFollowingCreateResponse**](UsersFollowingCreateResponse.md)

### Authorization

[OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersIdFollowers**
> Get2UsersIdFollowersResponse UsersIdFollowers(ctx, id, optional)
Followers by User ID

Returns a list of Users who are followers of the specified User ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the User to lookup. | 
 **optional** | ***UsersApiUsersIdFollowersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersIdFollowersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maxResults** | **optional.Int32**| The maximum number of results. | 
 **paginationToken** | [**optional.Interface of string**](.md)| This parameter is used to get a specified &#x27;page&#x27; of results. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 
 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **tweetFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Tweet fields to display. | 

### Return type

[**Get2UsersIdFollowersResponse**](Get2UsersIdFollowersResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken), [OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersIdFollowing**
> Get2UsersIdFollowingResponse UsersIdFollowing(ctx, id, optional)
Following by User ID

Returns a list of Users that are being followed by the provided User ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the User to lookup. | 
 **optional** | ***UsersApiUsersIdFollowingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersIdFollowingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maxResults** | **optional.Int32**| The maximum number of results. | 
 **paginationToken** | [**optional.Interface of string**](.md)| This parameter is used to get a specified &#x27;page&#x27; of results. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 
 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **tweetFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Tweet fields to display. | 

### Return type

[**Get2UsersIdFollowingResponse**](Get2UsersIdFollowingResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken), [OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersIdMute**
> MuteUserMutationResponse UsersIdMute(ctx, id, optional)
Mute User by User ID.

Causes the User (in the path) to mute the target User. The User (in the path) must match the User context authorizing the request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the authenticated source User that is requesting to mute the target User. | 
 **optional** | ***UsersApiUsersIdMuteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersIdMuteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of MuteUserRequest**](MuteUserRequest.md)|  | 

### Return type

[**MuteUserMutationResponse**](MuteUserMutationResponse.md)

### Authorization

[OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersIdMuting**
> Get2UsersIdMutingResponse UsersIdMuting(ctx, id, optional)
Returns User objects that are muted by the provided User ID

Returns a list of Users that are muted by the provided User ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the authenticated source User for whom to return results. | 
 **optional** | ***UsersApiUsersIdMutingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersIdMutingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maxResults** | **optional.Int32**| The maximum number of results. | [default to 100]
 **paginationToken** | [**optional.Interface of string**](.md)| This parameter is used to get the next &#x27;page&#x27; of results. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 
 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **tweetFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Tweet fields to display. | 

### Return type

[**Get2UsersIdMutingResponse**](Get2UsersIdMutingResponse.md)

### Authorization

[OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersIdUnblock**
> BlockUserMutationResponse UsersIdUnblock(ctx, sourceUserId, targetUserId)
Unblock User by User ID

Causes the source User to unblock the target User. The source User must match the User context authorizing the request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sourceUserId** | [**string**](.md)| The ID of the authenticated source User that is requesting to unblock the target User. | 
  **targetUserId** | [**string**](.md)| The ID of the User that the source User is requesting to unblock. | 

### Return type

[**BlockUserMutationResponse**](BlockUserMutationResponse.md)

### Authorization

[OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersIdUnfollow**
> UsersFollowingDeleteResponse UsersIdUnfollow(ctx, sourceUserId, targetUserId)
Unfollow User

Causes the source User to unfollow the target User. The source User must match the User context authorizing the request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sourceUserId** | [**string**](.md)| The ID of the authenticated source User that is requesting to unfollow the target User. | 
  **targetUserId** | [**string**](.md)| The ID of the User that the source User is requesting to unfollow. | 

### Return type

[**UsersFollowingDeleteResponse**](UsersFollowingDeleteResponse.md)

### Authorization

[OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersIdUnmute**
> MuteUserMutationResponse UsersIdUnmute(ctx, sourceUserId, targetUserId)
Unmute User by User ID

Causes the source User to unmute the target User. The source User must match the User context authorizing the request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sourceUserId** | [**string**](.md)| The ID of the authenticated source User that is requesting to unmute the target User. | 
  **targetUserId** | [**string**](.md)| The ID of the User that the source User is requesting to unmute. | 

### Return type

[**MuteUserMutationResponse**](MuteUserMutationResponse.md)

### Authorization

[OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

