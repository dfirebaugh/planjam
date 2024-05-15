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
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// moveCmd represents the move command
var moveCmd = &cobra.Command{
	Use:     "mv",
	Aliases: []string{"move"},
	Short:   "move a feature or lane",
	Long:    ``,
	Args:    cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			return
		}
		preMoveFeature(args[0])
	},
}

func init() {
	rootCmd.AddCommand(moveCmd)
}

func preMoveFeature(arg string) {
	b := readBoard()
	feature := b.GetFeatureRef(arg)

	if feature.ID == "" {
		fmt.Printf("Feature with ID [%s] does not exist.\n", arg)
		return
	}

	printBoardWithLaneNumbers()
	fmt.Printf("Which lane should we move [%s] to? ", feature.Label)
	laneIndex, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	li, err := strconv.Atoi(strings.TrimSpace(laneIndex))
	if err != nil {
		fmt.Printf("Error parsing lane index: %s\n", err)
		return
	}

	if li < 0 || li >= len(b.Lanes) {
		fmt.Printf("Invalid lane index: %d. Lane does not exist.\n", li)
		return
	}

	lane := b.GetLane(li)

	fmt.Printf("Moving feature '%s' to lane '%s'...\n", feature.Label, lane.Label)

	moveFeature(feature.ID, lane.Label)
}

func moveFeature(id string, lane string) {
	b := readBoard()
	b.Move(id, lane)
	writeBoard(b)
}
