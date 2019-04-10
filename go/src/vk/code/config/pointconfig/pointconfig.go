package pointconfig

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	vcfg "vk/cfg"
	vutils "vk/utils"
)

var CfgPoints CfgPointData

func init() {

	CfgPoints = CfgPointData{}
}

func (d CfgPointData) Save() (err error) {

	data, _ := json.Marshal(d)
	if err = ioutil.WriteFile(vcfg.Final.PointCfgFile, data, 0644); nil != err {
		return vutils.ErrFuncLine(err)
	}

	return
}

func CfgDefault() (d CfgPointData, err error) {

	if err = vutils.ReadJson(vcfg.Final.PointDefaultCfgFile, &d); nil != err {
		return CfgPointData{}, vutils.ErrFuncLine(err)
	}

	return
}

func PointCfg(chDone chan bool, chErr chan error) {

	locDone := make(chan bool)
	locErr := make(chan error)

	go preparePointCfg(locDone, locErr)

	select {
	case err := <-locErr:
		chErr <- vutils.ErrFuncLine(err)
	case <-locDone:
		// config read time to start rotation

		// xxx - vikulovs
		//vutils.StartRotate(vparam.Params.RotateRunCfg)

		chDone <- true
	}
}

func loadPointCfg() (data CfgPointData, err error) {

	if has, _ := vutils.PathExists(vcfg.Final.PointCfgFile); !has {
		if err := vutils.FileCopy(vcfg.Final.PointDefaultCfgFile, vcfg.Final.PointCfgFile); nil != err {
			return CfgPointData{}, vutils.ErrFuncLine(err)
		}
	}

	if err = vutils.ReadJson(vcfg.Final.PointCfgFile, &data); nil != err {
		return CfgPointData{}, vutils.ErrFuncLine(err)
	}

	return
}

func preparePointCfg(doneCh chan bool, errCh chan error) {

	var err error

	CfgPoints, err = loadPointCfg()
	if nil != err {
		errCh <- err
		return
	}

	// kaukenas
	//else {
	//	err = data.prepareCfg()
	//	if nil != err {
	//		errCh <- vutils.ErrFuncLine(err)
	//	}
	//}

	doneCh <- true
}

/* garanichev
func (d CfgPointData) prepareCfg() (err error) {
	if err = d.RelOnOffIntervJSON.prepare(); nil != err {
		return
	}

	return
}
*/
//##########################################################################################
//##########################################################################################
//##########################################################################################

var AllCfgData PointAllCfgs
var AllActivePoints activePoints

func init() {
	AllCfgData.RelayOnOffInterval = CfgRelOnOffIntervalPoints{}
	AllActivePoints = activePoints{}
}

func GetPointConfig(point string) (cfg PointCfgData) {
	for pnt, data := range AllCfgData.RelayOnOffInterval {
		fmt.Printf("SALIDZINU %s\n", pnt)
		if point == pnt {
			cfg = data
			fmt.Printf("  atradu SUTER %v\n", data)
			break
		} else {
			fmt.Printf("NEATRADU SUTER %s\n", point)
		}
	}

	return
}

func PointRunning(point string) (run bool) {
	for k, _ := range AllActivePoints {
		if point == k {
			return true
		}
	}

	return false
}
