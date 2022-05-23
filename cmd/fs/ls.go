package fs

import (
	"fmt"

	"github.com/databricks/bricks/project"
	"github.com/spf13/cobra"

	"github.com/databrickslabs/terraform-provider-databricks/storage"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls <dir-name>",
	Short: "Lists files",
	Long:  `Lists files`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		api := storage.NewDbfsAPI(cmd.Context(), project.Current.Client())
		files, err := api.List(args[0], false)
		if err != nil {
			panic(err)
		}
		// TODO: output formatting: JSON, CSV, tables and default
		for _, v := range files {
			fmt.Printf("[-] %s (%d, %v)\n", v.Path, v.FileSize, v.IsDir)
		}
	},
}

func init() {
	// TODO: pietern: conditionally register commands
	// fabianj: don't do it
	fsCmd.AddCommand(lsCmd)
}