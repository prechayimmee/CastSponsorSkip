package config

import "github.com/spf13/viper"

func Load() {
	InitViper()
	PausedIntervalValue = viper.GetDuration(PausedIntervalKey)
	PlayingIntervalValue = viper.GetDuration(PlayingIntervalKey)
	InterfaceValue = viper.GetString(InterfaceKey)
	CategoriesValue = viper.GetStringSlice(CategoriesKey)
}
