# Space

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | [**time.Time**](time.Time.md) | Creation time of the Space. | [optional] [default to null]
**CreatorId** | **string** |  | [optional] [default to null]
**EndedAt** | [**time.Time**](time.Time.md) | End time of the Space. | [optional] [default to null]
**HostIds** | **[]string** | The user ids for the hosts of the Space. | [optional] [default to null]
**Id** | **string** |  | [default to null]
**InvitedUserIds** | **[]string** | An array of user ids for people who were invited to a Space. | [optional] [default to null]
**IsTicketed** | **bool** | Denotes if the Space is a ticketed Space. | [optional] [default to null]
**Lang** | **string** | The language of the Space. | [optional] [default to null]
**ParticipantCount** | **int32** | The number of participants in a Space. | [optional] [default to null]
**ScheduledStart** | [**time.Time**](time.Time.md) | A date time stamp for when a Space is scheduled to begin. | [optional] [default to null]
**SpeakerIds** | **[]string** | An array of user ids for people who were speakers in a Space. | [optional] [default to null]
**StartedAt** | [**time.Time**](time.Time.md) | When the Space was started as a date string. | [optional] [default to null]
**State** | **string** | The current state of the Space. | [default to null]
**SubscriberCount** | **int32** | The number of people who have either purchased a ticket or set a reminder for this Space. | [optional] [default to null]
**Title** | **string** | The title of the Space. | [optional] [default to null]
**Topics** | [**[]SpaceTopics**](Space_topics.md) | The topics of a Space, as selected by its creator. | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | When the Space was last updated. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

