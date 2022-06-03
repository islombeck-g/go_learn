package cmd

//package main
import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Short: "simple archiever",
}

func Execute(){
	rootCmd.Execute();err != nil{
		handleErr(err)
	}
}

func handleErr(err error){
	_, _ = fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}