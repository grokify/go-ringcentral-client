# GlipPostEvent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Internal identifier of a post | [optional] [default to null]
**EventType** | **string** | Type of a post event | [optional] [default to null]
**GroupId** | **string** | Internal identifier of a group a post belongs to | [optional] [default to null]
**Type_** | **string** | Type of a post. &#39;TextMessage&#39; - an incoming text message; &#39;PersonJoined&#39; - a message notifying that person has joined a conversation; &#39;PersonsAdded&#39; - a message notifying that a person(s) were added to a conversation | [optional] [default to null]
**Text** | **string** | For &#39;TextMessage&#39; post type only. Message text | [optional] [default to null]
**CreatorId** | **string** | Internal identifier of a user - author of a post | [optional] [default to null]
**AddedPersonIds** | **[]string** | For PersonsAdded post type only. Identifiers of persons added to a group | [optional] [default to null]
**RemovedPersonIds** | **[]string** | For PersonsRemoved post type only. Identifiers of persons removed from a group | [optional] [default to null]
**Mentions** | [**[]GlipMentionsInfo**](GlipMentionsInfo.md) | List of at mentions in post text with names. | [optional] [default to null]
**CreationTime** | [**time.Time**](time.Time.md) | Post creation datetime in ISO 8601 format | [optional] [default to null]
**LastModifiedTime** | [**time.Time**](time.Time.md) | Post last change datetime in ISO 8601 format | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


