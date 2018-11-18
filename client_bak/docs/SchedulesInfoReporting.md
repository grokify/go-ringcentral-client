# SchedulesInfoReporting

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScheduleId** | **string** | Unique schedule identifier | [optional] 
**Recurrence** | [**[]RecurrenceInfoReporting**](RecurrenceInfoReporting.md) | Recurrence pattern of a schedule | [optional] 
**ViewType** | **string** | Type of report. Detailed reports include tables with data. Simple reports only include charts | [optional] 
**Attachments** | [**[]AttachmentInfoReporting**](AttachmentInfoReporting.md) | Set of optional attachments. Basically, every report email is in HTML format. Optionally, it can contain PDF or CSV files | [optional] 
**Pages** | **[]string** | List of pages to include to the report. If empty, all pages are included. Otherwise, only specified pages are included. API doesn&#39;t check validity of page names. Client application is responsible to do that | [optional] 
**Recipients** | **[]string** | List of emails to which to send rendered reports | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


