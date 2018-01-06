# \GlipApi

All URIs are relative to *https://api.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGlipFile**](GlipApi.md#CreateGlipFile) | **Post** /restapi/v1.0/glip/files | Upload File
[**CreateGroup**](GlipApi.md#CreateGroup) | **Post** /restapi/v1.0/glip/groups | Create Group
[**CreatePost**](GlipApi.md#CreatePost) | **Post** /restapi/v1.0/glip/groups/{groupId}/posts | Create Post
[**LoadCompany**](GlipApi.md#LoadCompany) | **Get** /restapi/v1.0/glip/companies/{companyId} | Get Company Info
[**LoadGlipFile**](GlipApi.md#LoadGlipFile) | **Get** /restapi/v1.0/glip/files/{fileId} | Get File Info
[**LoadGroup**](GlipApi.md#LoadGroup) | **Get** /restapi/v1.0/glip/groups/{groupId} | Get Group
[**LoadGroupList**](GlipApi.md#LoadGroupList) | **Get** /restapi/v1.0/glip/groups | Get User Groups
[**LoadPerson**](GlipApi.md#LoadPerson) | **Get** /restapi/v1.0/glip/persons/{personId} | Get Person
[**LoadPosts**](GlipApi.md#LoadPosts) | **Get** /restapi/v1.0/glip/groups/{groupId}/posts | Get Posts
[**LoadUnreadMessageCount**](GlipApi.md#LoadUnreadMessageCount) | **Get** /restapi/v1.0/glip/profile | Glip Unread Message Count
[**UpdateGroup**](GlipApi.md#UpdateGroup) | **Post** /restapi/v1.0/glip/groups/{groupId}/bulk-assign | Edit Group Members


# **CreateGlipFile**
> PostGlipFile CreateGlipFile(ctx, body, optional)
Upload File

<p style='font-style:italic;'>Since 1.0.31 (Release 9.2)</p><p>Posts a file.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>Glip</td><td>Availability of Glip</td></tr></tbody></table><h4>API Group</h4><p>Heavy</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | ***os.File**| The file to upload | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | ***os.File**| The file to upload | 
 **groupId** | **string**| Internal identifier of a group the post with file attached will be added to | 
 **name** | **string**| Name of a file attached | 

### Return type

[**PostGlipFile**](PostGlipFile.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateGroup**
> GlipGroupInfo CreateGroup(ctx, body)
Create Group

<p style='font-style:italic;'>Since 1.0.28 (Release 8.4)</p><p>Creates a group.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>Glip</td><td>Availability of Glip</td></tr></tbody></table><h4>API Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**GlipCreateGroup**](GlipCreateGroup.md)| JSON body | 

### Return type

[**GlipGroupInfo**](GlipGroupInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePost**
> GlipPostInfo CreatePost(ctx, groupId, body)
Create Post

<p style='font-style:italic;'>Since 1.0.28 (Release 8.4)</p><p>Creates a post.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>Glip</td><td>Availability of Glip</td></tr></tbody></table><h4>API Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **groupId** | **string**| Id of a group to send post | 
  **body** | [**GlipCreatePost**](GlipCreatePost.md)| JSON body | 

### Return type

[**GlipPostInfo**](GlipPostInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoadCompany**
> GlipCompany LoadCompany(ctx, companyId)
Get Company Info

<p style='font-style:italic;'>Since 1.0.28 (Release 8.4)</p><p>Returns a company by ID.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>Glip</td><td>Availability of Glip</td></tr></tbody></table><h4>API Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **companyId** | **string**| Internal identifier of an RC account/Glip company, or tilde (~) to indicate a company the current user belongs to | 

### Return type

[**GlipCompany**](GlipCompany.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoadGlipFile**
> PostGlipFile LoadGlipFile(ctx, fileId)
Get File Info

<p style='font-style:italic;'>Since 1.0.31 (Release 9.2)</p><p>Returns a file.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>Glip</td><td>Availability of Glip</td></tr></tbody></table><h4>API Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fileId** | **string**| Internal identifier of file. | 

### Return type

[**PostGlipFile**](PostGlipFile.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoadGroup**
> GlipGroupInfo LoadGroup(ctx, groupId)
Get Group

<p style='font-style:italic;'>Since 1.0.28 (Release 8.4)</p><p>Returns a group or few groups by ID(s). Batch request is supported.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>Glip</td><td>Availability of Glip</td></tr></tbody></table><h4>API Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **groupId** | **string**| Internal identifier of a group to be returned, the maximum number of IDs is 30 | 

### Return type

[**GlipGroupInfo**](GlipGroupInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoadGroupList**
> GlipGroupList LoadGroupList(ctx, optional)
Get User Groups

<p style='font-style:italic;'>Since 1.0.28 (Release 8.4)</p><p>Returns the list of groups associated with the user.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>Glip</td><td>Availability of Glip</td></tr></tbody></table><h4>API Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | [**[]string**](string.md)| Type of a group. &#39;PrivateChat&#39; is a group of 2 members. &#39;Group&#39; is a chat of 2 and more participants, the membership cannot be changed after group creation. &#39;Team&#39; is a chat of 1 and more participants, the membership can be modified in future | 
 **pageToken** | **string**| Token of a page to be returned, see Glip Navigation Info | 
 **recordCount** | **int64**| Max numbers of records to be returned. The default/maximum value is 250 | 

### Return type

[**GlipGroupList**](GlipGroupList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoadPerson**
> GlipPersonInfo LoadPerson(ctx, personId)
Get Person

<p style='font-style:italic;'>Since 1.0.28 (Release 8.4)</p><p>Returns a user or few users by ID(s). Batch request is supported.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>Glip</td><td>Availability of Glip</td></tr></tbody></table><h4>API Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **personId** | **string**| Internal identifier of a user to be returned, the maximum number of IDs is 30 | 

### Return type

[**GlipPersonInfo**](GlipPersonInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoadPosts**
> GlipPosts LoadPosts(ctx, groupId, optional)
Get Posts

<p style='font-style:italic;'>Since 1.0.28 (Release 8.4)</p><p>Returns list of posts.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>Glip</td><td>Availability of Glip</td></tr></tbody></table><h4>API Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **groupId** | **string**| Identifier of a group to filter posts | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **string**| Identifier of a group to filter posts | 
 **pageToken** | **string**| Token of a page to be returned, see Glip Navigation Info | 
 **recordCount** | **int64**| Max numbers of records to be returned. The default/maximum value is 250 | 

### Return type

[**GlipPosts**](GlipPosts.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoadUnreadMessageCount**
> GlipUnreadMessageCount LoadUnreadMessageCount(ctx, optional)
Glip Unread Message Count

<p style='font-style:italic;'>Since 1.0.30 (Release 9.1)</p><p>Returns Glip unread message count.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>Glip</td><td>Availability of Glip</td></tr></tbody></table><h4>API Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string**| List of attributes to be returned. To return unread message count &#39;unreadPostsCount&#39; and &#39;tooManyUnreadPosts&#39; should be specified | 

### Return type

[**GlipUnreadMessageCount**](GlipUnreadMessageCount.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateGroup**
> GlipGroupInfo UpdateGroup(ctx, groupId, body)
Edit Group Members

<p style='font-style:italic;'>Since 1.0.28 (Release 8.4)</p><p>Updates group members. Please note: Only groups of 'Team' type can be updated. Currently only one operation at a time (either adding or removal) is supported.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>Glip</td><td>Availability of Glip</td></tr></tbody></table><h4>API Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **groupId** | **string**| Internal identifier of a group to be edited | 
  **body** | [**EditGroupRequest**](EditGroupRequest.md)| JSON body | 

### Return type

[**GlipGroupInfo**](GlipGroupInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

