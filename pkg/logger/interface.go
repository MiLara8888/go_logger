package logger

var _ = [1]int{}[len(TypeLogg)-int(countType)]

// var _ = [1]int{}[len(NumLog)-int(countInfo)]

type LoggerType int

const (
	FILE LoggerType = iota
	SYSLOG
	MAIL

	countType
)

type LogType int

const (
	CRITICAL LogType = 2 //критическое событие
	ERROR    LogType = 3 //ошибка
	WARNING  LogType = 4 //предупреждение
	INFO     LogType = 6 //информационное событие
	DEBUG    LogType = 7 //отладочное событие

	countInfo // служебный
)

var (
	// уровни логгирования
	NumLog = map[string]LogType{
		"CRITICAL": CRITICAL,
		"ERROR":    ERROR,
		"WARNING":  WARNING, //2
		"INFO":     INFO,    //3
		"DEBUG":    DEBUG,   //4
	}

	TypeLogg = map[string]LoggerType{
		"FILE":   FILE,
		"SYSLOG": SYSLOG,
		"MAIL":   MAIL,
	}
)

type Logger interface {
	Debugf(patern string, s ...any) error
	Infof(patern string, s ...any) error
	Warningf(patern string, s ...any) error
	Errorf(patern string, s ...any) error
	Criticalf(patern string, s ...any) error

	Debug(s ...any) error
	Info(s ...any) error
	Warning(s ...any) error
	Error(s ...any) error
	Critical(s ...any) error
}

// // создание экземпляра логгера, точка входа
// func New(s *settings.Settings) (Logger, error) {
// 	//возвращает логи разных типов
// 	t, ok := TypeLogg[s.LoggerType] //реализация переклчателя типов
// 	if !ok {
// 		return nil, fmt.Errorf("укажите правильный тип лога")
// 	}
// 	switch t {
// 	case FILE:
// 		l, err := filelog.New(s)
// 		if err != nil {
// 			return nil, err
// 		}
// 		return l, nil

// 	case SYSLOG:
// 		l, err := syslog.New(s)
// 		if err != nil {
// 			return nil, err
// 		}
// 		return l, nil

// 	case MAIL:
// 		l, err := maillog.New(s)
// 		if err != nil {
// 			return nil, err
// 		}
// 		return l, nil

// 	default:
// 		return nil, fmt.Errorf("укажите правильный тип лога")
// 	}

// }
