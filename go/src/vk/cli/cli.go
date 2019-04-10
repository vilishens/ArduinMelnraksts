package cli

import (
	"flag"
	vomni "vk/omnibus"
)

var CliCfgPath, CliCfgFile, CliCfgAction string

func init() {
	flag.StringVar(&CliCfgPath, vomni.CfgFldPath, vomni.CfgFactoryPath, "configuration file path")
	flag.StringVar(&CliCfgAction, vomni.CfgFldAction, vomni.CfgFactoryAction, "action at start")

	flag.Parse()
}

func Init() (err error) {
	return
}
