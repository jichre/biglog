package config

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/larspensjo/config"
)

var (
	configFile = flag.String("configfile", "E:/gosrc/src/bitbucket.org/biglog/config.ini", "General configuration file")
)

var SysConfig = make(map[string]string)

func init() {
	flag.Parse()

	cfg, err := config.ReadDefault(*configFile)
	if err != nil {
		panic(err)
	}

	//Initialized topic from the configuration
	sections := cfg.Sections()
	for _, v := range sections {
		section, _ := cfg.SectionOptions(v)
		for _, v1 := range section {
			options, err := cfg.String(v, v1)
			fmt.Println(err)
			if err == nil {
				SysConfig[v1] = options
			}
		}
	}
	fmt.Println(SysConfig)
}

func GetMaxFileLen() int {
	if v, ok := SysConfig["max_len"]; ok {
		t, err := strconv.Atoi(v)

		if err == nil {
			return t
		}
	}

	return 12582912 //128M
}

func GetFilePath() string {
	return SysConfig["log_path"]
}
