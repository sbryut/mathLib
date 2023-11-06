package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/sbryut/mathLib/pkg/calculations"
	"github.com/sirupsen/logrus"
)

func setupLogger(level logrus.Level) {
	logrus.SetLevel(level)
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
	})
}

func main() {
	logFlag := flag.Bool("log", false, "Enable detailed logging with logrus")
	flag.Parse()

	if len(flag.Args()) < 1 {
		os.Exit(1)
	}

	n, err := strconv.ParseInt(flag.Arg(0), 10, 64)
	if err != nil {
		if *logFlag {
			setupLogger(logrus.InfoLevel)
			logrus.WithError(err).Fatal("Error converting argument to int64")
		} else {
			os.Exit(1)
		}
	}

	if *logFlag {
		setupLogger(logrus.InfoLevel)
		result := calculations.Calculate(n, true)
		logrus.WithFields(logrus.Fields{"result": result}).Info("Calculation result")
	} else {
		result := calculations.Calculate(n, false)
		fmt.Println(result)
	}
}
