// Copyright 2019 shimingyah. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// ee the License for the specific language governing permissions and
// limitations under the License.

package service

import (
	"bytes"
	"encoding/xml"
	"net/http"
)

// APIError for s3
type APIError struct {
	Code           string
	Description    string
	HTTPStatusCode int
}

// APIErrorResponse for s3
type APIErrorResponse struct {
	XMLName   xml.Name `xml:"Error" json:"-"`
	Code      string
	Message   string
	Resource  string
	RequestID string `xml:"RequestId"`
}

// APIErrorCode type of error status.
type APIErrorCode string

// Error codes, non exhaustive list - http://docs.aws.amazon.com/AmazonS3/latest/API/ErrorResponses.html
const (
	ErrNone          APIErrorCode = ""
	ErrInternalError              = "InternalError"
)

// error code to APIError structure, these fields carry respective
// descriptions for all the error responses.
var errorCodeResponse = map[APIErrorCode]APIError{
	ErrNone: {
		Code:           "Success",
		Description:    "Success",
		HTTPStatusCode: http.StatusOK,
	},
	ErrInternalError: {
		Code:           "InternalError",
		Description:    "We encountered an internal error, please try again.",
		HTTPStatusCode: http.StatusInternalServerError,
	},
}

func getAPIError(code APIErrorCode) APIError {
	return errorCodeResponse[code]
}

// getErrorResponse gets in standard error and resource value and
// provides a encodable populated response values
func getAPIErrorResponse(err APIError, resource string, requestId string) APIErrorResponse {
	return APIErrorResponse{
		Code:      err.Code,
		Message:   err.Description,
		Resource:  resource,
		RequestID: requestId,
	}
}

func encodeResponse(response interface{}) []byte {
	return encodeXMLResponse(response)
}

// encode response body by xml.
func encodeXMLResponse(response interface{}) []byte {
	var bytesBuffer bytes.Buffer
	bytesBuffer.WriteString(xml.Header)
	e := xml.NewEncoder(&bytesBuffer)
	e.Encode(response)
	return bytesBuffer.Bytes()
}
