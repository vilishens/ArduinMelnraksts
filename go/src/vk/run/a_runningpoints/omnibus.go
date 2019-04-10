package a_runningpoints

import (
	"net"
	"time"
	"vk/omnibus"
)

const (
	TypeRelayOnOffInterval = 0x0001
)

const (
	msgCdOutputHelloFromStation = 0x00000001 // Output <station name><msgCd><msgNbr><station UTC seconds><station time offset><stationIP><stationPort>
	msgCdInputHelloFromPoint    = 0x00000002 // Input  <point name><msgCd><msgNbr><pointIP><pointPort>
	msgCdOutputSetRelayGpio     = 0x00000004 // Output <point name><msgCd><msgNbr><Gpio><set value>
)

const (
	indexMsgSender = 0
	indexMsgCd     = 1
	indexMsgNbr    = 2
	lenPrefix      = 3
)

const (
	fieldSeparator = ":::"
)

type runningPoint struct {
	name      string
	pointType int

	udpAddr net.UDPAddr
	msg     chan string
}

type runPoints map[string]runningPoint

type Runner interface {
	GetUDPAddr() (addr net.UDPAddr)
	IsActive() (active bool)
	LetsGo(chGoOn chan bool, chErr chan error)
	LogPointStr(cd int, logStr string)
	RotateReAssign() (err error)
	Response(msg []string, chDelete chan bool, chErr chan error)
	SetUDPAddr(ip string, port int)
	WebPointData() (v omnibus.WPointData)
	ReceivedWebMsg(msg string, data interface{})
	Finish(done chan bool)
}

type PointItem struct {
	Name    string
	UDPAddr net.UDPAddr
	Cfg     Runner
	Msg     chan string
	Err     chan error
}

type SendMsg struct {
	UDPAddr net.UDPAddr
	MsgNbr  int // ja == "", tad jādzēš, jo nav ko sūtīt
	Repeat  int
	Msg     string
	Last    time.Time
}

type SendMsgArr []*SendMsg

const (
	indexHelloFromStationTime   = 0
	indexHelloFromStationOffset = 1
	indexHelloFromStationIP     = 2
	indexHelloFromStationPort   = 3
	lenHelloFromStation         = 4
)

const (
	indexHelloFromPointIP   = 3
	indexHelloFromPointPort = 4
	lenHelloFromPoint       = 5
)
