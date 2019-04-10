package step_finish

import (
	"fmt"
	"time"
	"vk/omnibus"
	"vk/steps/step"
)

type thisStep step.StepVars

var ThisStep thisStep
var stepErr = error(nil)

func init() {
	ThisStep.Name = omnibus.StepNameFinish
	ThisStep.Code = omnibus.StepStateNextOK
	ThisStep.GoOn = make(chan bool, 1)
	ThisStep.Return = make(chan bool, 1)
}

func (s *thisStep) StepCode() (rc int) {
	rc = s.Code
	return
}

func doStep() {
	// put code here to do what is necessary

	fmt.Println("FINISH STEP STARTED")

	// put code here to do what is necessary

	ThisStep.GoOn <- true // let know all done in this step

	//	ThisStep.Return <- true
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
			fmt.Println("StartFinish =====> Receive =====> GO-ON")
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
			fmt.Println("StartFinish =====> Receive =====> RETURN")

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
