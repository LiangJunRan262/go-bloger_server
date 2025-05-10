package core

import (
	"strings"

	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"github.com/sirupsen/logrus"

	utilsIp "bloger_server/utils/ip"
)

var (
	searcher *xdb.Searcher
)

func InitIPDB() {
	var dbPath = "init/ip2region.xdb"
	_searcher, err := xdb.NewWithFileOnly(dbPath)
	if err != nil {
		logrus.Fatalf("IP数据库加载失败: %s", err.Error())
		return
	}
	searcher = _searcher
}

func GetIPAddr(ip string) (addr string) {
	if utilsIp.HasLocalIPAddr(ip) {
		return "内网"
	}

	region, err := searcher.SearchByStr(ip)
	if err != nil {
		logrus.Warnf("错误的IP地址: %s", ip)
		return "异常地址"
	}
	addrList := strings.Split(region, "|")
	if len(addrList) != 5 {
		logrus.Errorf("异常的IP地址: %s", region)
		return "未知地址"
	}

	//
	// 国家 0 省份 市 运营商
	country := addrList[0]
	province := addrList[2]
	city := addrList[3]

	if province != "0" && city != "0" {
		return province + "·" + city
	}
	if country != "0" && province != "0" {
		return country + "·" + province
	}
	if country != "0" {
		return country
	}

	return region
}
