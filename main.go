package main

import (
	"fmt"
	"log"
	"time"

	"github.com/lazybark/lazyevent/v2/events"
	"github.com/lazybark/lazyevent/v2/logger"
	"github.com/lazybark/lazyevent/v2/lproc"
)

func main() {
	cli := logger.NewCLI(events.ErrorFlow)
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

	//Using LogErrOnly
	//Should be logged, but only to CLI
	p.LogErrOnly(GetErr(), events.ErrorFlow)
	p.LogErrOnly(GetErrEvent(), events.ErrorFlow)
	//Should NOT be logged
	p.LogErrOnly(GetString(), events.ErrorFlow)
	p.LogErrOnly(GetNonErrEvent(), events.ErrorFlow)

}

func GetErr() error {
	return fmt.Errorf("well, that's an error")
}

func GetErrEvent() events.Event {
	return events.Error("error event occured!")
}

func GetString() string {
	return "well, that's a string"
}

func GetNonErrEvent() events.Event {
	return events.Warning("it's just a warning")
}
