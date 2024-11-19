package main

import (
	"fmt"
	"os"
	"sort/internal/cli"
	"sort/internal/reader"
	"sort/internal/service"
)

func main() {
	cli := cli.NewCmd()
	reader := reader.NewReader()
	svc := service.NewService(cli, reader)

	if err := svc.Execute(); err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}

	os.Exit(0)
}
