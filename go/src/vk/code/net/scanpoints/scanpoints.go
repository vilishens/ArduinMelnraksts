package scanpoints

import (
	"fmt"
	"net"
	"time"

	vparam "vk/params"
	xmsg "vk/run/message"
	vutils "vk/utils"
)

// ScanPoints - goes through x.x.x.Start till x.x.x.End IP addresses and sends Hi message to points at these addresses
func ScanPoints(chDone chan bool, chErr chan error) {

	locDone := make(chan bool)
	locErr := make(chan error)

	go iterateIP(locDone, locErr)

	select {
	case err := <-locErr:
		chErr <- vutils.ErrFuncLine(err)
	case <-locDone:
		chDone <- true
	}
}

func iterateIP(chDone chan bool, chErr chan error) {

	for i := startIP; i <= endIP; i++ {

		ip := net.ParseIP(vparam.Params.InternalIPv4).To4()
		ip[3] = byte(i)

		if ip.String() == vparam.Params.InternalIPv4 {
			// don't send to the station itself (the same IP address)
			continue
		}

		locErr := make(chan error)
		locDone := make(chan bool)

		if i == endIP {

			fmt.Println("\n***\n\t***\n\t\t***\n\t\t\t***\n\t\t\t\t***\nPAKCRATION")
		}

		go tryPointConn(ip, locDone, locErr)

		select {
		case <-locDone:
			//
		case err := <-locErr:
			chErr <- vutils.ErrFuncLine(err)
			return
		}

		time.Sleep(200 * time.Millisecond)
	}

	fmt.Println("#################################################### Scan NET done...")
	for k, v := range xmsg.SendList {
		fmt.Printf("PEC VISIEM SCANIEM %2d --- IP %q NBR %d\n", k, v.UDPAddr.IP.String(), v.MsgNbr)
	}

	chDone <- true
}

func tryPointConn(ip net.IP, chDone chan bool, chErr chan error) {

	dstAddr := net.UDPAddr{IP: ip, Port: vparam.Params.UDPPort}

	locDone := make(chan bool)
	locErr := make(chan error)

	go xmsg.TryHello(dstAddr, locDone, locErr)

	select {
	case err := <-locErr:
		err = vutils.ErrFuncLine(err)
		chErr <- err
		break
	case <-locDone:
		chDone <- true
		break
	}
}
