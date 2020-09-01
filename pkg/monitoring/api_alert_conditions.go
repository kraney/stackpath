/*
 * Monitoring
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package monitoring

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// AlertConditionsApiService AlertConditionsApi service
type AlertConditionsApiService service

type apiBatchDeleteAlertConditionsRequest struct {
	ctx _context.Context
	apiService *AlertConditionsApiService
	stackId string
	monitorId string
	v2BatchDeleteAlertConditionsRequest *V2BatchDeleteAlertConditionsRequest
}


func (r apiBatchDeleteAlertConditionsRequest) V2BatchDeleteAlertConditionsRequest(v2BatchDeleteAlertConditionsRequest V2BatchDeleteAlertConditionsRequest) apiBatchDeleteAlertConditionsRequest {
	r.v2BatchDeleteAlertConditionsRequest = &v2BatchDeleteAlertConditionsRequest
	return r
}

/*
BatchDeleteAlertConditions Delete multiple alert conditions
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param stackId A stack ID or slug
 * @param monitorId A monitor ID
@return apiBatchDeleteAlertConditionsRequest
*/
func (a *AlertConditionsApiService) BatchDeleteAlertConditions(ctx _context.Context, stackId string, monitorId string) apiBatchDeleteAlertConditionsRequest {
	return apiBatchDeleteAlertConditionsRequest{
		apiService: a,
		ctx: ctx,
		stackId: stackId,
		monitorId: monitorId,
	}
}

