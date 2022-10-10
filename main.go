/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"git.tazi.ai/samet/rte-cli/environ"
)

func main() {
	// cmd.Execute()
	environ.AddZipEnv("tazitest", "zipTestEnv2.zip", "/Users/samet/tmp/envs")
}

// TODO: write test for addzipenv
