package utils

import (
	"fmt"
	"net"
	"os/exec"
	"strings"
	"time"
	vomni "vk/omnibus"
)

func ExternalIPv4() (ip string, err error) {

	cmds := []string{"dig +short myip.opendns.com @resolver1.opendns.com",
		"curl http://myip.dnsomatic.com"}

	tmpIP := ""

	ind := 0
	for "" == tmpIP && nil == err && ind < len(cmds) {
		tick := time.NewTicker(2 * time.Second)
		cmd := cmds[ind]
		chStr := make(chan string)
		chErr := make(chan error)

		go doCmd(cmd, chStr, chErr)
		select {
		case <-tick.C:
			ind++
		case tmpIP = <-chStr:
			return strings.Trim(tmpIP, "\n"), nil
		case err = <-chErr:
			// return "", ErrFuncLine(fmt.Errorf("Failed CMD \"%s\" --- %v", cmd, err))
			err1 := ErrFuncLine(fmt.Errorf("Failed CMD \"%s\" --- %v (index %d)", cmd, err, ind))

			vomni.LogErr.Println(err1)

			fmt.Println("########### OSET ### jāliek kļūdas ieraksts %v", err1)
			ind++
		}
	}

	if "" == tmpIP {
		return "", ErrFuncLine(fmt.Errorf("Couldn't get the external IP"))
	}

	return
}

func doCmd(cmd string, chStr chan string, chErr chan error) {
	res, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		chErr <- err
	} else {
		chStr <- string(res)
	}
}

// LocalIP returns the non loopback local IP of the host
func InternalIPv4() (ip string, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		err = ErrFuncLine(err)
		vomni.LogErr.Println(err)
		return
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
				return
			}
		}
	}
	return
}
