package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	uniq     bool
	filePath string
)

// Config will hold our parsed command line arguments
type Config struct {
	FilePath string
	Unique   bool
}

type Cli interface {
	Execute() (Config, error)
}

type Cmd struct {
	rootCmd *cobra.Command
}

func NewConfig(filePath string, uniq bool) Config {
	return Config{
		FilePath: filePath,
		Unique:   uniq,
	}
}

func NewCmd() *Cmd {
	return &Cmd{
		rootCmd: &cobra.Command{
			Use:   "sort",
			Short: "Replica of the sort utility",
			Long: `Replica of the sort utility for text files only.
				   it implements the unique flag to remove duplicate lines.`,
			Run: func(cmd *cobra.Command, args []string) {
				if len(args) > 0 {
					filePath = args[0]
				}
			},
		},
	}
}

func (c *Cmd) Execute() (Config, error) {
	c.rootCmd.PersistentFlags().BoolVarP(&uniq, "uniq", "u", false, "Only output unique lines")

	if err := c.rootCmd.Execute(); err != nil {
		return Config{}, err
	}
	fmt.Println(filePath)

	return NewConfig(filePath, uniq), nil
}
