package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/alihacks/sqlrun/common"
	"github.com/spf13/cobra"
)

/*var (
	cfgFile    string
	serverName string
)*/

var (
	runConfig common.SqlRunConfig
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sqlrun",
	Short: "Executes sql commands",
	Long:  `TODO: Describe details`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if "" == runConfig.ServerName {
			return errors.New("Server is required")
		}
		if "" == runConfig.Query {
			return errors.New("Query is required")
		}
		return common.RunSql(runConfig)
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	//cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVarP(&runConfig.ServerName, "server", "S", "", "DNS name / ip of server")
	rootCmd.PersistentFlags().Uint16VarP(&runConfig.Port, "port", "p", 0, "port of server if non standard")

	rootCmd.PersistentFlags().StringVarP(&runConfig.UserName, "username", "U", "", "username if not using integrated authentication")
	rootCmd.PersistentFlags().StringVarP(&runConfig.Password, "password", "P", "", "password")
	rootCmd.PersistentFlags().StringVarP(&runConfig.Database, "database", "d", "", "database name")

	rootCmd.PersistentFlags().StringVarP(&runConfig.Query, "query", "q", "", "SQL Query to execute")

	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.sqlrun.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
/*
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".sqlrun" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".sqlrun")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
*/
