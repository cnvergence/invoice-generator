package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/cnvergence/invoice-generator/invoice"
	"github.com/spf13/cobra"
)

// generaterCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate invoice",
	Long:  `Generate invoice to PDF file from YAML`,
	RunE: func(cmd *cobra.Command, args []string) error {
		yamlPath, _ := cmd.Flags().GetString("yaml")
		outputPath, err := cmd.Flags().GetString("out")
		if err != nil {
			fmt.Println("Could not get the value of flag", err)
			os.Exit(1)
		}
		err = generate(yamlPath, outputPath)

		return err
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	rootCmd.PersistentFlags().String("yaml", "invoice.yaml", "Path to the YAML file with invoice parameters")
	rootCmd.PersistentFlags().String("out", "invoice.pdf", "Path to where should it save pdf file")
}

func generate(sourcePath string, outputPath string) error {
	file, err := ioutil.ReadFile(sourcePath)
	if err != nil {
		fmt.Println("Could not read the file:", err)
		os.Exit(1)
	}

	inv := invoice.New(file)
	err = inv.SaveToPdf(outputPath)
	if err != nil {
		fmt.Println("Could not save PDF:", err)
		os.Exit(1)
	}
	return err
}
