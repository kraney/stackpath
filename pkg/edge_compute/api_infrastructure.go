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
)

// Linger please
var (
	_ _context.Context
)

// InfrastructureApiService InfrastructureApi service
type InfrastructureApiService service

type apiGetLocationsRequest struct {
	ctx _context.Context
	apiService *InfrastructureApiService
	pageRequestFirst *string
	pageRequestAfter *string
	pageRequestFilter *string
	pageRequestSortBy *string
}


func (r apiGetLocationsRequest) PageRequestFirst(pageRequestFirst string) apiGetLocationsRequest {
	r.pageRequestFirst = &pageRequestFirst
	return r
}

func (r apiGetLocationsRequest) PageRequestAfter(pageRequestAfter string) apiGetLocationsRequest {
	r.pageRequestAfter = &pageRequestAfter
	return r
}

func (r apiGetLocationsRequest) PageRequestFilter(pageRequestFilter string) apiGetLocationsRequest {
	r.pageRequestFilter = &pageRequestFilter
	return r
}

func (r apiGetLocationsRequest) PageRequestSortBy(pageRequestSortBy string) apiGetLocationsRequest {
	r.pageRequestSortBy = &pageRequestSortBy
	return r
}

/*
GetLocations Get compute locations
Retrieve information about the StackPath edge network that can host a compute workload
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return apiGetLocationsRequest
*/
func (a *InfrastructureApiService) GetLocations(ctx _context.Context) apiGetLocationsRequest {
	return apiGetLocationsRequest{
		apiService: a,
		ctx: ctx,
	}
}

/*
Execute executes the request
 @return V1GetLocationsResponse
*/
func (r apiGetLocationsRequest) Execute() (V1GetLocationsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  V1GetLocationsResponse
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "InfrastructureApiService.GetLocations")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/workload/v1/locations"

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
