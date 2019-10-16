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
	"github.com/gorilla/mux"
)

// RegistAPIRouter regist s3 router
func RegistAPIRouter() *mux.Router {
	ParamMux := mux.NewRouter().SkipClean(true)
	apiRouter := ParamMux.NewRoute().PathPrefix("/").Subrouter()

	// object routers
	bucket := apiRouter.PathPrefix("/{bucket}").Subrouter()

	// put object
	bucket.Methods("PUT").Path("/{object:.+}").HandlerFunc(PutObject)
	// get object
	bucket.Methods("GET").Path("/{object:.+}").HandlerFunc(GetObject)
	// delete object
	bucket.Methods("DELETE").Path("/{object:.+}").HandlerFunc(DeleteObject)

	return ParamMux
}
