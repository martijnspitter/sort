package service

import (
	"sort/internal/cli"
	"sort/internal/reader"
)

type Service interface {
	Execute() error
}

type SortService struct {
	cli    cli.Cli
	reader reader.Reader
}

func NewService(c cli.Cli, r reader.Reader) *SortService {
	return &SortService{
		cli:    c,
		reader: r,
	}
}

func (s *SortService) Execute() error {
	cfg, err := s.cli.Execute()
	if err != nil {
		return err
	}

	_, err = s.reader.Read(cfg.FilePath)
	if err != nil {
		return err
	}

	// Implement the sort logic here

	return nil
}
