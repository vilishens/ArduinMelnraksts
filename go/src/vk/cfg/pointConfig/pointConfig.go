package pointConfig

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	vomni "vk/omnibus"
	vparam "vk/params"
	vpoint "vk/points"
	vutils "vk/utils"
)

func PointCfg(chDone chan bool, chErr chan error) {

	locDone := make(chan bool)
	locErr := make(chan error)

	go loadModeCfgs(locDone, locErr)

	select {
	case err := <-locErr:
		chErr <- vutils.ErrFuncLine(err)
	case <-locDone:
		chDone <- true
	}
}

func loadModeCfgs(chDone chan bool, chErr chan error) {

	var err error

	for _, f := range vparam.Params.DevModes {
		switch f {
		case vomni.PointModeIntervalOnOff:
			err = intervalOnOffLoad()
		default:
			err = fmt.Errorf("Don't know what to do with point mode config \"%s\"", f)
		}
	}

	if nil != err {
		chErr <- err
	}

	chDone <- true

	return
}

func intervalOnOffLoad() (err error) {

	var data Intervals

	if err = intervalOnOffGetCfg(&data); nil != err {
		return vutils.ErrFuncLine(err)
	}

	if err = putIntervalOnOffCfg(data); nil != err {
		return vutils.ErrFuncLine(err)
	}

	return
}

func intervalOnOffGetCfg(dst *Intervals) (err error) {
	path := vparam.Params.PointModeFiles[vomni.PointModeIntervalOnOff]

	if ok, err := vutils.PathExists(path); !ok {
		return vutils.ErrFuncLine(err)
	}

	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return vutils.ErrFuncLine(err)
	}

	if err = json.Unmarshal(raw, dst); nil != err {
		return vutils.ErrFuncLine(err)
	}

	return
}

func putIntervalOnOffCfg(data Intervals) (err error) {

	for k, v := range data {
		vpoint.AllIntervalOnOff[k] = new(vpoint.PointIntervalOnOff)

		vpoint.AllIntervalOnOff[k].Point = k
		vpoint.AllIntervalOnOff[k].ChGoOn = make(chan bool)
		vpoint.AllIntervalOnOff[k].ChDone = make(chan bool)
		vpoint.AllIntervalOnOff[k].ChMsg = make(chan string)
		vpoint.AllIntervalOnOff[k].ChErr = make(chan error)

		for _, v1 := range v {

			newSeq := vpoint.IntervalOnOff{}

			if newSeq.Pin, err = strconv.Atoi(v1.Pin); nil != err {
				return vutils.ErrFuncLine(err)
			}
			if newSeq.State, err = strconv.Atoi(v1.State); nil != err {
				return vutils.ErrFuncLine(err)
			}
			if newSeq.Interval, err = vutils.DurationOf3PartStr(v1.Interval); nil != err {
				return vutils.ErrFuncLine(err)
			}

			vpoint.AllIntervalOnOff[k].Sequence = append(vpoint.AllIntervalOnOff[k].Sequence, newSeq)

		}
	}

	return
}
