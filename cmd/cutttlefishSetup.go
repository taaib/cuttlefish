package cmd

import (

	"fmt"
	"io"
	"log"
	"os"
	"runtime/debug"
	//"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/pkg/errors"
	ksLogger "github.com/stellar/kelp/support/logger"
	gsConfig "github.com/stellar/go/support/config"
	ksUtils "github.com/stellar/kelp/support/utils"
)
const cuttleExamples = `  cuttlefish claimable --conf ./path/config.cfg`

var cuttlefishSetupCmd = &cobra.Command{
	Use:     "claimable",
	Short:   "",
	Example: cuttleExamples,
}

func requiredFlag(flag string) {
	e := cuttlefishSetupCmd.MarkFlagRequired(flag)
	if e != nil {
		panic(e)
	}
}

func setLogFile(fileName string) error {
	f, e := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if e != nil {
		return fmt.Errorf("failed to set log file: %s", e)
	}
	mw := io.MultiWriter(os.Stdout, f)
	log.SetOutput(mw)
	return nil
}

func logPanic(l ksLogger.Logger) {
	if r := recover(); r != nil {
		st := debug.Stack()
		l.Errorf("PANIC!! recovered to log it in the file\npanic: %v\n\n%s\n", r, string(st))
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cuttlefish.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	configPath := cuttlefishSetupCmd.Flags().StringP("config", "c", "", "(required) cuttlefish bot's basic config file path")
	logPrefix := cuttlefishSetupCmd.Flags().StringP("log", "l", "", "log to a file (and stdout) with this prefix for the filename")

	requiredFlag("config")

	cuttlefishSetupCmd.Run = func(cmd *cobra.Command, args []string) { 
		//logger object
    	logger := ksLogger.MakeBasicLogger()
        //
		if *configPath == "" {
			ksLogger.Fatal(logger, fmt.Errorf("cmd-cuttlefishSetup: config file required"))
		}
		//cuttlefish config
		var cuttlefishConfig CuttlefishConfig
		//read config
		err := gsConfig.Read(*configPath, &cuttlefishConfig)
		if err != nil {
			ksLogger.Fatal(logger, errors.Cause(err))
		}
		//
		if *logPrefix != "" {
			t := time.Now().Format("20060102T150405MST")
			fileName := fmt.Sprintf("%s_%s.log", *logPrefix, t)
			err = setLogFile(fileName)
			if err != nil {
				ksLogger.Fatal(logger, err)
				return
			}
			logger.Infof("logging to file: %s\n", fileName)
			// we want to create a deferred recovery function here that will log panics to the log file and then exit
			defer logPanic(logger)
		}
		logger.Info(" ")
		logger.Info("Starting Cuttlefish")
		logger.Info(" ")
		//write config to log
		ksUtils.LogConfig(cuttlefishConfig)
		logger.Info(" ")
		

	}

}