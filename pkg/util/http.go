package util

import (
	"net"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/proxy"
)

var GithubMirror = []string{
	"https://ghgo.xyz",
	"https://gh.h233.eu.org",
	"https://gh.ddlc.top",
	"https://slink.ltd",
	"https://gh.con.sh",
	"https://cors.isteed.cc",
	"https://hub.gitmirror.com",
	"https://sciproxy.com",
	"https://ghproxy.cc",
	"https://cf.ghproxy.cc",
	"https://www.ghproxy.cc",
	"https://ghproxy.cn",
	"https://www.ghproxy.cn",
	"https://gh.jiasu.in",
	"https://dgithub.xyz",
	"https://download.ixnic.net",
	"https://download.nuaa.cf",
	"https://download.scholar.rr.nu",
	"https://download.yzuu.cf",
	"https://mirror.ghproxy.com",
	"https://ghproxy.net",
	"https://kkgithub.com",
	"https://gitclone.com",
	"https://hub.incept.pw",
	"https://github.moeyy.xyz",
	"https://gh.xiu2.us.kg",
	"https://dl.ghpig.top",
	"https://gh-proxy.com",
	"https://cors.isteed.cc",
	"https://sciproxy.com",
	"https://github.site",
	"https://github.store",
	"https://github.tmby.shop",
	"https://hub.whtrys.space",
	"https://gh-proxy.ygxz.in",
	"https://gitdl.cn",
	"https://ghp.ci",
	"https://githubfast.com",
	"https://ghproxy.net",
}

func ReqAgent(req *http.Request) {
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Sec-Fetch-Site", "same-origin")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/129.0.0.0")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("sec-ch-ua", "\"Microsoft Edge\";v=\"129\", \"Not=A?Brand\";v=\"8\", \"Chromium\";v=\"129\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Windows\"")
}

func NewProxyDial(prxoyUrls string) proxy.Dialer {
	var proxyDialer proxy.Dialer = proxy.Direct
	for _, prxoyUrl := range strings.Split(prxoyUrls, ",") {
		urlInfo, err := url.Parse(prxoyUrl)
		if err != nil {
			return proxyDialer
		}
		var auth *proxy.Auth = nil
		if urlInfo.User != nil {
			pwd, _ := urlInfo.User.Password()
			auth = &proxy.Auth{
				User:     urlInfo.User.Username(),
				Password: pwd,
			}
		}

		dialer, err := proxy.SOCKS5("tcp", urlInfo.Host, auth, proxyDialer)
		if err == nil {
			proxyDialer = dialer
		}
	}
	return proxyDialer
}

func GetProxyClient(proxystr string) *http.Client {
	client := http.Client{Transport: &http.Transport{
		Proxy: http.ProxyFromEnvironment,
	}}
	if proxystr != "" {
		proxyInfo, err := url.Parse(proxystr)
		if err == nil {
			if strings.HasPrefix(proxystr, "http") {
				client.Transport = &http.Transport{
					Proxy: http.ProxyURL(proxyInfo),
				}
			} else {
				client.Transport = &http.Transport{
					Dial: func(network, addr string) (net.Conn, error) {
						if strings.HasPrefix(addr, "127") {
							return net.Dial(network, addr)
						}
						proxyDialer := NewProxyDial(proxystr)
						return proxyDialer.Dial(network, addr)
					},
				}
			}

		}
	}
	return &client
}
