// Copyright 2021 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package conf

import (
	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
	log "unknwon.dev/clog/v2"
)

func init() {
	err := log.NewConsole()
	if err != nil {
		panic("init console logger: " + err.Error())
	}
}

type Conf struct {
	ICPs []*ICP `toml:"icp"`
}

var ICPs []*ICP

type ICP struct {
	URL string `toml:"url"`
	No  string `toml:"no"`
}

func Init(customConf string) error {
	if customConf == "" {
		customConf = "./conf/icppp.toml"
	}

	var conf Conf
	_, err := toml.DecodeFile(customConf, &conf)
	if err != nil {
		return errors.Wrap(err, "decode config file")
	}

	ICPs = conf.ICPs
	return nil
}

// GetICPByURL tries to find the ICP with the given URL.
func GetICPByURL(url string) string {
	for _, icp := range ICPs {
		if icp.URL == url {
			return icp.No
		}
	}
	return ""
}
