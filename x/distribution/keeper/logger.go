package keeper

import (
	"fmt"
	"os"

	kitlevel "github.com/go-kit/log/level"
	"github.com/go-kit/log/term"
	"github.com/tendermint/tendermint/libs/log"
)

var jl log.Logger

func init() {
	jl = log.NewTMLoggerWithColorFn(os.Stdout, colorFn)
}

// Color by level value
var colorFn = func(keyvals ...interface{}) term.FgBgColor {
	if keyvals[0] != kitlevel.Key() {
		panic(fmt.Sprintf("expected level key to be first, got %v", keyvals[0]))
	}
	switch keyvals[1].(kitlevel.Value).String() {
	case "debug":
		return term.FgBgColor{Fg: term.White}
	case "info":
		return term.FgBgColor{Fg: term.Blue}
	case "error":
		return term.FgBgColor{Fg: term.Red}
	default:
		return term.FgBgColor{}
	}
}
