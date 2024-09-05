package cli

import (
	"fmt"
	"os"

	"github.com/dejangegic/gravedigger/file"
	"github.com/spf13/cobra"
)

var rootCMD = &cobra.Command{
	Use:   "gravedigger [path]",
	Short: "GraveDigger is a cli tool for finding unused (dead) code in Go projects.",
	Long:  "GraveDigger is a cli tool for finding unused (dead) code in Go projects. The printout is limited to only the top-level unimplemented functions. Even if other functions are only included in that one dead function, they still count as alive. Remove the dead function and re-run the test to clean-up the smaller ones if needed",

	Run: func(cmd *cobra.Command, args []string) {
		var path string
		if len(args) == 0 {
			path, _ = os.Getwd()
		} else {
			path = args[0]
		}
		file.RunAll(path)
	},
}

func Execute() {
	if err := rootCMD.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}

func init() {
	rootCMD.Flags().BoolP("help", "h", false, "Help for Gravedigger")
}
