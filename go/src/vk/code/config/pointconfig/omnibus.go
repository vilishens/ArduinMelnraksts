package pointconfig

import (
	"net"
)

const (
	TypeRelayOnOffIntervals = 0x0001
)

type CfgPointData struct {
	RelOnOffIntervJSON CfgRelOnOffIntervalPoints `json:"relayOnOffIntervals"`
}

type PointAllCfgs struct {
	RelayOnOffInterval CfgRelOnOffIntervalPoints
}

type PointCfgData interface {
	CfgType() int
	CfgShow()
	CfgRun(point string, udp net.UDPAddr, msg chan string, err chan error)

	//	CfgData() interface{}
}

type runningPoint struct {
	name      string
	pointType int
	addr      net.UDPAddr
	msg       chan string
}

type activePoints map[string]*runningPoint
