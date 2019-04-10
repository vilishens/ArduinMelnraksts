package message

import "time"

const (
	MsgTypeLog    = "LOG"
	logLastNbr    = 5
	logFileEnding = ".log"
)

const (
	MsgTypeEvent = "EVENT"
)

const (
	MsgOutputHelloFromStation = 0x00000001 // Output <station name><msgCd><msgNbr><station time><stationIP><stationPort>
	MsgInputHelloFromPoint    = 0x00000002 // Input  <point name><msgCd><msgNbr><pointIP><pointPort>
	MsgOutputSetRelayGpio     = 0x00000004 // Output <point name><msgCd><msgNbr><Gpio><set value>
)

const (
	fieldSeperator = ":::"
)

/*
const (
	eventFileRecordMax       = 3
	eventMsgType             = "EVENT"
	eventRememberLastRecords = 5
	eventLogFileEnding       = ".event."
)
*/

const (
	//	dataLastNbr    = 5
	FieldSeparator    = ":::"
	UDPFieldSeparator = ":::" // Field separator
	FileSeparator     = "."
	PointSeparator    = "*"
	TimeFormat        = "2006-01-02 15:04:05 -07:00 MST"
)

//type logPoint struct {
//	flds pointFlds
//	last [logLastNbr]msgRecord
//}

//type eventPoint struct {
//	flds pointFlds
//	last [eventLogLastNbr]msgRecord
//}

type MsgRecord struct {
	When time.Time
	Msg  string
}

type PointFlds struct {
	Base            string
	Dir             string
	RecordCount     int
	CurrentFileNbr  int
	CurrentFileName string
}

type PointData struct {
	Flds        PointFlds
	LastRecords []MsgRecord
}

type dataAll map[string]*PointData

//var logActivePoints map[string]logPoint
var eventActivePoints = make(dataAll)

type separateType struct {
	Data          dataAll // all data of a type: record and file parameters, some records in memory
	fileRecordMax int     // max record number in a file
	MsgType       string  // message type !!! uppercase !!!
	memoryRecords int     // number of the records to keep in memory

	//	eventFileRecordMax       = 3
	//	eventMsgType             = "EVENT"
	//	eventRememberLastRecords = 5
	//	eventLogFileEnding       = ".event."

}

var TotalData []*separateType

type pointFiles map[int]string             // [fileNbr]path with file name
type allPointFiles map[string]pointFiles   // [point name] point file map
type allTypeFiles map[string]allPointFiles // [type] allPointFiles
