package checkNet

import (
	"fmt"
	"net"
	"time"
	vomni "vk/omnibus"
	vparam "vk/params"
	vsgrid "vk/sendgrid"
	vutils "vk/utils"
)

func NetInfo(chGoOn chan bool, chDone chan int, chErr chan error) {

	fmt.Println("BARGUZIN")

	locGoOn := make(chan bool)
	go checkIPv4(chDone, locGoOn)

	//	chGoOn <- true

	end := false
	for !end {
		select {
		case <-locGoOn:

			fmt.Println("==== Postanovka snookera =======================")

			chGoOn <- true
		case <-chDone:
			end = true
		}
	}
}

func checkIPv4(done chan int, goOn chan bool) {

	var netDuration = 10 * time.Second
	var first = true

	kika := 0

	repeatMax := checkNetRepeats

	for {
		if first {
			goOn <- true
			first = false
		} else {
			tick := time.NewTicker(netDuration)
			<-tick.C
		}

		intIP, extIP, err := getIPv4Addrs()
		if nil != err {
			repeatMax--
			if repeatMax >= 0 {

				str := fmt.Sprintf("\n=== %4d ##### ! ! ! ! ######=====> ATKĀRTOT INT %s  EXT %s\n\n", kika, intIP, extIP)
				vomni.LogErr.Println(str)
				continue
			}
			str := fmt.Sprintf("Couldn't find IP (int \"%s\", ext \"%s\", reapeat left %d)", intIP, extIP, repeatMax)
			vomni.LogFatal.Println(str)

			vomni.RootErr <- err
			done <- vomni.DoneOK
			return
		}

		repeatMax = checkNetRepeats

		if setCurrentNet(intIP, extIP) {
			if err = sendNetInfov4(); nil != err {
				err = vutils.ErrFuncLine(fmt.Errorf("Couldn't send new IPv4 - %v", err))

				vomni.LogFatal.Println(err)

				vomni.RootErr <- err
				done <- vomni.DoneOK
			}
		}

		kika++
		fmt.Printf("\n=== %4d =========================> CURRENT LAN Internal %s External %s\n\n", kika, vparam.Params.InternalIPv4, vparam.Params.ExternalIPv4)
	}
}

func setCurrentNet(intIP string, extIP string) (send bool) {
	if (nil != net.ParseIP(intIP)) && (vparam.Params.InternalIPv4 != intIP) {
		vparam.Params.InternalIPv4 = intIP
		send = true
	}

	if (nil != net.ParseIP(extIP)) && (vparam.Params.ExternalIPv4 != extIP) {
		vparam.Params.ExternalIPv4 = extIP
		send = true
	}

	if send {
		fmt.Printf("====== INT %s (%s) === EXT #\"%s\"# (*\"%s\"*) ====================+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++> BORTICH\n",
			intIP, vparam.Params.InternalIPv4, extIP, vparam.Params.ExternalIPv4)
	}

	return
}

func sendNetInfov4() (err error) {

	emails := vparam.Params.WebEmail

	subj := vparam.Params.Name + " --- " + vutils.TimeNow(vomni.TimeFormat1) + " --- NET"
	msg_txt := fmt.Sprintf("WEB: %s:%d\nSSH: %s:%d\n\n",
		vparam.Params.ExternalIPv4, 50177,
		vparam.Params.ExternalIPv4, 50354)
	msg_html := fmt.Sprintf("</h2>WEB: <strong>%s:%d</strong><br />SSH: <strong>%s:%d<strong><br /><br /></h2>",
		vparam.Params.ExternalIPv4, 50177, //vparam.Params.InternalPort,
		vparam.Params.ExternalIPv4, 50354) //vparam.Params.ExternalPort)

	return vsgrid.Send(emails, subj, msg_txt, msg_html)
}

func getIPv4Addrs() (intIP string, extIP string, err error) {
	if intIP, err = vutils.InternalIPv4(); nil != err {
		err = vutils.ErrFuncLine(fmt.Errorf("Couldn'get Internal IPv4 - %v", err))
		return
	}
	if extIP, err = vutils.ExternalIPv4(); nil != err {
		err = vutils.ErrFuncLine(fmt.Errorf("Couldn't get External IPv4 - %v", err))
		return
	}

	return
}
