/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/joshua-sajeev/infra-cli/internal/logger"
	"github.com/spf13/cobra"
)

var env string

// deployCmd represents the deploy command
var deployCmd = &cobra.Command{
	Use:   "deploy [stack-name]",
	Short: "Deploy a stack to an environment",

	Long: `Deploy a specified infrastructure stack to a target environment.

A stack represents a group of resources such as web servers or databases
defined in your configuration.

This command connects to target machines, installs dependencies, and
ensures the desired state is applied consistently.

You must provide the stack name as an argument.`,

	Example: `  infra-cli deploy webserver-stack
  infra-cli deploy webserver-stack --env production`,

	Args: cobra.ExactArgs(1),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if env != "dev" && env != "staging" && env != "prod" {
			logger.Error("Invalid environment: %s", env)
			return fmt.Errorf("invalid env")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		stackName := args[0]

		logger.Info("Deploying stack: %s", stackName)
		logger.Info("Environment: %s", env)

		// simulate success
		logger.Success("Deployment completed")
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)

	deployCmd.Flags().StringVarP(&env, "env", "e", "dev", "Environment (dev/staging/prod)")
}
