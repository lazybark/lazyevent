package main

import (
	"log"
	"time"

	"github.com/lazybark/lazyevent/v2/events"
	"github.com/lazybark/lazyevent/v2/logger"
	"github.com/lazybark/lazyevent/v2/lproc"
)

func main() {
	cli := logger.NewCLI(events.Any)
	ptl, err := logger.NewPlaintext("some.log", false, events.Any)
	if err != nil {
		log.Fatal(err)
	}
	js, err := logger.NewJSONtext("some.json", false, events.Any)
	if err != nil {
		log.Fatal(err)
	}
	csv, err := logger.NewCSVtext("some.csv", true, events.Any)
	if err != nil {
		log.Fatal(err)
	}
	p := lproc.New("", make(chan error), false, cli, ptl, js, csv)

	p.Log(events.Info("Some info"))

	time.Sleep(2 * time.Second)
	p.Log(events.Info("Some info2"))

	//Using of TimeFixed
	e := events.Error("Some error occured").FixTime()
	p.Log(e)
	time.Sleep(2 * time.Second)
	p.Log(e.SetText("And the same moment something else happened!"))

}
