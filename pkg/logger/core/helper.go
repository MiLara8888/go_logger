/* общие элементы, методы, настройки, структуры */
package core

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

type LogType int


const (
	CRITICAL LogType = 2 //критическое событие
	ERROR    LogType = 3 //ошибка
	WARNING  LogType = 4 //предупреждение
	INFO     LogType = 6 //информационное событие
	DEBUG    LogType = 7 //отладочное событие

	// countInfo // служебный
)


var NumLog = map[string]LogType{
		"CRITICAL": CRITICAL,
		"ERROR":    ERROR,
		"WARNING":  WARNING, //2
		"INFO":     INFO,    //3
		"DEBUG":    DEBUG,   //4
	}


//формирование строки - (сбора стека откуда пришла ошибка, названия функции, строки)
func StekFnLine(c int) string {
	pc, file, line, ok := runtime.Caller(c) //получения информации о месте откуда эта ошибка пришла
	if !ok {
		file = "?"
		line = 0
	}
	fn := runtime.FuncForPC(pc) //получения имени функции
	var fnName string
	if fn == nil { //если не удалось определить имя функции
		fnName = "?()"
	} else {
		dotName := filepath.Ext(fn.Name()) //если удалось, приводим к нормальному виду
		fnName = strings.TrimLeft(dotName, ".") + "()"
	}
	return fmt.Sprintf("%s:%d %s", filepath.Base(file), line, fnName)
}
