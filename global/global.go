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

package global

import "github.com/shimingyah/kv-sql/store"

// StoreType save meta
type StoreType int

const (
	// None not save meta
	None StoreType = iota
	// KV save meta by rocksdb
	KV
	// SQL save meta by mysql
	SQL
)

var backend StoreType

var (
	// ObjectSQL mysql backend
	ObjectSQL *store.ObjectMap

	// ValSize value size
	ValSize int
)

// SetStoreType set save meta type
func SetStoreType(storeType string) {
	switch storeType {
	case "kv":
		backend = KV
	case "sql":
		backend = SQL
	}
}

// GetStoreType return save meta type
func GetStoreType() string {
	switch backend {
	case None:
		return "none"
	case KV:
		return "kv"
	case SQL:
		return "sql"
	default:
		return "none"
	}
}

// IsKV is kv
func IsKV() bool {
	return backend == KV
}

// IsSQL is sql
func IsSQL() bool {
	return backend == SQL
}
