package points

import (
	"fmt"
	"net"
	"strconv"
	"time"
	vomni "vk/omnibus"
	vutils "vk/utils"
)

var AllIntervalOnOff map[string]*PointIntervalOnOff

var Mika chan bool

func init() {
	Mika = make(chan bool)
}

func runIntervalOnOff(chGoOn chan bool, chDone chan bool, chErr chan error) {

	chGoOn <- true
	return

	for k, v := range AllIntervalOnOff {
		AllPoints[k] = v
		go v.run(chDone, chErr)
	}

	chGoOn <- true

	end := false

	for !end {
		time.Sleep(vomni.PointExecDelay)
		select {
		case <-chDone:
			end = true
		case <-chErr:
			end = true
		}
	}
}

func (xxx *PointIntervalOnOff) DoCmd(cmd string) (back string, err error) {
	return
}

func (xxx *PointIntervalOnOff) Stop() (back string, err error) {

	xxx.Index = 0
	xxx.Active = false

	return
}

func (xxx *PointIntervalOnOff) Start(port string, addr net.UDPAddr) (back string, err error) {

	data := AllIntervalOnOff[xxx.Point]

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

func (xxx *PointIntervalOnOff) run(chDone chan bool, chErr chan error) {

	end := false

	xxx.Index = 0

	tic := new(time.Ticker)

	skaits := 0
	KOTJO := 0

	for !xxx.Active || (0 == len(xxx.Sequence)) {
		time.Sleep(vomni.PointExecDelay)
	}

	if err := xxx.sendMsg(); nil != err {
		chErr <- vutils.ErrFuncLine(fmt.Errorf("%s ---> %v", xxx.Point, err))
		return
	} else {
		tic = time.NewTicker(xxx.Sequence[xxx.Index].Interval)
	}

	for !end {
		var err error

		select {
		case err = <-xxx.ChErr:
			end = true
			chErr <- vutils.ErrFuncLine(fmt.Errorf("%s ---> %v", xxx.Point, err))
		case done := <-xxx.ChDone:
			end = true
			chDone <- done
		case str := <-xxx.ChMsg:
			skaits++
			fmt.Printf("%3d ************************** POINT %s received MSG %s\n", skaits, xxx.Point, str)
		case <-tic.C:
			if xxx.Point == "mahima" {
				KOTJO++
				t := time.Now()
				timeStr := t.Format("2006-01-02 15:04:05 -07:00 MST")
				fmt.Printf("%d -- %s zzzzzzzzzzzzzz Jāsūta komanda zzzzzzzzzzzzzzzzzzzz %s zzzzzzzzzzzzzzzz %v\n",
					KOTJO, timeStr, xxx.Point, xxx.Sequence[xxx.Index].Interval)
			}

			xxx.Index++
			if xxx.Index >= len(xxx.Sequence) {
				xxx.Index = 0
			}

			if err = xxx.sendMsg(); nil != err {
				end = true
				break
				chErr <- vutils.ErrFuncLine(fmt.Errorf("%s ---> %v", xxx.Point, err))
			}

			tic = time.NewTicker(xxx.Sequence[xxx.Index].Interval)
		}
	}
}

func (xxx *PointIntervalOnOff) sendMsg() (err error) {

	msg := xxx.intervalMsg(xxx.Index)
	return send2Address(xxx.UDPAddr, msg)
}

func (xxx *PointIntervalOnOff) intervalMsg(index int) (str string) {

	str = strconv.Itoa(xxx.Sequence[index].Pin) + vomni.UDPMessageSeparator
	str += strconv.Itoa(xxx.Sequence[index].State)

	return
}

func send2Address(addr net.UDPAddr, msg string) (err error) {
	addr_str := addr.IP.String() + ":" + strconv.Itoa(addr.Port)
	return sendToAddress(addr_str, msg)
}

func sendToAddress(addr string, msg string) (err error) {

	conn, err := net.Dial("udp", addr)
	if err != nil {
		err = vutils.ErrFuncLine(fmt.Errorf("Connection ERROR: %v\n", err))
		return
	}
	defer conn.Close()

	if _, err = conn.Write([]byte(msg)); nil != err {
		err = vutils.ErrFuncLine(fmt.Errorf("SentToAddress ERROR: %v\n", err))
	}

	return
}
