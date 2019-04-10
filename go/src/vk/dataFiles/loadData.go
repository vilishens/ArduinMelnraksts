package dataFiles

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	vutils "vk/utils"
)

func (xx typeData) loadTypeData(files map[string][]*dataFile) (err error) {

	typeNow := xx.dataType

	if _, has := allDataFiles[typeNow]; !has {
		allDataFiles[typeNow] = make(typeFilesData)
	}

	//	allDataFiles map[string]typeFilesData
	for k, v := range files {
		if err = xx.loadPointData(v); nil != err {
			return vutils.ErrFuncLine(fmt.Errorf("Couldn't load \"%s\" of \"%s\" data -- %v",
				k, typeNow, err))
		}
	}

	return
}

func (xx typeData) loadPointData(files []*dataFile) (err error) {

	arrFileNbrs := pointFileNbrs(files)
	sort.Sort(sort.Reverse(sort.IntSlice(arrFileNbrs)))

	if len := len(arrFileNbrs); 0 == len {
		return
	}

	xx.setNewFileData(files[0].name, files[0].path) // any iteme of point can be used as all pathes and names are equal
	// only file number differs
	dst := allDataFiles[xx.dataType][files[0].name]
	dst.currentFileNbr = arrFileNbrs[0]

	limit := xx.lastRecordLimit

	for k, v := range arrFileNbrs {
		full := files[0].path + "." + strconv.Itoa(v)

		records, err := getDataFileSlice(full)
		if nil != err {
			return vutils.ErrFuncLine(fmt.Errorf("Couldn't read \"%s\" -- %v",
				full, err))
		}
		count := len(records)
		if 0 == k {
			// remember the last file record count, it it the current file record count
			dst.currentRecordCount = len(records)
		}

		for i := count - 1; (0 <= i) && (len(dst.lastRecords) < limit); i-- {
			dst.recordPush(records[i], limit)
		}
	}

	return
}

func pointFileNbrs(files []*dataFile) (lst []int) {

	for _, v := range files {
		lst = append(lst, v.currentFileNbr)
	}

	return
}

func (xx typeData) loadAllTypeData() (err error) {

	files, err := xx.allTypeFiles()

	if 0 == len(files) {
		// no data type files found
		return
	}

	if nil != err {
		return vutils.ErrFuncLine(fmt.Errorf("Type \"%s\" file list failed -- %v", xx.dataType, err))
	}

	err = xx.loadTypeData(files)
	if nil != err {
		return vutils.ErrFuncLine(fmt.Errorf("Type \"%s\" file data load failed -- %v", xx.dataType, err))
	}

	koza := false
	if koza {
		for k, v := range files {
			fmt.Println("**** ", k, "*******")
			for k1, v1 := range v {
				fmt.Println("\t#### ", k1, "*******", v1)
				/*
					for _, v2 := range *(v1) {
						fmt.Println("\t\tNAME %s\n\t\tPATH %s\n\t\tNBR %d\n\t\tTYPE %s",
							v1.name, v1.path, v1.fileNbr, v1.dataType)

					}
				*/

			}
		}
	}

	return
}

func (xx typeData) allTypeFiles() (dst map[string][]*dataFile, err error) {

	has, err := vutils.PathExists(xx.path)
	if !has {
		//  data file directory doesn't exist, no need to continue
		return
	}

	files, err := ioutil.ReadDir(xx.path)
	if nil != err {
		return
	}

	dst = make(map[string][]*dataFile)

	thisDir := ""
	for _, f := range files {
		if f.IsDir() {
			newDir := filepath.Join(thisDir, f.Name())
			if err = xx.allDirFiles(newDir, dst); nil != err {
				err = vutils.ErrFuncLine(fmt.Errorf("Dir \"%s\" handling failed", newDir))
				return
			}
			continue
		}

		if err = xx.addNewTypeFile2List(thisDir, f.Name(), strings.ToLower(xx.dataType), dst); nil != err {
			err = vutils.ErrFuncLine(fmt.Errorf("Dir \"%s\" with \"%s\" handling failed -- %v", thisDir, f.Name(), err))
			return
		}
	}

	return
}

func (xx typeData) allDirFiles(sDir string, dst map[string][]*dataFile) (err error) {
	thisDir := filepath.Join(xx.path, sDir)

	files, err := ioutil.ReadDir(thisDir)
	if nil != err {
		return
	}

	for _, f := range files {
		if f.IsDir() {
			newDir := filepath.Join(sDir, f.Name())
			if err = xx.allDirFiles(newDir, dst); nil != err {
				return vutils.ErrFuncLine(fmt.Errorf("Dir \"%s\" handling failed -- %v", newDir, err))
			}
			continue
		}

		if err = xx.addNewTypeFile2List(sDir, f.Name(), strings.ToLower(xx.dataType), dst); nil != err {
			return vutils.ErrFuncLine(fmt.Errorf("Dir \"%s\" with \"%s\" handling failed -- %v", sDir, f.Name(), err))
		}
	}
	return
}

func (xx typeData) addNewTypeFile2List(sDir string, fName string, fEnding string, dst map[string][]*dataFile) (err error) {
	tFile := new(dataFile)

	nbr, err := xx.fileWithType(sDir, fName, fEnding, tFile)
	if 0 > nbr {
		return
	}
	if nil != err {
		return err
	}

	if 0 <= nbr {
		name := tFile.name
		if _, has := dst[name]; !has {
			dst[name] = []*dataFile{}
		}

		dst[name] = append(dst[name], tFile)
	}

	return
}

func (xx typeData) fileWithType(fDir string, fName string, fEnding string, dst *dataFile) (fileNbr int, err error) {

	fileNbr = -7 // set any negative (non valid) value

	parts := strings.Split(fName, ".")
	if (3 == len(parts)) && (fEnding == parts[1]) {
		// the file name contains all 3 parts <POINT-NAME>.<TYPE>.<FILE-NBR> with the required type
		fileNbr, err = strconv.Atoi(parts[2])
		if nil != err {
			// the file name doesn't end with number, it is not a valid message file
			return
		}

		dst.name = filepath.Join(fDir, parts[0])
		dst.path = filepath.Join(xx.path, fDir, parts[0]+"."+parts[1])
		dst.currentFileNbr = fileNbr
	}

	return
}

func (xx typeData) setNewFileData(name string, path string) {
	if _, has := allDataFiles[xx.dataType]; !has {
		allDataFiles[xx.dataType] = make(typeFilesData)
	}

	if _, has := allDataFiles[xx.dataType][name]; !has {
		newFile := new(dataFile)
		newFile.name = name
		newFile.path = path
		newFile.dataType = xx.dataType
		newFile.currentFileNbr = 0
		newFile.currentRecordCount = 0
		newFile.lastRecords = []string{}

		allDataFiles[newFile.dataType][name] = newFile
	}
}
