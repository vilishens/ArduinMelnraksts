package a_runningpoints

import (
	"fmt"
	"net"
	"strconv"

	vpointcfg "vk/code/config/pointconfig"
	vomni "vk/omnibus"
	vrunrelonoffinterv "vk/run/relayonoffinterval"
	vutils "vk/utils"
)

var Points map[string]Runner

var StopPts chan bool

func init() {
	msgSequenceNbr = 0

	Points = make(map[string]Runner)
	SendList = SendMsgArr{}

	StopPts = make(chan bool)
}

func Run(chGoOn chan bool, chDone chan int, chErr chan error) {

	// sameklēt visu punktus

	locDone := make(chan int)
	locGoOn := make(chan bool)
	locErr := make(chan error)

	go GetAllPoints(locGoOn, locDone, locErr)

	for {
		select {
		case <-locGoOn:
			chGoOn <- true
		case rc := <-locDone:
			chDone <- rc
			return
		case err := <-locErr:
			chErr <- err
			return
		}
	}
}

func getAllCfgsRelayOnOffInterv(d vpointcfg.CfgRelOnOffIntervalPoints) (err error) {
	for name, data := range d {
		if _, has := Points[name]; has {
			continue
		}

		var cfgRun, cfgSaved vrunrelonoffinterv.RunRelOnOffIntervalStruct
		if cfgRun, err = vrunrelonoffinterv.PutRelOnOffIntervalConfiguration(data); nil != err {
			return vutils.ErrFuncLine(err)
		}
		if cfgSaved, err = vrunrelonoffinterv.PutRelOnOffIntervalConfiguration(data); nil != err {
			return vutils.ErrFuncLine(err)
		}

		p := new(vrunrelonoffinterv.PointItem)

		p.Point = name
		p.ChErr = make(chan error)
		p.ChMsg = make(chan string)
		p.ChDone = make(chan int)
		p.CfgRun = cfgRun
		p.CfgSaved = cfgSaved
		p.Index.Start = -1
		p.Index.Base = -1
		p.Index.Finish = -1
		p.Type = vomni.PointTypeRelayOnOffInterval

		Points[name] = p
	}

	/*
		for name, data := range d {
			if _, has := xrun.Points[name]; has {
				continue
			}

			var cfgRun, cfgSaved xrelonoffinterv.RunRelOnOffIntervalStruct
			if cfgRun, err = data.putRelOnOffIntervalConfiguration(); nil != err {
				return vutils.ErrFuncLine(err)
			}

			if cfgSaved, err = data.putRelOnOffIntervalConfiguration(); nil != err {
				return vutils.ErrFuncLine(err)
			}
			//		item = xrelonoffinterv.RunRelOnOffIntervalStruct{Def: runDef, Seq: runSeq}
			p := new(xrelonoffinterv.PointItem)

			p.Point = name
			p.ChErr = make(chan error)
			p.ChMsg = make(chan string)
			p.CfgRun = cfgRun
			p.CfgSaved = cfgSaved
			p.Index.Start = -1
			p.Index.Base = -1
			p.Index.Finish = -1
			p.Type = vomni.PointTypeRelayOnOffInterval

			xrun.Points[name] = p
		}
	*/
	//	fmt.Printf("==============================> Čainij Grib:\n%s\n", vrun.AllRunningCfgs)

	return
}

func getAllCfgs() (err error) {
	if err = getAllCfgsRelayOnOffInterv(vpointcfg.CfgPoints.RelOnOffIntervJSON); nil != err {
		err = vutils.ErrFuncLine(err)
	}

	return
}

func GetAllPoints(chGoOn chan bool, chDone chan int, chErr chan error) {

	// ssameklēt visus punktus

	if err := getAllCfgs(); nil != err {
		chErr <- vutils.ErrFuncLine(err)
		return
	}

	chGoOn <- true
	locDone := make(chan int)
	go RunTillEnd(locDone)

	select {
	case rc := <-locDone:
		chDone <- rc
		return
	}

}

func RunTillEnd(chDone chan int) {
	select {
	case <-StopPts:
		// Stop all points

		chDone <- vomni.DoneOK
	}
}

