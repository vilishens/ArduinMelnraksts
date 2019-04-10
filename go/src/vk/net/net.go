package net

import (
	"fmt"
	"net"
	"os/exec"
	"strconv"
	"time"

	vomni "vk/omnibus"
	vparam "vk/params"
	vutils "vk/utils"

	vgrid "vk/sendgrid"
)

var checkNet = true
var checkWEBInterval = 5 * time.Minute

func ExternalIP() (ip string, err error) {

	cmdIP := "curl -s checkip.dyndns.org | sed -e 's/.*Current IP Address: //' -e 's/<.*$//'"

	ipTmp, err := exec.Command("bash", "-c", cmdIP).Output()
	if err != nil {
		err = vutils.ErrFuncLine(err)
	} else {
		ip = string(ipTmp)
	}

	return
}

// LocalIP returns the non loopback local IP of the host
func InternalIP() (ip string, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		err = vutils.ErrFuncLine(err)
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

func HandleWEB(chGoOn chan bool, chDone chan bool, chErr chan error) {

	fmt.Println("Pomeshatj sudjbe")

	lDone := make(chan bool)
	lErr := make(chan error)

	goAhead := make(chan bool)
	go checkWEBInfo(lDone, lErr, goAhead)

	fmt.Println("******************************** HAMUrapi ")

	select {
	case canContinue := <-goAhead:
		chGoOn <- canContinue
		break
	case <-lDone:
		break
	case err := <-lErr:
		vomni.RootErr <- err
		break
	}

}

func checkWEBInfo(done chan bool, chErr chan error, goAhead chan bool) {
	addrTick := new(time.Ticker)
	alive := new(time.Ticker)

	tmpsLaiks := time.Duration(10 * time.Second)
	tmps := time.NewTicker(tmpsLaiks)

	aliveDuration := time.Duration(vparam.Params.WebAliveInterval) * time.Second
	if vparam.Params.WebAliveInterval > 0 {
		alive = time.NewTicker(aliveDuration)
	}

	sendAddrEmail := true
	addrTick = time.NewTicker(checkWEBInterval)

	fmt.Printf("TAGAD pirms cikla\n")

	for {
		if sendAddrEmail {

			err := sendNewAddrEmail()
			//var err error

			if nil != err {
				chErr <- vutils.ErrFuncLine(err)
			}

			sendAddrEmail = false
			goAhead <- true
		}

		select {
		case <-tmps.C:
			t := time.Now()
			timeStr := t.Format("2006-01-02 15:04:05 -07:00 MST")
			fmt.Printf("%s >>>>>>>>>>>>>>>> Melioranska TMP <<<<<<<<<<<<<<<<\n", timeStr)
			tmps = time.NewTicker(tmpsLaiks)

		case <-addrTick.C:
			t := time.Now()
			timeStr := t.Format("2006-01-02 15:04:05 -07:00 MST")
			fmt.Printf("%s ################ MEZOTICH jābauda IP #######################\n", timeStr)

			err := error(nil)
			if sendAddrEmail, err = newWEBAddrs(); nil != err {
				chErr <- vutils.ErrFuncLine(err)
			}

			addrTick = time.NewTicker(checkWEBInterval)
		case <-alive.C:
			t := time.Now()
			timeStr := t.Format("2006-01-02 15:04:05 -07:00 MST")
			fmt.Printf("%s XXXXXXXXXXXXXXXX BAZVANS alive XXXXXXXXXXXXXXXXXXXXXXXXXXXXX\n", timeStr)
			err := sendAliveEmail()
			if nil != err {
				chErr <- vutils.ErrFuncLine(err)
			}
			alive = time.NewTicker(aliveDuration)
		}
	}
}

func newWEBAddrs() (isNew bool, err error) {

	intIP := ""
	extIP := ""

	if intIP, err = InternalIP(); nil != err {
		err = vutils.ErrFuncLine(err)
		return
	}

	if extIP, err = ExternalIP(); nil != err {
		vomni.RootErr <- err
	}

	if vparam.Params.ExternalIPv4 != extIP {
		vparam.Params.ExternalIPv4 = extIP
		isNew = true
	}

	if vparam.Params.InternalIPv4 != intIP {
		vparam.Params.InternalIPv4 = intIP
		isNew = true
	}

	return
}

func sendNewAddrEmail() (err error) {

	t := time.Now()
	timeStr := t.Format("2006-01-02 15:04:05 -07:00 MST")

	txtMsg := fmt.Sprintf("%s --- Current effective contact parameters:     \nWEB %s:%s    \nSSH %s:%s\n",
		vparam.Params.Name,
		vparam.Params.ExternalIPv4,
		strconv.Itoa(50177),
		vparam.Params.ExternalIPv4,
		strconv.Itoa(50354))

	htmlMsg := fmt.Sprintf("<strong><h2>%s --- Current effective contact parameters:<br /><br />WEB %s:%s<br />SSH %s:%s</h2></strong>",
		vparam.Params.Name,
		vparam.Params.ExternalIPv4,
		strconv.Itoa(50177),
		vparam.Params.ExternalIPv4,
		strconv.Itoa(50354))

	subject := fmt.Sprintf("%s - \"%s\" Raspberry Web Contact Info", timeStr, vparam.Params.Name)
	addr := vparam.Params.WebEmail
	//file := vparam.Params.WebEmailMutt

	err = vgrid.Send(addr, subject, txtMsg, htmlMsg)
	res := ""
	//	res, err = vmutt.SendMail(addr, subject, txtMsg, htmlMsg, file)

	fmt.Printf("xxxxxxxxxxxxxxxxxx Emailas resultata:\n%s\n", res)

	return
}

func sendAliveEmail() (err error) {

	t := time.Now()
	timeStr := t.Format("2006-01-02 15:04:05 -07:00 MST")

	txtMsg := fmt.Sprintf("%s --- I'm alive!\n", vparam.Params.Name)

	htmlMsg := fmt.Sprintf("<strong>%s --- I'm alive!<br />", vparam.Params.Name)

	subject := fmt.Sprintf("ALIVE!!! %s - \"%s\" Raspberry", timeStr, vparam.Params.Name)
	addr := vparam.Params.WebEmail

	err = vgrid.Send(addr, subject, txtMsg, htmlMsg)

	return
}
