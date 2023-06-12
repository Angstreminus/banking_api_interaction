package xmlparcer

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/html/charset"
)

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Text    string   `xml:",chardata"`
	Date    string   `xml:"Date,attr"`
	Name    string   `xml:"name,attr"`
	Valute  []struct {
		Text     string `xml:",chardata"`
		ID       string `xml:"ID,attr"`
		NumCode  string `xml:"NumCode"`
		CharCode string `xml:"CharCode"`
		Nominal  string `xml:"Nominal"`
		Name     string `xml:"Name"`
		Value    string `xml:"Value"`
	} `xml:"Valute"`
}

func SendAndParse(dateArr []string) []ValCurs {
	httpClient := &http.Client{}
	var ParsedResult []ValCurs

	for _, date := range dateArr {
		URL := fmt.Sprintf("http://www.cbr.ru/scripts/XML_daily.asp?date_req=%s", date)
		req, err := http.NewRequest("GET", URL, nil)

		if err != nil {
			fmt.Println("Unable to create request!", err)
		}

		req.Header.Add("User-Agent", "test-task-app")

		resp, err := httpClient.Do(req)

		if err != nil {
			fmt.Println("Can`t accept response", err)
		}

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			fmt.Println("Error while reading respone body", err)
		}

		resp.Body.Close()

		decoder := xml.NewDecoder(bytes.NewReader(body))
		decoder.CharsetReader = charset.NewReaderLabel

		var valCurs ValCurs

		err = decoder.Decode(&valCurs)

		if err != nil {
			fmt.Println("Error while deserialize data ", err)
		}
		ParsedResult = append(ParsedResult, valCurs)
	}
	return ParsedResult
}
