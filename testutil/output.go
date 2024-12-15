package testutil

import (
	"fmt"
	"strings"
)

var debugOutput bool

type debugObj struct {
	bullet string
}

func SetDebugOutput(output bool) {
	debugOutput = output
}

//goland:noinspection GoExportedFuncWithUnexportedType
func Depth(depth int) debugObj {
	var bullet string
	if depth > 0 {
		bullet = strings.Repeat("—", depth) + " "
	}
	return debugObj{bullet}
}
func (io debugObj) Debugf(format string, args ...any) {
	Debugf(io.bullet+format, args...)
}
func (io debugObj) Debug(args ...any) {
	if len(args) > 0 {
		Debug(fmt.Sprintf("%s%v", io.bullet, args[0]), args[1:]...)
	}
}
func (io debugObj) Warnf(format string, args ...any) {
	Warnf(io.bullet+format, args...)
}
func (io debugObj) Warn(args ...any) {
	if len(args) > 0 {
		Warn(fmt.Sprintf("%s%v", io.bullet, args[0]), args[1:]...)
	}
}
func Debugf(format string, args ...any) {
	if debugOutput {
		fmt.Printf("+++ DEBUG: "+format+"\n", args...)
	}
}
func Debug(format string, args ...any) {
	if debugOutput {
		fmt.Println(append([]any{"+++ DEBUG: " + format}, args...)...)
	}
}
func Warnf(format string, args ...any) {
	fmt.Printf("+++ WARN: "+format+"\n", args...)
}
func Warn(format string, args ...any) {
	fmt.Println(append([]any{"+++ WARN: " + format}, args...)...)
}
func Infof(format string, args ...any) {
	fmt.Printf("+++ INFO: "+format+"\n", args...)
}
func Info(format string, args ...any) {
	fmt.Println(append([]any{"+++ INFO: " + format}, args...)...)
}
