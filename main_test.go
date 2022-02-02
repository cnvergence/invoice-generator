package main

import (
	"os"
	"testing"

	"github.com/cnvergence/invoice-generator/cmd"
)

func Test_Generate(t *testing.T) {
	os.Args = []string{"invoice-generator", "generate", "--yaml=example/test-invoice.yaml", "--out=example/test.pdf"}

	cmd.Execute()
}
