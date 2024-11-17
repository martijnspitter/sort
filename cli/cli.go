package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	uniq bool
)

// Config will hold our parsed command line arguments
type Config struct {
	FilePath string
	Unique   bool
}

type Cmd struct {
	rootCmd *cobra.Command
}

func ParseFlags(filePath string, uniq bool) (Config, error) {
	if filePath == "" {
		return Config{}, fmt.Errorf("filepath is required")
	}

	return Config{
		FilePath: filePath,
		Unique:   uniq,
	}, nil
}

func NewCmd() *Cmd {
	return &Cmd{
		rootCmd: &cobra.Command{
			Use:   "sort",
			Short: "Replica of the sort utility",
			Long: `Replica of the sort utility for text files only.
				   it implements the unique flag to remove duplicate lines.`,
			Run: func(cmd *cobra.Command, args []string) {
				if len(args) < 1 {
					cmd.Usage()
					os.Exit(1)
				}
				config, err := ParseFlags(args[0], uniq)
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					os.Exit(1)
				}
				fmt.Println(config)
				os.Exit(0)
			},
		},
	}
}

func (c *Cmd) Execute() {
	c.rootCmd.PersistentFlags().BoolVarP(&uniq, "uniq", "u", false, "Only output unique lines")

	if err := c.rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
