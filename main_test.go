package main

import (
	"os"
	"testing"

	"github.com/cnvergence/invoice-generator/cmd"
)

func Test_Generate(t *testing.T) {
	os.Args = []string{"invoice", "generate", "--yaml=test/test-invoice.yaml", "--out=test/test.pdf"}

	cmd.Execute()
}
