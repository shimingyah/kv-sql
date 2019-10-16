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

var kv = &cobra.Command{
	Use:   "kv",
	Short: "Start a kv server.",
	Long:  `Start a kv server.`,
	Run:   runKV,
}

type kvType struct {
	Endpoint string
	Dir      string
	ValSize  int
}

var kvParam kvType

func init() {
	rootCmd.AddCommand(kv)
	kv.PersistentFlags().StringVar(&kvParam.Endpoint, "endpoint", "0.0.0.0:8080", "kv-sql server endpoint")
	kv.PersistentFlags().IntVar(&kvParam.ValSize, "size", 512, "value size")
	kv.PersistentFlags().StringVar(&kvParam.Dir, "dir", "/tmp/rocksdb", "rocksdb dir")
}

func runKV(cmd *cobra.Command, args []string) {
	glog.Infoln("start a kv server")
	runtime.GOMAXPROCS(runtime.NumCPU())

	global.SetStoreType("kv")
	global.ValSize = kvParam.ValSize
	objKV, err := store.NewKVMap(kvParam.Dir)
	if err != nil {
		glog.Fatalf("failed to init rocksdb: %v", err)
	}
	global.KVSQL = objKV

	router := service.RegistAPIRouter()

	if err := http.ListenAndServe(kvParam.Endpoint, router); err != nil {
		glog.Fatalf("service fail to serve: %v", err)
	}
}
