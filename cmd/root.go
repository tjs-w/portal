/*
Copyright © 2021 Tejas Wanjari

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
	"bufio"
	"fmt"
	"log"
	"os"

	tm "github.com/buger/goterm"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"

	"github.com/tejas-w/portal/portal"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "portal",
	Short: "A portal into scrolling text",
	Long: `
Portal provides a convenient  way to scroll through humongous text by 
overwriting the set number of lines in the terminal, instead of creating huge 
terminal output.

It is also a sleek way to track the progress of certain workloads. For example, 
compilation logs where mostly we are just tracking the progress.

Portal can also redirect the output of the to a log file while displaying it on 
stdout.
`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		stat, err := os.Stdin.Stat()
		if err != nil {
			log.Fatalln(err)
		}
		if (stat.Mode() & os.ModeCharDevice) != 0 {
			return
		}
		var s, w int
		var o string
		if s, err = cmd.Flags().GetInt("size"); err != nil {
			log.Fatalln(err)
		}
		if s == 0 {
			s = 8
		}
		if w, err = cmd.Flags().GetInt("width"); err != nil {
			log.Fatalln(err)
		}
		if o, err = cmd.Flags().GetString("out-file"); err != nil {
			log.Fatalln(err)
		}
		p := portal.New(&portal.Options{
			Height:  s,
			Width:   w,
			OutFile: o,
		})
		defer p.Close()

		inCh := p.Open()
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			inCh <- scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			log.Fatalln(err)
		}

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "",
		"config file (default is $HOME/.portal.yaml)")

	rootCmd.Flags().StringP("out-file", "o", "", "Log file to fork the writes to")
	rootCmd.Flags().IntP("size", "s", 8, "Size is the height of the portal in number of lines.")
	rootCmd.Flags().IntP("width", "w", tm.Width(), "Set the width for the output, terminal width as default")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
		return
	}

	// Find home directory.
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	// Search config in home directory with name ".portal" (without extension).
	viper.AddConfigPath(home)
	viper.SetConfigType("yaml")
	viper.SetConfigName(".portal")

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		_, _ = fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
