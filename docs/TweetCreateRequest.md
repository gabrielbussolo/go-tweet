# TweetCreateRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardUri** | **string** | Card Uri Parameter. This is mutually exclusive from Quote Tweet Id, Poll, Media, and Direct Message Deep Link. | [optional] [default to null]
**DirectMessageDeepLink** | **string** | Link to take the conversation from the public timeline to a private Direct Message. | [optional] [default to null]
**ForSuperFollowersOnly** | **bool** | Exclusive Tweet for super followers. | [optional] [default to false]
**Geo** | [***TweetCreateRequestGeo**](TweetCreateRequest_geo.md) |  | [optional] [default to null]
**Media** | [***TweetCreateRequestMedia**](TweetCreateRequest_media.md) |  | [optional] [default to null]
**Nullcast** | **bool** | Nullcasted (promoted-only) Tweets do not appear in the public timeline and are not served to followers. | [optional] [default to false]
**Poll** | [***TweetCreateRequestPoll**](TweetCreateRequest_poll.md) |  | [optional] [default to null]
**QuoteTweetId** | **string** |  | [optional] [default to null]
**Reply** | [***TweetCreateRequestReply**](TweetCreateRequest_reply.md) |  | [optional] [default to null]
**ReplySettings** | **string** | Settings to indicate who can reply to the Tweet. | [optional] [default to null]
**Text** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

