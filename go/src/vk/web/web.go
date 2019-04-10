package web

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strconv"
	vomni "vk/omnibus"
	vparam "vk/params"
	vutils "vk/utils"
)

//var webPrefix = "/xK-@eRty7/"

func GoWeb(chGoOn chan bool, chDone chan int, chErr chan error) {

	os.Mkdir(backendfile, 0755)
	defer os.Remove(backendfile)

	fmt.Println("********* SITJO")

	//	func Start() {

	//	v_tmpls.SetTemplates()
	setMux()

	//		go startListen()
	//	}

	fmt.Println("MELHIORS")

	go startListen(chGoOn)

	chGoOn <- true

}

func startListen(chGoOn chan bool) {

	prefix := vomni.WebPrefix //v_cli.Param(v_cli.CliFileServPrefix)
	//	f_root := vomni.RootPath //v_cli.Param(v_cli.CliFileRoot)
	f_static := vutils.FileAbsPath(vomni.WebStaticPath, "")
	l_port := vparam.Params.WebPort //v_cli.Param(v_cli.CliWebPort)
	net_addr := ""                  // v_cli.Param(v_cli.CliNetAddr)

	/*
		prefix := vomni.WebPrefix
		f_root := vutils. filepath.Join(wom)
	*/
	//http.Handle(prefix, http.StripPrefix(prefix, http.FileServer(http.Dir(f_root))))
	//http.Handle(prefix, http.StripPrefix(prefix, http.FileServer(http.Dir(f_static))))

	_ = prefix

	//http.Handle("/sap/", http.FileServer(http.Dir("static")))
	//http.Handle(prefix, http.StripPrefix(prefix, http.FileServer(http.Dir(f_static))))

	//http.Handle("/", http.FileServer(http.Dir(f_static)))

	listen_addr := ":" + strconv.Itoa(l_port)

	//	fmt.Printf("PREFIX: %s\nFILE ROOT: %s\nIP PORT: %s\nNET ADDR: %s\nLISTEN ADDR: %s\n",
	//		prefix, f_root, l_port, net_addr, listen_addr)

	//	fmt.Printf("PLESANS: %v\n", http.StripPrefix(prefix, http.FileServer(http.Dir(f_root))))

	fmt.Println("Listen...", listen_addr, " static ", f_static)

	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/static"))
	mux.Handle("/fima/", http.StripPrefix("/fima/", files))

	if net_addr != "" {
		net_type := "tcp"                                  // v_cli.Param(v_cli.CliNet)
		app_env := "dev"                                   // v_cli.Param(v_cli.CliAppEnv)
		addr_out := "/home/vilis/tmp/ip-address-found.txt" // v_cli.Param(v_cli.CliAddrOut)
		l, err := net.Listen(net_type, net_addr)
		if err != nil {
			err = fmt.Errorf("Error! %s (env '%s')", err.Error(), app_env)

			panic(err)
		}

		tmp_str := l.Addr().Network() + ":" + l.Addr().String()

		err = ioutil.WriteFile(addr_out, []byte(tmp_str), 0644)
		if err != nil {
			err = fmt.Errorf("Error! %s (env '%s')", err.Error(), app_env)

			panic(err)
		}

		fmt.Println("Volodja Dakšs...", tmp_str)

		s := &http.Server{}

		//		fmt.Println("Volodja Dakšs...")

		panic(s.Serve(l))

	} else {
		panic(http.ListenAndServe(listen_addr, nil))
	}

	//	chGoOn <- true
}
