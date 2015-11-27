package useragent

import (
	"github.com/springCat/user_agent"
	"net/http"
	"strings"
)
//android
//"Mozilla/5.0 (Linux; U; Android 4.4.2; zh-cn; Coolpad 8297W Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko)Version/4.0 MQQBrowser/6.1 Mobile Safari/537.36"
//mac
//"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/46.0.2490.80 Safari/537.36"
//ios itouch
// "Mozilla/5.0 (iPod touch; CPU iPhone OS 8_1_2 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Version/8.0 Mobile/12B440 Safari/600.1.4"
//微信
//"Mozilla/5.0 (Linux; U; Android 4.4.2; zh-cn; Coolpad 8297W Build/KOT49H) AppleWebKit/533.1 (KHTML, like Gecko)Version/4.0 MQQBrowser/5.4 TBS/025478 Mobile Safari/533.1 MicroMessenger/6.3.7.51_rbb7fa12.660 NetType/WIFI Language/zh_CN"
//ios iphone
// "Mozilla/5.0 (iPhone; CPU iPhone OS 9_1 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Version/9.0 Mobile/13B143 Safari/601.1"
//windows
//Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/36.0.1985.125 Safari/537.36
//linux
//"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.97 Safari/537.11"

func UserAgent(r *http.Request) *user_agent.UserAgent {
	uaStr := r.Header.Get("User-Agent")
	return user_agent.New(uaStr)
}

func IsMobile(ua *user_agent.UserAgent) bool {
	return ua.Mobile()
}

func IsWeChat(ua *user_agent.UserAgent) bool {
	return ua.Mobile() && ua.Wecaht()
}

func IsIOS(ua *user_agent.UserAgent) bool {
	os := ua.OS()
	return ua.Mobile() && strings.Contains(os, "iPhone OS")
}

func IsAndroid(ua *user_agent.UserAgent) bool {
	os := ua.OS()
	return ua.Mobile() && strings.Contains(os, "Android")
}

func IsWeb(ua *user_agent.UserAgent) bool {
	return !ua.Mobile()
}
