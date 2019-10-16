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
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"
)

// ErrObjectNotExist object is not exist
var ErrObjectNotExist = errors.New("object is not exit")

// ObjectMap object op for mysql
type ObjectMap struct {
	db *sql.DB
}

// NewObjectMap return object map
func NewObjectMap(addr string, maxConn, maxIdle int) (*ObjectMap, error) {
	db, err := sql.Open("mysql", addr)
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(maxIdle)
	db.SetMaxOpenConns(maxConn)
	if err = db.Ping(); err != nil {
		return nil, errors.New("ping db fail:" + err.Error())
	}

	return &ObjectMap{
		db: db,
	}, nil
}

// PutObject put object
func (omap *ObjectMap) PutObject(bucket, object, meta string) error {
	sql := "insert into object(bucket_name, object_key, meta) values(?, ?, ?)"
	result, err := omap.db.Exec(sql, bucket, object, meta)
	if err != nil {
		return err
	}

	_, err = result.RowsAffected()
	return err
}

// DeleteObject delete object
func (omap *ObjectMap) DeleteObject(bucket, object string) error {
	sql := "delete from object where bucket_name=? and object_key=?"
	result, err := omap.db.Exec(sql, bucket, object)
	if err != nil {
		return err
	}

	_, err = result.RowsAffected()
	return err
}

// GetObject get object
func (omap *ObjectMap) GetObject(bucket, object string) (obj *Object, err error) {
	var rows *sql.Rows
	sql := "select bucket_name, object_key, meta from object where bucket_name=? and object_key=?"
	rows, err = omap.db.Query(sql, bucket, object)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	objects := []*Object{}
	for rows.Next() {
		var res *Object

		res, err = omap.NewQueryObject(rows)
		if err != nil {
			return nil, err
		}
		objects = append(objects, res)
	}
	if len(objects) > 0 {
		return objects[0], nil
	}
	return nil, ErrObjectNotExist
}

// NewQueryObject new query object
func (omap *ObjectMap) NewQueryObject(rows *sql.Rows) (object *Object, err error) {
	object = new(Object)
	err = rows.Scan(
		&object.BucketName,
		&object.ObjectKey,
		&object.Meta,
	)

	return object, err
}
