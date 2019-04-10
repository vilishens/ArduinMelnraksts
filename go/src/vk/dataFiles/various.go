package dataFiles

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	//	vmsg "vk/message"
)

func (xx *dataFile) String() (str string) {
	str = fmt.Sprintf("%15s : %s\n", "TYPE", xx.dataType)
	str += fmt.Sprintf("%15s : %s\n", "NAME", xx.name)
	str += fmt.Sprintf("%15s : %s\n", "PATH", xx.path)
	str += fmt.Sprintf("%15s : %s\n", "REC-NR", strconv.Itoa(xx.currentRecordCount))
	str += fmt.Sprintf("%15s : %s\n", "FILE-NR", strconv.Itoa(xx.currentFileNbr))

	return
}

func getDataFileSlice(name string) (data []string, err error) {
	fullF := name // vutils.FileAbsPath(vparam.Params.LogPath, name)

	dBytes, err := ioutil.ReadFile(fullF) // just pass the file name
	if err != nil {
		return
	}

	data = strings.Split(string(dBytes), "\n")
	lastInd := len(data) - 1
	if "" == data[lastInd] {
		data = data[0:lastInd]
	}

	return
}

func (xx *dataFile) recordPush(record string, limit int) {

	count := len(xx.lastRecords)
	if count >= limit {
		return
	}

	if 0 == count {
		xx.lastRecords = append(xx.lastRecords, record)
	} else {
		xx.lastRecords = append([]string{record}, xx.lastRecords[:count]...)
	}
}

func checkLoaded() (pass bool) {

	pass = false

	/*


		for _, v := range vmsg.TotalData {
			tips := v.MsgType
			if _, ok := allDataFiles[tips]; !ok {
				fmt.Println("Neatradu datu tipu ***", tips)
				return
			} else {
				fmt.Println("Data type Velengurovs", tips)
			}
			for k1, v1 := range v.Data {
				_ = v1
				points := k1
				nPoints := strings.Replace(points, "*", "/", -1)
				point, ok := allDataFiles[tips][nPoints]
				if !ok {
					fmt.Println("Neatradu datu tipa ***", tips, " POINTU ", points, "(", nPoints, ")vajag ", filepath.Join(v1.Flds.Dir, v1.Flds.Base))

					for z, _ := range v.Data {
						_, ir := allDataFiles[tips][nPoints]
						fmt.Printf("Esošais \"%s\" ---> %t\n", z, ir)
					}

					for z, _ := range allDataFiles[tips] {
						fmt.Printf("#### Jaunais \"%s\" --->\n", z)
					}

					return
				} else {
					fmt.Println("\tPoint ", points)
				}

				name := filepath.Join(v1.Flds.Dir, v1.Flds.Base)
				newName := strings.Replace(points, "*", "/", -1)
				if point.name != newName {
					fmt.Println("Neatradu datu tipa ***", tips, " POINTA ", points, " vārdu ", name)
					return
				}
				if point.dataType != tips {
					fmt.Println("Neatradu datu tipa ***", tips, " POINTA ", points, " slikts tips ", tips)
					return
				}
				/ *
					taka := filepath.Join(v1.Flds.Dir, v1.Flds.Base) //+ "." + strings.ToLower(tips)
					if point.path != taka {
						fmt.Println("Neatradu datu tipa ***", tips, " POINTA ", points, " slikta taka ", taka, "(vajag ", point.path, ")")
						return
					}
				* /
				if point.currentFileNbr != v1.Flds.CurrentFileNbr {
					fmt.Println("Neatradu datu tipa ***", tips, " POINTA ", points, " files NBRa ", point.currentFileNbr)
					return
				}
				if point.currentRecordCount != v1.Flds.RecordCount {
					return
					fmt.Println("Neatradu datu tipa ***", tips, " POINTA ", points, " Recordes NBRa ", point.currentRecordCount, "(vajag ", v1.Flds.RecordCount, ")")
					return
				}

				//			continue

				if len(point.lastRecords) != len(v1.LastRecords) {
					fmt.Println("Neatradu datu tipa ***", tips, " POINTA ", points, " Recoržu skaits ", len(point.lastRecords), "(vajag ", len(v1.LastRecords), ")")
					return
				}

				for i, l := range v1.LastRecords {
					msg := l.When.Format(vomni.TimeFormat1)
					msg += vomni.UDPMessageSeparator + l.Msg
					if msg != point.lastRecords[i] {
						fmt.Println("Neatradu datu tipa ***", tips, "*** POINTA ", points, " Indexu ", i)
						fmt.Printf("ORI MSG %s (%d ??? ori %d)\n", msg, len(point.lastRecords), len(v1.LastRecords))
						for j, z := range point.lastRecords {
							fmt.Printf("mizja %2d ==> %s\n", j, z)
						}
						//					fmt.Printf("MSG %s\n\n\ng %s\n", msg, point.lastRecords[i])
						return
					}
				}
			}
		}
	*/

	return true
}
