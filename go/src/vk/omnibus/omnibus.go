package omnibus

import (
	"log"
	"os"
	"time"
)

var RootPath string
var RootErr = make(chan error)
var RootDone = make(chan int)

var (
	LogMainFile *os.File
	LogErr      *log.Logger
	LogInfo     *log.Logger
	LogData     *log.Logger
	LogFatal    *log.Logger
)

const (
	MsgCdOutputHelloFromStation = 0x00000001 // Output <station name><msgCd><msgNbr><station UTC seconds><station offset seconds><stationIP><stationPort>
	MsgCdInputHelloFromPoint    = 0x00000002 // Input  <point name><msgCd><msgNbr><pointIP><pointPort>
	MsgCdOutputSetRelayGpio     = 0x00000004 // Output <point name><msgCd><msgNbr><Gpio><set value>
	MsgCdInputSuccess           = 0x00000008 // Input  <point name><msgCd><msgNbr>
	MsgCdInputFailed            = 0x00000010 // Input  <point name><msgCd><msgNbr><pointIP><pointPort>
	MsgCdFakeLimitExceeded      = 0x00000020 // Fake   <station name><msgCd><msgNbr>
)

const (
	MsgIndexMsgSender          = 0
	MsgIndexMsgCd              = 1
	MsgIndexMsgNbr             = 2
	MsgIndexHelloFromPointIP   = 3
	MsgIndexHelloFromPointPort = 4
)

const (
	LogFileFlags   = os.O_RDWR | os.O_CREATE | os.O_APPEND
	LogUserPerms   = os.FileMode(0666)
	LogMainPath    = "../log/main/logMain.log"
	LogLoggerFlags = log.LstdFlags | log.LUTC
	LogPrefixData  = "==== DATA === "
	LogPrefixErr   = "!!! ERROR !!! "
	LogPrefixInfo  = "**** INFO *** "
	LogPrefixFatal = "xxx FATAL xxx "
)

const (
	PointTypeRelayOnOffInterval = 0x000001
)

const (
	//	WebPrefix     = "/xK-@eRty7yZ/"
	//	WebStaticPath = "static/"

	WebPrefix     = "/xK-@eRty$Wj7yZ/"
	WebStaticPath = "webstatic/"
)

//#####################################################

const (
	LogStatusFile    = "logstatus.status"
	PointLogDataFile = "data.log"
	PointLogInfoFile = "info.log"
)

const (
	DoneOK      = 0x0000001
	DoneError   = 0x0000002
	DoneStop    = 0x0000004
	DoneRestart = 0x0000008
)

const (
	CfgFactoryPath   = "../cfg/factory/factory.cfg"
	CfgFactoryAction = ""
	CfgFldAction     = "action"
	CfgFldPath       = "path"
)

const (
	MessageTypeCmd   = "CMD"
	MessageTypeError = "ERROR"
	MessageTypeEvent = "EVENT"
	MessageTypeStart = "START"
	MessageTypeStop  = "STOP"
	MessageOmnibus   = "OMNIBUS"
)

var MessageTypeLimits = map[string]int{
	MessageOmnibus: 3} // type,point, cmd

const (
	PointModeIntervalOnOff = "time-on-off"
	PointIntervalOnOff     = "intervalOnOff"
)

const (
	Duration3PartSplitter = ":"
	UDPMessageSeparator   = ":::"
)

const (
	TimeFormat1 = "2006-01-02 15:04:05 -07:00 MST"
)

const (
	StepNameCheckNet     = "step-check-net"
	StepNameConfig       = "step-config"
	StepNameDataFiles    = "step-data-files"
	StepNameFinish       = "step-finish"
	StepNameLoadMsg      = "step-load-msg"
	StepNameNetInfo      = "step-net-info"
	StepNameParams       = "step-params"
	StepNamePoints       = "step-points"
	StepNamePointConfig  = "step-point-config"
	StepNamePointPrepare = "step-point-prepare"
	StepNamePointScan    = "step-point-scan"
	StepNameRunPoints    = "step-run-points"
	StepNameStart        = "step-start"
	StepNameUDP          = "step-udp"
	StepNameWEB          = "step-web"
	StepNameWEBInfo      = "step-web-info"
)

const (
	PointExecDelay = 10 * time.Millisecond
	StepExecDelay  = 10 * time.Millisecond
)

const (
	DIR_PERMISSIONS = 0744

//FILE_PERMISSIONS       = 0644
//GAMMU_CFG_FILE_INDEX   = 3
//GAMMU_PIN_SUBMIT_INDEX = 5
)

type WPointData struct {
	Point    string
	Active   bool
	Frozen   bool
	Descr    string
	Type     int
	CfgRun   interface{}
	CfgSaved interface{}
	Index    interface{}
	State    int
}
