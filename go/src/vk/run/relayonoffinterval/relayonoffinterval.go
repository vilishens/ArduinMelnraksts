package relayonoffinterval

import (
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	vpointcfg "vk/code/config/pointconfig"
	"vk/omnibus"
	vomni "vk/omnibus"
	vparam "vk/params"
	xmsg "vk/run/message"
	vutils "vk/utils"
)

var zima = 0

func init() {
}

func (d *PointItem) Finish(done chan bool) {

	fmt.Println("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$ Cepo 0 $$$$$$$$$$$$$$$$$$$$$$$")
	fmt.Println("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$ Cepo 0 $$$$$$$$$$$$$$$$$$$$$$$")
	fmt.Println("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$ Cepo 0 $$$$$$$$$$$$$$$$$$$$$$$")
	fmt.Println("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$ Cepo 0 $$$$$$$$$$$$$$$$$$$$$$$")
	fmt.Println("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$ Cepo 0 $$$$$$$$$$$ SAPUNOV $$$ LEN ") //, len(d.ChDone))

	d.ChMsg <- "stopnow"

	select {
	case rc := <-d.ChDone:
		if vomni.DoneStop == rc {
			done <- true
			return
		}
	}

	olsen := <-d.ChDone

	_ = olsen

	//	if olsen == vomni.DoneStop {
	//		*done = true
	//		fmt.Println(d.Point, " ***** TOMINGAS ", *done)
	//	}

	/*
		stop, ok := <-d.ChDone
		if !ok {
			fmt.Println("############################## FIRS #######################")
			fmt.Println("############################## FIRS #######################")
			fmt.Println("############################## FIRS #######################")
			fmt.Println("############################## FIRS #######################")
		}

		_ = stop
	*/
	//zu := 0
	for {
		//zu++
		select {
		case stop, ok := <-d.ChDone:
			fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ Cepo Z ^^^^^^^^^^^^^^^^^^^^^^^", ok)
			fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ Cepo Z ^^^^^^^^^^^^^^^^^^^^^^^")
			fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ Cepo Z ^^^^^^^^^^^^^^^^^^^^^^^")
			fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ Cepo Z ^^^^^^^^^^^^^^^^^^^^^^^")
			fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ Cepo Z ^^^^^^^^^^^^^^^^^^^^^^^")
			fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ Cepo Z ^^^^^^^^^^^^^^^^^^^^^^^")

			if vomni.DoneStop == stop {
				//				*done = true
				return
			} else {
				vomni.RootErr <- vutils.ErrFuncLine(fmt.Errorf("Kautkāda iemesla pēc ZURABS"))
				return
			}
			//		default:
			//			fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ Cepo XX ^^^^^^^^^^^^^^^^^^^^^^^", zu)
			//			fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ Cepo XX ^^^^^^^^^^^^^^^^^^^^^^^", zu)

		}

		time.Sleep(time.Millisecond)

	}

	return
}

func (d *PointItem) Response(flds []string, chDelete chan bool, chErr chan error) {

	cd, err := strconv.Atoi(flds[vomni.MsgIndexMsgCd])
	if nil != err {
		chErr <- vutils.ErrFuncLine(err)
		return
	}

	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")

	switch cd {
	case vomni.MsgCdInputHelloFromPoint:
		str := fmt.Sprintf("Point %q (address %s:%s) SIGNED IN",
			flds[vomni.MsgIndexMsgSender], flds[vomni.MsgIndexHelloFromPointIP], flds[vomni.MsgIndexHelloFromPointPort])
		d.LogPointStr(cd, str)
	case vomni.MsgCdInputSuccess, vomni.MsgCdInputFailed:
		res := "Success"
		if vomni.MsgCdInputFailed == cd {
			res = "Failure"
		}
		str := fmt.Sprintf("%15s of the message #%s", res, flds[vomni.MsgIndexMsgNbr])
		d.LogPointStr(cd, str)
	default:
		chErr <- vutils.ErrFuncLine(fmt.Errorf("Unknown message code 0x%08X", cd))
		return

	}

	chDelete <- true
}

func (d *PointItem) GetUDPAddr() (addr net.UDPAddr) {

	fmt.Printf("\n\n\n<============== ZIBROV =======================> %s!!!\n\n", d.Point)

	return d.UDPAddr
}

