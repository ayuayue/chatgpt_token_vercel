package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetToken(w http.ResponseWriter, r *http.Request) {
	url := "https://auth0.openai.com/oauth/token"
	method := "POST"
	code := r.PostFormValue("code")
	payload := `{
		"redirect_uri": "com.openai.chat://auth0.openai.com/ios/com.openai.chat/callback",
		"grant_type": "authorization_code",
		"client_id": "pdlLIX2Y72MIl2rhLhTE9VV9bN905kBh",
		"code": "` + code + `",
		"code_verifier": "yGrXROHx_VazA0uovsxKfE263LMFcrSrdm4SlC-rob8"
	  }`
	client := &http.Client{}
	req, err := http.NewRequest(method, url, strings.NewReader(payload))

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
