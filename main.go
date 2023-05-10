package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Content struct {
	InternetAccess bool     `json:"internet_access"`
	Conversation   []string `json:"conversation"`
	Parts          []Part   `json:"parts"`
}

type Part struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Meta struct {
	Content Content `json:"content"`
}

type Data struct {
	Action    string `json:"action"`
	Model     string `json:"model"`
	Jailbreak string `json:"jailbreak"`
	Meta      Meta   `json:"meta"`
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your message: ")
	message, _ := reader.ReadString('\n')

	data := Data{
		Action:    "_ask",
		Model:     "gpt-3.5-turbo-0301",
		Jailbreak: "default",
		Meta: Meta{
			Content: Content{
				InternetAccess: false,
				Conversation:   []string{},
				Parts: []Part{
					{
						Role:    "system",
						Content: "You are GPT-3.5 Turbo, a large language model trained by OpenAI. Knowledge cutoff: 2021-09-01. Current date: 2023-05-10.",
					},
					{
						Role:    "user",
						Content: message,
					},
				},
			},
		},
	}

	jsonData, _ := json.Marshal(data)

	resp, err := http.Post("http://shuttleproxy.com:6999/backend-api/v2/conversation", "application/json", bytes.NewBuffer(jsonData))

	if err != nil {
		fmt.Println(err)
	} else {
		defer resp.Body.Close()
		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)
		newStr := buf.String()

		fmt.Println(newStr)
	}
}
