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
	"os"
	"path/filepath"
	"strings"

	"github.com/dfirebaugh/planjam/pkg/config"
	"github.com/dfirebaugh/planjam/pkg/plan"
	"github.com/spf13/cobra"
)

// boardCmd represents the board command
var boardCmd = &cobra.Command{
	Use:   "board [name of board]",
	Short: "select a board to work with",
	Long:  ``,
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		boardFile := filepath.Join(
			filepath.Join(planDir, boardDir),
			args[0]+".yaml")

		conf := config.Config{
			CurrentBoard: boardFile,
		}

		makePlanDir()
		writeConfig(conf)
		if _, err := os.Stat(boardFile); errors.Is(err, os.ErrNotExist) {
			println("empty board created")
			writeBoard(&plan.Board{})
		}

		b := readBoard()
		writeBoard(b)
	},
}

func init() {
	makePlanDir()
	rootCmd.AddCommand(boardCmd)
}

func makePlanDir() error {
	if _, err := os.Stat(planDir); errors.Is(err, os.ErrNotExist) {
		err = os.MkdirAll(planDir, 0755)
		if err != nil {
			return err
		}
	}
	if _, err := os.Stat(filepath.Join(planDir, boardDir)); errors.Is(err, os.ErrNotExist) {
		err = os.MkdirAll(filepath.Join(planDir, boardDir), 0755)
		if err != nil {
			return err
		}
	}

	if _, err := os.Stat(filepath.Join(planDir, featureDir)); errors.Is(err, os.ErrNotExist) {
		err = os.MkdirAll(filepath.Join(planDir, featureDir), 0755)
		if err != nil {
			return err
		}
	}

	configFile := filepath.Join(planDir, defaultConfig)
	if _, err := os.Stat(configFile); errors.Is(err, os.ErrNotExist) {
		file, err := os.Create(configFile)
		if err != nil {
			return err
		}
		defer file.Close()

		if err := os.Chmod(configFile, 0644); err != nil {
			return err
		}
	}

	return nil
}

func readBoard() *plan.Board {
	c := readConfig()
	file, err := os.ReadFile(c.CurrentBoard)
	if err != nil {
		panic(err)
	}

	p, err := plan.GetBoard(file)
	if err != nil {
		panic(err)
	}
	return p
}

func getBoardName() string {
	c := readConfig()
	return strings.TrimSuffix(filepath.Base(c.CurrentBoard), filepath.Ext(c.CurrentBoard))
}

func writeBoard(b *plan.Board) {
	c := readConfig()
	os.Remove(c.CurrentBoard)
	boardFile, err := os.OpenFile(c.CurrentBoard, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer boardFile.Close()

	b.Write(boardFile)
	printBoard()

	for _, l := range b.Lanes {
		for _, fr := range l.Features {
			writeFeature(fr.Feature())
		}
	}
}

func writeFeature(f plan.Feature) {
	featureFilePath := f.FileName(filepath.Join(
		planDir,
		featureDir,
	))
	os.Remove(featureFilePath)

	featureFile, err := os.OpenFile(featureFilePath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer featureFile.Close()

	f.Write(featureFile)
}

func printBoardWithIssueIDs() {
	println("board: ", getBoardName())
	b := readBoard()
	b.ShowFeatureIDs = true
	b.ShowLaneNumbers = false
	println(b.String())
}

func printBoardWithLaneNumbers() {
	println("board: ", getBoardName())
	b := readBoard()
	b.ShowLaneNumbers = true
	b.ShowFeatureIDs = false
	println(b.String())
}

func printBoard() {
	printBoardWithIssueIDs()
}

func printBoardClean() {
	println("board: ", getBoardName())
	b := readBoard()
	println(b.String())
}
