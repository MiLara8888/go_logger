package maillog

import (
	"errors"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strings"

	"github.com/milara8888/logger_go/pkg/logger"
	"github.com/milara8888/logger_go/pkg/logger/core"
	"github.com/milara8888/logger_go/pkg/settings"
)

type MailLogger struct {
	LVL          core.LogType
	From         string
	SmtpHostPort string
	To           string
	ServiceName  string
	HostName     string
}

func New(s *settings.Settings) (logger.Logger, error) {

	lvl, ok := core.NumLog[s.LogLevel]
	if !ok {
		log.Fatal(fmt.Errorf("уровень логироваания неверен"))
		return nil, errors.New("Ошибка уровня логгирования")
	}

	m := fmt.Sprintf("%v:%v", s.EmailLog.SmtpHost, s.EmailLog.SmtpPort)
	hostName, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &MailLogger{
		LVL:          lvl,
		From:         s.EmailLog.From,
		SmtpHostPort: m,
		To:           s.EmailLog.To,
		ServiceName:  s.ServiceName,
		HostName:     hostName,
	}, nil
}

func (l *MailLogger) mailSend(str, p, lvl string) error {

	from := l.From
	to := strings.Split(l.To, " ")

	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("Subject: %s %s %s %s\r\n", lvl, "из", l.ServiceName, l.HostName)
	msg += fmt.Sprintf("\r\n %s  :  %s\r\n", str, p)

	// Отправка почты.
	err := smtp.SendMail(l.SmtpHostPort, nil, from, to, []byte(msg))
	if err != nil {
		return err
	}
	return nil
}

func (l *MailLogger) Debugf(pattern string, s ...any) error {
	if l.LVL < core.DEBUG { //если не подходит по общему, ничего не делаем
		return nil
	}
	p := core.StekFnLine(2)
	str := fmt.Sprintf("%s %s", pattern, s)
	err := l.mailSend(p, str, "debug")
	if err != nil {
		return err
	}
	return nil
}

func (l *MailLogger) Infof(pattern string, s ...any) error {
	if l.LVL < core.INFO { //если не подходит по общему, ничего не делаем
		return nil
	}
	p := core.StekFnLine(2)
	str := fmt.Sprintf("%s %s", pattern, s)
	err := l.mailSend(p, str, "info")
	if err != nil {
		return err
	}
	return nil
}

func (l *MailLogger) Warningf(pattern string, s ...any) error {
	if l.LVL < core.WARNING { //если не подходит по общему, ничего не делаем
		return nil
	}
	p := core.StekFnLine(2)
	str := fmt.Sprintf("%s %s", pattern, s)
	err := l.mailSend(p, str, "warning")
	if err != nil {
		return err
	}
	return nil
}

func (l *MailLogger) Errorf(pattern string, s ...any) error {
	if l.LVL < core.ERROR { //если не подходит по общему, ничего не делаем
		return nil
	}
	p := core.StekFnLine(2)
	str := fmt.Sprintf("%s %s", pattern, s)
	err := l.mailSend(p, str, "error")
	if err != nil {
		return err
	}
	return nil
}

func (l *MailLogger) Criticalf(pattern string, s ...any) error {
	if l.LVL < core.CRITICAL { //если не подходит по общему, ничего не делаем
		return nil
	}
	p := core.StekFnLine(2)
	str := fmt.Sprintf("%s %s", pattern, s)
	err := l.mailSend(p, str, "critical")
	if err != nil {
		return err
	}
	return nil
}


func (l *MailLogger) Debug(s ...any) error {
	if l.LVL < core.DEBUG { //если не подходит по общему, ничего не делаем
		return nil
	}
	p := core.StekFnLine(2)
	str := fmt.Sprint(s...)
	err := l.mailSend(p, str, "debug")
	if err != nil {
		return err
	}
	return nil
}

func (l *MailLogger) Info(s ...any) error {
	if l.LVL < core.INFO { //если не подходит по общему, ничего не делаем
		return nil
	}
	p := core.StekFnLine(2)
	str := fmt.Sprint(s...)
	err := l.mailSend(p, str, "info")
	if err != nil {
		return err
	}
	return nil
}

func (l *MailLogger) Warning(s ...any) error {
	if l.LVL < core.WARNING { //если не подходит по общему, ничего не делаем
		return nil
	}
	p := core.StekFnLine(2)
	str := fmt.Sprint(s...)
	err := l.mailSend(p, str, "warning")
	if err != nil {
		return err
	}
	return nil
}

func (l *MailLogger) Error(s ...any) error {
	if l.LVL < core.ERROR { //если не подходит по общему, ничего не делаем
		return nil
	}
	p := core.StekFnLine(2)
	str := fmt.Sprint(s...)
	err := l.mailSend(p, str, "error")
	if err != nil {
		return err
	}
	return nil
}

func (l *MailLogger) Critical(s ...any) error {
	if l.LVL < core.CRITICAL { //если не подходит по общему, ничего не делаем
		return nil
	}
	p := core.StekFnLine(2)
	str := fmt.Sprint(s...)
	err := l.mailSend(p, str, "critical")
	if err != nil {
		return err
	}
	return nil
}