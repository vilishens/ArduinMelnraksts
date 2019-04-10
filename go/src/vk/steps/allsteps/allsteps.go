package allsteps

import (
	"fmt"
	"runtime"
	vcfg "vk/cfg"
	vomni "vk/omnibus"
	vutils "vk/utils"

	"vk/steps/step"
	schecknet "vk/steps/stepchecknet"
	scfg "vk/steps/stepconfig"
	sparam "vk/steps/stepparams"
	spointcfg "vk/steps/steppointconfig"
	spointscan "vk/steps/steppointscan"
	srunpoints "vk/steps/steprunpoints"
	sstart "vk/steps/stepstart"
	sudp "vk/steps/stepudp"
	sweb "vk/steps/stepweb"
)

var steps = make(map[string]step.Step)
var stepSequence []string

func init() {
	initLogs()
	initSteps()
}

func initLogs() {
	//	allErr = v_log.InitAllLogs()
}

func initSteps() {
	addStep(&(sstart.ThisStep))
	addStep(&(scfg.ThisStep))
	addStep(&(sparam.ThisStep))
	addStep(&(schecknet.ThisStep))
	addStep(&(sweb.ThisStep)) // WEB step must be before point steps
	addStep(&(spointcfg.ThisStep))
	addStep(&(srunpoints.ThisStep))
	addStep(&(sudp.ThisStep))
	addStep(&(spointscan.ThisStep))
}

func addStep(s step.Step) {
	sName := s.StepName()

	if _, exists := steps[sName]; exists {
		panic(fmt.Sprintf("ALERT! Step '%s' used more than once (ind %d)", sName, len(stepSequence)))
	}

	stepSequence = append(stepSequence, sName)
	steps[sName] = s
}

func DoSteps(chDone chan int) {

	locDone := make(chan int)

	go doAllSteps(locDone)

	done := <-locDone

	fmt.Println("MAXIMOVS-MAXIMOVS")

	chDone <- done
}

func doAllSteps(chanDone chan int) {

	chErr := make(chan error)
	chDone := make(chan int)
	chGoOn := make(chan bool)

	handled := -1
	stop := false
	err := error(nil)
	done := 0

	ind := 0

	for _, s := range stepSequence {
		this_s := steps[s]

		ind++

		vutils.LogStr(vomni.LogInfo, fmt.Sprintf("===== Step - %s", this_s.StepName()))

		kol := runtime.NumGoroutine()

		fmt.Printf("%d/%d ===== vk-XXX *********** RECEIVED GO-ON of %s\n", kol, runtime.NumGoroutine(), this_s.StepName())

		go step.Execute(this_s, chDone, chGoOn, chErr)

		select {
		case <-chGoOn:

			fmt.Printf("%d/%d ===== vk-YYY *********** RECEIVED GO-ON of %s\n", kol, runtime.NumGoroutine(), this_s.StepName())

		case err = <-chErr:
			stop = true
		case done = <-chDone:
			stop = true
		}

		fmt.Printf("%d/%d ***** vk-ZZZ FINISHED GO-ON of %s GRANICIN %s\n", kol, runtime.NumGoroutine(),
			this_s.StepName(),
			vcfg.Final.LogPath)

		if stop {
			break
		}

		if this_s.StepName() == vomni.StepNamePointScan {
			fmt.Println("JATURPINA....")
		}

		handled++

		// beigās jāieliek lastStep, kas aizturēs ciklu līdz pašām beigām

	}

	fmt.Println("ooooooooooooooooooooooooooooooooooooooooooooo END-APP")
	fmt.Println("ooooooooooooooooooooooooooooooooooooooooooooo END-APP")
	fmt.Println("ooooooooooooooooooooooooooooooooooooooooooooo END-APP")

	if !stop {

		fmt.Println("********************************************* END-APP")
		fmt.Println("********************************************* END-APP")
		fmt.Println("********************************************* END-APP")

		select {
		case err = <-chErr:
		case done = <-vomni.RootDone:
			stop = true

			fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++ END-APP")
			fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++ END-APP")
			fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++ END-APP")
			fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++ END-APP")
			fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++ END-APP")
			fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++ END-APP")
			fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++ END-APP")
			fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++ END-APP")
		}
	}

	fmt.Println("===== vk-xxx PIRCHILAVA")

	for ; handled >= 0; handled-- {
		// sākot no beigām katram soli veicu POST darbības
		this_s := steps[stepSequence[handled]]
		micha := runtime.NumGoroutine()
		jogurt := len(steps) - handled
		fmt.Printf("%d ==================== of %25s\n", jogurt, this_s.StepName())
		this_s.StepPost()
		fmt.Printf("%d ### %5d (%5d) ### %d ***** vk-ZZZ FINISHES GO-POST of %25s\n",
			jogurt, runtime.NumGoroutine(), micha, len(stepSequence), this_s.StepName())
	}

	fmt.Println("===== vk-xxx RECEIVED STURDEVANTS", stop, " DONE ", done)

	if stop {
		if nil != err {
			vomni.RootErr <- err
		}
		if 0 != done {
			fmt.Println("Iz luzy...")

			chanDone <- done
		}
		return
	}
}
