package relayonoffinterval

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	vpointcfg "vk/code/config/pointconfig"
	vutils "vk/utils"
)

func webInterface2Struct(inter interface{}) (str RunRelOnOffIntervalStruct) {
	// WEB struct
	web := webPointStruct{}
	for part, v := range inter.(map[string]interface{}) { // list add configuration parts
		d := webPointArr{}                     // array for the configuration part records
		for _, v1 := range v.([]interface{}) { // fill part record array
			rec := webPoint{} // storage for a record data
			for k2, v2 := range v1.(map[string]interface{}) {
				switch strings.ToUpper(k2) {
				case "GPIO":
					rec.Gpio = v2.(string)
				case "STATE":
					rec.State = v2.(string)
				case "INTERVAL":
					rec.Interval = v2.(string)
				default:
					log.Fatal(fmt.Sprintf("Unknow WEB interface record field \"%s\"", k2))
				}
			}
			d = append(d, rec)
		}

		switch strings.ToUpper(part) {
		case "START":
			web.Start = d
		case "BASE":
			web.Base = d
		case "FINISH":
			web.Finish = d
		default:
			log.Fatal(fmt.Sprintf("Unknow WEB interface part \"%s\"", part))
		}
	}

	// from the WEB structure to the regular one
	str = RunRelOnOffIntervalStruct{}
	str.Start = web.Start.webArray2Regular()
	str.Base = web.Base.webArray2Regular()
	str.Finish = web.Finish.webArray2Regular()

	return
}

/*
func webInterface2CfgStruct(inter interface{}) (web vpointcfg.CfgRelOnOffIntervalStruct) {
	// WEB struct
	web = vpointcfg.CfgRelOnOffIntervalStruct{}

	for part, v := range inter.(map[string]interface{}) { // list add configuration parts
		d := webPointSaveArr{}                 // array for the configuration part records
		for _, v1 := range v.([]interface{}) { // fill part record array
			rec := webPointSave{} // storage for a record data

			for k2, v2 := range v1.(map[string]interface{}) {
				switch strings.ToUpper(k2) {
				case "GPIO":
					rec.Gpio = v2.(string)
				case "STATE":
					rec.State = v2.(string)
				case "INTERVAL":
					rec.Interval = vutils.DurationOf3PartMinimizeStr(v2.(string))
				default:
					log.Fatal(fmt.Sprintf("Unknow WEB interface record field \"%s\"", k2))
				}
			}

			d = append(d, rec)
		}

		switch strings.ToUpper(part) {
		case "START":
			web.Start = d
		case "BASE":
			web.Base = d
		case "FINISH":
			web.Finish = d
		default:
			log.Fatal(fmt.Sprintf("Unknow WEB interface part \"%s\"", part))
		}
	}

	/ *
		// from the WEB structure to the regular one
		str = RunRelOnOffIntervalStruct{}
		str.Start = web.Start.webArray2Regular()
		str.Base = web.Base.webArray2Regular()
		str.Finish = web.Finish.webArray2Regular()
	* /
	return
}
*/

func webInterface2SaveCfg(inter interface{}) (web vpointcfg.CfgRelOnOffIntervalStruct) {
	// WEB struct
	web = vpointcfg.CfgRelOnOffIntervalStruct{}

	for part, v := range inter.(map[string]interface{}) { // list add configuration parts
		d := vpointcfg.CfgRelOnOffIntervalArr{} // array for the configuration part records
		for _, v1 := range v.([]interface{}) {  // fill part record array
			rec := vpointcfg.CfgRelOnOffInterval{} // storage for a record data

			for k2, v2 := range v1.(map[string]interface{}) {
				switch strings.ToUpper(k2) {
				case "GPIO":
					rec.Gpio = v2.(string)
				case "STATE":
					rec.State = v2.(string)
				case "INTERVAL":
					rec.Interval = vutils.DurationOf3PartMinimizeStr(v2.(string))
				default:
					log.Fatal(fmt.Sprintf("Unknow WEB interface record field \"%s\"", k2))
				}
			}

			d = append(d, rec)
		}

		switch strings.ToUpper(part) {
		case "START":
			web.Start = d
		case "BASE":
			web.Base = d
		case "FINISH":
			web.Finish = d
		default:
			log.Fatal(fmt.Sprintf("Unknow WEB interface part \"%s\"", part))
		}
	}

	return
}

func webDefaultCfgStruct(point string) (d RunRelOnOffIntervalStruct, err error) {

	whole := vpointcfg.CfgPointData{}
	if whole, err = vpointcfg.CfgDefault(); nil != err {
		return RunRelOnOffIntervalStruct{}, vutils.ErrFuncLine(err)
	}

	if d, err = PutRelOnOffIntervalConfiguration(whole.RelOnOffIntervJSON[point]); nil != err {
		return RunRelOnOffIntervalStruct{}, vutils.ErrFuncLine(err)
	}

	return
}

func webSavePointCfg(point string, data interface{}) (err error) {

	whole := vpointcfg.CfgPoints
	newData := webInterface2SaveCfg(data)

	whole.RelOnOffIntervJSON[point] = newData

	return whole.Save()
}

func (d webPointArr) webArray2Regular() (newStr RunRelOnOffIntervalArr) {

	newStr = RunRelOnOffIntervalArr{}

	for _, v := range d {
		newR := RunRelOnOffInterval{}
		newR.Gpio, _ = strconv.Atoi(v.Gpio)
		newR.State, _ = strconv.Atoi(v.State)
		newR.Seconds, _ = vutils.DurationOf3PartStr(v.Interval)

		newStr = append(newStr, newR)
	}

	return
}
