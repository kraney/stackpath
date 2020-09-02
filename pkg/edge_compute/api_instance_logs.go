/*
 * Edge Compute
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package edge_compute

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ _context.Context
)

// InstanceLogsApiService InstanceLogsApi service
type InstanceLogsApiService service

// GetLogsOpts Optional parameters for the method 'GetLogs'
type GetLogsOpts struct {
    ContainerName optional.String
    Follow optional.Bool
    Previous optional.Bool
    SinceSeconds optional.String
    SinceTime optional.Time
    Timestamps optional.Bool
    TailLines optional.String
    LimitBytes optional.String
}

/*
GetLogs Get log stream
Retrieve a stream of logs generated by a workload instance&#39;s containers. Logs are generated by the containers and are not modified by StackPath. Log contents vary by the application running in the container, though many containerized applications are configured to log to STDOUT and STDERR.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param stackId A stack ID or slug
 * @param workloadId An EdgeCompute workload ID
 * @param instanceName An EdgeCompute workload instance name
 * @param optional nil or *GetLogsOpts - Optional Parameters:
 * @param "ContainerName" (optional.String) -  The name of the container to obtain logs for. This defaults to first container in the instance.
 * @param "Follow" (optional.Bool) -  Whether or not to follow the instance's log stream. This defaults to false.
 * @param "Previous" (optional.Bool) -  Whether or not to return log entries made by previously terminated containers. This defaults to false.
 * @param "SinceSeconds" (optional.String) -  A relative time in seconds before the current time from which to show logs. If this value precedes the time an instance was started, only logs since the instance's start time will be returned. If this value is in the future, no logs will be returned.  Only one of since_seconds or since_time may be specified
 * @param "SinceTime" (optional.Time) -  An RFC3339 timestamp from which to show logs. If this value precedes the time an instance was started, only logs since the instance's start time will be returned. If this value is in the future, no logs will be returned.  Only one of since_seconds or since_time may be specified
 * @param "Timestamps" (optional.Bool) -  Whether or not to add an RFC3339 or RFC3339Nano timestamp at the beginning of every line of log output. This defaults to false.
 * @param "TailLines" (optional.String) -  The number of lines from the end of the logs to show. If not specified, logs are shown from the creation of the container or since_seconds or since_time.
 * @param "LimitBytes" (optional.String) -  The number of bytes to read from the server before terminating log output. This may not display a complete final line of logging, and may return slightly more or slightly less than the specified limit.
@return V1LogChunk
*/
func (a *InstanceLogsApiService) GetLogs(ctx _context.Context, stackId string, workloadId string, instanceName string, localVarOptionals *GetLogsOpts) (V1LogChunk, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  V1LogChunk
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/workload/v1/stacks/{stack_id}/workloads/{workload_id}/instances/{instance_name}/logs"
	localVarPath = strings.Replace(localVarPath, "{"+"stack_id"+"}", _neturl.QueryEscape(parameterToString(stackId, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"workload_id"+"}", _neturl.QueryEscape(parameterToString(workloadId, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"instance_name"+"}", _neturl.QueryEscape(parameterToString(instanceName, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.ContainerName.IsSet() {
		localVarQueryParams.Add("container_name", parameterToString(localVarOptionals.ContainerName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Follow.IsSet() {
		localVarQueryParams.Add("follow", parameterToString(localVarOptionals.Follow.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Previous.IsSet() {
		localVarQueryParams.Add("previous", parameterToString(localVarOptionals.Previous.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SinceSeconds.IsSet() {
		localVarQueryParams.Add("since_seconds", parameterToString(localVarOptionals.SinceSeconds.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SinceTime.IsSet() {
		localVarQueryParams.Add("since_time", parameterToString(localVarOptionals.SinceTime.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Timestamps.IsSet() {
		localVarQueryParams.Add("timestamps", parameterToString(localVarOptionals.Timestamps.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TailLines.IsSet() {
		localVarQueryParams.Add("tail_lines", parameterToString(localVarOptionals.TailLines.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.LimitBytes.IsSet() {
		localVarQueryParams.Add("limit_bytes", parameterToString(localVarOptionals.LimitBytes.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v StackpathapiStatus
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v StackpathapiStatus
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
			var v StackpathapiStatus
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}