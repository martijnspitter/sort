package cli

// Config will hold our parsed command line arguments
type Config struct {
	FilePath string
	Unique   bool
}

func ParseFlags(args []string) (Config, error) {
	return Config{}, nil
}
