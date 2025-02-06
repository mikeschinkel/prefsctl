package cmds

import (
	"reflect"
)

func DomainsFlag(assignFunc func(any)) *CmdFlag {
	return &CmdFlag{
		Name:       DomainsFlagName,
		Type:       reflect.Slice,
		Subtype:    reflect.String,
		Descr:      "Filter to include only specified domains",
		Default:    []string{},
		Required:   false,
		AssignFunc: assignFunc,
	}
}
