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

import (
	"fmt"

	"github.com/tecbot/gorocksdb"
)

// KVMap struct
type KVMap struct {
	DB *gorocksdb.DB
}

// NewKVMap return rocksdb db
func NewKVMap(dir string) (*KVMap, error) {
	opts := gorocksdb.NewDefaultOptions()
	opts.SetCreateIfMissing(true)
	db, err := gorocksdb.OpenDb(opts, dir)
	return &KVMap{DB: db}, err
}

// PutObject put object
func (kvmap *KVMap) PutObject(bucket, object, meta string) error {
	key := fmt.Sprintf("%s_%s", bucket, object)
	woption := gorocksdb.NewDefaultWriteOptions()
	return kvmap.DB.Put(woption, []byte(key), []byte(meta))
}

// DeleteObject delete object
func (kvmap *KVMap) DeleteObject(bucket, object string) error {
	key := fmt.Sprintf("%s_%s", bucket, object)
	woption := gorocksdb.NewDefaultWriteOptions()
	return kvmap.DB.Delete(woption, []byte(key))
}

// GetObject get object
func (kvmap *KVMap) GetObject(bucket, object string) (obj *Object, err error) {
	key := fmt.Sprintf("%s_%s", bucket, object)
	roption := gorocksdb.NewDefaultReadOptions()
	value, err := kvmap.DB.Get(roption, []byte(key))
	if err != nil {
		return nil, err
	}
	defer value.Free()

	return &Object{
		BucketName: bucket,
		ObjectKey:  object,
		Meta:       string(value.Data()),
	}, nil
}
