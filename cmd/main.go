package main

import (
	logrusHook "github.com/go-udy/logrus-hook"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.AddHook(logrusHook.NewHook())

	for i := 0; i < 100; i++ {
		logrus.Infof("Hello info %d", i)
		logrus.Errorf("Hello error %d", i)
		logrus.Debugf("Hello debug %d", i)
		logrus.Tracef("Hello trace %d", i)
		logrus.Warnf("Hello warn %d", i)
		//logrus.Fatalf("Hello fatal %d", i)
		//logrus.Panicf("Hello panic %d", i)
		//time.Sleep(1 * time.Second)
	}
}
