package pointconfig

import (
	"fmt"
	"net"
	"strconv"
	"time"

	//	 vrun "vk/code/runningpoints"

	//	xrun "vk/run/a_runningpoints"
	//	xrelonoffinterv "vk/run/relayonoffinterval"
	vutils "vk/utils"
)

/* garanichev
func (d CfgRelOnOffInterval) putCfg(dst *xrelonoffinterv.RunRelOnOffInterval) (err error) {

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
*/
/* garanichev
func (d CfgRelOnOffIntervalArr) putArrCfg() (back xrelonoffinterv.RunRelOnOffIntervalArr, err error) {

	back = xrelonoffinterv.RunRelOnOffIntervalArr{}

	for _, val := range d {
		interv := xrelonoffinterv.RunRelOnOffInterval{}
		err = val.putCfg(&interv)
		if nil != err {
			return
		}

		back = append(back, interv)
	}

	return
}
*/
//##############################################################
/* garanichev
func (d CfgRelOnOffIntervalPoints) prepare() (err error) {
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
		/ *
			var runDef, runSeq xrelonoffinterv.RunRelOnOffIntervalArr
			var savedDef, savedSeq xrelonoffinterv.RunRelOnOffIntervalArr

			if runDef, err = data.Def.xputArrCfg(); nil != err {
				return vutils.ErrFuncLine(err)
			}

			if savedDef, err = data.Def.xputArrCfg(); nil != err {
				return vutils.ErrFuncLine(err)
			}

			if runSeq, err = data.Seq.xputArrCfg(); nil != err {
				return vutils.ErrFuncLine(err)
			}

			if savedSeq, err = data.Seq.xputArrCfg(); nil != err {
				return vutils.ErrFuncLine(err)
			}
		* /
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

	//	fmt.Printf("==============================> Čainij Grib:\n%s\n", vrun.AllRunningCfgs)

	return
}
*/
/* garanichev
func (d CfgRelOnOffIntervalStruct) putRelOnOffIntervalConfiguration() (back xrelonoffinterv.RunRelOnOffIntervalStruct, err error) {

	back = xrelonoffinterv.RunRelOnOffIntervalStruct{}

	var rStart, rBase, rFinish xrelonoffinterv.RunRelOnOffIntervalArr

	if rStart, err = d.Start.putArrCfg(); nil != err {
		return back, vutils.ErrFuncLine(err)
	}

	if rBase, err = d.Base.putArrCfg(); nil != err {
		return back, vutils.ErrFuncLine(err)
	}

	if rFinish, err = d.Finish.putArrCfg(); nil != err {
		return back, vutils.ErrFuncLine(err)
	}

	back = xrelonoffinterv.RunRelOnOffIntervalStruct{Start: rStart, Base: rBase, Finish: rFinish}

	return
}
*/
//##############################################################

//******************************************************************************
//******************************************************************************
//******************************************************************************

/*
map[string]struct {
	Def []cfgRelOnOffInterval `json:"default"`
	Seq []cfgRelOnOffInterval `json:"sequence"`
} `json:"relayOnOffIntervals"`
*/

type CfgRelOnOffInterval struct {
	Gpio     string `json:"Gpio"`
	State    string `json:"State"`
	Interval string `json:"Interval"`
}

type CfgRelOnOffIntervalArr []CfgRelOnOffInterval

type CfgRelOnOffIntervalStruct struct {
	Start  CfgRelOnOffIntervalArr `json:"Start"`  // array of the point relay default settings (used at the start and exit)
	Base   CfgRelOnOffIntervalArr `json:"Base"`   // array of the point relay setting sequences (used between the start and exit)
	Finish CfgRelOnOffIntervalArr `json:"Finish"` // array of the point relay setting sequences (used between the start and exit)
}

type CfgRelOnOffIntervalPoints map[string]CfgRelOnOffIntervalStruct

//*** cfg *** Start ************************************************************

//type cfgRelOnOffIntervalArr []*cfgRelOnOffInterval

//type cfgRelOnOffIntervalStruct struct {
//	def *cfgRelOnOffIntervalArr `json:"default"`  // array of the point relay default settings (used at the start and exit)
//	seq *cfgRelOnOffIntervalArr `json:"sequence"` // array of the point relay setting sequences (used between the start and exit)
//}

//type cfgRelOnOffIntervalData map[string]cfgRelOnOffIntervalStruct

//*** cfg *** End **************************************************************

//type intervalOnOff struct {
//	Gpio     string `json:"gpio"`
//	State    string `json:"state"`
//	Interval string `json:"interval"`
//}

//type RelayIntervalStruct struct {
//}

//type ArrayRelayIntervalOnOff []*RelayIntervalOnOff

type runRelayOnOffInterval struct {
	Gpio    int
	State   int
	Seconds time.Duration
	Type    int
	Active  bool
}

type RelayOnOffInterval struct {
	pointName string
	udpAddr   net.UDPAddr
	params    CfgRelOnOffIntervalStruct
	active    int
	msg       chan string
	err       chan error
}

var gpioValue []string

func init() {
	gpioValue = []string{"2", "1"}
}

//*** PointCfgData interface *** Start *****************************************

