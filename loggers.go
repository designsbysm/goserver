package main

import (
	"os"

	"github.com/designsbysm/logger/v2"
	"github.com/designsbysm/loggeremail"
	"github.com/designsbysm/loggerfile"
	"github.com/spf13/viper"
)

func loggers() error {
	if err := cli(); err != nil {
		return err
	}

	if err := email(); err != nil {
		return err
	}

	if err := file(); err != nil {
		return err
	}

	return nil
}

func cli() error {
	return logger.New(
		os.Stdout,
		viper.GetInt("logger.cli.level"),
		viper.GetString("logger.cli.timestamp"),
		viper.GetInt("logger.cli.flags"),
	)
}

func email() error {
	w := loggeremail.New(
		viper.GetString("logger.email.subject"),
		viper.GetString("logger.email.from"),
		"",
		[]string{viper.GetString("logger.email.to")},
		viper.GetString("logger.email.host"),
		viper.GetInt("logger.email.port"),
	)

	return logger.New(
		w,
		viper.GetInt("logger.email.level"),
		viper.GetString("logger.email.timestamp"),
		viper.GetInt("logger.email.flags"),
	)
}

func file() error {
	w := loggerfile.New(
		viper.GetString("logger.file.path"),
	)

	return logger.New(
		w,
		viper.GetInt("logger.file.level"),
		viper.GetString("logger.file.timestamp"),
		viper.GetInt("logger.file.flags"),
	)
}
