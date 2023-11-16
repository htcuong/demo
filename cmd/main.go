package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/htcuong/demo/cmd/api"
	"github.com/htcuong/demo/config"
	"github.com/htcuong/demo/pkg/log"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	configFile = "config.yml"
)

// run server with CLI
var rootCmd = &cobra.Command{
	Use:   "server",
	Short: "server CLI",
	Long:  "run server with CLI",
}

func init() {
	initEnv()
	configs := initConfig()
	logger := initLogger(configs.LocalDevelopment)
	apiCmd := api.NewServerCmd(configs, logger)
	rootCmd.AddCommand(apiCmd)
}

// main is the entry point for the run command.
func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("run command has failed with error: %v\n", err)
		os.Exit(1)
	}
}

// initEnv initializes the environment variables from .env file if it exist
func initEnv() {
	_ = godotenv.Load()
}

// initLogger creates a new logger
func initLogger(isLocalDevelopment bool) *log.Logger {
	logger := log.NewLogger()
	// Set report caller and debug level for local development
	if isLocalDevelopment {
		logger.Logger.SetReportCaller(true)
		logger.Logger.SetLevel(logrus.DebugLevel)
	} else {
		logger.Logger.SetLevel(logrus.InfoLevel)
	}

	return logger
}

// initConfig initializes the config.
func initConfig() *config.Configurations {
	viper.SetConfigType("yaml")

	// Expand environment variables inside the config file
	b, err := os.ReadFile(configFile)
	if err != nil {
		fmt.Printf("read config has failed failed with error: %v\n", err)
		os.Exit(1)
	}

	expand := os.ExpandEnv(string(b))
	configReader := strings.NewReader(expand)

	viper.AutomaticEnv()

	if err := viper.ReadConfig(configReader); err != nil {
		fmt.Printf("read config has failed with error: %v\n", err)
		os.Exit(1)
	}

	configs := config.Configurations{}
	if err := viper.Unmarshal(&configs); err != nil {
		fmt.Printf("read config has failed failed with error: %v\n", err)
		os.Exit(1)
	}

	return &configs
}
