package a_runningpoints

import (
	"sort"
	vomni "vk/omnibus"
)

var allPointDescr map[int]string

func init() {

	allPointDescr = make(map[int]string)

	allPointDescr[vomni.PointTypeRelayOnOffInterval] = "Relay On/Off Interval"

}

func RunAndActive() (list []string, active map[string]vomni.WPointData) {

	list = []string{}
	active = make(map[string]vomni.WPointData)

	// garanichev
	//for k, v := range Points {
	for k, v := range Points {
		list = append(list, k)

		actNew := vomni.WPointData{}

		// garanichev
		actNew = v.WebPointData()
		actNew.Descr = allPointDescr[actNew.Type]

		active[k] = actNew

	}

	sort.Strings(list)

	return
}

func ReceivedWebMsg(point string, msg string, cfg interface{}) {
	Points[point].ReceivedWebMsg(msg, cfg)
}
