/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"github.com/spf13/cobra"

	"github.com/hi019/fiber-boilerplate/pkg/api"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serves the boilerplate with default settings",
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetString("port")


		dbDialect, _ := cmd.Flags().GetString("dbDialect")
		dbURL, _ := cmd.Flags().GetString("dbURL")

		cfg := &api.Config{Port: port, DriverName: dbDialect, DataSourceName: dbURL}
		app, db, err := api.Start(cfg)
		if err != nil {
			panic(err)
		}

		defer db.Close()

		app.Listen(port)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	serveCmd.PersistentFlags().String("port", ":3000", "Port to run server on")
	serveCmd.PersistentFlags().String("dbDialect", "sqlite3", "Database dialect")
	serveCmd.PersistentFlags().String("dbURL", "file:db.sqlite?_fk=1", "Database url")

}
