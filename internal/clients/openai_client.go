package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// OpenAIRequest represents the request payload sent to the OpenAI API.
type OpenAIRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// OpenAIResponse represents the response payload received from the OpenAI API.
type OpenAIResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		FinishReason string `json:"finish_reason"`
		Message      struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		Index int `json:"index"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	}
}

type OpenAIError struct {
	Error struct {
		Message string `json:"message"`
		Type    string `json:"type"`
		Param   string `json:"param"`
		Code    string `json:"code"`
	} `json:"error"`
}

// OpenAIClient is a wrapper for the OpenAI API.
type OpenAIClient struct {
	ApiKey string
	ApiURL string
}

// NewOpenAIClient creates a new OpenAIClient.
func NewOpenAIClient(apiURL string, apiKey string) *OpenAIClient {
	return &OpenAIClient{ApiURL: apiURL, ApiKey: apiKey}
}

// Summarize summarizes the given text to a given level.
func (client *OpenAIClient) Summarize(content string, level string) (string, error) {

	if level != "short" && level != "medium" && level != "long" {
		return "", fmt.Errorf("invalid level: %s", level)
	}

	messages := []Message{
		{
			Role: "system",
			Content: "You are a content summarizer. Given a piece of text and a summarization level " +
				"you should summarize the given text to the appropriate level. The summarization levels are:\n" +
				"\"short\": <250 characters\n\"medium\": <500 characters\n\"long\": <800 characters.",
		},
		{
			Role:    "user",
			Content: "level: \"" + level + "\"\ncontent: \"" + content + "\"",
		},
	}
	// Create the request body
	reqBody := OpenAIRequest{
		Model:    "gpt-3.5-turbo",
		Messages: messages,
	}

	// Send the request
	jsonReq, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", client.ApiURL, bytes.NewBuffer(jsonReq))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+client.ApiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		var apiError OpenAIError
		err = json.Unmarshal(respBody, &apiError)
		if err != nil {
			return "", err
		}

		return "", fmt.Errorf("error message: %s, error type: %s, error code: %s",
			apiError.Error.Message, apiError.Error.Type, apiError.Error.Code)
	}

	var apiResponse OpenAIResponse
	err = json.Unmarshal(respBody, &apiResponse)
	if err != nil {
		return "", err
	}

	if apiResponse.Choices == nil || len(apiResponse.Choices) == 0 {
		return "", fmt.Errorf("unexpected response: %s", string(respBody))
	}

	return apiResponse.Choices[0].Message.Content, nil
}
