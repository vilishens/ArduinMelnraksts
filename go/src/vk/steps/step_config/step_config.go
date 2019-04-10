package step_config

import (
	"fmt"
	"time"
	"vk/config"
	"vk/omnibus"
	"vk/steps/step"
)

type thisStep step.StepVars

var ThisStep thisStep
var stepErr = error(nil)

func init() {
	ThisStep.Name = omnibus.StepNameConfig
	ThisStep.Code = omnibus.StepStateNextOK
	ThisStep.GoOn = make(chan bool, 1)
	ThisStep.Return = make(chan bool, 1)
}

func (s *thisStep) StepCode() (rc int) {
	rc = s.Code
	return
}

func doStep() {
	stepErr = config.Run()
	// put code here to do what is necessary
	//if stepErr = config.GetConfig(); nil != stepErr {
	//	return
	//}

	ThisStep.GoOn <- true // let know all done in this step
}

func (s *thisStep) StepName() string {
	return s.Name
}

func (s *thisStep) StepPost() {
	s.Code = omnibus.StepStateNextOK
}

func (s *thisStep) StepPre() {
	s.Code = omnibus.StepStateNextOK
}

func (s *thisStep) StepExec() {
	go doStep()

	fmt.Println("StartXXX =====> Receive =====> ALL-DONE")

	if nil != stepErr {
		fmt.Println("Error!", stepErr)
		s.Code = omnibus.StepStateNextError
		s.Return <- true
		return
	} else {
		s.Code = omnibus.StepStateNextOK
	}
}

func (s *thisStep) StepWaitGoOn() {
	for {
		time.Sleep(omnibus.StepGoOnWaitSleep)

		select {
		case _ = <-s.GoOn:
			fmt.Println("StartXXX =====> Receive =====> GO-ON")
			// the signal all done in this step received and time to go on
			return
		default:
			// just placeholder
		}
	}
}

func (s *thisStep) StepWaitReturn() (ret bool) {
	for {
		time.Sleep(omnibus.StepGoOnWaitSleep)

		select {
		case ret = <-s.Return:
			fmt.Println("StartXXX =====> Receive =====> RETURN")

			// the signal to return from this step received and time close the step
			// RequiredBeforeClose() // Do what is required before closing the step
			return
		default:
			// just placeholder
		}
	}

	return
}

func (s *thisStep) StepErr() (err error) {
	err = stepErr
	return
}
