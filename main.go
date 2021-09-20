package main

import (
	"github.com/diazharizky/codebase-grpc-golang/cmd"
	"github.com/diazharizky/codebase-grpc-golang/configs"
)

func main() {
	configs.LoadConfig()
	cmd.Execute()
}
