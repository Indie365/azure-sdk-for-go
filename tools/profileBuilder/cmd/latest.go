// +build go1.9

// Copyright © 2018 Microsoft Corporation
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
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/tools/profileBuilder/model"
	"github.com/marstr/randname"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	previewLongName    = "preview"
	previewShortName   = "p"
	previewDefault     = false
	previewDescription = "Include preview API Versions."
)

const (
	rootLongName    = "root"
	rootShortName   = "r"
	rootDescription = "The location of the API Version folders which should be considered for `latest`."
)

var rootDefault = model.DefaultInputRoot()

// latestCmd represents the latest command
var latestCmd = &cobra.Command{
	Use:   "latest",
	Short: "Reflects on the available packages, choosing the most recent ones.",
	Long: `Scans through the availabe API Versions, and chooses only the most 
recent functionality.

By default, this command ignores API versions that are in preview.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		logWriter := ioutil.Discard
		if viper.GetBool("verbose") {
			logWriter = os.Stdout
		}

		outputLog := log.New(logWriter, "[STATUS] ", 0)
		errLog := log.New(os.Stderr, "[ERROR] ", 0)

		packageStrategy := model.LatestStrategy{
			Root:          viper.GetString(rootLongName),
			Predicate:     model.IgnorePreview,
			VerboseOutput: outputLog,
		}

		if viper.GetBool(previewLongName) {
			packageStrategy.Predicate = model.AcceptAll
			outputLog.Println("Using preview versions.")
		}

		model.BuildProfile(
			packageStrategy,
			*nameToUse,
			*outputLocationName,
			outputLog,
			errLog)
	},
}

// This is technical debt. It is being tracked by issue:
// https://github.com/Azure/azure-sdk-for-go/issues/1060
var outputLocationName *string
var nameToUse *string

func init() {
	rootCmd.AddCommand(latestCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// latestCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// latestCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	latestCmd.Flags().BoolP(previewLongName, previewShortName, previewDefault, previewDescription)

	outputLocationName = latestCmd.Flags().StringP(outputLocationLongName, outputLocationShortName, outputLocationDefault, outputLocationDescription)

	nameToUse = latestCmd.Flags().StringP(nameLongName, nameShortName, nameDefault, nameDescription)

	latestCmd.Flags().StringP(rootLongName, rootShortName, rootDefault, rootDescription)

	viper.BindPFlags(latestCmd.Flags())
	viper.SetDefault(nameLongName, randname.Generate())
}
