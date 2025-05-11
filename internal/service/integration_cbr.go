package service

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/beevik/etree"
)

func GetKeyRate() (float64, error) {
	from := time.Now().AddDate(0, 0, -30).Format("2006-01-02")
	to := time.Now().Format("2006-01-02")
	req := fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?>
    <soap12:Envelope xmlns:soap12="http://www.w3.org/2003/05/soap-envelope">
      <soap12:Body>
        <KeyRate xmlns="http://web.cbr.ru/">
          <fromDate>%s</fromDate>
          <ToDate>%s</ToDate>
        </KeyRate>
      </soap12:Body>
    </soap12:Envelope>`, from, to)

	resp, err := http.Post(
		"https://www.cbr.ru/DailyInfoWebServ/DailyInfo.asmx",
		"application/soap+xml; charset=utf-8",
		bytes.NewBufferString(req),
	)
	if err != nil {
		return 0, err
	}
	body, _ := io.ReadAll(resp.Body)
	doc := etree.NewDocument()
	doc.ReadFromBytes(body)
	els := doc.FindElements("//diffgram/KeyRate/KR/Rate")
	if len(els) == 0 {
		return 0, errors.New("rate not found")
	}
	var rate float64
	fmt.Sscanf(els[0].Text(), "%f", &rate)
	return rate, nil
}
