package cmd

import (
	"fmt"
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
		yamlPath, err := cmd.Flags().GetString("yaml")
		if err != nil {
			fmt.Println("Could not get the value of yaml input flag", err)
			os.Exit(1)
		}
		outputPath, err := cmd.Flags().GetString("out")
		if err != nil {
			fmt.Println("Could not get the value of output flag", err)
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
	file, err := os.ReadFile(sourcePath)
	if err != nil {
		return fmt.Errorf("could not read the file: %s", err)
	}

	inv, err := invoice.New(file)
	if err != nil {
		return fmt.Errorf("could not prepare the invoice: %s", err)
	}

	err = inv.SaveToPdf(outputPath)
	if err != nil {
		return fmt.Errorf("could not save PDF: %s", err)
	}
	return err
}
