package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)


func RevokeToken(w http.ResponseWriter, r *http.Request) {
	url := "https://auth0.openai.com/oauth/revoke"
	method := "POST"
	token := r.PostFormValue("token")
	payload := `{
		"client_id": "pdlLIX2Y72MIl2rhLhTE9VV9bN905kBh",
		"token": "` + token + `"
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
