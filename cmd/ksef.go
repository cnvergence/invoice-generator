package cmd

import (
	"fmt"
	"os"

	"github.com/cnvergence/invoice-generator/invoice"
	"github.com/cnvergence/invoice-generator/ksef"
	"github.com/spf13/cobra"
)

var ksefCmd = &cobra.Command{
	Use:   "ksef",
	Short: "Generate KSeF FA_3 XML",
	Long:  `Generate a Polish KSeF FA_3 e-invoice XML from a YAML invoice file`,
	RunE: func(cmd *cobra.Command, args []string) error {
		yamlPath, err := cmd.Flags().GetString("yaml")
		if err != nil {
			fmt.Println("Could not get the value of yaml input flag", err)
			os.Exit(1)
		}
		outputPath, err := cmd.Flags().GetString("ksef-out")
		if err != nil {
			fmt.Println("Could not get the value of ksef-out flag", err)
			os.Exit(1)
		}
		return generateKSeF(yamlPath, outputPath)
	},
}

func init() {
	rootCmd.AddCommand(ksefCmd)
	ksefCmd.Flags().String("ksef-out", "invoice.xml", "Path to the output KSeF XML file")
}

func generateKSeF(sourcePath, outputPath string) error {
	file, err := os.ReadFile(sourcePath)
	if err != nil {
		return fmt.Errorf("could not read the file: %w", err)
	}

	inv, err := invoice.New(file)
	if err != nil {
		return fmt.Errorf("could not prepare the invoice: %w", err)
	}

	xmlBytes, err := ksef.Generate(inv)
	if err != nil {
		return fmt.Errorf("could not generate KSeF XML: %w", err)
	}

	if err := os.WriteFile(outputPath, xmlBytes, 0o644); err != nil {
		return fmt.Errorf("could not write XML file: %w", err)
	}

	fmt.Printf("KSeF XML written to %s\n", outputPath)
	return nil
}
