package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CrtSh []struct {
	IssuerCaID     int    `json:"issuer_ca_id"`
	IssuerName     string `json:"issuer_name"`
	CommonName     string `json:"common_name"`
	NameValue      string `json:"name_value"`
	ID             int64  `json:"id"`
	EntryTimestamp string `json:"entry_timestamp"`
	NotBefore      string `json:"not_before"`
	NotAfter       string `json:"not_after"`
	SerialNumber   string `json:"serial_number"`
}

func CRTSH(URL string) string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", URL, nil)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("can't requet to url", err)
	}
	defer res.Body.Close()
	resBody, _ := ioutil.ReadAll(res.Body)
	var result CrtSh
	if err := json.Unmarshal(resBody, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	return PrettyPrint(result)
}

func JldcMe(URL string) string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", URL, nil)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("can't requet to url", err)
	}
	defer res.Body.Close()
	resBody, _ := ioutil.ReadAll(res.Body)
	return string(resBody)
}


func AbuseDB(URL string) string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", URL, nil)
	req.Header.Set("user-agent", "Chrome")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("can't requet to url", err)
	}
	defer res.Body.Close()
	resBody, _ := ioutil.ReadAll(res.Body)
	return string(resBody)
}

// PrettyPrint to print struct in a readable way
func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
