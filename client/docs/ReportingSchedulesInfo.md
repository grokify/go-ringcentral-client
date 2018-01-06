# ReportingSchedulesInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScheduleId** | **string** | Unique schedule identifier | [optional] [default to null]
**Recurrence** | [**[]ReportingRecurrenceInfo**](ReportingRecurrenceInfo.md) | Recurrence pattern of a schedule | [optional] [default to null]
**ViewType** | **string** | Type of report. Detailed reports include tables with data. Simple reports only include charts | [optional] [default to null]
**Attachments** | [**[]ReportingAttachmentInfo**](ReportingAttachmentInfo.md) | Set of optional attachments. Basically, every report email is in HTML format. Optionally, it can contain PDF or CSV files | [optional] [default to null]
**Pages** | **[]string** | List of pages to include to the report. If empty, all pages are included. Otherwise, only specified pages are included. API doesn&#39;t check validity of page names. Client application is responsible to do that | [optional] [default to null]
**Recipients** | **[]string** | List of emails to which to send rendered reports | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


