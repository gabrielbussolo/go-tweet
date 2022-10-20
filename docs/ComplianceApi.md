# {{classname}}

All URIs are relative to *https://api.twitter.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBatchComplianceJob**](ComplianceApi.md#CreateBatchComplianceJob) | **Post** /2/compliance/jobs | Create compliance job
[**GetBatchComplianceJob**](ComplianceApi.md#GetBatchComplianceJob) | **Get** /2/compliance/jobs/{id} | Get Compliance Job
[**GetTweetsComplianceStream**](ComplianceApi.md#GetTweetsComplianceStream) | **Get** /2/tweets/compliance/stream | Tweets Compliance stream
[**GetTweetsLabelStream**](ComplianceApi.md#GetTweetsLabelStream) | **Get** /2/tweets/label/stream | Tweets Label stream
[**GetUsersComplianceStream**](ComplianceApi.md#GetUsersComplianceStream) | **Get** /2/users/compliance/stream | Users Compliance stream
[**ListBatchComplianceJobs**](ComplianceApi.md#ListBatchComplianceJobs) | **Get** /2/compliance/jobs | List Compliance Jobs

# **CreateBatchComplianceJob**
> CreateComplianceJobResponse CreateBatchComplianceJob(ctx, body)
Create compliance job

Creates a compliance for the given job type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateComplianceJobRequest**](CreateComplianceJobRequest.md)|  | 

### Return type

[**CreateComplianceJobResponse**](CreateComplianceJobResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBatchComplianceJob**
> Get2ComplianceJobsIdResponse GetBatchComplianceJob(ctx, id, optional)
Get Compliance Job

Returns a single Compliance Job by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the Compliance Job to retrieve. | 
 **optional** | ***ComplianceApiGetBatchComplianceJobOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ComplianceApiGetBatchComplianceJobOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **complianceJobFields** | [**optional.Interface of []string**](string.md)| A comma separated list of ComplianceJob fields to display. | 

### Return type

[**Get2ComplianceJobsIdResponse**](Get2ComplianceJobsIdResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTweetsComplianceStream**
> TweetComplianceStreamResponse GetTweetsComplianceStream(ctx, partition, optional)
Tweets Compliance stream

Streams 100% of compliance data for Tweets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **partition** | **int32**| The partition number. | 
 **optional** | ***ComplianceApiGetTweetsComplianceStreamOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ComplianceApiGetTweetsComplianceStreamOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **backfillMinutes** | **optional.Int32**| The number of minutes of backfill requested. | 
 **startTime** | **optional.Time**| YYYY-MM-DDTHH:mm:ssZ. The earliest UTC timestamp from which the Tweet Compliance events will be provided. | 
 **endTime** | **optional.Time**| YYYY-MM-DDTHH:mm:ssZ. The latest UTC timestamp to which the Tweet Compliance events will be provided. | 

### Return type

[**TweetComplianceStreamResponse**](TweetComplianceStreamResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTweetsLabelStream**
> TweetLabelStreamResponse GetTweetsLabelStream(ctx, optional)
Tweets Label stream

Streams 100% of labeling events applied to Tweets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ComplianceApiGetTweetsLabelStreamOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ComplianceApiGetTweetsLabelStreamOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **backfillMinutes** | **optional.Int32**| The number of minutes of backfill requested. | 
 **startTime** | **optional.Time**| YYYY-MM-DDTHH:mm:ssZ. The earliest UTC timestamp from which the Tweet labels will be provided. | 
 **endTime** | **optional.Time**| YYYY-MM-DDTHH:mm:ssZ. The latest UTC timestamp from which the Tweet labels will be provided. | 

### Return type

[**TweetLabelStreamResponse**](TweetLabelStreamResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsersComplianceStream**
> UserComplianceStreamResponse GetUsersComplianceStream(ctx, partition, optional)
Users Compliance stream

Streams 100% of compliance data for Users

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **partition** | **int32**| The partition number. | 
 **optional** | ***ComplianceApiGetUsersComplianceStreamOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ComplianceApiGetUsersComplianceStreamOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **backfillMinutes** | **optional.Int32**| The number of minutes of backfill requested. | 
 **startTime** | **optional.Time**| YYYY-MM-DDTHH:mm:ssZ. The earliest UTC timestamp from which the User Compliance events will be provided. | 
 **endTime** | **optional.Time**| YYYY-MM-DDTHH:mm:ssZ. The latest UTC timestamp from which the User Compliance events will be provided. | 

### Return type

[**UserComplianceStreamResponse**](UserComplianceStreamResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListBatchComplianceJobs**
> Get2ComplianceJobsResponse ListBatchComplianceJobs(ctx, type_, optional)
List Compliance Jobs

Returns recent Compliance Jobs for a given job type and optional job status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**| Type of Compliance Job to list. | 
 **optional** | ***ComplianceApiListBatchComplianceJobsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ComplianceApiListBatchComplianceJobsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **optional.String**| Status of Compliance Job to list. | 
 **complianceJobFields** | [**optional.Interface of []string**](string.md)| A comma separated list of ComplianceJob fields to display. | 

### Return type

[**Get2ComplianceJobsResponse**](Get2ComplianceJobsResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

