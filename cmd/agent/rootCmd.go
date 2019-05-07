package agent

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	cmdMode    bool
	uri        string
	methodType string
	path       string
	verbose    bool
)

var agentCmd = &cobra.Command{
	Use:   "perph",
	Short: "Runs the perph agent",
	Long:  `Runs the perph agent`,
}

var requestCmd = &cobra.Command{
	Use:   "request [uri of target service]",
	Short: "Runs a single request",
	Long:  `Runs the perph agent against the target endpoint`,
	Run: func(cmd *cobra.Command, args []string) {
		request()
	},
}

var batchCmd = &cobra.Command{
	Use:   "batch [uri of target service]",
	Short: "Runs a batch request",
	Long:  `Runs a batch request against the target endpoint`,
	Run: func(cmd *cobra.Command, args []string) {
		batch()
	},
}

var loadCmd = &cobra.Command{
	Use:   "load [uri of target service]",
	Short: "Runs a load test",
	Long:  `Runs a load test against the target endpoint`,
	Run: func(cmd *cobra.Command, args []string) {
		load()
	},
}

var validateCmd = &cobra.Command{
	Use:   "validate [uri of target service] [validation template name]",
	Short: "Validates the response of an endpoint",
	Long:  `Validates the response of a target endpoint`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
		validate()
	},
}

func init() {
	agentCmd.Flags().BoolVarP(&cmdMode, "cmdMode", "c", false, "run as a cmd")
	agentCmd.PersistentFlags().StringVarP(&methodType, "type", "t", "rest", "the type of request to make: REST or GRPC")
	agentCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "run as a command with verbose logging")
}

func Execute() {
	agentCmd.AddCommand(requestCmd)
	agentCmd.AddCommand(batchCmd)
	agentCmd.AddCommand(loadCmd)
	agentCmd.AddCommand(validateCmd)
	if err := agentCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
