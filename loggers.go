package main

import (
	"os"

	"github.com/designsbysm/timber/v2"
	"github.com/designsbysm/timberemail"
	"github.com/designsbysm/timberfile"
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
	return timber.New(
		os.Stdout,
		timber.StringToLevel(viper.GetString("timber.cli.level")),
		viper.GetString("timber.cli.timestamp"),
		timber.StringToFlags(viper.GetString("timber.cli.flags")),
	)
}

func email() error {
	w := timberemail.New(
		viper.GetString("timber.email.subject"),
		viper.GetString("timber.email.from"),
		"",
		[]string{viper.GetString("timber.email.to")},
		viper.GetString("timber.email.host"),
		viper.GetInt("timber.email.port"),
	)

	return timber.New(
		w,
		timber.StringToLevel(viper.GetString("timber.email.level")),
		viper.GetString("timber.email.timestamp"),
		timber.StringToFlags(viper.GetString("timber.email.flags")),
	)
}

func file() error {
	w := timberfile.New(
		viper.GetString("timber.file.path"),
	)

	return timber.New(
		w,
		timber.StringToLevel(viper.GetString("timber.file.level")),
		viper.GetString("timber.file.timestamp"),
		timber.StringToFlags(viper.GetString("timber.file.flags")),
	)
}
