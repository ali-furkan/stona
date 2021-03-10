package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/ali-furkqn/stona/internal/pkg/defaults"
	"github.com/fatih/color"
)

type logConfig struct {
	err         bool
	clrTime     color.Attribute
	clrContext  color.Attribute
	clrText     color.Attribute
	hasTime     bool
	hasContext  bool
	hasDuration bool
	writeFile   bool
}

var lastTime = time.Now()

func logFactory(cfg *logConfig) func(msg ...interface{}) {
	var (
		clrContext = defaults.Get(cfg.clrContext, color.FgWhite).(color.Attribute)
		clrTime    = defaults.Get(cfg.clrTime, color.FgHiBlack).(color.Attribute)
	)

	sprintTime := defaults.GetWithCond(
		cfg.hasTime,
		color.New(clrTime).SprintFunc(),
		fmt.Sprint).(func(a ...interface{}) string)

	sprintCtx := defaults.GetWithCond(
		cfg.hasContext,
		color.New(clrContext).SprintFunc(),
		fmt.Sprint).(func(a ...interface{}) string)

	sprintMsg := defaults.GetWithCond(
		cfg.clrText != 0,
		color.New(cfg.clrText).SprintFunc(),
		fmt.Sprint).(func(a ...interface{}) string)

	return func(msg ...interface{}) {
		currTime := time.Now()

		var logArr []interface{}

		if cfg.hasTime {
			logArr = append(logArr, sprintTime(currTime.Format(time.Stamp)))
		}

		if cfg.hasContext && len(msg) > 1 {
			logArr = append(logArr, sprintCtx(fmt.Sprintf("[%s]", msg[0])))
			logArr = append(logArr, sprintMsg(msg[1:]...))
		} else {
			logArr = append(logArr, sprintMsg(msg...))
		}

		if cfg.hasDuration {
			logArr = append(logArr, color.HiYellowString(currTime.Sub(lastTime).String()))
			lastTime = currTime
		}

		fmt.Println(logArr...)

		if cfg.err {
			os.Exit(1)
		}
	}
}

// Error Log
var Error = logFactory(&logConfig{
	err:        true,
	hasTime:    true,
	hasContext: true,
	clrContext: color.FgYellow,
	clrText:    color.FgHiMagenta,
})

// Log is default log
var Log = logFactory(&logConfig{
	hasTime:    true,
	hasContext: true,
	clrContext: color.FgYellow,
	clrText:    color.FgGreen,
})

// Debug Log
var Debug = logFactory(&logConfig{
	hasTime:    true,
	hasContext: true,
	clrContext: color.FgMagenta,
	clrText:    color.FgHiCyan,
})
