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
	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:     "rm",
	Aliases: []string{"remove"},
	Short:   "remove a feature, lane, or field",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Usage()
			return
		}
	},
}

var removeLaneCmd = &cobra.Command{
	Use:     "lane",
	Aliases: []string{"l"},
	Short:   "remove lane",
	Args:    cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		removeLane(args[0])
	},
}

var removeFeatureCmd = &cobra.Command{
	Use:     "feature",
	Aliases: []string{"ft", "f", "feat"},
	Short:   "remove feature",
	Args:    cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		removeFeature(args[0])
	},
}

var removeFieldCmd = &cobra.Command{
	Use:     "field",
	Aliases: []string{"fld", "attr"},
	Short:   "remove field",
	Args:    cobra.MatchAll(cobra.ExactArgs(2), cobra.OnlyValidArgs),
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		removeField(args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
	removeCmd.AddCommand(removeLaneCmd)
	removeCmd.AddCommand(removeFeatureCmd)
	removeCmd.AddCommand(removeFieldCmd)
}

func removeField(id string, label string) {
	b := readBoard()
	f := b.GetFeatureRef(id).Feature()

	f.RemoveField(label)

	writeFeature(f)
}

func removeFeature(id string) {
	b := readBoard()
	b.RemoveFeature(id)
	writeBoard(b)
}

func removeLane(label string) {
	b := readBoard()

	for i, l := range b.Lanes {
		if l.Label != label {
			continue
		}

		b.Lanes = append(b.Lanes[:i], b.Lanes[i+1:]...)
	}

	writeBoard(b)
}
