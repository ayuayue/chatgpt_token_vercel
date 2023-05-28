package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func PandoraToken(w http.ResponseWriter, r *http.Request) {
	url := "https://auth0.openai.com/oauth/token"
	method := "POST"
	code := r.URL.Query().Get("code")
	payload := strings.NewReader(`{
    "redirect_uri": "com.openai.chat://auth0.openai.com/ios/com.openai.chat/callback",
    "grant_type": "authorization_code",
    "client_id": "pdlLIX2Y72MIl2rhLhTE9VV9bN905kBh",
    "code": "` + code + `",
    "code_verifier": "IkrrBD89CBmwwzM-csfBnWKLMan5uE7laCMd2YTcPWE"
}`)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	req.Header.Add("User-Agent", "apifox/1.0.0 (https://www.apifox.cn)")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Fprint(w, err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	fmt.Fprint(w, string(body))
}
