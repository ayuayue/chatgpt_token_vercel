package api

import (
	"fmt"
	"net/http"
	"strings"
	"regexp"
	"os"
	"net/url"
)

func GetLoginToken(w http.ResponseWriter, r *http.Request) {
	loginUrl := os.Getenv("LoginURL")+"/auth/login"
	method := "POST"
	params := url.Values{}
	params.Set("username", r.PostFormValue("username"))
	params.Set("password", r.PostFormValue("password"))
	params.Set("mfa_code", r.PostFormValue("mfa_code"))
	params.Set("action", "default")
	params.Set("next", "")
	payload := strings.NewReader(params.Encode())
	fmt.Println(params.Encode())

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	req, err := http.NewRequest(method, loginUrl, payload)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	req.Header.Add("User-Agent", "apifox/1.0.0 (https://www.apifox.cn)")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, err := client.Do(req)
	if err != nil {
		fmt.Fprint(w, err)
	}
	defer res.Body.Close()


	// 获取重定向返回的 header
	redirectHeaders := res.Header

	// 将 header 转换为字符串
	var sb strings.Builder
	for key, values := range redirectHeaders {
		for _, value := range values {
			sb.WriteString(key)
			sb.WriteString(": ")
			sb.WriteString(value)
			sb.WriteString("\r\n")
		}
	}
	redirectHeadersStr := sb.String()


	// 使用正则表达式匹配 access_token 的内容
	regex := regexp.MustCompile(`access-token=(.*?);`)
	matches := regex.FindStringSubmatch(redirectHeadersStr)

	if len(matches) > 1 {
		accessToken := matches[1]
		fmt.Fprint(w, "Access Token:"+accessToken)
	} else {
		fmt.Fprint(w, "Access Token not found in the response.")
	}
}
