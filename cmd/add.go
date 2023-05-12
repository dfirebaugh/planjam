/*
Copyright © 2023 Dustin Firebaugh<dafirebaugh@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a feature, lane, or field to the board",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Usage()
		}
	},
}

var addLaneCmd = &cobra.Command{
	Use:   "lane [label]",
	Short: "Add a lane to the currently selected board",
	Long:  ``,
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		addLane(args[0])
	},
}

var addFeatureCmd = &cobra.Command{
	Use:   "feature [label]",
	Short: "Add a feature to the first lane",
	Long:  ``,
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		if err := addFeature(args[0]); err != nil {
			println(err.Error())
		}
	},
}

var addFieldCmd = &cobra.Command{
	Use:   "field [feature id] [field label] [field value]",
	Short: "Add a field to a feature",
	Long:  ``,
	Args:  cobra.MatchAll(cobra.ExactArgs(3), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		if err := addField(args[0], args[1], args[2]); err != nil {
			println(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.AddCommand(addLaneCmd)
	addCmd.AddCommand(addFeatureCmd)
	addCmd.AddCommand(addFieldCmd)
}

func addFeature(label string) error {
	b := readBoard()

	if len(b.Lanes) == 0 {
		return errors.New("you must first create a lane (e.g. `planjam add lane todo`)")
	}

	b.AddFeature(label)

	writeBoard(b)
	return nil
}

func addField(id string, label string, value string) error {
	b := readBoard()

	feature := b.GetFeatureRef(id).Feature()
	feature.AddField(label, value)

	writeFeature(feature)

	return nil
}

func addLane(label string) {
	b := readBoard().AddLane(label)
	writeBoard(b)
}
