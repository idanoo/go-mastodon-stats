package gomastodonstats

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type MatrixWebhook struct {
	Body string `json:"body"`
	Key  string `json:"key"`
}

func sendToMatrix(m []metric) {
	if MATRIX_WEBHOOK_URL == "" || len(m) == 0 {
		log.Println("Skipping posting to Matrix as missing env vars")
		return
	}

	startOfDay := getStartofDay()
	sendMatrixWebhook(fmt.Sprintf("*User counts for %s*", startOfDay.String()), MATRIX_WEBHOOK_CHANNEL)
	for _, m := range m {
		sendMatrixWebhook(getPrintableString(m, false), MATRIX_WEBHOOK_CHANNEL)
	}
}

// sendMatrixWebhook - takes msg, sends to matrix
func sendMatrixWebhook(msgText string, channel string) error {
	// log.Println(msgText)
	data := MatrixWebhook{
		Key: MATRIX_WEBHOOK_API_KEY,
	}
	data.Body = msgText
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", MATRIX_WEBHOOK_URL+"/"+channel, bytes.NewBuffer(b))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Print(err)
		return nil
	}

	defer resp.Body.Close()

	return nil
}
