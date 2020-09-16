// Package cmd /*
package cmd

import (
	"fmt"
	"os"
	"sigma-vega/derivatives"
	"sigma-vega/logger"
	"sigma-vega/plot"
	"strconv"

	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sigma-vega",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//    Run: func(cmd *cobra.Command, args []string) { },
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
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.AddCommand(plotCmd)
	rootCmd.AddCommand(optionCmd)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.sigma-vega.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
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

		// Search config in home directory with name ".sigma-vega" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".sigma-vega")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

// plotCmd represents the plot command
var plotCmd = &cobra.Command{
	Use:   "plot",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		plot.Plot()
	},
}

// plotCmd represents the plot command
var optionCmd = &cobra.Command{
	Use:   "option",
	Short: "option s k r t q v",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		option(args)

	},
}

func asFloat(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}

	return f
}

func option(args []string) {
	stock := derivatives.Stock{
		Vol:      asFloat(args[5]),
		Price:    asFloat(args[0]),
		Dividend: asFloat(args[4]),
		RFrate:   asFloat(args[2]),
	}

	o := derivatives.NewOptionBuilder().WithUnderlying(stock).Strike(asFloat(args[1])).Rate(asFloat(args[2])).TTE(asFloat(args[3])).Put()
	derivatives.PrintOption(o)
	logger.Logz.Info("",
		zap.Float64("price", o.Price()),
	)
	o.GreekValues().String()
}
