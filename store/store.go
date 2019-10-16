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

package store

// Object struct
type Object struct {
	BucketName string
	ObjectKey  string
	Meta       string
}

// Store interface
type Store interface {
	PutObject(bucket, object, meta string) error
	DeleteObject(bucket, object string) error
	GetObject(bucket, object string) (obj *Object, err error)
}
