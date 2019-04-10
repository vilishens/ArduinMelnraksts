package message

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"

	vomni "vk/omnibus"
	vparam "vk/params"

	//	xrun "vk/run/a_runningpoints"
	vutils "vk/utils"
)

var msgSequenceNbr = 0
var SendList SendMsgArr

func init() {
	msgSequenceNbr = 0
	SendList = SendMsgArr{}
}

func MessageGet(msg string, chErr chan error) {

	var err error
	flds := strings.Split(msg, vomni.UDPMessageSeparator)

	msgNbr, err := strconv.Atoi(flds[indexMsgNbr])
	if nil != err {
		chErr <- vutils.ErrFuncLine(err)
		return
	}

	locErr := make(chan error)
	locDelete := make(chan bool)

	// garanichev

	fmt.Printf("!!!!!!!!!!!!!!!!!!!!!! %s\n", msg)
	//go xrun.MessageGet(flds, locDelete, locErr)
	select {
	case <-locDelete:
		SendList.minusNbr(msgNbr)
	case err = <-locErr:
		break
	}

	chErr <- err
}

func messageLocal(flds []string, chErr chan error) {
	msgCd, err := strconv.Atoi(flds[indexMsgCd])
	if nil != err {
		chErr <- vutils.ErrFuncLine(err)
		return
	}

	msgNbr := -1

	switch msgCd {
	case vomni.MsgCdInputSuccess, vomni.MsgCdInputFailed:
		msgNbr, err = strconv.Atoi(flds[indexMsgNbr])
		if nil != err {
			chErr <- vutils.ErrFuncLine(err)
			return
		}

		if msgCd == vomni.MsgCdInputFailed {
			fmt.Println("FAILED <<< >>>", msgNbr, "<<< >>> FAILED")

		}
		if msgCd == vomni.MsgCdInputSuccess {
			fmt.Println("SUCCES <<< >>>", msgNbr, "<<< >>> SUCCES")

		}

		SendList.minusNbr(msgNbr)
	}

	chErr <- err

}

func TryHello(dst net.UDPAddr, chDone chan bool, chErr chan error) {

	msgData := stationHelloData()

	fmt.Println("================== try ================================> _______ SITKO _______ DST ", dst)

	if err := MessageSend(vomni.MsgCdOutputHelloFromStation, dst, msgData, "", ""); nil != err {
		err = vutils.ErrFuncLine(err)
		vomni.LogErr.Println(err)
		chErr <- err
	} else {
		chDone <- true
	}
}

func stationHelloData() (d []string) {

	_, tzSecs := time.Now().Zone()

	d = make([]string, lenHelloFromStation)

	d[indexHelloFromStationTime] = strconv.Itoa(int(time.Now().Unix()))
	d[indexHelloFromStationOffset] = strconv.Itoa(tzSecs)
	d[indexHelloFromStationIP] = vparam.Params.InternalIPv4
	d[indexHelloFromStationPort] = strconv.Itoa(vparam.Params.UDPPort)

	return
}

func MessageSend(cd int, addr net.UDPAddr, data []string, point string, logFormat string) (err error) {

	nbr, prefix := messagePrefix(cd)

	var d []string
	d = append(d, prefix...)
	d = append(d, data...)

	msgStr := messageString(d)

	addSendList(addr, nbr, msgStr)

	fmt.Printf("==================================> ADDR %v NBR %v MSG %s\n", addr, nbr, msgStr)

	if "" != logFormat {
		strLog := fmt.Sprintf(logFormat, nbr)

		// garanichev
		_ = strLog
		//xrun.PointLogAdd(point, cd, strLog)
	}

	return
}

/*
func PointLogAdd(point string, cd int, logString string) {

	fmt.Println("\n#\n #\n  #\n   #\nROZETKA Log from A")

	Points[point].LogPointStr(cd, logString)
}
*/
func messagePrefix(cd int) (nbr int, d []string) {

	d = make([]string, lenPrefix)

	lock := new(sync.Mutex)
	lock.Lock()

	msgSequenceNbr++

	nbr = msgSequenceNbr

	lock.Unlock()

	d[indexMsgSender] = vparam.Params.Name
	d[indexMsgCd] = strconv.Itoa(cd)
	d[indexMsgNbr] = strconv.Itoa(nbr)

	return
}

func addSendList(dst net.UDPAddr, nbr int, msg string) {
	newMsg := new(SendMsg)
	newMsg.UDPAddr = dst
	newMsg.MsgNbr = nbr
	newMsg.Repeat = -1
	newMsg.Msg = msg

	SendList.plus(newMsg)
}

func (d SendMsgArr) plus(msg *SendMsg) {
	lock := new(sync.Mutex)

	lock.Lock()
	defer lock.Unlock()

	SendList = append(SendList, msg)
}

func (d SendMsgArr) minusNbr(nbr int) {
	ind := -1
	for key, val := range d {
		if val.MsgNbr == nbr {
			ind = key
			break
		}
	}

	if -1 == ind {
		fmt.Printf("Received MSG #%d without record\n", nbr)
		return
	}

	chDone := make(chan bool)
	go d.MinusIndex(ind, chDone)
	<-chDone

	fmt.Printf(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>> Minus NBR %d (IND %d)\n", nbr, ind)
}

func (d SendMsgArr) MinusIndex(ind int, chDone chan bool) {

	fmt.Printf("MALIKORN IND %d LEN %d\n", ind, len(d))

	if ind < len(d) {
		lock := new(sync.Mutex)
		lock.Lock()
		defer lock.Unlock()

		SendList = append(SendList[:ind], SendList[ind+1:]...)
	}

	chDone <- true
}

func MessageFake(newCd int, msg string) (err error) {

	flds := strings.Split(msg, vomni.UDPMessageSeparator)
	flds[indexMsgCd] = strconv.Itoa(newCd)

	msgStr := messageString(flds)

	chErr := make(chan error)
	go MessageGet(msgStr, chErr)

	err = <-chErr

	return
}

func messageString(flds []string) (mStr string) {
	mStr = ""
	last := len(flds) - 1
	for ind, str := range flds {
		mStr += str
		if ind < last {
			mStr += vomni.UDPMessageSeparator
		}
	}

	return
}
