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
	"fmt"
	"net/http"
)

func writeErrorResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Server", "kv-sql-server")
	apiError := getAPIError(ErrInternalError)
	errorResponse := getAPIErrorResponse(apiError, r.URL.Path, "requestID")
	encodedErrorResponse := encodeResponse(errorResponse)

	if encodedErrorResponse != nil {
		w.Header().Set("Content-Type", "application/xml")
		w.Header().Set("Content-Length", fmt.Sprintf("%d", len(encodedErrorResponse)))
		w.WriteHeader(apiError.HTTPStatusCode)
		w.Write(encodedErrorResponse)
		w.(http.Flusher).Flush()
	}
}

func writeSuccessResponse(w http.ResponseWriter, r *http.Request, data []byte) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Server", "kv-sql-server")

	if data != nil {
		w.Header().Set("Content-Type", "application/xml")
		w.Header().Set("Content-Length", fmt.Sprintf("%d", len(data)))
		w.WriteHeader(http.StatusOK)
		w.Write(data)
		w.(http.Flusher).Flush()
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
