package points

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	vomni "vk/omnibus"
	vutils "vk/utils"
)

var AllPoints map[string]HandlePoint

func init() {
	AllPoints = make(map[string]HandlePoint)
	AllIntervalOnOff = make(map[string]*PointIntervalOnOff)
}

func Run(chGoOn chan bool, chDone chan int, chErr chan error) {

	chGoOnIntervalOnOff := make(chan bool)
	chDoneIntervalOnOff := make(chan bool)
	chErrIntervalOnOff := make(chan error)

	chGoOn <- true
	return

	go runIntervalOnOff(chGoOnIntervalOnOff, chDoneIntervalOnOff, chErrIntervalOnOff)

	end := false
	for !end {
		select {
		case <-chGoOnIntervalOnOff:
			chGoOn <- true
		case <-chDoneIntervalOnOff:
			end = true
			chDone <- vomni.DoneOK
		case err := <-chErrIntervalOnOff:
			end = true
			chErr <- vutils.ErrFuncLine(err)
		}
	}
}

func StartPoint(msg string, addr net.UDPAddr) (back string, err error) {

	flds := strings.Split(msg, vomni.UDPMessageSeparator)
	// type:::point-name:::UDP-server-port
	if 3 != len(flds) {
		err = vutils.ErrFuncLine(fmt.Errorf("Incorrect data format of \"%s\"", msg))
		return
	}

	point := strings.Trim(flds[1], " ")
	port := strings.Trim(flds[2], " ")
	if iFace, ok := AllPoints[point]; !ok {
		err = vutils.ErrFuncLine(fmt.Errorf("Invalid point name \"%s\"", point))
		return
	} else {
		iFace.Start(port, addr)
	}

	//back, err = handlePointStart(point, port, addr)

	return
}

func Cmd(msg string, addr net.UDPAddr) (back string, err error) {
	flds := strings.Split(msg, vomni.UDPMessageSeparator)
	// type:::point-name:::UDP-server-port
	if 3 != len(flds) {
		err = vutils.ErrFuncLine(fmt.Errorf("Incorrect data format of \"%s\"", msg))
		return
	}

	point := strings.Trim(flds[1], " ")
	if p, ok := AllIntervalOnOff[point]; !ok {
		err = vutils.ErrFuncLine(fmt.Errorf("Invalid point name \"%s\"", point))
		return
	} else if !p.Active {
		back = fmt.Sprintf("Point \"%s\" isn't active", point)
	} else {
		p.ChMsg <- "Getmanov"
	}

	return
}

func handlePointStart(point string, port string, addr net.UDPAddr) (back string, err error) {

	data := AllIntervalOnOff[point]

	back = "Init failed"

	if !data.Active {
		data.Active = true
		data.UDPAddr = addr
		if data.UDPAddr.Port, err = strconv.Atoi(port); nil != err {
			err = vutils.ErrFuncLine(fmt.Errorf("Wrong UPD Port number \"%s\" --- %v", port, err))
			return
		}
	}

	back = "OK"

	return
}

func GetPointUDPAddr(point string) (addr net.UDPAddr, err error) {
	val, ok := AllIntervalOnOff[point]
	if !ok {
		err = vutils.ErrFuncLine(fmt.Errorf("Invalid point name \"%s\"", point))
		return
	}

	return val.UDPAddr, nil
}
