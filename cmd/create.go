/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"kubeSecretManage/services"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "creates a kubernetes secret",
	Long: `creates kubernetes secrets in the namespace specified. Requires value of the secret as an argument`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Printf("ns->%s\ndata->%s", ns, data)
		err := services.CreateSecret(ns, "testKey", data, "testName")
		if err != nil {
			logrus.WithError(err).Fatal("Error performing create secret action")
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//createCmd.Flags().BoolP("namespace", "ns", false, "Help message for toggle")
	//createCmd.Flags().StringP("namespace", "n", "", "Namespace where secret is to be stored")
	//createCmd.Flags().StringP("data", "d", "", "Data to be stored as a secret")
}