func (d *PointItem) SetUDPAddr(ip string, port int) {
	d.UDPAddr = net.UDPAddr{IP: net.ParseIP(ip), Port: port}
}

func (d *PointItem) LogPointStr(cd int, logStr string) {

	//	fmt.Println("********************************<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
	//	fmt.Println("********************** MAXIMOFF <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
	//	fmt.Println("********************************<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<")

	if nil != d.LogInfo {
		vutils.LogStr(d.LogInfo, logStr)
	}
}

func (d *PointItem) LetsGo(chGoOn chan bool, chErr chan error) {

	locErr := make(chan error)
	locDone := make(chan int)

	// prepare all log files before to start
	if err := d.logFiles(); nil != err {
		chErr <- vutils.ErrFuncLine(err)
		return
	}

	// start the point
	parts := []listPart{}

	parts = append(parts, listPart{last: false, once: true, name: listPartStart, data: d.CfgRun.Start})
	parts = append(parts, listPart{last: false, once: false, name: listPartBase, data: d.CfgRun.Base})
	go d.run(parts, locDone, locErr)

	chGoOn <- true
	end := false
	for !end {
		select {
		case err := <-locErr:
			chErr <- err
			return
		case <-locDone:

			fmt.Println("&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&")

			end = true
		}
	}

	parts = []listPart{}
	parts = append(parts, listPart{last: false, once: true, name: listPartFinish, data: d.CfgRun.Finish})

	fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", len(parts))
	fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", len(parts))
	fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", len(parts))
	fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", len(parts))
	fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", len(parts))

	go d.run(parts, locDone, locErr)

	fmt.Println("????????????? FETYASKA ?????????????????????????????????????????????????????????")
	fmt.Println("????????????? FETYASKA ?????????????????????????????????????????????????????????")
	fmt.Println("????????????? FETYASKA ?????????????????????????????????????????????????????????")
	fmt.Println("????????????? FETYASKA ?????????????????????????????????????????????????????????")

	end = false
	for !end {
		select {
		case err := <-locErr:
			chErr <- err
			return
		case rc := <-locDone:

			fmt.Println("??????????????????????????????????????????????????????????????????????", rc)
			fmt.Println("??????????????????????????????????????????????????????????????????????", rc)
			fmt.Println("??????????????????????????????????????????????????????????????????????", rc)
			fmt.Println("??????????????????????????????????????????????????????????????????????", rc)

			end = true
		}
	}

	d.ChDone <- vomni.DoneStop
}

func (d *PointItem) run(parts []listPart, chDone chan int, chErr chan error) {

	stopRun := false
	doneCd := doZero

	for !stopRun {
		for zzz := 0; zzz < len(parts); zzz++ {
			v := parts[zzz]
			locDone := make(chan int)
			locErr := make(chan error)
			locFinish := make(chan bool)
			stop := false
			go d.runList(v, locDone, locErr, locFinish) //
			for !stop {
				select {
				case err := <-locErr:
					chErr <- vutils.ErrFuncLine(err)
					return
				case doneCd = <-locDone:
					switch doneCd {
					case doFreeze:
					case doUnfreeze:
					case doRestart:
						d.Index = RunRelOnOffIntervalIndex{Start: -1, Base: -1, Finish: -1}
						zzz = len(parts) + 4
						stop = true

					case doStopNow:
						// stop before restart or exit

						fmt.Println("______________________________ Cepo DO STOP NOW _______________________", zzz)
						fmt.Println("______________________________ Cepo DO STOP NOW _______________________", zzz)

						// need to run the finish part and stop
						d.Index = RunRelOnOffIntervalIndex{Start: -1, Base: -1, Finish: -1}
						stop = true
						stopRun = true
						zzz = len(parts) + 5
					case doListEnd:
						// the regular finish of the list
						// don't do anything, just move to the next list if it exists
						stop = true
						//						if strings.ToUpper(parts[zzz].name) == "FINISH" {
						fmt.Println("______________________________ Received Cepo LIST END _______________________", zzz, "???", len(parts))
						fmt.Println("______________________________ Received Cepo LIST END _______________________", zzz, "???", len(parts))
						//						}

						if zzz+1 == len(parts) {
							stopRun = true
						}

					}
				}
			}

			fmt.Println("______________________________ LAUKA NO STOP  _______________________", zzz, "???", len(parts))
			fmt.Println("______________________________ LAUKA NO STOP  _______________________", zzz, "???", len(parts))

		}
	}

	fmt.Println("______________________________ Cepo 7 _______________________")
	fmt.Println("______________________________ Cepo 7 _______________________")
	fmt.Println("______________________________ Cepo 7 _______________________")

	chDone <- doneCd
	//	d.ChDone <- vomni.DoneStop

	fmt.Println("______________________________ Cepo Pti7 _______________________")
	fmt.Println("______________________________ Cepo Pti7 _______________________")
	fmt.Println("______________________________ Cepo Pti7 _______________________")
}

