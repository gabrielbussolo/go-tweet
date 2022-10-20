# {{classname}}

All URIs are relative to *https://api.twitter.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOrDeleteRules**](TweetsApi.md#AddOrDeleteRules) | **Post** /2/tweets/search/stream/rules | Add/Delete rules
[**CreateTweet**](TweetsApi.md#CreateTweet) | **Post** /2/tweets | Creation of a Tweet
[**DeleteTweetById**](TweetsApi.md#DeleteTweetById) | **Delete** /2/tweets/{id} | Tweet delete by Tweet ID
[**FindTweetById**](TweetsApi.md#FindTweetById) | **Get** /2/tweets/{id} | Tweet lookup by Tweet ID
[**FindTweetsById**](TweetsApi.md#FindTweetsById) | **Get** /2/tweets | Tweet lookup by Tweet IDs
[**FindTweetsThatQuoteATweet**](TweetsApi.md#FindTweetsThatQuoteATweet) | **Get** /2/tweets/{id}/quote_tweets | Retrieve Tweets that quote a Tweet.
[**GetRules**](TweetsApi.md#GetRules) | **Get** /2/tweets/search/stream/rules | Rules lookup
[**GetTweetsFirehoseStream**](TweetsApi.md#GetTweetsFirehoseStream) | **Get** /2/tweets/firehose/stream | Firehose stream
[**GetTweetsSample10Stream**](TweetsApi.md#GetTweetsSample10Stream) | **Get** /2/tweets/sample10/stream | Sample 10% stream
[**HideReplyById**](TweetsApi.md#HideReplyById) | **Put** /2/tweets/{tweet_id}/hidden | Hide replies
[**ListsIdTweets**](TweetsApi.md#ListsIdTweets) | **Get** /2/lists/{id}/tweets | List Tweets timeline by List ID.
[**SampleStream**](TweetsApi.md#SampleStream) | **Get** /2/tweets/sample/stream | Sample stream
[**SearchStream**](TweetsApi.md#SearchStream) | **Get** /2/tweets/search/stream | Filtered stream
[**SpaceBuyers**](TweetsApi.md#SpaceBuyers) | **Get** /2/spaces/{id}/buyers | Retrieve the list of Users who purchased a ticket to the given space
[**SpaceTweets**](TweetsApi.md#SpaceTweets) | **Get** /2/spaces/{id}/tweets | Retrieve Tweets from a Space.
[**TweetCountsFullArchiveSearch**](TweetsApi.md#TweetCountsFullArchiveSearch) | **Get** /2/tweets/counts/all | Full archive search counts
[**TweetCountsRecentSearch**](TweetsApi.md#TweetCountsRecentSearch) | **Get** /2/tweets/counts/recent | Recent search counts
[**TweetsFullarchiveSearch**](TweetsApi.md#TweetsFullarchiveSearch) | **Get** /2/tweets/search/all | Full-archive search
[**TweetsRecentSearch**](TweetsApi.md#TweetsRecentSearch) | **Get** /2/tweets/search/recent | Recent search
[**UsersIdLike**](TweetsApi.md#UsersIdLike) | **Post** /2/users/{id}/likes | Causes the User (in the path) to like the specified Tweet
[**UsersIdLikedTweets**](TweetsApi.md#UsersIdLikedTweets) | **Get** /2/users/{id}/liked_tweets | Returns Tweet objects liked by the provided User ID
[**UsersIdMentions**](TweetsApi.md#UsersIdMentions) | **Get** /2/users/{id}/mentions | User mention timeline by User ID
[**UsersIdRetweets**](TweetsApi.md#UsersIdRetweets) | **Post** /2/users/{id}/retweets | Causes the User (in the path) to retweet the specified Tweet.
[**UsersIdTimeline**](TweetsApi.md#UsersIdTimeline) | **Get** /2/users/{id}/timelines/reverse_chronological | User home timeline by User ID
[**UsersIdTweets**](TweetsApi.md#UsersIdTweets) | **Get** /2/users/{id}/tweets | User Tweets timeline by User ID
[**UsersIdUnlike**](TweetsApi.md#UsersIdUnlike) | **Delete** /2/users/{id}/likes/{tweet_id} | Causes the User (in the path) to unlike the specified Tweet
[**UsersIdUnretweets**](TweetsApi.md#UsersIdUnretweets) | **Delete** /2/users/{id}/retweets/{source_tweet_id} | Causes the User (in the path) to unretweet the specified Tweet

# **AddOrDeleteRules**
> AddOrDeleteRulesResponse AddOrDeleteRules(ctx, body, optional)
Add/Delete rules

Add or delete rules from a User's active rule set. Users can provide unique, optionally tagged rules to add. Users can delete their entire rule set or a subset specified by rule ids or values.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AddOrDeleteRulesRequest**](AddOrDeleteRulesRequest.md)|  | 
 **optional** | ***TweetsApiAddOrDeleteRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TweetsApiAddOrDeleteRulesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dryRun** | **optional.**| Dry Run can be used with both the add and delete action, with the expected result given, but without actually taking any action in the system (meaning the end state will always be as it was when the request was submitted). This is particularly useful to validate rule changes. | 

### Return type

[**AddOrDeleteRulesResponse**](AddOrDeleteRulesResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateTweet**
> TweetCreateResponse CreateTweet(ctx, body)
Creation of a Tweet

Causes the User to create a Tweet under the authorized account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TweetCreateRequest**](TweetCreateRequest.md)|  | 

### Return type

[**TweetCreateResponse**](TweetCreateResponse.md)

### Authorization

[OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTweetById**
> TweetDeleteResponse DeleteTweetById(ctx, id)
Tweet delete by Tweet ID

Delete specified Tweet (in the path) by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the Tweet to be deleted. | 

### Return type

[**TweetDeleteResponse**](TweetDeleteResponse.md)

### Authorization

[OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindTweetById**
> Get2TweetsIdResponse FindTweetById(ctx, id, optional)
Tweet lookup by Tweet ID

Returns a variety of information about the Tweet specified by the requested ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A single Tweet ID. | 
 **optional** | ***TweetsApiFindTweetByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TweetsApiFindTweetByIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tweetFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Tweet fields to display. | 
 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **mediaFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Media fields to display. | 
 **pollFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Poll fields to display. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 
 **placeFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Place fields to display. | 

### Return type

[**Get2TweetsIdResponse**](Get2TweetsIdResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken), [OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindTweetsById**
> Get2TweetsResponse FindTweetsById(ctx, ids, optional)
Tweet lookup by Tweet IDs

Returns a variety of information about the Tweet specified by the requested ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ids** | [**[]string**](string.md)| A comma separated list of Tweet IDs. Up to 100 are allowed in a single request. | 
 **optional** | ***TweetsApiFindTweetsByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TweetsApiFindTweetsByIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tweetFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Tweet fields to display. | 
 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **mediaFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Media fields to display. | 
 **pollFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Poll fields to display. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 
 **placeFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Place fields to display. | 

### Return type

[**Get2TweetsResponse**](Get2TweetsResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken), [OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindTweetsThatQuoteATweet**
> Get2TweetsIdQuoteTweetsResponse FindTweetsThatQuoteATweet(ctx, id, optional)
Retrieve Tweets that quote a Tweet.

Returns a variety of information about each Tweet that quotes the Tweet specified by the requested ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A single Tweet ID. | 
 **optional** | ***TweetsApiFindTweetsThatQuoteATweetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TweetsApiFindTweetsThatQuoteATweetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maxResults** | **optional.Int32**| The maximum number of results to be returned. | [default to 10]
 **paginationToken** | [**optional.Interface of string**](.md)| This parameter is used to get a specified &#x27;page&#x27; of results. | 
 **exclude** | [**optional.Interface of []string**](string.md)| The set of entities to exclude (e.g. &#x27;replies&#x27; or &#x27;retweets&#x27;). | 
 **tweetFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Tweet fields to display. | 
 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **mediaFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Media fields to display. | 
 **pollFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Poll fields to display. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 
 **placeFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Place fields to display. | 

### Return type

[**Get2TweetsIdQuoteTweetsResponse**](Get2TweetsIdQuoteTweetsResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken), [OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRules**
> RulesLookupResponse GetRules(ctx, optional)
Rules lookup

Returns rules from a User's active rule set. Users can fetch all of their rules or a subset, specified by the provided rule ids.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TweetsApiGetRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TweetsApiGetRulesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | [**optional.Interface of []string**](string.md)| A comma-separated list of Rule IDs. | 
 **maxResults** | **optional.Int32**| The maximum number of results. | [default to 1000]
 **paginationToken** | **optional.String**| This value is populated by passing the &#x27;next_token&#x27; returned in a request to paginate through results. | 

### Return type

[**RulesLookupResponse**](RulesLookupResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTweetsFirehoseStream**
> StreamingTweetResponse GetTweetsFirehoseStream(ctx, partition, optional)
Firehose stream

Streams 100% of public Tweets.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **partition** | **int32**| The partition number. | 
 **optional** | ***TweetsApiGetTweetsFirehoseStreamOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TweetsApiGetTweetsFirehoseStreamOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **backfillMinutes** | **optional.Int32**| The number of minutes of backfill requested. | 
 **startTime** | **optional.Time**| YYYY-MM-DDTHH:mm:ssZ. The earliest UTC timestamp to which the Tweets will be provided. | 
 **endTime** | **optional.Time**| YYYY-MM-DDTHH:mm:ssZ. The latest UTC timestamp to which the Tweets will be provided. | 
 **tweetFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Tweet fields to display. | 
 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **mediaFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Media fields to display. | 
 **pollFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Poll fields to display. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 
 **placeFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Place fields to display. | 

### Return type

[**StreamingTweetResponse**](StreamingTweetResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTweetsSample10Stream**
> Get2TweetsSample10StreamResponse GetTweetsSample10Stream(ctx, partition, optional)
Sample 10% stream

Streams a deterministic 10% of public Tweets.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **partition** | **int32**| The partition number. | 
 **optional** | ***TweetsApiGetTweetsSample10StreamOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TweetsApiGetTweetsSample10StreamOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **backfillMinutes** | **optional.Int32**| The number of minutes of backfill requested. | 
 **startTime** | **optional.Time**| YYYY-MM-DDTHH:mm:ssZ. The earliest UTC timestamp to which the Tweets will be provided. | 
 **endTime** | **optional.Time**| YYYY-MM-DDTHH:mm:ssZ. The latest UTC timestamp to which the Tweets will be provided. | 
 **tweetFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Tweet fields to display. | 
 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **mediaFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Media fields to display. | 
 **pollFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Poll fields to display. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 
 **placeFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Place fields to display. | 

### Return type

[**Get2TweetsSample10StreamResponse**](Get2TweetsSample10StreamResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HideReplyById**
> TweetHideResponse HideReplyById(ctx, tweetId, optional)
Hide replies

Hides or unhides a reply to an owned conversation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tweetId** | [**string**](.md)| The ID of the reply that you want to hide or unhide. | 
 **optional** | ***TweetsApiHideReplyByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TweetsApiHideReplyByIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of TweetHideRequest**](TweetHideRequest.md)|  | 

### Return type

[**TweetHideResponse**](TweetHideResponse.md)

### Authorization

[OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListsIdTweets**
> Get2ListsIdTweetsResponse ListsIdTweets(ctx, id, optional)
List Tweets timeline by List ID.

Returns a list of Tweets associated with the provided List ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the List. | 
 **optional** | ***TweetsApiListsIdTweetsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TweetsApiListsIdTweetsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maxResults** | **optional.Int32**| The maximum number of results. | [default to 100]
 **paginationToken** | [**optional.Interface of string**](.md)| This parameter is used to get the next &#x27;page&#x27; of results. | 
 **tweetFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Tweet fields to display. | 
 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **mediaFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Media fields to display. | 
 **pollFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Poll fields to display. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 
 **placeFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Place fields to display. | 

### Return type

[**Get2ListsIdTweetsResponse**](Get2ListsIdTweetsResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken), [OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SampleStream**
> StreamingTweetResponse SampleStream(ctx, optional)
Sample stream

Streams a deterministic 1% of public Tweets.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TweetsApiSampleStreamOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TweetsApiSampleStreamOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **backfillMinutes** | **optional.Int32**| The number of minutes of backfill requested. | 
 **tweetFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Tweet fields to display. | 
 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **mediaFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Media fields to display. | 
 **pollFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Poll fields to display. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 
 **placeFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Place fields to display. | 

### Return type

[**StreamingTweetResponse**](StreamingTweetResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchStream**
> FilteredStreamingTweetResponse SearchStream(ctx, optional)
Filtered stream

Streams Tweets matching the stream's active rule set.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TweetsApiSearchStreamOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TweetsApiSearchStreamOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **backfillMinutes** | **optional.Int32**| The number of minutes of backfill requested. | 
 **startTime** | **optional.Time**| YYYY-MM-DDTHH:mm:ssZ. The earliest UTC timestamp from which the Tweets will be provided. | 
 **endTime** | **optional.Time**| YYYY-MM-DDTHH:mm:ssZ. The latest UTC timestamp to which the Tweets will be provided. | 
 **tweetFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Tweet fields to display. | 
 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **mediaFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Media fields to display. | 
 **pollFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Poll fields to display. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 
 **placeFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Place fields to display. | 

### Return type

[**FilteredStreamingTweetResponse**](FilteredStreamingTweetResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken)

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
 **optional** | ***TweetsApiSpaceBuyersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TweetsApiSpaceBuyersOpts struct
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
 **optional** | ***TweetsApiSpaceTweetsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TweetsApiSpaceTweetsOpts struct
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

# **TweetCountsFullArchiveSearch**
> Get2TweetsCountsAllResponse TweetCountsFullArchiveSearch(ctx, query, optional)
Full archive search counts

Returns Tweet Counts that match a search query.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **query** | **string**| One query/rule/filter for matching Tweets. Refer to https://t.co/rulelength to identify the max query length. | 
 **optional** | ***TweetsApiTweetCountsFullArchiveSearchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TweetsApiTweetCountsFullArchiveSearchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startTime** | **optional.Time**| YYYY-MM-DDTHH:mm:ssZ. The oldest UTC timestamp (from most recent 7 days) from which the Tweets will be provided. Timestamp is in second granularity and is inclusive (i.e. 12:00:01 includes the first second of the minute). | 
 **endTime** | **optional.Time**| YYYY-MM-DDTHH:mm:ssZ. The newest, most recent UTC timestamp to which the Tweets will be provided. Timestamp is in second granularity and is exclusive (i.e. 12:00:01 excludes the first second of the minute). | 
 **sinceId** | [**optional.Interface of string**](.md)| Returns results with a Tweet ID greater than (that is, more recent than) the specified ID. | 
 **untilId** | [**optional.Interface of string**](.md)| Returns results with a Tweet ID less than (that is, older than) the specified ID. | 
 **nextToken** | [**optional.Interface of string**](.md)| This parameter is used to get the next &#x27;page&#x27; of results. The value used with the parameter is pulled directly from the response provided by the API, and should not be modified. | 
 **paginationToken** | [**optional.Interface of string**](.md)| This parameter is used to get the next &#x27;page&#x27; of results. The value used with the parameter is pulled directly from the response provided by the API, and should not be modified. | 
 **granularity** | **optional.String**| The granularity for the search counts results. | [default to hour]
 **searchCountFields** | [**optional.Interface of []string**](string.md)| A comma separated list of SearchCount fields to display. | 

### Return type

[**Get2TweetsCountsAllResponse**](Get2TweetsCountsAllResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TweetCountsRecentSearch**
> Get2TweetsCountsRecentResponse TweetCountsRecentSearch(ctx, query, optional)
Recent search counts

Returns Tweet Counts from the last 7 days that match a search query.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **query** | **string**| One query/rule/filter for matching Tweets. Refer to https://t.co/rulelength to identify the max query length. | 
 **optional** | ***TweetsApiTweetCountsRecentSearchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TweetsApiTweetCountsRecentSearchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startTime** | **optional.Time**| YYYY-MM-DDTHH:mm:ssZ. The oldest UTC timestamp (from most recent 7 days) from which the Tweets will be provided. Timestamp is in second granularity and is inclusive (i.e. 12:00:01 includes the first second of the minute). | 
 **endTime** | **optional.Time**| YYYY-MM-DDTHH:mm:ssZ. The newest, most recent UTC timestamp to which the Tweets will be provided. Timestamp is in second granularity and is exclusive (i.e. 12:00:01 excludes the first second of the minute). | 
 **sinceId** | [**optional.Interface of string**](.md)| Returns results with a Tweet ID greater than (that is, more recent than) the specified ID. | 
 **untilId** | [**optional.Interface of string**](.md)| Returns results with a Tweet ID less than (that is, older than) the specified ID. | 
 **nextToken** | [**optional.Interface of string**](.md)| This parameter is used to get the next &#x27;page&#x27; of results. The value used with the parameter is pulled directly from the response provided by the API, and should not be modified. | 
 **paginationToken** | [**optional.Interface of string**](.md)| This parameter is used to get the next &#x27;page&#x27; of results. The value used with the parameter is pulled directly from the response provided by the API, and should not be modified. | 
 **granularity** | **optional.String**| The granularity for the search counts results. | [default to hour]
 **searchCountFields** | [**optional.Interface of []string**](string.md)| A comma separated list of SearchCount fields to display. | 

### Return type

[**Get2TweetsCountsRecentResponse**](Get2TweetsCountsRecentResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TweetsFullarchiveSearch**
> Get2TweetsSearchAllResponse TweetsFullarchiveSearch(ctx, query, optional)
Full-archive search

Returns Tweets that match a search query.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **query** | **string**| One query/rule/filter for matching Tweets. Refer to https://t.co/rulelength to identify the max query length. | 
 **optional** | ***TweetsApiTweetsFullarchiveSearchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TweetsApiTweetsFullarchiveSearchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startTime** | **optional.Time**| YYYY-MM-DDTHH:mm:ssZ. The oldest UTC timestamp from which the Tweets will be provided. Timestamp is in second granularity and is inclusive (i.e. 12:00:01 includes the first second of the minute). | 
 **endTime** | **optional.Time**| YYYY-MM-DDTHH:mm:ssZ. The newest, most recent UTC timestamp to which the Tweets will be provided. Timestamp is in second granularity and is exclusive (i.e. 12:00:01 excludes the first second of the minute). | 
 **sinceId** | [**optional.Interface of string**](.md)| Returns results with a Tweet ID greater than (that is, more recent than) the specified ID. | 
 **untilId** | [**optional.Interface of string**](.md)| Returns results with a Tweet ID less than (that is, older than) the specified ID. | 
 **maxResults** | **optional.Int32**| The maximum number of search results to be returned by a request. | [default to 10]
 **nextToken** | [**optional.Interface of string**](.md)| This parameter is used to get the next &#x27;page&#x27; of results. The value used with the parameter is pulled directly from the response provided by the API, and should not be modified. | 
 **paginationToken** | [**optional.Interface of string**](.md)| This parameter is used to get the next &#x27;page&#x27; of results. The value used with the parameter is pulled directly from the response provided by the API, and should not be modified. | 
 **sortOrder** | **optional.String**| This order in which to return results. | 
 **tweetFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Tweet fields to display. | 
 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **mediaFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Media fields to display. | 
 **pollFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Poll fields to display. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 
 **placeFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Place fields to display. | 

### Return type

[**Get2TweetsSearchAllResponse**](Get2TweetsSearchAllResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TweetsRecentSearch**
> Get2TweetsSearchRecentResponse TweetsRecentSearch(ctx, query, optional)
Recent search

Returns Tweets from the last 7 days that match a search query.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **query** | **string**| One query/rule/filter for matching Tweets. Refer to https://t.co/rulelength to identify the max query length. | 
 **optional** | ***TweetsApiTweetsRecentSearchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TweetsApiTweetsRecentSearchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startTime** | **optional.Time**| YYYY-MM-DDTHH:mm:ssZ. The oldest UTC timestamp from which the Tweets will be provided. Timestamp is in second granularity and is inclusive (i.e. 12:00:01 includes the first second of the minute). | 
 **endTime** | **optional.Time**| YYYY-MM-DDTHH:mm:ssZ. The newest, most recent UTC timestamp to which the Tweets will be provided. Timestamp is in second granularity and is exclusive (i.e. 12:00:01 excludes the first second of the minute). | 
 **sinceId** | [**optional.Interface of string**](.md)| Returns results with a Tweet ID greater than (that is, more recent than) the specified ID. | 
 **untilId** | [**optional.Interface of string**](.md)| Returns results with a Tweet ID less than (that is, older than) the specified ID. | 
 **maxResults** | **optional.Int32**| The maximum number of search results to be returned by a request. | [default to 10]
 **nextToken** | [**optional.Interface of string**](.md)| This parameter is used to get the next &#x27;page&#x27; of results. The value used with the parameter is pulled directly from the response provided by the API, and should not be modified. | 
 **paginationToken** | [**optional.Interface of string**](.md)| This parameter is used to get the next &#x27;page&#x27; of results. The value used with the parameter is pulled directly from the response provided by the API, and should not be modified. | 
 **sortOrder** | **optional.String**| This order in which to return results. | 
 **tweetFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Tweet fields to display. | 
 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **mediaFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Media fields to display. | 
 **pollFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Poll fields to display. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 
 **placeFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Place fields to display. | 

### Return type

[**Get2TweetsSearchRecentResponse**](Get2TweetsSearchRecentResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken), [OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersIdLike**
> UsersLikesCreateResponse UsersIdLike(ctx, id, optional)
Causes the User (in the path) to like the specified Tweet

Causes the User (in the path) to like the specified Tweet. The User in the path must match the User context authorizing the request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the authenticated source User that is requesting to like the Tweet. | 
 **optional** | ***TweetsApiUsersIdLikeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TweetsApiUsersIdLikeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UsersLikesCreateRequest**](UsersLikesCreateRequest.md)|  | 

### Return type

[**UsersLikesCreateResponse**](UsersLikesCreateResponse.md)

### Authorization

[OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersIdLikedTweets**
> Get2UsersIdLikedTweetsResponse UsersIdLikedTweets(ctx, id, optional)
Returns Tweet objects liked by the provided User ID

Returns a list of Tweets liked by the provided User ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the User to lookup. | 
 **optional** | ***TweetsApiUsersIdLikedTweetsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TweetsApiUsersIdLikedTweetsOpts struct
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

[**Get2UsersIdLikedTweetsResponse**](Get2UsersIdLikedTweetsResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken), [OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersIdMentions**
> Get2UsersIdMentionsResponse UsersIdMentions(ctx, id, optional)
User mention timeline by User ID

Returns Tweet objects that mention username associated to the provided User ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the User to lookup. | 
 **optional** | ***TweetsApiUsersIdMentionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TweetsApiUsersIdMentionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sinceId** | [**optional.Interface of string**](.md)| The minimum Tweet ID to be included in the result set. This parameter takes precedence over start_time if both are specified. | 
 **untilId** | [**optional.Interface of string**](.md)| The maximum Tweet ID to be included in the result set. This parameter takes precedence over end_time if both are specified. | 
 **maxResults** | **optional.Int32**| The maximum number of results. | 
 **paginationToken** | [**optional.Interface of string**](.md)| This parameter is used to get the next &#x27;page&#x27; of results. | 
 **startTime** | **optional.Time**| YYYY-MM-DDTHH:mm:ssZ. The earliest UTC timestamp from which the Tweets will be provided. The since_id parameter takes precedence if it is also specified. | 
 **endTime** | **optional.Time**| YYYY-MM-DDTHH:mm:ssZ. The latest UTC timestamp to which the Tweets will be provided. The until_id parameter takes precedence if it is also specified. | 
 **tweetFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Tweet fields to display. | 
 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **mediaFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Media fields to display. | 
 **pollFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Poll fields to display. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 
 **placeFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Place fields to display. | 

### Return type

[**Get2UsersIdMentionsResponse**](Get2UsersIdMentionsResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken), [OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersIdRetweets**
> UsersRetweetsCreateResponse UsersIdRetweets(ctx, id, optional)
Causes the User (in the path) to retweet the specified Tweet.

Causes the User (in the path) to retweet the specified Tweet. The User in the path must match the User context authorizing the request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the authenticated source User that is requesting to retweet the Tweet. | 
 **optional** | ***TweetsApiUsersIdRetweetsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TweetsApiUsersIdRetweetsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UsersRetweetsCreateRequest**](UsersRetweetsCreateRequest.md)|  | 

### Return type

[**UsersRetweetsCreateResponse**](UsersRetweetsCreateResponse.md)

### Authorization

[OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersIdTimeline**
> Get2UsersIdTimelinesReverseChronologicalResponse UsersIdTimeline(ctx, id, optional)
User home timeline by User ID

Returns Tweet objects that appears in the provided User ID's home timeline

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the authenticated source User to list Reverse Chronological Timeline Tweets of. | 
 **optional** | ***TweetsApiUsersIdTimelineOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TweetsApiUsersIdTimelineOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sinceId** | [**optional.Interface of string**](.md)| The minimum Tweet ID to be included in the result set. This parameter takes precedence over start_time if both are specified. | 
 **untilId** | [**optional.Interface of string**](.md)| The maximum Tweet ID to be included in the result set. This parameter takes precedence over end_time if both are specified. | 
 **maxResults** | **optional.Int32**| The maximum number of results. | 
 **paginationToken** | [**optional.Interface of string**](.md)| This parameter is used to get the next &#x27;page&#x27; of results. | 
 **exclude** | [**optional.Interface of []string**](string.md)| The set of entities to exclude (e.g. &#x27;replies&#x27; or &#x27;retweets&#x27;). | 
 **startTime** | **optional.Time**| YYYY-MM-DDTHH:mm:ssZ. The earliest UTC timestamp from which the Tweets will be provided. The since_id parameter takes precedence if it is also specified. | 
 **endTime** | **optional.Time**| YYYY-MM-DDTHH:mm:ssZ. The latest UTC timestamp to which the Tweets will be provided. The until_id parameter takes precedence if it is also specified. | 
 **tweetFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Tweet fields to display. | 
 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **mediaFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Media fields to display. | 
 **pollFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Poll fields to display. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 
 **placeFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Place fields to display. | 

### Return type

[**Get2UsersIdTimelinesReverseChronologicalResponse**](Get2UsersIdTimelinesReverseChronologicalResponse.md)

### Authorization

[OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersIdTweets**
> Get2UsersIdTweetsResponse UsersIdTweets(ctx, id, optional)
User Tweets timeline by User ID

Returns a list of Tweets authored by the provided User ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the User to lookup. | 
 **optional** | ***TweetsApiUsersIdTweetsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TweetsApiUsersIdTweetsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sinceId** | [**optional.Interface of string**](.md)| The minimum Tweet ID to be included in the result set. This parameter takes precedence over start_time if both are specified. | 
 **untilId** | [**optional.Interface of string**](.md)| The maximum Tweet ID to be included in the result set. This parameter takes precedence over end_time if both are specified. | 
 **maxResults** | **optional.Int32**| The maximum number of results. | 
 **paginationToken** | [**optional.Interface of string**](.md)| This parameter is used to get the next &#x27;page&#x27; of results. | 
 **exclude** | [**optional.Interface of []string**](string.md)| The set of entities to exclude (e.g. &#x27;replies&#x27; or &#x27;retweets&#x27;). | 
 **startTime** | **optional.Time**| YYYY-MM-DDTHH:mm:ssZ. The earliest UTC timestamp from which the Tweets will be provided. The since_id parameter takes precedence if it is also specified. | 
 **endTime** | **optional.Time**| YYYY-MM-DDTHH:mm:ssZ. The latest UTC timestamp to which the Tweets will be provided. The until_id parameter takes precedence if it is also specified. | 
 **tweetFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Tweet fields to display. | 
 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **mediaFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Media fields to display. | 
 **pollFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Poll fields to display. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 
 **placeFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Place fields to display. | 

### Return type

[**Get2UsersIdTweetsResponse**](Get2UsersIdTweetsResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken), [OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersIdUnlike**
> UsersLikesDeleteResponse UsersIdUnlike(ctx, id, tweetId)
Causes the User (in the path) to unlike the specified Tweet

Causes the User (in the path) to unlike the specified Tweet. The User must match the User context authorizing the request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the authenticated source User that is requesting to unlike the Tweet. | 
  **tweetId** | [**string**](.md)| The ID of the Tweet that the User is requesting to unlike. | 

### Return type

[**UsersLikesDeleteResponse**](UsersLikesDeleteResponse.md)

### Authorization

[OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersIdUnretweets**
> UsersRetweetsDeleteResponse UsersIdUnretweets(ctx, id, sourceTweetId)
Causes the User (in the path) to unretweet the specified Tweet

Causes the User (in the path) to unretweet the specified Tweet. The User must match the User context authorizing the request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the authenticated source User that is requesting to retweet the Tweet. | 
  **sourceTweetId** | [**string**](.md)| The ID of the Tweet that the User is requesting to unretweet. | 

### Return type

[**UsersRetweetsDeleteResponse**](UsersRetweetsDeleteResponse.md)

### Authorization

[OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

