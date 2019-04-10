package relayonoffinterval

import (
	"log"
	"net"
	"os"
	"time"
)

const (
	indexSetRelayGpioPin = 0
	indexSetRelayGpioSet = 1
	lenSetRelayGpio      = 2
)

const (
	doZero     = 0x0000
	doListEnd  = 0x0001
	doRestart  = 0x0002
	doFreeze   = 0x0004
	doUnfreeze = 0x0008
	doStopNow  = 0x0010
)

const (
	stateNone   = 0x0000
	stateActive = 0x0001
	stateFreeze = 0x0002
)

//*** Start - Running Configuration
type RunRelOnOffInterval struct {
	Gpio    int
	State   int
	Seconds time.Duration
}

type RunRelOnOffIntervalArr []RunRelOnOffInterval

type RunRelOnOffIntervalStruct struct {
	Start  RunRelOnOffIntervalArr
	Base   RunRelOnOffIntervalArr
	Finish RunRelOnOffIntervalArr
}

type RunRelOnOffIntervalPoints map[string]RunRelOnOffIntervalStruct

type runningCfgs struct {
	RelayOnOffInterval RunRelOnOffIntervalPoints
}

//*** End - Running Configuration

type PointItem struct {
	Point       string
	UDPAddr     net.UDPAddr
	CfgRun      RunRelOnOffIntervalStruct
	CfgSaved    RunRelOnOffIntervalStruct
	ChMsg       chan string
	ChErr       chan error
	ChDone      chan int
	LogData     *log.Logger
	LogDataFile *os.File
	LogInfo     *log.Logger
	LogInfoFile *os.File
	Index       RunRelOnOffIntervalIndex
	Type        int
	State       int
}

type RunRelOnOffIntervalIndex struct {
	Start  int
	Base   int
	Finish int
}

type webPoint struct {
	Gpio     string
	State    string
	Interval string
}

type webPointArr []webPoint

type webPointStruct struct {
	Start  webPointArr
	Base   webPointArr
	Finish webPointArr
}

type webPointSaveStruct struct {
	Start  webPointSaveArr
	Base   webPointSaveArr
	Finish webPointSaveArr
}

type webPointSave struct {
	Gpio     string
	State    string
	Interval string
}

type webPointSaveArr []webPointSave

type listPart struct {
	last bool   // is this part the last one in the list?
	once bool   // run records once only?
	name string // the name of the part
	data RunRelOnOffIntervalArr
}

const (
	listPartStart  = "START"
	listPartBase   = "BASE"
	listPartFinish = "FINISH"
)
