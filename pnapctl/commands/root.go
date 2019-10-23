package commands

import (
	"fmt"
	"os"

	"phoenixnap.com/pnap-cli/pnapctl/client"
	"phoenixnap.com/pnap-cli/pnapctl/commands/create"
	"phoenixnap.com/pnap-cli/pnapctl/commands/delete"
	"phoenixnap.com/pnap-cli/pnapctl/commands/get"
	"phoenixnap.com/pnap-cli/pnapctl/commands/poweroff"
	"phoenixnap.com/pnap-cli/pnapctl/commands/poweron"
	"phoenixnap.com/pnap-cli/pnapctl/commands/reboot"
	"phoenixnap.com/pnap-cli/pnapctl/commands/reset"
	"phoenixnap.com/pnap-cli/pnapctl/commands/shutdown"
	"phoenixnap.com/pnap-cli/pnapctl/fileprocessor"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "pnapctl",
	Short: "pnapctl creates new and manages existing bare metal servers.",
	Long: `pnapctl creates new and manages existing bare metal servers provided by the PhoenixNAP Bare Metal Cloud service.

Find More information at: INSERT_LINK_HERE`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

// Execute adds all child commands to the root command, setting flags appropriately.
// Called by main.main(), only needing to happen once.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		var _ = fmt.Errorf("%s", err)
		os.Exit(1)
	}
}

func init() {
	// add flags here when needed
	rootCmd.AddCommand(get.GetCmd)
	rootCmd.AddCommand(create.CreateCmd)
	rootCmd.AddCommand(reset.ResetCmd)
	rootCmd.AddCommand(delete.DeleteCmd)
	rootCmd.AddCommand(poweroff.PowerOffCmd)
	rootCmd.AddCommand(poweron.PowerOnCmd)
	rootCmd.AddCommand(shutdown.ShutdownCmd)
	rootCmd.AddCommand(reboot.RebootCmd)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file defaults to environment variable \"PNAPCTL_HOME\" or \"pnap.yaml\" in the home directory.")

	cobra.OnInitialize(initConfig)
}

func initConfig() {
	envHome := os.Getenv("PNAPCTL_HOME")
	if envHome != "" && cfgFile == "" {
		cfgFile = envHome
	}

	if cfgFile != "" {
		// Use config file from the flag
		fileprocessor.ExpandPath(&cfgFile)
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name "pnap" (without extension)
		viper.AddConfigPath(home)
		viper.SetConfigName("pnap")
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file:", err)
		os.Exit(1)
	} else if viper.GetString("clientId") == "" || viper.GetString("clientSecret") == "" {
		fmt.Println("Client ID and Client Secret in config file should not be empty")
		os.Exit(1)
	} else {
		client.MainClient = client.NewHTTPClient(viper.GetString("clientId"), viper.GetString("clientSecret"))
	}
}
