package main

import "os"

func commandExit(cfg *config, extra string) error {
	os.Exit(0)
	return nil
}
