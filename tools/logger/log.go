package logger

import (
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
)

type LogFactoryConfig struct {
	err       bool
	Date      bool
	CtxColor  color.Attribute
	TextColor color.Attribute
}

var lastT = time.Now()

func logFactory(config *LogFactoryConfig) func(ctx string, text string) {
	var (
		tColor color.Attribute
		cColor color.Attribute
	)

	if config.CtxColor != 0 {
		cColor = config.CtxColor
	}
	if config.TextColor != 0 {
		tColor = config.TextColor
	}

	cLog := color.New(cColor).SprintFunc()
	tLog := color.New(tColor).SprintFunc()

	return func(ctx string, text string) {
		curT := time.Now()
		var t string
		if config.Date {
			t = time.Now().Format(time.RFC850)
		}
		p := []string{color.HiBlackString(t), cLog("[", ctx, "]"), tLog(text), color.HiYellowString("+") + color.HiYellowString(curT.Sub(lastT).String())}
		if config.err {
			fmt.Errorf(strings.Join(p, " "))
		} else {
			fmt.Println(strings.Join(p, " "))
		}
		lastT = curT
	}
}

var Error = logFactory(&LogFactoryConfig{
	err:       true,
	Date:      true,
	CtxColor:  color.FgHiRed,
	TextColor: color.FgHiMagenta,
})

var Log = logFactory(&LogFactoryConfig{
	Date:      true,
	CtxColor:  color.FgYellow,
	TextColor: color.FgGreen,
})

var Debug = logFactory(&LogFactoryConfig{
	Date:      true,
	CtxColor:  color.FgHiBlue,
	TextColor: color.FgHiCyan,
})
