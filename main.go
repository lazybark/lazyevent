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
	//Create  loggers (non-default loggers would just need to implement logger.ILogger interface)
	//CLI logger for all events
	cli := logger.NewCLI(events.Any)
	//CLI for errors only
	cli2 := logger.NewCLI(events.ErrorFlow)
	//Text loggers for all event types
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
	//New LogProcessor to rule them all
	p := lproc.New("", make(chan error), false, cli, cli2, ptl, js, csv)

	//Log simple messages with default parameters
	//.Verbose() here to avoid doubles in CLI output as we have two CLI loggers above
	p.Log(events.Info("Some info").Verbose())
	p.Log(events.Warning("Some warning").Verbose())
	p.Log(events.Error("Some error").Verbose())
	time.Sleep(2 * time.Second)
	p.Log(events.Info("Some info2").Verbose())

	//Now add some default event sources
	p.Log(events.Error("Some error").Src(events.EvsDebug).Verbose())
	p.Log(events.Error("Some error").Src(events.EvsMain).Verbose())
	//And create non-default source
	weird := events.Source{
		Text:  "WEIRD SOURCE",
		Open:  "<",
		Close: ">",
	}
	p.Log(events.Error("That one has come from weird place").Src(weird).Verbose())

	//Using of TimeFixed
	e := events.Error("Some error occured").FixTime().Verbose()
	p.Log(e)
	time.Sleep(2 * time.Second)
	p.Log(e.SetText("And the same moment something else happened!"))

	//Set one of default formats (non-default can be processes only by custom loggers)
	p.Log(events.Info("Red info message").Red())

	//Using LogErrOnly
	//Should be logged, but will have doubles in CLI
	p.LogErrOnly(GetErr(), events.ErrorFlow)
	p.LogErrOnly(GetErrEvent(), events.ErrorFlow)
	//Should NOT be logged
	p.LogErrOnly(GetString(), events.ErrorFlow)
	p.LogErrOnly(GetNonErrEvent(), events.ErrorFlow)

	ed := events.EvDefault{
		Source: events.EvsMain,
		Type:   events.Debug,
		Format: events.None,
	}
	p.Log(ed.Note("Event from default template"))

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
