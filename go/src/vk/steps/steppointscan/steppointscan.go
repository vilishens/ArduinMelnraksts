package steppointscan

import (
	"fmt"
	vscanpoints "vk/code/net/scanpoints"
	vomni "vk/omnibus"
	"vk/steps/step"
)

type thisStep step.StepVars

var ThisStep thisStep

func init() {
	ThisStep.Name = vomni.StepNamePointScan
	ThisStep.Err = make(chan error)
	ThisStep.GoOn = make(chan bool)
	ThisStep.Done = make(chan int)
}

func (s *thisStep) stepDo() {
	// put code here to do what is necessary

	chErr := make(chan error)
	chDone := make(chan bool)
	//	chGoOn := make(chan bool)

	//	vscannet.ScanNet(net.ParseIP(vparam.Params.InternalIPv4), 0, 255)
	//	vudp.Server(chGoOn, chDone, chErr) // put the right call here

	//	if ip, errIP := net.ResolveIPAddr("ip", vparam.Params.InternalIPv4); nil == errIP {
	//		fmt.Printf("STEPA IP %s\n", ip.String())
	//	}

	fmt.Println("____________________________________________________ FILARETAI")

	go vscanpoints.ScanPoints(chDone, chErr)

	// 	go vscanpoints.ScanPoints(net.ParseIP(vparam.Params.InternalIPv4).To4(), 8, 15, chGoOn, &wg, filar)
	//fmt.Printf("&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&& SISKIMO\n %s\n", resp)

	//	for {
	select {
	case err := <-chErr:
		s.Err <- err
		//		case done := <-chDone:
		//			s.Done <- done
	case <-chDone:
		s.GoOn <- true
	}
	//	}

	fmt.Println("____________________________________________________ FABIJONISKIS")

	//chGoOn <- true

	//	fmt.Printf("########################################################################################### Scan NET done... %d", ki)
	//	chGoOn <- true

}

func (s *thisStep) StepName() string {
	return ThisStep.Name
}

func (s *thisStep) StepPre(chDone chan int, chGoOn chan bool, chErr chan error) {
	// do if something is required before the step execution

	// seit vajag pārbaudīt, vai vispār ir punkti

	//	err := fmt.Errorf("ERO MENTIRANTA")
	//	chErr <- err
	//  chDone <- true 2.variants - nokonstatēt sākumā, ka viss izdarīts

	chGoOn <- true
}

func (s *thisStep) StepExec(chDone chan int, chGoOn chan bool, chErr chan error) {

	fmt.Printf("STEP %s EXEC-000\n", s.Name)

	// do what you would like
	go s.stepDo()

	fmt.Printf("STEP %s EXEC-111\n", s.Name)
	//!!!!!	for {
	select {
	case locErr := <-s.Err:
		chErr <- locErr
	case locDone := <-s.Done:
		chDone <- locDone
	case locGoOn := <-s.GoOn:
		chGoOn <- locGoOn
	}
	//!!!!!	}
}

func (s *thisStep) StepPost() {
	// may be something needs to be done before leave the step
}
