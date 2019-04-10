package points

import (
	"net"
	"time"
)

type HandlePoint interface {
	Start(port string, addr net.UDPAddr) (back string, err error)
	DoCmd(cmd string) (back string, err error)
	Stop() (back string, err error)
}

type PointIntervalOnOff struct {
	Point    string
	Active   bool
	UDPAddr  net.UDPAddr
	Index    int
	ChGoOn   chan bool
	ChDone   chan bool
	ChMsg    chan string
	ChErr    chan error
	Sequence []IntervalOnOff
}

/*
type OnePoint struct {
	Point   string
	Active  bool
	UDPAddr net.UDPAddr
	ChDone  chan bool
	ChMsg   chan string
	ChErr   chan error
}
*/

type IntervalOnOff struct {
	Pin      int
	State    int
	Interval time.Duration
}
