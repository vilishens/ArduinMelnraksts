package relayonoffinterval

import (
	"fmt"
	"strings"
	"time"
)

var msgCounter = 0

func PrintRunningCfgs() {
	fmt.Printf("\n\n*************************************************************\n")
	fmt.Printf("******************* %20s *******************\n", "Running Configuration")
	fmt.Printf("*************************************************************\n\n")
	//fmt.Printf("%s\n", Cfgs)
}

func (d *PointItem) CfgOut(off int) (str string) {
	offNew := off + 1

	strOld := strings.Repeat("\t", off)
	strNew := strings.Repeat("\t", offNew)

	str = ""
	str += fmt.Sprintf("%s##### Relay On/Off Intervals ##### %s\n", strOld, d.Point)
	str += fmt.Sprintf("%s%15s : %s:%d\n", strNew, "UDP address", d.UDPAddr.IP.String(), d.UDPAddr.Port)
	str += fmt.Sprintf("%s\n", d.CfgRun.string(offNew))
	return
}

func (d RunRelOnOffIntervalStruct) string(off int) (str string) {
	offNew := off + 1
	strOld := strings.Repeat("\t", off)

	str = ""
	str += fmt.Sprintf("%s##### Start  \n%s\n", strOld, d.Start.string(offNew))
	str += fmt.Sprintf("%s##### Base   \n%s\n", strOld, d.Base.string(offNew))
	str += fmt.Sprintf("%s##### Finish \n%s\n", strOld, d.Finish.string(offNew))

	//	return

	return
}

func (d RunRelOnOffIntervalArr) string(off int) (str string) {
	offNew := off

	str = ""
	for _, data := range d {
		str += data.string(offNew)
	}

	return
}

func (d RunRelOnOffInterval) string(off int) (str string) {
	offNew := off + 1

	//	strOld := strings.Repeat("\t", off)
	strNew := strings.Repeat("\t", offNew)

	str = ""
	str = fmt.Sprintf("%sGPIO: % 5d STATE % 2d SECS % 6d\n", strNew, d.Gpio, d.State, d.Seconds/time.Second)

	return
}

func devStr(ind int, dala string, d RunRelOnOffInterval) (str string) {

	msgCounter++
	timeStr := time.Now().Format("2006-01-02 15:04:05 -07:00 MST")

	str = fmt.Sprintf("==> % 6d <== %s === Index % 3d === %s === GPIO % 2d State % 2d Secs %d",
		msgCounter,
		timeStr,
		ind,
		dala, d.Gpio, d.State, d.Seconds/time.Second)

	return
}

//###########################
//???????????????????????????
//###########################

func (d RunRelOnOffIntervalPoints) String() (str string) {

	offN := 1
	offset := strings.Repeat("\t", offN)

	str = ""
	str += fmt.Sprintf("##### Relay On/Off Intervals\n")

	for key, val := range d {
		str += fmt.Sprintf("%s***** NEUSPELA Point %20s:\n%s", offset, key, val)
	}

	if len(str) == 0 {
		str += fmt.Sprintf("%s***** No Relay On/Off Interval data found\n", offset)
	}

	return
}

func (d RunRelOnOffIntervalStruct) String() (str string) {

	offN := 2
	offset := strings.Repeat("\t", offN)

	str = ""
	str += fmt.Sprintf("%s##### Start \n%s\n", offset, d.Start)
	str += fmt.Sprintf("%s##### Base  \n%s\n", offset, d.Base)
	str += fmt.Sprintf("%s##### Finish\n%s\n", offset, d.Finish)

	return
}

func (d RunRelOnOffIntervalArr) String() (str string) {
	str = ""
	for _, data := range d {
		str += data.String()
	}

	return
}

func (d RunRelOnOffInterval) String() (str string) {
	offN := 3
	offset := strings.Repeat("\t", offN)

	str = ""
	str = fmt.Sprintf("%sGPIO: % 5d STATE % 2d SECS % 6d\n", offset, d.Gpio, d.State, d.Seconds/time.Second)

	return
}
