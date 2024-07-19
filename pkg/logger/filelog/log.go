package filelog

import (
	"errors"
	"fmt"
	"log"

	"github.com/milara8888/logger_go/pkg/logger"
	"github.com/milara8888/logger_go/pkg/logger/core"
)

type Options struct {
	LogDir string
}

type Logger struct {
	LVL      core.LogType //числовой тип ошибки
	critical *log.Logger
	err      *log.Logger
	warning  *log.Logger
	info     *log.Logger
	debug    *log.Logger
}

// создатель файловой структуры и логгера для записи в файлы
func New(logLevel string, opts ...func(*Options)) (logger.Logger, error) {

	options := &Options{}
	for _, fn := range opts {
		fn(options)
	}

	lvl, ok := core.NumLog[logLevel]

	if !ok {
		log.Fatal(fmt.Errorf("уровень логироваания неверен"))
		return nil, errors.New("Ошибка уровня логгирования")
	}
	lg := &Logger{
		LVL: lvl,
	}
	if core.CRITICAL <= lvl {
		lg.critical = fileLog(options.LogDir, "critical")
	}
	if core.ERROR <= lvl {
		lg.err = fileLog(options.LogDir, "err")
	}
	if core.WARNING <= lvl {
		lg.warning = fileLog(options.LogDir, "warning")
	}
	if core.INFO <= lvl {
		lg.info = fileLog(options.LogDir, "info")
	}
	if core.DEBUG == lvl {
		lg.debug = fileLog(options.LogDir, "debug")
	}
	return lg, nil
}

func (l *Logger) Debugf(pattern string, s ...any) error {
	if l.LVL < core.DEBUG {
		return nil
	}
	p := core.StekFnLine(2)
	str := fmt.Sprintf("%s %s", pattern, s)
	l.debug.Printf("%s:%s", p, str)
	return nil
}

func (l *Logger) Infof(pattern string, s ...any) error {
	if l.LVL < core.INFO {
		return nil
	}
	p := core.StekFnLine(2)
	str := fmt.Sprintf("%s %s", pattern, s)
	l.info.Printf("%s:%s", p, str)
	return nil
}

func (l *Logger) Warningf(pattern string, s ...any) error {
	if l.LVL < core.WARNING {
		return nil
	}
	p := core.StekFnLine(2)
	str := fmt.Sprintf("%s %s", pattern, s)
	l.warning.Printf("%s:%s", p, str)
	return nil
}

func (l *Logger) Errorf(pattern string, s ...any) error {
	if l.LVL < core.ERROR {
		return nil
	}
	p := core.StekFnLine(2)
	str := fmt.Sprintf("%s %s", pattern, s)
	l.err.Printf("%s:%s", p, str)
	return nil
}

func (l *Logger) Criticalf(pattern string, s ...any) error {
	if l.LVL < core.CRITICAL {
		return nil
	}
	p := core.StekFnLine(2)
	str := fmt.Sprintf("%s %s", pattern, s)
	l.critical.Printf("%s:%s", p, str)
	return nil
}



func (l *Logger) Debug(s ...any) error {
	if l.LVL < core.DEBUG {
		return nil
	}
	p := core.StekFnLine(2)
	str := fmt.Sprint(s...)
	l.debug.Printf("%s:%s", p, str)
	return nil
}

func (l *Logger) Info(s ...any) error {
	if l.LVL < core.INFO {
		return nil
	}
	p := core.StekFnLine(2)
	str := fmt.Sprint(s...)
	l.info.Printf("%s:%s", p, str)
	return nil
}

func (l *Logger) Warning(s ...any) error {
	if l.LVL < core.WARNING {
		return nil
	}
	p := core.StekFnLine(2)
	str := fmt.Sprint(s...)
	l.warning.Printf("%s:%s", p, str)
	return nil
}

func (l *Logger) Error(s ...any) error {
	if l.LVL < core.ERROR {
		return nil
	}
	p := core.StekFnLine(2)
	str := fmt.Sprint(s...)
	l.err.Printf("%s:%s", p, str)
	return nil
}

func (l *Logger) Critical(s ...any) error {
	if l.LVL < core.CRITICAL {
		return nil
	}
	p := core.StekFnLine(2)
	str := fmt.Sprint(s...)
	l.critical.Printf("%s:%s", p, str)
	return nil
}
