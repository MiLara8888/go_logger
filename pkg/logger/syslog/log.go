package syslog

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/milara8888/logger_go/pkg/logger"
	"github.com/milara8888/logger_go/pkg/logger/core"
)

type Options struct {
	LogLevel string

	SysLogHost string
	SysLogPort string

	ServiceName string
	HostName    string
	Facility    int

	ToConsole bool
}

func (o *Options) valid() error {
	var err error
	if o.ServiceName == "" {
		err = errors.Join(err, errors.New("имя сервиса обязательно"))
	}
	if o.SysLogHost == "" || o.SysLogPort == "" {
		err = errors.Join(err, errors.New("укажите порт и имя хоста сервера syslog"))
	}

	return err
}

type ServerLogger struct {
	*net.TCPConn
	lvl        core.LogType //числовой тип ошибки
	hostFor    string
	serviceFor string
	facility   int
	toConsole  bool
}

func New(opts ...func(*Options)) (logger.Logger, error) {
	o := &Options{
		Facility: 16,
	}
	for _, fn := range opts {
		fn(o)
	}
	err := o.valid()
	if err != nil {
		return nil, err
	}

	if o.HostName == "" {
		h, err := os.Hostname()
		if err != nil {
			return nil, err
		}
		o.HostName = h
	}

	lvl, ok := core.NumLog[o.LogLevel]
	if !ok {
		log.Fatal(fmt.Errorf("уровень логироваания неверен"))
		return nil, errors.New("Ошибка уровня логгирования")
	}

	m := fmt.Sprintf("%v:%v", o.SysLogHost, o.SysLogPort)
	tcpAddr, err := net.ResolveTCPAddr("tcp", m)

	if err != nil {
		return nil, err
	}

	srv, err := net.DialTCP("tcp", nil, tcpAddr) //установление соединения
	if err != nil {
		return nil, err
	}

	return &ServerLogger{
		TCPConn:    srv,
		lvl:        lvl,
		facility:   o.Facility,
		hostFor:    o.HostName,
		serviceFor: o.ServiceName,
		toConsole:  o.ToConsole,
	}, nil
}

// Обработка форматированных строк
func (l *ServerLogger) Debugf(pattern string, s ...any) error {
	if l.lvl < core.DEBUG {
		return nil
	}
	str := fmt.Sprintf(pattern, s...)
	err := l.serverSend(str, core.DEBUG)
	if err != nil {
		return err
	}
	if l.toConsole{
		log.Printf(pattern, s...)
	}
	return nil
}

func (l *ServerLogger) Infof(pattern string, s ...any) error {
	if l.lvl < core.INFO {
		return nil
	}
	str := fmt.Sprintf(pattern, s...)
	err := l.serverSend(str, core.INFO)
	if err != nil {
		return err
	}
	if l.toConsole{
		log.Printf(pattern, s...)
	}
	return nil
}

func (l *ServerLogger) Warningf(pattern string, s ...any) error {
	if l.lvl < core.WARNING {
		return nil
	}
	str := fmt.Sprintf(pattern, s...)
	err := l.serverSend(str, core.WARNING)
	if err != nil {
		return err
	}
	if l.toConsole{
		log.Printf(pattern, s...)
	}
	return nil

}

func (l *ServerLogger) Errorf(pattern string, s ...any) error {
	if l.lvl < core.ERROR {
		return nil
	}
	str := fmt.Sprintf(pattern, s...)
	err := l.serverSend(str, core.ERROR)
	if err != nil {
		return err
	}
	if l.toConsole{
		log.Printf(pattern, s...)
	}
	return nil
}

func (l *ServerLogger) Criticalf(pattern string, s ...any) error {
	if l.lvl < core.CRITICAL {
		return nil
	}
	str := fmt.Sprintf(pattern, s...)
	err := l.serverSend(str, core.CRITICAL)
	if err != nil {
		return err
	}
	if l.toConsole{
		log.Printf(pattern, s...)
	}
	return nil
}

// обработка без форматирования
func (l *ServerLogger) Debug(s ...any) error {
	if l.lvl < core.DEBUG {
		return nil
	}
	str := fmt.Sprint(s...)
	err := l.serverSend(str, core.DEBUG)
	if err != nil {
		return err
	}
	if l.toConsole{
		log.Println(str)
	}
	return nil
}

func (l *ServerLogger) Info(s ...any) error {
	if l.lvl < core.INFO {
		return nil
	}
	str := fmt.Sprint(s...)
	err := l.serverSend(str, core.INFO)
	if err != nil {
		return err
	}
	if l.toConsole{
		log.Println(str)
	}
	return nil
}

func (l *ServerLogger) Warning(s ...any) error {
	if l.lvl < core.WARNING {
		return nil
	}
	str := fmt.Sprint(s...)
	err := l.serverSend(str, core.WARNING)
	if err != nil {
		return err
	}
	if l.toConsole{
		log.Println(str)
	}
	return nil
}

func (l *ServerLogger) Error(s ...any) error {
	if l.lvl < core.ERROR {
		return nil
	}
	str := fmt.Sprint(s...)
	err := l.serverSend(str, core.ERROR)
	if err != nil {
		return err
	}
	if l.toConsole{
		log.Println(str)
	}
	return nil
}

func (l *ServerLogger) Critical(s ...any) error {
	if l.lvl < core.CRITICAL {
		return nil
	}
	str := fmt.Sprint(s...)
	err := l.serverSend(str, core.CRITICAL)
	if err != nil {
		return err
	}
	if l.toConsole{
		log.Println(str)
	}
	return nil
}

func (l *ServerLogger) serverSend(str string, lvl core.LogType) error {

	flvl := int(lvl) + (l.facility * 8)

	res_str := fmt.Sprintf("%s %s", core.StekFnLine(3), str)

	n := time.Now().Format("Jan 02 15:04:05")

	st := fmt.Sprintf("<%v>%v %v %v: %v\n", flvl, n, l.hostFor, l.serviceFor, res_str)
	_, err := l.Write([]byte(st))
	if err != nil {
		return err
	}

	return nil
}
