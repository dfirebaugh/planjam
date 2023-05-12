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
	"io/ioutil"
	"path/filepath"

	"github.com/dfirebaugh/planjam/pkg/plan"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use: "ls [ lane | feature | feature id ]",
	Aliases: []string{
		"show",
		"ls",
		"print",
	},
	Short: "print out of the board",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		isClean, _ := cmd.LocalFlags().GetBool("clean")
		if isClean {
			printBoardClean()
			return
		}
		isASCIIDoc, _ := cmd.LocalFlags().GetBool("asciidoc")
		if isASCIIDoc {
			printASCIIDocTable()
			return
		}
		if len(args) == 0 {
			printBoard()
			return
		}

		if len(args) > 1 {
			return
		}

		listWithArg(args[0])
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolP("clean", "c", false, "prints the clean mdtable with no annotations")
	listCmd.Flags().BoolP("asciidoc", "a", false, "prints an asciidoc with no annotations")
}

func listWithArg(arg string) {
	switch arg {
	case "lane":
		printLanes()
	case "feature":
		printFeatureLabels()
	default:
		b := readBoard()
		ref := b.GetFeatureRef(arg)
		if ref.ID == "" {
			println("could not find an feature of id: ", arg)
			return
		}

		filePath := ref.Feature().FileName(filepath.Join(
			planDir,
			featureDir,
		))
		file, err := ioutil.ReadFile(filePath)
		if err != nil {
			panic(err)
		}

		println(filePath)

		feature, err := plan.GetFeature(file)
		if err != nil {
			println(err.Error())
			return
		}

		println(feature.String())
	}
}

func printLanes() {
	for _, l := range readBoard().Lanes {
		println(l.Label)
	}
}

func printFeatureLabels() {
	for _, l := range readBoard().Lanes {
		for _, f := range l.Features {
			println(f.Label)
		}
	}
}

func printASCIIDocTable() {
	println(readBoard().ASCIIDocTable())
}