func (d *PointItem) runList(pd listPart, chDone chan int, chErr chan error, chFinish chan bool) {

	data := pd.data
	ind := new(int)

	switch pd.name {
	case listPartStart:
		ind = &(d.Index.Start)
	case listPartBase:
		ind = &(d.Index.Base)
	case listPartFinish:
		ind = &(d.Index.Finish)

		fmt.Println("______________________________ Cepo A _______________________ LEN ", len(data))
		fmt.Println("______________________________ Cepo A _______________________ LEN ", len(data))

	default:
		chErr <- fmt.Errorf("RelayOnOffInterval tried to use par \"%s\" with no logic", pd.name)
		return
	}

	if 0 == len(data) {
		// this part with no data, send it is end of the list and leave now
		chDone <- doListEnd
		*ind = -1
		return
	}

	index := 0
	*ind = index

	doCd := doZero
	finishList := false
	d.IsActive()

	var currentInterv time.Duration

	va := false

	allRecs := false

	for !finishList {

		time.Sleep(time.Millisecond)

		if strings.ToUpper(pd.name) == "FINISH" {
			fmt.Println("______________________________ Cepo B _______________________", index)
			fmt.Println("______________________________ Cepo B _______________________", index)
		}

		interv := data[index]
		currentInterv = interv.Seconds

		if va {
			fmt.Println("#################### KILIMANDZARO ##################")
		}

		timer := time.NewTimer(currentInterv)

		if 0 < (d.State & stateFreeze) {
			// stop timer the point is in FREEZE state
			if !timer.Stop() {
				<-timer.C
			}
		}

		fmt.Println(devStr(index, pd.name, data[index]))

		ms := []string{strconv.Itoa(data[index].Gpio), strconv.Itoa(data[index].State)}

		addr := d.UDPAddr

		fmt.Printf("Message #%%d --- %q (%s:%d) --- set GPIO %2d in State %d Duration %v\n", d.Point,
			d.UDPAddr.IP.String(), d.UDPAddr.Port,
			data[index].Gpio, data[index].State, data[index].Seconds)

		logFormat := fmt.Sprintf("Message #%%d --- %q (%s:%d) --- set GPIO %2d in State %d Duration %v", d.Point,
			d.UDPAddr.IP.String(), d.UDPAddr.Port,
			data[index].Gpio, data[index].State, data[index].Seconds)

		if err := xmsg.MessageSend(vomni.MsgCdOutputSetRelayGpio, addr, ms, d.Point, logFormat); nil != err {
			chErr <- vutils.ErrFuncLine(err)
			return
		}

		//time.Sleep(time.Millisecond)

		//		if strings.ToUpper(name) == "FINISH" {
		//			fmt.Println("______________________________ Cepo C _______________________", index)
		//			fmt.Println("______________________________ Cepo C _______________________", index)
		//		}

		select {
		case <-timer.C:
			index = d.nextIndex(data, ind)
			if index+1 == len(data) {

				if strings.ToUpper(pd.name) == "FINISH" {
					fmt.Println("______________________________ ALL RECS _______________________", index)
					fmt.Println("______________________________ ALL RECS _______________________", index)
				}

				allRecs = true
			}
		case msg := <-d.ChMsg:
			doCd = getDoCode(msg)
			switch doCd {
			case doRestart, doStopNow:

				fmt.Println("______________________________ REC STOP _______________________", doCd)
				fmt.Println("______________________________ REC STOP _______________________", doCd)

				finishList = true
				*ind = -1
			case doFreeze:
				d.State |= stateFreeze
			case doUnfreeze:
				d.State &^= stateFreeze
			}
		}

		if pd.once && allRecs && (0 == index) {
			// This is the list to run once only
			// if the current index is 0 it means the 1st run has done
			fmt.Printf("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!! bratec\n")

			doCd = doListEnd
			*ind = -1
			finishList = true

			if strings.ToUpper(pd.name) == "FINISH" {
				fmt.Println("______________________________ Cepo Bb _______________________", index)
				fmt.Println("______________________________ Cepo Bb _______________________", index)
			}

		}
	}

	chDone <- doCd
}