//##########################
func messageGet(flds []string, chDelete chan bool, chErr chan error) {

	//	point := flds[indexMsgSender]

	msgCd := -1
	var err error

	if msgCd, err = strconv.Atoi(flds[indexMsgCd]); nil != err {
		chErr <- vutils.ErrFuncLine(err)
	}

	locErr := make(chan error)
	locDone := make(chan bool)
	locDelete := make(chan bool)

	if msgCd == vomni.MsgCdInputHelloFromPoint {

		fmt.Println("RUNNING HELLO")

		go handleHelloFromPoint(flds, locDone, locErr)

		fmt.Println("RUNNING TI KUDA???")

		select {
		case <-locDone:
			//
		case err = <-locErr:
			chErr <- err
			return
		}
	}

	point := flds[vomni.MsgIndexMsgSender]
	item, ok := Points[point]
	if !ok {
		chErr <- vutils.ErrFuncLine(fmt.Errorf("\nThe message received from the unknown point %q", point))
		return
	}

	go item.Response(flds, locDelete, locErr)

	select {
	case <-locDelete:
		chDelete <- true
		return
	case err = <-locErr:
		chErr <- vutils.ErrFuncLine(err)
		return
	}

	//	default:
	//		err = fmt.Errorf("\nNo logic for the received message code 0x%08X\n", msgCd)
	//		break
	//	}

	//	chErr <- err
}

func RunFinish(chDone chan bool) {

	fmt.Printf("---------------------------------------------- EXPORT RunFinish\n")
	fmt.Printf("---------------------------------------------- EXPORT RunFinish\n")
	fmt.Printf("---------------------------------------------- EXPORT RunFinish\n")
	fmt.Printf("---------------------------------------------- EXPORT RunFinish\n")

	type finishType struct {
		point string
		done  bool
	}

	list := []finishType{}

	fmt.Println("######################### NANESEN UDAR --- 0")

	pointCount := 0

	for k, v := range Points {
		if nil != v.GetUDPAddr().IP {
			list = append(list, finishType{point: k, done: false})
			pointCount++
		}
	}
	fmt.Println("######################### NANESEN UDAR --- 1")

	chPoint := make(chan bool)

	for _, v := range list {
		go Points[v.point].Finish(chPoint)
	}
	fmt.Println("######################### NANESEN UDAR --- 2 --- LEN ", len(list))

	for 0 < pointCount {

		//fmt.Println("************* MIRONOVA ****************")
		select {
		case <-chPoint:
			pointCount--
		}
	}

	fmt.Println("######################### NANESEN UDAR --- 3")

	fmt.Printf("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@ EXPORT RunFinish DONE\n")
	fmt.Printf("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@ EXPORT RunFinish DONE\n")
	fmt.Printf("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@ EXPORT RunFinish DONE\n")
	fmt.Printf("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@ EXPORT RunFinish DONE\n")

	chDone <- true
}

//##########################
func handleHelloFromPoint(flds []string, chDone chan bool, chErr chan error) {

	point := flds[indexMsgSender]
	//var err error
	//	if msgNbr, err = strconv.Atoi(flds[indexMsgNbr]); nil != err {
	//		chErr <- vutils.ErrFuncLine(err)
	//		return
	//	}

	item, ok := Points[point]
	if ok {
		newItem := false
		if item.GetUDPAddr().IP == nil {
			newItem = true
		}

		portStr := flds[indexHelloFromPointPort]
		ipStr := flds[indexHelloFromPointIP]
		port, err := strconv.Atoi(portStr)
		if ip := net.ParseIP(ipStr); nil == ip {
			//		if nil != err {
			err = fmt.Errorf("Wrong port '%s' in the message --- %s", portStr, err.Error())
			chErr <- vutils.ErrFuncLine(err)
			return
		}

		if newItem {
			item.SetUDPAddr(ipStr, port)

			//			chMsg := make(chan string)

			fmt.Println("@@@@@@@@@@@@@@@@@@@ VILODJA MOSSABIT @@@@@@@@@@@@@@@@@@@@@@@@@@@")

			chGoOn := make(chan bool)

			go item.LetsGo(chGoOn, vomni.RootErr)
			<-chGoOn

		} else {
			item.SetUDPAddr(ipStr, port)
		}

		chDone <- true
	} else {
		err := fmt.Errorf("Received message from the unknown point '%s'", point)
		chErr <- vutils.ErrFuncLine(err)
		return
	}
}

func PointLogAdd(point string, cd int, logString string) {

	fmt.Println("\n#\n #\n  #\n   #\nROZETKA Log from A")

	Points[point].LogPointStr(cd, logString)
}

func RotateReAssign() (err error) {

	fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXx")

	for k, v := range Points {

		fmt.Println("XXXXXXXXXXXXXx ", k, " XXXXXXXXXXXXXXXXXXXXX")
		if nil == v.GetUDPAddr().IP {
			fmt.Println("-------------------------- ", k, " no UPDAddr")

		} else {
			fmt.Println("++++++++++++++++++++++++++ ", k, " no UPDAddr")

			if err = v.RotateReAssign(); nil != err {
				return vutils.ErrFuncLine(fmt.Errorf("Point check rotate failure --- %v", k, err))
			}
		}
	}

	return
}
