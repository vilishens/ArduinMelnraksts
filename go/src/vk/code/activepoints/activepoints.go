package activepoints

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	vcfg "vk/code/config/pointconfig"
	vmsg "vk/message"
	vomni "vk/omnibus"
)

func HandleReceived(msg string, chErr chan error) {

	flds := strings.Split(msg, vmsg.FieldSeparator)

	cfg := vcfg.GetPointConfig(flds[0])
	fmt.Println("*** SEKUNDOMER *** SEKUNDOMER *** SEKUNDOMER *** SEKUNDOMER ", flds, cfg)

	if nil == cfg {

		fmt.Println("==================== NEATRADU CFG")

		return

	} else {

		fmt.Println("!!!!!!!!!!!!!!!! Atradu cfg !!!!!!!!!!!!!!!!!!!!!!!")
		if vcfg.PointRunning(flds[0]) {

			fmt.Println("######################## HELLO FROM POINT", flds[0])

			//cfg.NewMsg(msg)
		} else if flds[1] == strconv.Itoa(vmsg.MsgInputHelloFromPoint) {
			fmt.Println("@@@@@@@@@@@@@@@@@@@ HELLO FROM POINT")

			port, _ := strconv.Atoi(flds[3])

			udpAddr := net.UDPAddr{IP: net.ParseIP(flds[2]), Port: port}

			chMsg := make(chan string)

			go cfg.CfgRun(flds[0], udpAddr, chMsg, vomni.RootErr)
		}

		//for k, v := range cfg {
		//		fmt.Println(k, "---> ", v)
		//}

		cfg.CfgShow()

		fmt.Println("SITKOVETSKY naff NULL", cfg.CfgType())

	}

	chErr <- nil
}
