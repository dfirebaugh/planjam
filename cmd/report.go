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
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/dfirebaugh/planjam/pkg/imggen"
	"github.com/spf13/cobra"
)

// reportCmd represents the report command
var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "generate a report of your plans",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		genReport()
	},
}

func init() {
	rootCmd.AddCommand(reportCmd)
}

func genReport() {
	makeASCIIDoc(genASCIIDoc())
	img := imggen.Gen(getBoardName(), readBoard())

	reportPath := filepath.Join(planDir, reportPathIMG)
	os.Remove(reportPath)
	file, err := os.Create(reportPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	println("report created at: ", reportPath)

	err = png.Encode(file, img)
	if err != nil {
		panic(err)
	}
}

func genASCIIDoc() string {
	var s strings.Builder

	b := readBoard()
	s.WriteString(".Board: " + getBoardName())
	s.WriteString(b.ASCIIDocTable())

	return s.String()
}

func makeASCIIDoc(contents string) {
	filePath := filepath.Join(planDir, reportPath)
	os.Remove(filePath)
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.Write([]byte(contents))
}
