package web

import (
	"encoding/json"
	"fmt"
	"html/template"
	"strings"

	vmsg "vk/message"
	vomni "vk/omnibus"
	vparam "vk/params"
	xrun "vk/run/a_runningpoints"
	vutils "vk/utils"
)

func raspName() string {
	return vparam.Params.Name
}

func pointList() (stra interface{}) {

	runL, runAct := xrun.RunAndActive()

	type str struct {
		List []string
		Data map[string]vomni.WPointData
	}

	tmplD := str{List: runL, Data: runAct}

	back, err := json.Marshal(tmplD)
	if nil != err {
		panic(err)
		return
	}

	json.Unmarshal(back, &stra)

	return
}

func increment1(n int) (inc int) {
	return n + 1
}

func pointCfg(point string) (cfg interface{}) {

	_, runAct := xrun.RunAndActive()

	//fmt.Println("MIMINO DEF ", len(runAct[point].Cfg.Def), "	SEQ ", runAct[point].Cfg.Seq)

	back, err := json.Marshal(runAct[point])
	if nil != err {
		panic(err)
		return
	}

	json.Unmarshal(back, &cfg)

	return runAct[point]
}

func tmpDataFromInterface(data interface{}) (back interface{}) {
	return
}

func pointLastEvents(name string) (list []string) {
	var ind int

	list = []string{}

	if ind, _ = vmsg.FindMsgIndex(vmsg.MsgTypeEvent); 0 > ind {
		return
	}

	//	fmt.Println("IERAKSTI ", len(vmsg.TotalData[ind].Data["fornarina"].LastRecords), " >>> ", eventPointName)

	for _, val := range vmsg.TotalData[ind].Data[name].LastRecords {
		//		for _, val := range vmsg.TotalData[ind].Data[eventPointName].LastRecords {
		//	for _, val := range vmsg.TotalData[ind].Data["fornarina"].LastRecords {
		timeStr := val.When.Format(vmsg.TimeFormat)
		list = append(list, fmt.Sprintf("%s --> \"%s\"", timeStr, val.Msg))
	}

	return

}

func pointCfgJsFile() (str string) {

	path := vutils.FileAbsPath("marketa", "a.js")

	str, err := vutils.FileReadString(path)
	if nil != err {
		panic(err)
		return
	}

	str = strings.Trim(str, "\n")

	return template.JSEscapeString(str)
}

func webPrefix() (prefix string) {
	return vomni.WebPrefix
}

func webPref() (prefix string) {
	return vomni.WebPrefix
}