func getDoCode(msg string) (rc int) {

	switch strings.ToUpper(msg) {
	case "LOADCFG", "LOADSAVEDCFG", "LOADDEFAULTCFG":
		return doRestart
	case "FREEZE":
		return doFreeze
	case "UNFREEZE":
		return doUnfreeze
	case "STOPNOW":
		return doStopNow
	default:
		log.Fatal(fmt.Sprintf("Don't know what to do with \"%s\"", msg))
	}

	return doZero
}

func (d *PointItem) nextIndex(data RunRelOnOffIntervalArr, index *int) (newInd int) {

	i := *index

	newInd = *index + 1

	if newInd >= len(data) {
		newInd = 0
	}

	*index = newInd

	fmt.Printf("===> BIJA %d JAUNS %d ---> LEN %d\n", i, *index, len(data))

	return
}

func (d *PointItem) IsActive() (active bool) {

	active = (d.Index.Start >= 0) || (d.Index.Base >= 0) || (d.Index.Finish >= 0)

	d.State &^= stateActive
	if active {
		d.State |= stateActive
	}

	return
}

func (d *PointItem) ReceivedWebMsg(msg string, data interface{}) {

	str := strings.ToUpper(msg)

	switch str {
	case "LOADCFG":
		// load the new configuration
		dNew := webInterface2Struct(data)

		d.CfgRun = dNew
		d.ChMsg <- str
	case "SAVECFG":
		dNew := webInterface2Struct(data)
		if err := webSavePointCfg(d.Point, data); nil != err {
			vomni.LogErr.Println(vutils.ErrFuncLine(err))
		} else {
			d.CfgSaved = dNew
		}
		d.CfgRun = dNew

	case "LOADDEFAULTCFG":
		if dNew, err := webDefaultCfgStruct(d.Point); nil != err {
			vomni.LogErr.Println(vutils.ErrFuncLine(err))
		} else {
			d.CfgRun = dNew
		}

		d.ChMsg <- str

	case "LOADSAVEDCFG":
		d.CfgRun = d.CfgSaved

		d.ChMsg <- str

	case "FREEZE", "UNFREEZE":
		d.ChMsg <- str
	default:
		log.Fatal(fmt.Sprintf("Missed logic for \"%s\"", str))
	}

	//d.ChMsg <- str
}

//###################################
//###################################
//###################################

func PutRelOnOffIntervalConfiguration(d vpointcfg.CfgRelOnOffIntervalStruct) (back RunRelOnOffIntervalStruct, err error) {

	back = RunRelOnOffIntervalStruct{}

	var rStart, rBase, rFinish RunRelOnOffIntervalArr

	if rStart, err = putArrCfg(d.Start); nil != err {
		return back, vutils.ErrFuncLine(err)
	}

	if rBase, err = putArrCfg(d.Base); nil != err {
		return back, vutils.ErrFuncLine(err)
	}

	if rFinish, err = putArrCfg(d.Finish); nil != err {
		return back, vutils.ErrFuncLine(err)
	}

	back = RunRelOnOffIntervalStruct{Start: rStart, Base: rBase, Finish: rFinish}

	return
}

func putArrCfg(d vpointcfg.CfgRelOnOffIntervalArr) (back RunRelOnOffIntervalArr, err error) {

	back = RunRelOnOffIntervalArr{}

	for _, val := range d {
		interv, err := putCfg(val)
		if nil != err {
			return RunRelOnOffIntervalArr{}, vutils.ErrFuncLine(err)
		}

		back = append(back, interv)
	}

	return
}

func putCfg(d vpointcfg.CfgRelOnOffInterval) (dst RunRelOnOffInterval, err error) {

	dst = RunRelOnOffInterval{}

	if nil == err {
		dst.Gpio, err = strconv.Atoi(d.Gpio)
	}
	if nil == err {
		dst.State, err = strconv.Atoi(d.State)
	}
	if nil == err {
		dst.Seconds, err = vutils.DurationOf3PartStr(d.Interval)
	}

	return
}

