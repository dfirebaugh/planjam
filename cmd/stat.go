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
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// statCmd represents the stat command
var statCmd = &cobra.Command{
	Use:   "stat",
	Short: "print stats",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		printStats()
	},
}

func init() {
	rootCmd.AddCommand(statCmd)
}

const (
	barWidth = 50
	green    = "\033[32m"
	white    = "\033[37m"
	reset    = "\033[0m"
)

func printStats() {
	b := readBoard()
	total := 0

	for _, l := range b.Lanes {
		total += len(l.Features)
	}
	if total == 0 {
		total = 1
	}
	for _, l := range b.Lanes {
		completed := (len(l.Features) * barWidth) / total
		remaining := barWidth - completed

		fmt.Printf(
			"# %s \n- [%d]: %s%s%s\n",
			l.Label,
			len(l.Features),
			green+strings.Repeat("█", completed)+reset,
			white+strings.Repeat("█", remaining)+reset,
			reset,
		)
	}
	println()
}
