package stepfinish

import (
	"fmt"
	"time"
	"vk/omnibus" //	"vk/start"
	vomni "vk/omnibus"
	"vk/steps/step"
)

type thisStep step.StepVars

var ThisStep thisStep

func init() {
	ThisStep.Name = omnibus.StepNameFinish
	ThisStep.Err = make(chan error)
	ThisStep.GoOn = make(chan bool)
	ThisStep.Done = make(chan int)
}

func (s *thisStep) stepDo() {
	// put code here to do what is necessary

	s.GoOn <- true
	//	for {
	//		time.Sleep(3 * vomni.StepExecDelay)
	//	}

	/*
		chErr := make(chan error)
		chDone := make(chan bool)
		chGoOn := make(chan bool)
		//	go vudp.Server(chGoOn, chDone, chErr) // put the right call here

		select {
		case err := <-chErr:
			s.Err <- err
		case done := <-chDone:
			s.Done <- done
		case <-chGoOn:
			s.GoOn <- true
		}
	*/
}

func (s *thisStep) StepName() string {
	return ThisStep.Name
}

func (s *thisStep) StepPre(chDone chan int, chGoOn chan bool, chErr chan error) {
	// do if something is required before the step execution
	chGoOn <- true
}

func (s *thisStep) StepExec(chDone chan int, chGoOn chan bool, chErr chan error) {

	fmt.Println("Alexandr Hamilton @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")

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

			fmt.Println("Bojok @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")

			chGoOn <- locGoOn
		}
		time.Sleep(vomni.StepExecDelay)
	}
}

func (s *thisStep) StepPost() {
	// may be something needs to be done before leave the step
}
