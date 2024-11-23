package cliutil

import (
	"reflect"
)

type OptionType int

type OptionsMap map[OptionType]Option

//goland:noinspection GoUnusedExportedFunction
func NewBoolOptions(opts ...OptionType) OptionsMap {
	m := make(OptionsMap, len(opts))
	for _, opt := range opts {
		m[opt] = Option{
			Type:  reflect.Bool,
			Value: true,
		}
	}
	return m
}

func (m OptionsMap) BoolOption(opt OptionType) (ok bool) {
	option, found := m[opt]
	return option.Type == reflect.Bool && found && option.Value.(bool)
}

type Option struct {
	Type  reflect.Kind
	Value any
}
