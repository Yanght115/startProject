package logging

import (
	"github.com/cihub/seelog"
)

const loggerFormat = `
<seelog type="asynctimer" asyncinterval="10000" minlevel="info">
    <outputs formatid="common">
        <rollingfile type="date" filename="log/common.log" datepattern="20060102" maxrolls="120"/>
        <console/>
        <filter levels="warn,error,critical">
            <rollingfile formatid="error" type="date" filename="log/error.log" datepattern="2006010215" maxrolls="120"/>
        </filter>
    </outputs>
    <formats>
        <format id="common" format="%Date %Time [%Level] %Msg%n"/>
        <format id="error" format="%Date %Time [%Level] %RelFile %Line %Msg%n"/>
    </formats>
</seelog>
`

// Logger -
var Logger seelog.LoggerInterface

func InitLogger(f string) {
	DisableLog()
	loadAppConfig(f)
}

func loadAppConfig(formatFile string) {
	var logger seelog.LoggerInterface
	var err error
	if formatFile == "" {
		logger, err = seelog.LoggerFromConfigAsString(loggerFormat)
	} else {
		logger, err = seelog.LoggerFromConfigAsFile(formatFile)
	}

	if err != nil {
		panic(err)
	}
	UseLogger(logger)
}

// DisableLog disables all library log output
func DisableLog() {
	Logger = seelog.Disabled
}

// UseLogger uses a specified seelog.LoggerInterface to output library log.
// Use this func if you are using Seelog logging system in your app.
func UseLogger(newLogger seelog.LoggerInterface) {
	Logger = newLogger
}
