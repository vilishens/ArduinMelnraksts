package stepCheckNet

import (
	"fmt"
	"time"

	vchecknet "vk/checkNet"
	vomni "vk/omnibus"
	"vk/steps/step"
)

type thisStep step.StepVars

var ThisStep thisStep

func init() {
	ThisStep.Name = vomni.StepNameCheckNet
	ThisStep.Err = make(chan error)
	ThisStep.GoOn = make(chan bool)
	ThisStep.Done = make(chan int)
}

func (s *thisStep) stepDo() {
	// put code here to do what is necessary

	chErr := make(chan error)
	chDone := make(chan int)
	chGoOn := make(chan bool)

	go vchecknet.NetInfo(chGoOn, chDone, chErr)

	for {
		select {
		case err := <-chErr:
			s.Err <- err
			return
		case <-chDone:
			fmt.Println("MAXACOW------------- STEPDO")
			return
		case <-chGoOn:
			fmt.Println("MAXACOW------------- spaceman")

			s.GoOn <- true
		}
	}
}

func (s *thisStep) StepName() string {
	return ThisStep.Name
}

func (s *thisStep) StepPre(chDone chan int, chGoOn chan bool, chErr chan error) {
	// do if something is required before the step execution

	chGoOn <- true
}

func (s *thisStep) StepExec(chDone chan int, chGoOn chan bool, chErr chan error) {

	// do what you would like
	go s.stepDo()

	for {
		select {
		case locErr := <-s.Err:
			chErr <- locErr
			return
		case locDone := <-s.Done:
			fmt.Println("MAXACOW------------- EXEC")
			chDone <- locDone
			return
		case locGoOn := <-s.GoOn:
			fmt.Println("===================== DOMINIK")
			chGoOn <- locGoOn
		}
		time.Sleep(vomni.StepExecDelay)
	}
}

func (s *thisStep) StepPost() {
	// may be something needs to be done before leave the step
	fmt.Println("===================== POST")
	s.Done <- vomni.DoneOK
}
