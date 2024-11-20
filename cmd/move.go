package cmd

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type MoveRequest struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func MovePlayer(url string, token string) {
	move := MoveRequest{
		X: 0,
		Y: 1,
	}

	body, err := json.Marshal(move)
	if err != nil {
		log.Fatal("Error marshaling JSON:", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		log.Fatal("Error creating request:", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error sending request:", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal("Error closing connection:", err)
		}
	}(resp.Body)

	var responseData struct {
		Data interface{} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
		log.Fatal("Error decoding response:", err)
	}

	log.Println(responseData.Data)
}
