package message

import (
	"fmt"
	"net"
	"path/filepath"
	"strconv"
	"time"
	vdatafiles "vk/dataFiles"
	vomni "vk/omnibus"
	vparam "vk/params"
	vpoint "vk/points"
	vutils "vk/utils"
)

var msgTypeList map[string]bool

var gpioValue []string

var msgSequenceNbr int

func init() {
	gpioValue = []string{"2", "1"}
	msgSequenceNbr = 0
}

func msg2Send(msgCd int, data []string) (msg string, err error) {
	msgSequenceNbr++
	prefix := []string{vparam.Params.Name,
		strconv.Itoa(msgCd),
		strconv.Itoa(msgSequenceNbr)}

	switch msgCd {
	case MsgOutputHelloFromStation:
		prefix = append(prefix, data...)
		last := len(prefix) - 1
		for ind, str := range prefix {
			msg += str
			if ind < last {
				msg += fieldSeperator
			}
		}

		fmt.Println("MOHENDZODARO |", msg, "|")

	default:
		err = vutils.ErrFuncLine(fmt.Errorf("\nThere is no logic for the message code %d", msgCd))
	}

	return
}

func MsgOutHelloFromStation() (msg string, err error) {
	//<station name><msgCd><msgNbr><station time><stationIP><stationPort>
	data := []string{strconv.Itoa(int(time.Now().Unix())),
		vparam.Params.InternalIPv4,
		strconv.Itoa(vparam.Params.UDPPort)}

	msg, err = msg2Send(MsgOutputHelloFromStation, data)

	return
}

func MsgSetRelayGpio(gpio int, set int) (msg string, err error) {
	// Output <point name><msgCd><msgNbr><Gpio><set value>
	data := []string{strconv.Itoa(gpio), gpioValue[set]}

	msg, err = msg2Send(MsgOutputSetRelayGpio, data)

	/*
		msg = vparam.Params.Name
		msg += UDPFieldSeparator
		msg += strconv.Itoa(MsgOutputSetRelayGpio)
		msg += UDPFieldSeparator
		msg += strconv.Itoa(gpio)
		msg += UDPFieldSeparator
		msg += gpioValue[set]
	*/
	return
}

func FindMsgIndex(strType string) (ind int, err error) {

	ind = -3

	for k, v := range TotalData {
		if strType == v.MsgType {
			ind = k
		}
	}

	if 0 > ind {
		err = fmt.Errorf("The unknown message type \"%s\" received", strType)
	}

	return
}

func MsgReceived(conn *net.UDPConn, addr *net.UDPAddr, msg string) (point string, back string, err error) {
	msgType, subpath, name, point, txt, err := vutils.SplitMsg(msg)
	if nil != err {
		err = vutils.ErrFuncLine(err)
		return
	}

	if nil == err {
		switch msgType {
		case vomni.MessageTypeCmd:
			back, err = vpoint.Cmd(msg, *addr)
		case vomni.MessageTypeStart:
			back, err = vpoint.StartPoint(msg, *addr)
		case vomni.MessageTypeEvent, vomni.MessageTypeError:
			err = vdatafiles.AddRecord(msgType, subpath, point, txt)

		default:
			err = vutils.ErrFuncLine(fmt.Errorf("Unsupported message type \"%s\"", msgType))
		}
	}

	point = filepath.Join(subpath, name)

	return
}

/*
func HandleReceived(msg string, chErr chan error) {

	flds := strings.Split(msg, fieldSeperator)

	cfg := vcfg.GetPointConfig(flds[0])
	fmt.Println("*** SEKUNDOMER *** SEKUNDOMER *** SEKUNDOMER *** SEKUNDOMER ", flds, cfg)

	if nil == cfg {

		fmt.Println("==================== NEATRADU CFG")

		return

	} else {

		fmt.Println("!!!!!!!!!!!!!!!! Atradu cfg !!!!!!!!!!!!!!!!!!!!!!!")
		if vcfg.PointRunning(flds[0]) {

			fmt.Println("######################## HELLO FROM POINT", flds[0])

			//cfg.NewMsg(msg)
		} else if flds[1] == strconv.Itoa(MsgInputHelloFromPoint) {
			fmt.Println("@@@@@@@@@@@@@@@@@@@ HELLO FROM POINT")

			port, _ := strconv.Atoi(flds[3])

			udpAddr := net.UDPAddr{IP: net.ParseIP(flds[2]), Port: port}

			chMsg := make(chan string)

			go cfg.CfgRun(flds[0], udpAddr, chMsg, vomni.RootErr)
		}

		//for k, v := range cfg {
		//		fmt.Println(k, "---> ", v)
		//}

		cfg.CfgShow()

		fmt.Println("SITKOVETSKY naff NULL", cfg.CfgType())

	}

	chErr <- nil
}
*/
