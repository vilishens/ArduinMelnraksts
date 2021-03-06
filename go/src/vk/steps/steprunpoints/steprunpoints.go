package steprunpoints

import (
	"time"
	"vk/steps/step"

	vomni "vk/omnibus"
	vrunpoints "vk/run/a_runningpoints"
)

type thisStep step.StepVars

var ThisStep thisStep

func init() {
	ThisStep.Name = vomni.StepNameRunPoints
	ThisStep.Err = make(chan error)
	ThisStep.GoOn = make(chan bool)
	ThisStep.Done = make(chan int)
}

func (s *thisStep) stepDo() {
	// put code here to do what is necessary

	chErr := make(chan error)
	chDone := make(chan int)
	chGoOn := make(chan bool)
	go vrunpoints.Run(chGoOn, chDone, chErr) // put the right call here

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

func (s *thisStep) StepPre(chDone chan int, chGoOn chan bool, chErr chan error) {
	// do if something is required before the step execution

	chGoOn <- true
}

func (s *thisStep) StepExec(chDone chan int, chGoOn chan bool, chErr chan error) {

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
	loc := make(chan bool)
	go vrunpoints.RunFinish(loc)

	<-loc
	// sleep 2 seconds to allow finish all WEB stuff
	time.Sleep(2 * time.Second)

	s.Done <- vomni.DoneOK
}
