package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type IsRequired bool

const (
	Required    IsRequired = true
	NotRequired IsRequired = false
)

func addPersistentStringFlag(cmd *cobra.Command, name, shorthand, usage string, required IsRequired) {
	cmd.PersistentFlags().StringP(name, shorthand, "", usage)
	if required {
		_ = cmd.MarkFlagRequired(name)
	}
	_ = viper.BindPFlag(name, cmd.PersistentFlags().Lookup(name))
}

func addStringFlag(cmd *cobra.Command, name, shorthand, usage string, required IsRequired) {
	cmd.Flags().StringP(name, shorthand, "", usage)
	if required {
		_ = cmd.MarkFlagRequired(name)
	}
	_ = viper.BindPFlag(name, cmd.Flags().Lookup(name))
}

func addPersistentBoolFlag(cmd *cobra.Command, name, shorthand, usage string, defaultVal bool, required IsRequired) {
	cmd.PersistentFlags().BoolP(name, shorthand, defaultVal, usage)
	if required {
		_ = cmd.MarkFlagRequired(name)
	}
	_ = viper.BindPFlag(name, cmd.PersistentFlags().Lookup(name))
}
