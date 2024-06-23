package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type GPT3Request struct {
	Messages  []GPTMessage `json:"messages"`
	MaxTokens int          `json:"max_tokens"`
	Model     string       `json:"model"`
}

type GPTMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type GPT3Response struct {
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}

func checkCodingPrinciples(code string) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	message := GPTMessage{Role: "system", Content: fmt.Sprintf("Check the following code for basic coding principles and suggest improvements:")}
	userMessage := GPTMessage{Role: "user", Content: code}
	reqBody := GPT3Request{
		Messages:  []GPTMessage{message, userMessage},
		MaxTokens: 1500,
		Model:     "gpt-3.5-turbo",
	}
	reqBodyJson, _ := json.Marshal(reqBody)
	fmt.Println(string(reqBodyJson))
	req, _ := http.NewRequest("POST", "https://api.openai.com/v1/completions", bytes.NewBuffer(reqBodyJson))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request to OpenAI:", err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var gptResp GPT3Response
	err = json.Unmarshal(body, &gptResp)
	if err != nil {
		fmt.Println("Error reading request to OpenAI:", err)
		return
	}
	fmt.Println(gptResp.Choices)
	if len(gptResp.Choices) > 0 {
		fmt.Println("GPT-3 Analysis and Suggestions:")
		fmt.Println(gptResp.Choices[0].Text)
	}
}
