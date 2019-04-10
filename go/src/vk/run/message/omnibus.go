package message

import (
	"net"
	"time"
)

/*
const (
	MsgCdOutputHelloFromStation = 0x00000001 // Output <station name><msgCd><msgNbr><station UTC seconds><station offset seconds><stationIP><stationPort>
	MsgCdInputHelloFromPoint    = 0x00000002 // Input  <point name><msgCd><msgNbr><pointIP><pointPort>
	MsgCdOutputSetRelayGpio     = 0x00000004 // Output <point name><msgCd><msgNbr><Gpio><set value>
	MsgCdInputSuccess           = 0x00000008 // Input  <point name><msgCd><msgNbr>
	MsgCdInputFailed            = 0x00000010 // Input  <point name><msgCd><msgNbr><pointIP><pointPort>
	MsgCdFakeLimitExceeded      = 0x00000020 // Fake   <station name><msgCd><msgNbr>
)
*/

const (
	indexMsgSender = 0
	indexMsgCd     = 1
	indexMsgNbr    = 2
	lenPrefix      = 3
)

const (
	fieldSeparator = ":::"
)

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