func (d CfgRelOnOffIntervalStruct) CfgRun(pointName string, addr net.UDPAddr, msg chan string, err chan error) {
	active := new(runningPoint)
	active.name = pointName
	active.addr = addr
	active.pointType = d.CfgType()
	active.msg = msg
	AllActivePoints[active.name] = active

	newP := RelayOnOffInterval{}
	newP.pointName = pointName
	newP.udpAddr = addr
	newP.params = AllCfgData.RelayOnOffInterval[newP.pointName]
	newP.active = -1

	locMsg := AllActivePoints[active.name].msg
	locErr := make(chan error)

	newP.run(locMsg, locErr)

	end := false
	for !end {
		select {
		case err1 := <-locErr:
			err <- err1
			end = true
		}
	}

	fmt.Println("PUSKEPALIS ", active)
}

func (d CfgRelOnOffIntervalStruct) CfgType() (cfgCd int) {
	return TypeRelayOnOffIntervals
}

func (d CfgRelOnOffIntervalStruct) CfgShow() {
}

//*** PointCfgData interface *** End *******************************************

func (d *runRelayOnOffInterval) prepareCfg(cfg CfgRelOnOffInterval) (err error) {

	if nil == err {
		d.Gpio, err = strconv.Atoi(cfg.Gpio)
	}
	if nil == err {
		d.State, err = strconv.Atoi(cfg.State)
	}
	if nil == err {
		d.Seconds, err = vutils.DurationOf3PartStr(cfg.Interval)
	}
	if nil == err {
		d.Active = false
	}

	return
}

//func (d ArrayRelayIntervalOnOff) CfgType() (cfgCd int) {
//	return TypeRelayOnOffIntervals
//}

//func (d cfgRelOnOffIntervalStruct) CfgShow() {
//
//	for k, v := range d {
//		fmt.Println("INDEX ", k, " DOYL ", v)
//
//	}
//}

/*
func (d cfgRelOnOffIntervalStruct) CfgRun(pointName string, addr net.UDPAddr, msg chan string, err chan error) {

	active := new(runningPoint)
	active.name = pointName
	active.addr = addr
	active.pointType = d.CfgType()
	active.msg = msg
	AllActivePoints[active.name] = active

	newP := RunRelayIntervalOnOff{}
	newP.point = pointName
	newP.udpAddr = addr
	newP.params = WholeCfgData.RelayIntervalOnOffData[newP.point]
	newP.active = -1

	locMsg := AllActivePoints[active.name].msg
	locErr := make(chan error)

	newP.run(locMsg, locErr)

	end := false
	for !end {
		select {
		case err1 := <-locErr:
			err <- err1
			end = true
		}
	}

	fmt.Println("PUSKEPALIS ", active)
}
*/

func (d RelayOnOffInterval) run(msg chan string, err chan error) {

	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!! RUN jāatjauno !!!!!!!!!!!!!!!!!!!!!")

	/*
		index := -1

		if (d.active < 0) && (len(d.params) > 0) {
			index = 0
		}

		if index < 0 {
			fmt.Println("KAUT AKS GREIZI....")
			return
		}

		//	tick := time.NewTicker(time.Duration(d.params[index].Seconds * time.Second))
		//	index := d.nextIndex(d.active)
		fmt.Println("************* jāsūta sākums ***************", d.params[index].Seconds)

		//	addrTick := new(time.Ticker)
		//	alive := new(time.Ticker)

		secs := d.params[index].Seconds
		timer := time.NewTimer(secs)

		//	a := <-timer.C
		fmt.Println("************* aizšūtīju sākums *************** ", d.params[index].Seconds, "mezoy")

		end := false

		for !end {
			if index < 0 {

				fmt.Printf("NEDERIGS actives indekss")

				return
			}

			select {
			case nowTime := <-timer.C:
				timeStr := nowTime.Format("2006-01-02 15:04:05 -07:00 MST")

				fmt.Printf("%s ************* MAKSIMOVS *************** %v\n", timeStr, d.params[index].Seconds)

				index = d.nextIndex(index)

				timer = time.NewTimer(d.params[index].Seconds)

				msgStr, _ := vmsg.MsgSetRelayGpio(d.params[index].Gpio, d.params[index].State)

				sendToAddress(AllActivePoints[d.point].addr, msgStr)

				//vudp.SendToAddress(d.params[index].udpAddr, "Sitkoveckij")

			}

		}
	*/
}

func (d *CfgRelOnOffIntervalArr) nextIndex(ind int) (newInd int) {
	newInd = ind + 1

	if newInd >= len(*d) {
		newInd = 0
	}

	return
}

//func (d RelOnOffIntervJSON) prepare() (err error) {

func sendToAddress(addr net.UDPAddr, msg string) (err error) {

	addrStr := addr.IP.String() + ":" + strconv.Itoa(addr.Port)

	conn, err := net.Dial("udp", addrStr)
	if err != nil {
		err = vutils.ErrFuncLine(fmt.Errorf("Connection ERROR: %v", err))
		return
	}
	defer conn.Close()

	if _, err = conn.Write([]byte(msg)); nil != err {
		err = vutils.ErrFuncLine(fmt.Errorf("SentToAddress ERROR: %v", err))
	}

	return
}
