package stepstart

import (
	"fmt"
	vcli "vk/cli"
	vomni "vk/omnibus" //	"vk/start"
	"vk/steps/step"
)

type thisStep step.StepVars

var ThisStep thisStep

var count = -1

func init() {
	ThisStep.Name = vomni.StepNameStart
	ThisStep.Err = make(chan error)
	ThisStep.GoOn = make(chan bool)
	ThisStep.Done = make(chan int)
}

func (s *thisStep) stepDo() {
	// put code here to do what is necessary

	err := vcli.Init()
	if nil != err {
		s.Err <- err
	}

	//stepErr = start.Run()
	//fmt.Println("ESMU ieks STEP DOOOO")
	//	ti := time.NewTicker(5 * time.Second)

	//	for {
	/*
		select {

		case <-ti.C:
			fmt.Println("################## LOGOFET stepDO #################################")
			ti = time.NewTicker(5 * time.Second)
			s.GoOn <- true // let know all done in this step
		}*/

	//		s.GoOn <- true
	//	}

	s.GoOn <- true

}

func (s *thisStep) StepName() string {
	return ThisStep.Name
}

func (s *thisStep) StepPost() {
	s.Done <- vomni.DoneOK
	return
}

func (s *thisStep) StepPre(chDone chan int, chGoOn chan bool, chErr chan error) {

	fmt.Println("ESMU ieks PRE")

	chGoOn <- true
}

func (s *thisStep) StepExec(chDone chan int, chGoOn chan bool, chErr chan error) {

	count++
	fmt.Println("ESMU ieks EXEC ", count)

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
	}
}
