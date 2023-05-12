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
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/dfirebaugh/planjam/pkg/config"
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "modify or print config",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			if args[0] == "ls" {
				printConfig()
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}

func readConfig() config.Config {
	confFile, err := ioutil.ReadFile(filepath.Join(planDir, defaultConfig))
	if err != nil {
		panic(err)
	}

	c, err := config.New(confFile)
	if err != nil {
		panic(err)
	}

	return c
}

func writeConfig(c config.Config) {
	filePath := filepath.Join(planDir, defaultConfig)
	os.Remove(filePath)
	confFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer confFile.Close()

	c.Write(confFile)
}

func printConfig() {
	fmt.Printf("%v", readConfig())
}
