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

package cmd

import (
	"net/http"
	"runtime"

	"github.com/golang/glog"
	"github.com/shimingyah/kv-sql/global"
	"github.com/shimingyah/kv-sql/service"
	"github.com/shimingyah/kv-sql/store"
	"github.com/spf13/cobra"
)

var sql = &cobra.Command{
	Use:   "sql",
	Short: "Start a sql server.",
	Long:  `Start a sql server.`,
	Run:   runSQL,
}

type sqlType struct {
	Endpoint string
	Addr     string
	ValSize  int
}

var sqlParam sqlType

func init() {
	rootCmd.AddCommand(sql)
	sql.PersistentFlags().StringVar(&sqlParam.Endpoint, "endpoint", "0.0.0.0:8080", "kv-sql server endpoint")
	sql.PersistentFlags().IntVar(&sqlParam.ValSize, "size", 512, "value size")
	sql.PersistentFlags().StringVar(&sqlParam.Addr, "mysql", "root:123456@tcp(127.0.0.1:3306)/kv?charset=utf8", "mysql address")
}

func runSQL(cmd *cobra.Command, args []string) {
	glog.Infoln("start a sql server")
	runtime.GOMAXPROCS(runtime.NumCPU())

	global.SetStoreType("sql")
	global.ValSize = sqlParam.ValSize
	objSQL, err := store.NewObjectMap(sqlParam.Addr, 200, 10)
	if err != nil {
		glog.Fatalf("failed to init mysql: %v", err)
	}
	global.ObjectSQL = objSQL

	router := service.RegistAPIRouter()

	if err := http.ListenAndServe(sqlParam.Endpoint, router); err != nil {
		glog.Fatalf("service fail to serve: %v", err)
	}
}
