package main

import "os"

func callbackExit(cfg *config, args ...string) error {
	os.Exit(0)
	return nil
}