//#####################################
//###################################
//###################################

//??????????????????????????????????????????????????????????????????????????????
//?????????????? LOG FAILU KODU ŠAJĀ PAKĀ NELIETO ??????????????????????????????
//??????????????????????????????????????????????????????????????????????????????
func (d *PointItem) logFiles() (err error) {

	path := filepath.Join(vparam.Params.PointLogPath, d.Point)
	full := vutils.FileAbsPath(path, vomni.PointLogDataFile)
	if d.LogDataFile, err = d.logOneFile(full); nil != err {
		return vutils.ErrFuncLine(fmt.Errorf("Point %q data log setting failure --- %v", d.Point, err))
	}
	d.LogData = vutils.LogNewPoint(d.LogDataFile)

	full = vutils.FileAbsPath(path, vomni.PointLogInfoFile)
	if d.LogInfoFile, err = d.logOneFile(full); nil != err {
		return vutils.ErrFuncLine(fmt.Errorf("Point %q info log setting failure --- %v", d.Point, err))
	}
	d.LogInfo = vutils.LogNewPoint(d.LogInfoFile)

	return
}

func (d *PointItem) logOneFile(path string) (f *os.File, err error) {
	if f, err = vutils.OpenFile(path, vomni.LogFileFlags, vomni.LogUserPerms); nil != err {
		return nil, vutils.ErrFuncLine(fmt.Errorf("Point %q data log file failure --- %v", d.Point, err))
	}
	if err = vutils.SetRotateCfg(path, vparam.Params.RotatePointCfg, vparam.Params.RotateRunCfg, false); nil != err {
		return nil, vutils.ErrFuncLine(fmt.Errorf("Point %q log rotate file failure --- %v", d.Point, err))
	}

	return
}

func (d *PointItem) RotateReAssign() (err error) {

	path := filepath.Join(vparam.Params.PointLogPath, d.Point)

	full := vutils.FileAbsPath(path, vomni.PointLogDataFile)
	if d.LogDataFile, err = vutils.LogReAssign(d.LogDataFile, full); nil != err {
		return vutils.ErrFuncLine(fmt.Errorf("\nPoint %q data file reaasign failure -- %v", d.Point, err))
	}

	if nil == d.LogData {
		d.LogData = vutils.LogNewPoint(d.LogDataFile)
	} else {
		d.LogData.SetOutput(d.LogDataFile)
	}

	full = vutils.FileAbsPath(path, vomni.PointLogInfoFile)
	if d.LogInfoFile, err = vutils.LogReAssign(d.LogInfoFile, full); nil != err {
		return vutils.ErrFuncLine(fmt.Errorf("\nPoint %q info file reaasign failure -- %v", d.Point, err))
	}

	if nil == d.LogInfo {
		d.LogInfo = vutils.LogNewPoint(d.LogInfoFile)
	} else {
		d.LogInfo.SetOutput(d.LogInfoFile)
	}

	return
}

func (d *PointItem) WebPointData() (v omnibus.WPointData) {

	v.Point = d.Point
	v.Active = d.IsActive()
	v.Frozen = 0 < (d.State & stateFreeze)
	v.Descr = d.Point
	v.Type = d.Type
	v.State = d.State

	v.Index = d.Index

	v.CfgRun = d.running2Web(d.CfgRun)
	v.CfgSaved = d.running2Web(d.CfgSaved)

	return
}

func (d *PointItem) running2Web(s RunRelOnOffIntervalStruct) (x webPointStruct) {

	x = webPointStruct{}

	for _, v := range s.Start {
		x.Start = append(x.Start, webPoint{Gpio: strconv.Itoa(v.Gpio),
			State: strconv.Itoa(v.State), Interval: vutils.DurationTo3PartStr(v.Seconds, true)})
	}

	for _, v := range s.Base {
		x.Base = append(x.Base, webPoint{Gpio: strconv.Itoa(v.Gpio),
			State: strconv.Itoa(v.State), Interval: vutils.DurationTo3PartStr(v.Seconds, true)})
	}

	for _, v := range s.Finish {
		x.Finish = append(x.Finish, webPoint{Gpio: strconv.Itoa(v.Gpio),
			State: strconv.Itoa(v.State), Interval: vutils.DurationTo3PartStr(v.Seconds, true)})
	}

	return
}
