package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	vomni "vk/omnibus"
	sall "vk/steps/allsteps"
	vutils "vk/utils"
)

var Bashkatoff = 0

func init() {
	root()
}

func main() {

	end := false

	endCd := -1

	for !end {
		if 0 > endCd {
			vutils.LogStr(vomni.LogInfo, "***** App - START *****")
		}

		endCd := runApp()

		switch endCd {
		case vomni.DoneRestart, vomni.DoneStop, vomni.DoneError:
			os.Exit(endCd)
			end = true
		}

		fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> END-APP")
		fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> END-APP")
		fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> END-APP")
		fmt.Printf(">>>>>>>>>>>>>>>>>>>>>>> GOROUTINES %5d >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> END-APP\n", runtime.NumGoroutine())
		fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> END-APP")
		fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> END-APP")

	}
}

func runApp() (cd int) {

	Bashkatoff++

	fmt.Println("Roots", vomni.RootPath)

	fmt.Printf("================================ PIRMS brembisy %d\n", runtime.NumGoroutine())

	chDone := make(chan int)

	go sall.DoSteps(chDone)

	select {
	case err := <-vomni.RootErr:
		fmt.Printf("App finished due to an error ---> %v\n", err)
		cd = vomni.DoneError
		break
	case cd = <-chDone:
		fmt.Printf("App finished with Done code %d\n", cd)
		break
	}

	fmt.Printf("================================ PIRMS RETURNA %d\n", runtime.NumGoroutine())
	return
}

func root() {
	rootPath()
	rootLog()
}

func rootPath() {
	// It is necessary to keep the root caller path to create
	// correct file paths further
	if _, rootFile, _, ok := runtime.Caller(0); !ok {
		err := fmt.Errorf("Could not get Root Path")
		log.Fatal(err)
	} else {
		vomni.RootPath = filepath.Dir(rootFile)
	}

	return
}

func rootLog() {
	var err error
	vomni.LogMainFile, err = vutils.OpenFile(vomni.LogMainPath, vomni.LogFileFlags, vomni.LogUserPerms)
	if nil != err {
		log.Fatal(fmt.Errorf("Could not open the main log file --- %v", err))
	}

	plusStr := vomni.UDPMessageSeparator + " "

	vomni.LogData = log.New(vomni.LogMainFile, vomni.LogPrefixData+plusStr, vomni.LogLoggerFlags)
	vomni.LogErr = log.New(vomni.LogMainFile, vomni.LogPrefixErr+plusStr, vomni.LogLoggerFlags)
	vomni.LogFatal = log.New(vomni.LogMainFile, vomni.LogPrefixFatal+plusStr, vomni.LogLoggerFlags)
	vomni.LogInfo = vutils.LogNew(vomni.LogMainFile, vomni.LogPrefixInfo)
}
