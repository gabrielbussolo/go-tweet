# Tweet

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attachments** | [***TweetAttachments**](Tweet_attachments.md) |  | [optional] [default to null]
**AuthorId** | **string** |  | [optional] [default to null]
**ContextAnnotations** | [**[]ContextAnnotation**](ContextAnnotation.md) |  | [optional] [default to null]
**ConversationId** | **string** |  | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | Creation time of the Tweet. | [optional] [default to null]
**EditControls** | [***TweetEditControls**](Tweet_edit_controls.md) |  | [optional] [default to null]
**EditHistoryTweetIds** | **[]string** | A list of Tweet Ids in this Tweet chain. | [default to null]
**Entities** | [***FullTextEntities**](FullTextEntities.md) |  | [optional] [default to null]
**Geo** | [***TweetGeo**](Tweet_geo.md) |  | [optional] [default to null]
**Id** | **string** |  | [default to null]
**InReplyToUserId** | **string** |  | [optional] [default to null]
**Lang** | **string** | Language of the Tweet, if detected by Twitter. Returned as a BCP47 language tag. | [optional] [default to null]
**NonPublicMetrics** | [***TweetNonPublicMetrics**](Tweet_non_public_metrics.md) |  | [optional] [default to null]
**OrganicMetrics** | [***TweetOrganicMetrics**](Tweet_organic_metrics.md) |  | [optional] [default to null]
**PossiblySensitive** | **bool** | Indicates if this Tweet contains URLs marked as sensitive, for example content suitable for mature audiences. | [optional] [default to null]
**PromotedMetrics** | [***TweetPromotedMetrics**](Tweet_promoted_metrics.md) |  | [optional] [default to null]
**PublicMetrics** | [***TweetPublicMetrics**](Tweet_public_metrics.md) |  | [optional] [default to null]
**ReferencedTweets** | [**[]TweetReferencedTweets**](Tweet_referenced_tweets.md) | A list of Tweets this Tweet refers to. For example, if the parent Tweet is a Retweet, a Quoted Tweet or a Reply, it will include the related Tweet referenced to by its parent. | [optional] [default to null]
**ReplySettings** | [***ReplySettings**](ReplySettings.md) |  | [optional] [default to null]
**Source** | **string** | The name of the app the user Tweeted from. | [optional] [default to null]
**Text** | **string** |  | [default to null]
**Withheld** | [***TweetWithheld**](TweetWithheld.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

