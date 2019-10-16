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
	"fmt"

	"github.com/spf13/cobra"
)

// Version the version of openkvss
const Version = "0.0.1"

var version = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of kv-sql",
	Long:  `All software has versions. This is kv-sql's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("kv-sql version %s\n", Version)
	},
}

func init() {
	rootCmd.AddCommand(version)
}
