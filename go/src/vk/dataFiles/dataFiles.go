package dataFiles

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
	vomni "vk/omnibus"
	vparam "vk/params"
	vutils "vk/utils"
)

var allDataFiles map[string]typeFilesData // [type][path bez numura]faila dati
var knownTypes map[string]typeData

func init() {
	allDataFiles = make(map[string]typeFilesData)
	knownTypes = make(map[string]typeData)
}

func LoadAll(chGoOn chan bool, chDone chan int, chErr chan error) {
	for k, v := range knownTypes {
		if err := v.loadAllTypeData(); err != nil {
			chErr <- vutils.ErrFuncLine(fmt.Errorf("Couldn't load data of type \"%s\" %s -- %v", v.dataType, k, err))
			return
		}

	}

	testIt := true
	if testIt {
		fmt.Println("#####################################################################")
		fmt.Println("#####################################################################", vparam.Params.EventPath)
		fmt.Println("#####################################################################")
		if checkLoaded() {
			fmt.Println("=====================================================================")
			fmt.Println("=================== OK OK OK ========================================")
			fmt.Println("=====================================================================")
		} else {
			chErr <- fmt.Errorf("Artuziff")
		}
	}
	chGoOn <- true
}

func SetKnownTypes() {
	setEventType()
	setErrorType()
}

func setEventType() {
	src := typeData{}
	src.dataType = vomni.MessageTypeEvent
	src.fileLimit = omniFileLimit
	src.lastRecordLimit = omniLastRecordLimit
	src.path = vparam.Params.EventPath
	src.recordLimit = omniRecordLimit

	knownTypes[src.dataType] = src
}

func setErrorType() {
	src := typeData{}
	src.dataType = vomni.MessageTypeError
	src.fileLimit = omniFileLimit
	src.lastRecordLimit = omniLastRecordLimit
	src.path = vparam.Params.ErrorPath
	src.recordLimit = omniRecordLimit

	knownTypes[src.dataType] = src
}

func AddRecord(msgType string, subpath string, point string, txt string) (err error) {

	record := time.Now().Format(vomni.TimeFormat1) + vomni.UDPMessageSeparator + txt

	typeObj, has := knownTypes[msgType]
	if !has {
		return vutils.ErrFuncLine(fmt.Errorf("Unknown message type \"%s\"", msgType))
	}

	typePath := typeObj.path

	fileKey := filepath.Join(subpath, point)                                  // subpath of the file with the point name at the end
	filePath := filepath.Join(typePath, fileKey+"."+strings.ToLower(msgType)) // file path without file number

	fmt.Println("======================> JAPIEVIENO recorde ", filePath, "PARAMS", vparam.Params.ErrorPath)

	typeObj.setNewFileData(fileKey, filePath)

	recLimit := typeObj.recordLimit
	fileLimit := typeObj.fileLimit
	memoryLimit := typeObj.lastRecordLimit

	fullPath := ""
	dst := allDataFiles[msgType][fileKey]
	recNow := dst.currentRecordCount
	fileNbr := dst.currentFileNbr
	fullPath = dst.path + "." + strconv.Itoa(fileNbr)

	newFile := false
	if recNow >= recLimit {
		newFile = true
		fullPath = dst.path + "." + strconv.Itoa(fileNbr+1)
	}

	fmt.Println("0C !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	fmt.Printf("****** FILE %2d *************************************** IERAKSTI %2d Limits %2d NEW %t\n",
		dst.currentFileNbr, recNow, recLimit, newFile)
	fmt.Println("0C !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")

	if err = vutils.FileAppend(fullPath, record+"\n"); nil != err {
		return vutils.ErrFuncLine(fmt.Errorf("Couldn't add the \"%s\" record \"%s\" --- %v", msgType, record, err))
	}

	dst.inMemoryRecord(record, memoryLimit)

	dst.currentRecordCount++

	if newFile {
		fmt.Println("0K !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
		fmt.Println("NEW FILE ", newFile, "REC NBR ", dst.currentRecordCount, "REC LIMIT", recLimit)
		fmt.Println("0K !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")

		//		fmt.Printf("   TYPE: %s\nSUBPATH: %s\n  POINT: %s\n    TXT: %s\n RECORD: %s\n   FILE: %s\n   PATH: %s\n",
		//			msgType, subpath, point, txt, record, file, vparam.Params.EventPath)

		fmt.Printf("################# Curr %d **** Limit %d\n", dst.currentFileNbr, recLimit)

		dst.currentFileNbr++
		dst.currentRecordCount = 1

		fmt.Println("PARNAKIVI")

		if dst.currentFileNbr >= fileLimit {
			if err = limitFiles(allDataFiles[msgType][fileKey].path, fileLimit); nil != err {
				return vutils.ErrFuncLine(fmt.Errorf("Couldn't clean data files of  \"%s\" --- %v", msgType, err))
			}
		}
	}

	//if
	fmt.Println("0Z !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	fmt.Println("NEW FILE ", newFile, "REC NBR ", dst.currentRecordCount, "REC LIMIT", recLimit)
	fmt.Println("0Z !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")

	fmt.Println("oooooooooooooooooooooo", dst.name, "ooooooooooooooooooooooooooo")
	for k, v := range dst.lastRecords {
		fmt.Println(k+1, " ====> ", v)
	}

	fmt.Printf("$$$$$$$$$$$$$$$$$$$$$$$$$$$ VASHUKOFF \n%s", dst)

	return
}

func (xx *dataFile) inMemoryRecord(record string, limit int) {

	count := len(xx.lastRecords)

	if count >= limit {
		xx.lastRecords = xx.lastRecords[1:]
	}

	xx.lastRecords = append(xx.lastRecords, record)
}

func limitFiles(path string, limit int) (err error) {

	fmt.Println("ALONA SVIRIDOVA FLAMINGO")

	thisDir := filepath.Dir(path)
	thisBase := filepath.Base(path)

	fileData, err := ioutil.ReadDir(thisDir)
	if nil != err {
		return vutils.ErrFuncLine(err)
	}

	nbrLst := []int{}

	fmt.Println("ALONA SVIRIDOVA Path", path, " LIMIT", limit)

	for _, v := range fileData {
		if !v.IsDir() {
			name := v.Name()
			pos := strings.LastIndex(name, ".")

			if 0 <= pos {
				base := name[:pos]
				fmt.Printf("BASE %s ??? ReadBASE %s\n", thisBase, base)
				if base == thisBase {
					fNbr, err := strconv.Atoi(name[pos+1:])
					if nil != err {
						return vutils.ErrFuncLine(err)
					}
					nbrLst = append(nbrLst, fNbr)
				}
			}
		}
	}

	if len(nbrLst) <= limit {
		return
	}

	fmt.Println("Vaka", nbrLst)

	sort.Sort(sort.Reverse(sort.IntSlice(nbrLst)))

	fmt.Println("Satu", nbrLst)

	for i, v := range nbrLst {
		if i >= limit {
			newPath := path + "." + strconv.Itoa(v)
			if err = vutils.FileDelete(newPath); nil != err {
				return vutils.ErrFuncLine(err)
			}
		}
	}

	return
}
