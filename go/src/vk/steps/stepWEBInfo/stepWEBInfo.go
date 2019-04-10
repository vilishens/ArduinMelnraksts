package stepWEBInfo

import (
	"time"
	"vk/omnibus"
	//	"vk/start"
	"vk/steps/step"

	vnet "vk/net"
	vomni "vk/omnibus"
)

type thisStep step.StepVars

var ThisStep thisStep

func init() {
	ThisStep.Name = omnibus.StepNameWEBInfo
	ThisStep.Err = make(chan error)
	ThisStep.GoOn = make(chan bool)
	ThisStep.Done = make(chan bool)
}

func (s *thisStep) stepDo() {
	// put code here to do what is necessary

	chErr := make(chan error)
	chDone := make(chan bool)
	chGoOn := make(chan bool)
	go vnet.HandleWEB(chGoOn, chDone, chErr)

	for {
		select {
		case err := <-chErr:
			s.Err <- err
		case done := <-chDone:
			s.Done <- done
		case <-chGoOn:
			s.GoOn <- true
		}
	}
}

func (s *thisStep) StepName() string {
	return ThisStep.Name
}

func (s *thisStep) StepPre(chDone chan bool, chGoOn chan bool, chErr chan error) {
	// do if something is required before the step execution

	chGoOn <- true
}

func (s *thisStep) StepExec(chDone chan bool, chGoOn chan bool, chErr chan error) {

	// do what you would like
	go s.stepDo()

	stop := false
	for !stop {
		select {
		case locErr := <-s.Err:
			chErr <- locErr
			stop = true
		case locDone := <-s.Done:
			chDone <- locDone
			stop = true
		case locGoOn := <-s.GoOn:
			chGoOn <- locGoOn
		}
		time.Sleep(vomni.StepExecDelay)
	}
}

func (s *thisStep) StepPost() {
	// may be something needs to be done before leave the step
}
