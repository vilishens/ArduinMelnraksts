package mail

import (
	"fmt"
	"time"

	vomni "vk/omnibus"
	vparam "vk/params"
	vsgrid "vk/sendgrid"
	vutils "vk/utils"
)

var checkNet = true
var netDone = make(chan bool)
var netDuration = 5 * time.Second

func NetInfo(chGoOn chan bool, chDone chan int, chErr chan error) {

	locGoOn := make(chan bool)
	go checkIPv4(chDone, locGoOn)

	chGoOn <- true

	end := false
	for !end {
		select {
		case <-locGoOn:
			chGoOn <- true
		case <-chDone:
			end = true
		}
	}
}

func checkIPv4(done chan int, goOn chan bool) {

	//err := error(nil)

	skif := 0

	for {
		if checkNet {
			intIP, extIP, err := getIPv4Addrs()
			//locIP, err := vnet.LocalIP()
			if nil != err {
				vomni.RootErr <- err
				done <- vomni.DoneError
			}

			if (vparam.Params.InternalIPv4 != intIP) || (vparam.Params.ExternalIPv4 != extIP) {

				fmt.Println("=============================> GARNIZOV")
				skif++

				vparam.Params.InternalIPv4 = intIP
				vparam.Params.ExternalIPv4 = extIP

				if err = sendNetInfov4(); nil != err {
					err = vutils.ErrFuncLine(fmt.Errorf("Couldn't send new IPv4 - %v", err))
					vomni.RootErr <- err
					done <- vomni.DoneError
				}
			} else {
				fmt.Printf("=============================> MUSKATIN %d INT %s  EXT %s\n", skif, vparam.Params.InternalIPv4, vparam.Params.ExternalIPv4)
			}

			//			time.Sleep(20 * time.Second)

			goOn <- true

			//extIP := ""
			//			if extIP, err = vnet.ExternalIP(); nil != err {
			//				vomni.RootErr <- err
			//			}

			//			if (vomni.CfgExternalIP != extIP) || (vomni.CfgInternalIP != locIP) {

			//sendNetInfo(locIP, extIP)
			//			}

			//fmt.Println("LOC IP ", locIP, "EXTERN IP ", extIP)
			/*
			   //			msg := fmt.Sprintf("Current effective contact parameters:\nWEB address %s:%s\nInternal address %s:%s\n",
			   //				extIP, strconv.Itoa(vomni.CfgExternalPort), locIP, strconv.Itoa(vomni.CfgInternalPort))

			   //			html := fmt.Sprintf("<strong>Current effective contact parameters:<br /><br />WEB address %s:%s<br />Internal address %s:%s<br /></strong>",
			   //				extIP, strconv.Itoa(vomni.CfgExternalPort), locIP, strconv.Itoa(vomni.CfgInternalPort))

			   			subject := "Raspberry of Arduino Web Contact Info"

			   			sendgrid.Send(subject, msg, html, "vilis@mail.com")
			*/
			tick := time.NewTicker(netDuration)
			<-tick.C

			fmt.Println("=============================> AIVAZO")

		}
	}
}

func sendNetInfov4() (err error) {

	emails := vparam.Params.WebEmail

	//emails := "vilisX@hotmail.com"

	fmt.Println("=============================> KUKSOV  @", emails)

	subj := vparam.Params.Name + " --- " + vutils.TimeNow(vomni.TimeFormat1) + " --- NET"
	msg_txt := fmt.Sprintf("Int: %s:%d\nExt: %s:%d\n\n", vparam.Params.InternalIPv4, vparam.Params.InternalPort,
		vparam.Params.ExternalIPv4, vparam.Params.ExternalPort)
	msg_html := fmt.Sprintf("</h2>Int: <strong>%s:%d</strong><br />Ext: <strong>%s:%d<strong><br /><br /></h2>",
		vparam.Params.InternalIPv4, vparam.Params.InternalPort,
		vparam.Params.ExternalIPv4, vparam.Params.ExternalPort)

	fmt.Println("########################################################################## MAIL", emails)
	fmt.Println("########################################################################## MAIL", emails)
	fmt.Println("########################################################################## MAIL", emails)
	fmt.Println("########################################################################## MAIL", emails)
	fmt.Println("########################################################################## MAIL", emails)
	fmt.Println("########################################################################## MAIL", emails)
	fmt.Println("########################################################################## MAIL", emails)
	fmt.Println("########################################################################## MAIL", emails)
	fmt.Println("########################################################################## MAIL", emails)
	fmt.Println("########################################################################## MAIL", emails)
	fmt.Println("########################################################################## MAIL", emails)
	fmt.Println("########################################################################## MAIL", emails)

	//return vsgrid.Send(vparam.Params.WebEmail, subj, msg_txt, msg_html)
	//return vsgrid.Send("vilis@mail.com", subj, msg_txt, msg_html)
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
