package dataFiles

const (
	omniFileLimit       = 2
	omniRecordLimit     = 3
	omniLastRecordLimit = 5
)

//type dataExist map[string]bool

type typeFilesData map[string]*dataFile

type dataFile struct {
	name           string // path without extention as name
	path           string // path wihtout the file number and dot
	dataType       string
	currentFileNbr int // the last file number
	//	recordLimit      int
	currentRecordCount int // record count in the file
	//	fileLimit        int
	lastRecords []string // records im memory
}

type typeData struct {
	dataType        string // data type path
	recordLimit     int    // limit of records in a file
	fileLimit       int    // limit of data type file to keep
	path            string // path of the without the
	lastRecordLimit int    // limit of last records to keep in memory
}
