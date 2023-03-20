package main

import (
	"github.com/D3xt3rrrr/verificationTax-service/data"
	"io"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

func pstVerification(tax *data.Tax) {
	res, err := http.Get(qstToken)
	if err != nil {
		log.Panic(err)
	}

	bodyBytes, err := io.ReadAll(res.Body)
	bodyString := string(bodyBytes)
	token := getToken(bodyString)

	cookies := res.Cookies()
	sendForm(cookies, token, tax)

	tax.IsQstValid = getResp(cookies)
}

func sendForm(cookies []*http.Cookie, token string, tax *data.Tax) {
	client := &http.Client{}
	form := url.Values{}
	form.Add("struts.token.name", "token")
	form.Add("token", token)
	form.Add("businessNumber", tax.PstNumber)
	form.Add("businessName", tax.Enterprise)
	form.Add("requestDate", "2019-04-08")
	form.Add("reg.label.submit", "Search")

	req, err := http.NewRequest("POST", qstRequest, strings.NewReader(form.Encode()))
	if err != nil {
		log.Panic(err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	for i := range cookies {
		req.AddCookie(cookies[i])
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Panic(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

}

func getToken(bodyString string) string {
	regex := "token..value=\"(.*).*\""
	r := regexp.MustCompile(regex)
	l := r.FindAllStringSubmatch(bodyString, -1)
	return l[0][1]
}

func getResp(cookies []*http.Cookie) bool {
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, qstResponse, nil)
	if err != nil {
		log.Panic(err)
	}

	for i := range cookies {
		req.AddCookie(cookies[i])
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Panic(err)
	}

	respByte, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}
	return strings.Contains(string(respByte), "GST/HST number registered on this transaction date.")
}
