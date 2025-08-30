package os

import (
	"github.com/lxinyucn/goefun/coreUtil"
	. "github.com/lxinyucn/goefun/os/cqjzb"
	. "github.com/lxinyucn/goefun/os/ehttp"
)

func E获取本机ip() string {
	http := NewHttp()
	ret, flag := http.Get("https://www.taobao.com/help/getip.php")
	if flag {
		return ""
	}
	return E.StrCut(ret, `{ip:"$"}`)
}

func E获取ip信息(ip string) string {
	http := NewHttp()
	ret, flag := http.Get("http://ip.taobao.com/service/getIpInfo.php?ip=" + ip)
	if flag {
		return ""
	}
	json := New存取键值表()
	json.LoadFromJsonString(ret)
	return E.E格式化文本("%s %s %s %s",
		json.GetString("data.country"),
		json.GetString("data.region"),
		json.GetString("data.city"),
		json.GetString("data.isp"),
	)
}