/*
Execute executes the request

*/
func (r apiBatchDeleteAlertConditionsRequest) Execute() (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "AlertConditionsApiService.BatchDeleteAlertConditions")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/monitoring/v2/stacks/{stack_id}/monitors/{monitor_id}/alerts/conditions/batch_delete"
	localVarPath = strings.Replace(localVarPath, "{"+"stack_id"+"}", _neturl.QueryEscape(parameterToString(r.stackId, "")) , -1)
	localVarPath = strings.Replace(localVarPath, "{"+"monitor_id"+"}", _neturl.QueryEscape(parameterToString(r.monitorId, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	
	
	
	if r.v2BatchDeleteAlertConditionsRequest == nil {
		return nil, reportError("v2BatchDeleteAlertConditionsRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.v2BatchDeleteAlertConditionsRequest
	req, err := r.apiService.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := r.apiService.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v StackpathapiStatus
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v StackpathapiStatus
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
			var v StackpathapiStatus
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
type apiCreateAlertConditionRequest struct {
	ctx _context.Context
	apiService *AlertConditionsApiService
	stackId string
	monitorId string
	v2CreateAlertConditionRequest *V2CreateAlertConditionRequest
}


func (r apiCreateAlertConditionRequest) V2CreateAlertConditionRequest(v2CreateAlertConditionRequest V2CreateAlertConditionRequest) apiCreateAlertConditionRequest {
	r.v2CreateAlertConditionRequest = &v2CreateAlertConditionRequest
	return r
}

/*
CreateAlertCondition Create an alert condition
An alert condition defines when to be alerted by a change in the monitored service.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param stackId A stack ID or slug
 * @param monitorId A monitor ID
@return apiCreateAlertConditionRequest
*/
func (a *AlertConditionsApiService) CreateAlertCondition(ctx _context.Context, stackId string, monitorId string) apiCreateAlertConditionRequest {
	return apiCreateAlertConditionRequest{
		apiService: a,
		ctx: ctx,
		stackId: stackId,
		monitorId: monitorId,
	}
}

/*
Execute executes the request
 @return V2CreateAlertConditionResponse
*/
func (r apiCreateAlertConditionRequest) Execute() (V2CreateAlertConditionResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  V2CreateAlertConditionResponse
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "AlertConditionsApiService.CreateAlertCondition")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/monitoring/v2/stacks/{stack_id}/monitors/{monitor_id}/alerts/conditions"
	localVarPath = strings.Replace(localVarPath, "{"+"stack_id"+"}", _neturl.QueryEscape(parameterToString(r.stackId, "")) , -1)
	localVarPath = strings.Replace(localVarPath, "{"+"monitor_id"+"}", _neturl.QueryEscape(parameterToString(r.monitorId, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	
	
	
	if r.v2CreateAlertConditionRequest == nil {
		return localVarReturnValue, nil, reportError("v2CreateAlertConditionRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.v2CreateAlertConditionRequest
	req, err := r.apiService.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := r.apiService.client.callAPI(req)
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
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v StackpathapiStatus
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
			var v StackpathapiStatus
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = r.apiService.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
type apiDeleteAlertConditionRequest struct {
	ctx _context.Context
	apiService *AlertConditionsApiService
	stackId string
	monitorId string
	conditionId string
}


/*
DeleteAlertCondition Delete an alert condition
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param stackId A stack ID or slug
 * @param monitorId A monitor ID
 * @param conditionId A monitoring alert condition ID
@return apiDeleteAlertConditionRequest
*/
func (a *AlertConditionsApiService) DeleteAlertCondition(ctx _context.Context, stackId string, monitorId string, conditionId string) apiDeleteAlertConditionRequest {
	return apiDeleteAlertConditionRequest{
		apiService: a,
		ctx: ctx,
		stackId: stackId,
		monitorId: monitorId,
		conditionId: conditionId,
	}
}

/*
Execute executes the request

*/
func (r apiDeleteAlertConditionRequest) Execute() (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "AlertConditionsApiService.DeleteAlertCondition")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/monitoring/v2/stacks/{stack_id}/monitors/{monitor_id}/alerts/conditions/{condition_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"stack_id"+"}", _neturl.QueryEscape(parameterToString(r.stackId, "")) , -1)
	localVarPath = strings.Replace(localVarPath, "{"+"monitor_id"+"}", _neturl.QueryEscape(parameterToString(r.monitorId, "")) , -1)
	localVarPath = strings.Replace(localVarPath, "{"+"condition_id"+"}", _neturl.QueryEscape(parameterToString(r.conditionId, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	
	
	

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
	req, err := r.apiService.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := r.apiService.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v StackpathapiStatus
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v StackpathapiStatus
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
			var v StackpathapiStatus
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
type apiDisableAlertConditionRequest struct {
	ctx _context.Context
	apiService *AlertConditionsApiService
	stackId string
	monitorId string
	conditionId string
}


/*
DisableAlertCondition Disable an alert condition
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param stackId A stack ID or slug
 * @param monitorId A monitor ID
 * @param conditionId A monitoring alert condition ID
@return apiDisableAlertConditionRequest
*/
func (a *AlertConditionsApiService) DisableAlertCondition(ctx _context.Context, stackId string, monitorId string, conditionId string) apiDisableAlertConditionRequest {
	return apiDisableAlertConditionRequest{
		apiService: a,
		ctx: ctx,
		stackId: stackId,
		monitorId: monitorId,
		conditionId: conditionId,
	}
}

/*
Execute executes the request

*/
func (r apiDisableAlertConditionRequest) Execute() (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "AlertConditionsApiService.DisableAlertCondition")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/monitoring/v2/stacks/{stack_id}/monitors/{monitor_id}/alerts/conditions/{condition_id}/disable"
	localVarPath = strings.Replace(localVarPath, "{"+"stack_id"+"}", _neturl.QueryEscape(parameterToString(r.stackId, "")) , -1)
	localVarPath = strings.Replace(localVarPath, "{"+"monitor_id"+"}", _neturl.QueryEscape(parameterToString(r.monitorId, "")) , -1)
	localVarPath = strings.Replace(localVarPath, "{"+"condition_id"+"}", _neturl.QueryEscape(parameterToString(r.conditionId, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	
	
	

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
	req, err := r.apiService.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := r.apiService.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v StackpathapiStatus
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v StackpathapiStatus
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
			var v StackpathapiStatus
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
type apiEnableAlertConditionRequest struct {
	ctx _context.Context
	apiService *AlertConditionsApiService
	stackId string
	monitorId string
	conditionId string
}


/*
EnableAlertCondition Enable an alert condition
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param stackId A stack ID or slug
 * @param monitorId A monitor ID
 * @param conditionId A monitoring alert condition ID
@return apiEnableAlertConditionRequest
*/
func (a *AlertConditionsApiService) EnableAlertCondition(ctx _context.Context, stackId string, monitorId string, conditionId string) apiEnableAlertConditionRequest {
	return apiEnableAlertConditionRequest{
		apiService: a,
		ctx: ctx,
		stackId: stackId,
		monitorId: monitorId,
		conditionId: conditionId,
	}
}

/*
Execute executes the request

*/
func (r apiEnableAlertConditionRequest) Execute() (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "AlertConditionsApiService.EnableAlertCondition")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/monitoring/v2/stacks/{stack_id}/monitors/{monitor_id}/alerts/conditions/{condition_id}/enable"
	localVarPath = strings.Replace(localVarPath, "{"+"stack_id"+"}", _neturl.QueryEscape(parameterToString(r.stackId, "")) , -1)
	localVarPath = strings.Replace(localVarPath, "{"+"monitor_id"+"}", _neturl.QueryEscape(parameterToString(r.monitorId, "")) , -1)
	localVarPath = strings.Replace(localVarPath, "{"+"condition_id"+"}", _neturl.QueryEscape(parameterToString(r.conditionId, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	
	
	

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
	req, err := r.apiService.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := r.apiService.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v StackpathapiStatus
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v StackpathapiStatus
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
			var v StackpathapiStatus
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
type apiGetAlertConditionRequest struct {
	ctx _context.Context
	apiService *AlertConditionsApiService
	stackId string
	monitorId string
	conditionId string
}


/*
GetAlertCondition Get an alert condition
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param stackId A stack ID or slug
 * @param monitorId A monitor ID
 * @param conditionId A monitoring alert condition ID
@return apiGetAlertConditionRequest
*/
func (a *AlertConditionsApiService) GetAlertCondition(ctx _context.Context, stackId string, monitorId string, conditionId string) apiGetAlertConditionRequest {
	return apiGetAlertConditionRequest{
		apiService: a,
		ctx: ctx,
		stackId: stackId,
		monitorId: monitorId,
		conditionId: conditionId,
	}
}

/*
Execute executes the request
 @return V2GetAlertConditionResponse
*/
func (r apiGetAlertConditionRequest) Execute() (V2GetAlertConditionResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  V2GetAlertConditionResponse
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "AlertConditionsApiService.GetAlertCondition")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/monitoring/v2/stacks/{stack_id}/monitors/{monitor_id}/alerts/conditions/{condition_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"stack_id"+"}", _neturl.QueryEscape(parameterToString(r.stackId, "")) , -1)
	localVarPath = strings.Replace(localVarPath, "{"+"monitor_id"+"}", _neturl.QueryEscape(parameterToString(r.monitorId, "")) , -1)
	localVarPath = strings.Replace(localVarPath, "{"+"condition_id"+"}", _neturl.QueryEscape(parameterToString(r.conditionId, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	
	
	

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
	req, err := r.apiService.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := r.apiService.client.callAPI(req)
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
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v StackpathapiStatus
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
			var v StackpathapiStatus
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = r.apiService.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
type apiGetAlertConditionsRequest struct {
	ctx _context.Context
	apiService *AlertConditionsApiService
	stackId string
	monitorId string
	pageRequestFirst *string
	pageRequestAfter *string
	pageRequestFilter *string
	pageRequestSortBy *string
}


func (r apiGetAlertConditionsRequest) PageRequestFirst(pageRequestFirst string) apiGetAlertConditionsRequest {
	r.pageRequestFirst = &pageRequestFirst
	return r
}

func (r apiGetAlertConditionsRequest) PageRequestAfter(pageRequestAfter string) apiGetAlertConditionsRequest {
	r.pageRequestAfter = &pageRequestAfter
	return r
}

func (r apiGetAlertConditionsRequest) PageRequestFilter(pageRequestFilter string) apiGetAlertConditionsRequest {
	r.pageRequestFilter = &pageRequestFilter
	return r
}

func (r apiGetAlertConditionsRequest) PageRequestSortBy(pageRequestSortBy string) apiGetAlertConditionsRequest {
	r.pageRequestSortBy = &pageRequestSortBy
	return r
}

/*
GetAlertConditions Get all alert conditions
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param stackId A stack ID or slug
 * @param monitorId A monitor ID
@return apiGetAlertConditionsRequest
*/
func (a *AlertConditionsApiService) GetAlertConditions(ctx _context.Context, stackId string, monitorId string) apiGetAlertConditionsRequest {
	return apiGetAlertConditionsRequest{
		apiService: a,
		ctx: ctx,
		stackId: stackId,
		monitorId: monitorId,
	}
}

/*
Execute executes the request
 @return V2GetAlertConditionsResponse
*/
func (r apiGetAlertConditionsRequest) Execute() (V2GetAlertConditionsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  V2GetAlertConditionsResponse
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "AlertConditionsApiService.GetAlertConditions")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/monitoring/v2/stacks/{stack_id}/monitors/{monitor_id}/alerts/conditions"
	localVarPath = strings.Replace(localVarPath, "{"+"stack_id"+"}", _neturl.QueryEscape(parameterToString(r.stackId, "")) , -1)
	localVarPath = strings.Replace(localVarPath, "{"+"monitor_id"+"}", _neturl.QueryEscape(parameterToString(r.monitorId, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	
	
				
	if r.pageRequestFirst != nil {
		localVarQueryParams.Add("page_request.first", parameterToString(*r.pageRequestFirst, ""))
	}
	if r.pageRequestAfter != nil {
		localVarQueryParams.Add("page_request.after", parameterToString(*r.pageRequestAfter, ""))
	}
	if r.pageRequestFilter != nil {
		localVarQueryParams.Add("page_request.filter", parameterToString(*r.pageRequestFilter, ""))
	}
	if r.pageRequestSortBy != nil {
		localVarQueryParams.Add("page_request.sort_by", parameterToString(*r.pageRequestSortBy, ""))
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
	req, err := r.apiService.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := r.apiService.client.callAPI(req)
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
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v StackpathapiStatus
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
			var v StackpathapiStatus
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = r.apiService.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
type apiUpdateAlertConditionRequest struct {
	ctx _context.Context
	apiService *AlertConditionsApiService
	stackId string
	monitorId string
	conditionId string
	v2UpdateAlertConditionRequest *V2UpdateAlertConditionRequest
}


func (r apiUpdateAlertConditionRequest) V2UpdateAlertConditionRequest(v2UpdateAlertConditionRequest V2UpdateAlertConditionRequest) apiUpdateAlertConditionRequest {
	r.v2UpdateAlertConditionRequest = &v2UpdateAlertConditionRequest
	return r
}

/*
UpdateAlertCondition Update an alert condition
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param stackId A stack ID or slug
 * @param monitorId A monitor ID
 * @param conditionId A monitoring alert condition ID
@return apiUpdateAlertConditionRequest
*/
func (a *AlertConditionsApiService) UpdateAlertCondition(ctx _context.Context, stackId string, monitorId string, conditionId string) apiUpdateAlertConditionRequest {
	return apiUpdateAlertConditionRequest{
		apiService: a,
		ctx: ctx,
		stackId: stackId,
		monitorId: monitorId,
		conditionId: conditionId,
	}
}

/*
Execute executes the request
 @return V2UpdateAlertConditionResponse
*/
func (r apiUpdateAlertConditionRequest) Execute() (V2UpdateAlertConditionResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  V2UpdateAlertConditionResponse
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "AlertConditionsApiService.UpdateAlertCondition")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/monitoring/v2/stacks/{stack_id}/monitors/{monitor_id}/alerts/conditions/{condition_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"stack_id"+"}", _neturl.QueryEscape(parameterToString(r.stackId, "")) , -1)
	localVarPath = strings.Replace(localVarPath, "{"+"monitor_id"+"}", _neturl.QueryEscape(parameterToString(r.monitorId, "")) , -1)
	localVarPath = strings.Replace(localVarPath, "{"+"condition_id"+"}", _neturl.QueryEscape(parameterToString(r.conditionId, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	
	
	
	
	if r.v2UpdateAlertConditionRequest == nil {
		return localVarReturnValue, nil, reportError("v2UpdateAlertConditionRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.v2UpdateAlertConditionRequest
	req, err := r.apiService.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := r.apiService.client.callAPI(req)
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
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v StackpathapiStatus
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
			var v StackpathapiStatus
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = r.apiService.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}