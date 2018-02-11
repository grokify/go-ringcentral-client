# GlipGroupsEvent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Internal identifier of a group | [optional] [default to null]
**Type_** | **string** | Type of a group. &#39;PrivateChat&#39; is a group of 2 members. &#39;Group&#39; is a chat of 2 and more participants, the membership cannot be changed after group creation. &#39;Team&#39; is a chat of 1 and more participants, the membership can be modified in future. &#39;PersonalChat&#39; is a private chat thread of a user | [optional] [default to null]
**IsPublic** | **bool** | For &#39;Team&#39; group type only. Team access level | [optional] [default to null]
**Name** | **string** | For &#39;Team&#39; group type only. Team name | [optional] [default to null]
**Description** | **string** | For &#39;Team&#39; group type only. Team description | [optional] [default to null]
**Members** | **[]string** | Identifier(s) of group members | [optional] [default to null]
**CreationTime** | [**time.Time**](time.Time.md) | Group creation datetime in ISO 8601 format | [optional] [default to null]
**LastModifiedTime** | [**time.Time**](time.Time.md) | Group last change datetime in ISO 8601 format | [optional] [default to null]
**EventType** | **string** | Type of a group event. Only the person who joined/was added to a group will receive &#39;GroupJoined&#39; notification. Only the person who left/was removed from a group will receive &#39;GroupLeft&#39; notification | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


