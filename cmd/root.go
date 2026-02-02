package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/udaykumar-dhokia/Gopherledger/internal/expense"
)

var (
	amount   float64
	note     string
	expenses []expense.Expense
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "Gopherledger",
	Short: "A CLI based expense tracker",
	Long: `
	                                                                         
 ▄▄▄▄▄▄▄              ▄▄                ▄▄          ▄▄                   
███▀▀▀▀▀              ██                ██          ██                   
███       ▄███▄ ████▄ ████▄ ▄█▀█▄ ████▄ ██ ▄█▀█▄ ▄████ ▄████ ▄█▀█▄ ████▄ 
███  ███▀ ██ ██ ██ ██ ██ ██ ██▄█▀ ██ ▀▀ ██ ██▄█▀ ██ ██ ██ ██ ██▄█▀ ██ ▀▀ 
▀██████▀  ▀███▀ ████▀ ██ ██ ▀█▄▄▄ ██    ██ ▀█▄▄▄ ▀████ ▀████ ▀█▄▄▄ ██    
                ██                                        ██             
                ▀▀                                      ▀▀▀              
	Gopherledger is a CLI based expense tracker that helps you track your daily expenses.
It allows you to add and delete expenses easily from your terminal.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
