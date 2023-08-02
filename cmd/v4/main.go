package main

import (
	"fmt"
	"log"
	"time"

	logger "github.com/lazybark/lazyevent/v4"
)

func main() {
	//Create  loggers (non-default loggers would just need to implement logger.ILogger interface)
	//CLI logger for all events
	cli := logger.NewCLI(true, logger.Any)
	//CLI for errors only
	cli2 := logger.NewCLI(true, logger.ErrorFlow)
	//Text loggers for all event types
	ptl, err := logger.NewPlaintext("some", false, false, 1, nil, logger.Any)
	if err != nil {
		log.Fatal(err)
	}
	js, err := logger.NewJSONtext("some", false, 1, nil, logger.Any)
	if err != nil {
		log.Fatal(err)
	}
	csv, err := logger.NewCSVtext("some", true, 1, nil, logger.Any)
	if err != nil {
		log.Fatal(err)
	}
	//New LogProcessor to rule them all
	p := logger.New(true, "", make(chan error), false, cli, cli2, ptl, js, csv)

	//Log simple messages with default parameters
	//.Verbose() here to avoid doubles in CLI output as we have two CLI loggers above
	p.Log(logger.Info("Some info").Verbose())
	p.Log(logger.Warning("Some warning").Verbose())
	p.Log(logger.Error("Some error").Verbose())
	time.Sleep(2 * time.Second)
	p.Log(logger.Info("Some info2").Verbose())

	//Now add some default event sources
	p.Log(logger.Error("Some error").Src(logger.EvsDebug).Verbose())
	p.Log(logger.Error("Some error").Src(logger.EvsMain).Verbose())
	//And create non-default source
	weird := logger.Source{
		Text:  "WEIRD SOURCE",
		Open:  "<",
		Close: ">",
	}
	p.Log(logger.Error("That one has come from weird place").Src(weird).Verbose())

	//Using of TimeFixed
	e := logger.Error("Some error occured").FixTime().Verbose()
	p.Log(e)
	time.Sleep(2 * time.Second)
	p.Log(e.SetText("And the same moment something else happened!"))

	//Set one of default formats (non-default can be processes only by custom loggers)
	p.Log(logger.Info("Red info message").Red())

	//Using LogErrOnly
	//Should be logged, but will have doubles in CLI
	p.LogErrOnly(GetErr(), logger.ErrorFlow)
	p.LogErrOnly(GetErrEvent(), logger.ErrorFlow)
	//Should NOT be logged
	p.LogErrOnly(GetString(), logger.ErrorFlow)
	p.LogErrOnly(GetNonErrEvent(), logger.ErrorFlow)

	ed := logger.EvDefault{
		Source: logger.EvsMain,
		Type:   logger.Debug,
		Format: logger.None,
	}
	p.Log(ed.Note("Event from default template"))

	time.Sleep(time.Minute)
	p.Log(logger.Info("Some info after a minute").Verbose())

}

func GetErr() error {
	return fmt.Errorf("well, that's an error")
}

func GetErrEvent() logger.Event {
	return logger.Error("error event occured!")
}

func GetString() string {
	return "well, that's a string"
}

func GetNonErrEvent() logger.Event {
	return logger.Warning("it's just a warning")
}
