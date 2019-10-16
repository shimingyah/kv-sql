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
	"net/http"

	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"github.com/shimingyah/kv-sql/global"
	"github.com/shimingyah/kv-sql/pkg"
)

// PutObject put object
func PutObject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bucket := vars["bucket"]
	object := vars["object"]
	meta := pkg.RandomString(global.ValSize)

	if err := global.KVSQL.PutObject(bucket, object, meta); err != nil {
		glog.Errorf("failed to put object, backend: %s, error: %v", global.GetStoreType(), err)
		writeErrorResponse(w, r)
		return
	}

	writeSuccessResponse(w, r, nil)
}

// GetObject get object
func GetObject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bucket := vars["bucket"]
	object := vars["object"]

	obj, err := global.KVSQL.GetObject(bucket, object)
	if err != nil {
		glog.Errorf("failed to get object, backend: %s, error: %v", global.GetStoreType(), err)
		writeErrorResponse(w, r)
		return
	}

	writeSuccessResponse(w, r, []byte(obj.Meta))

}

// DeleteObject delete object
func DeleteObject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bucket := vars["bucket"]
	object := vars["object"]

	if err := global.KVSQL.DeleteObject(bucket, object); err != nil {
		glog.Errorf("failed to delete object, backend: %s, error: %v", global.GetStoreType(), err)
		writeErrorResponse(w, r)
		return
	}

	writeSuccessResponse(w, r, nil)
}
