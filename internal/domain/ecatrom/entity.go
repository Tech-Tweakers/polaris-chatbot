package ecatrom

import "time"

type ReplyEntity struct {
	CreatedAt time.Time `json:"createdAt"`
	ChatID    string    `json:"chatId"`
	EntryID   string    `json:"entryId"`
	Name      string    `json:"name"`
	AIReply   string    `json:"reply"`
}

type ChatPersistence struct {
	Key       string    `json:"key"`
	EntryID   float64   `json:"entryId"`
	ChatID    string    `json:"chatId"`
	CreatedAt time.Time `json:"createdAt"`
	Role      string    `json:"role"`
	Content   string    `json:"content"`
}

type AiQuestion struct {
	Model          string      `json:"model"`
	File           string      `json:"file"`
	Language       string      `json:"language"`
	ResponseFormat string      `json:"response_format"`
	Size           string      `json:"size"`
	Prompt         interface{} `json:"prompt"`
	Instruction    string      `json:"instruction"`
	Input          interface{} `json:"input"`
	Stop           interface{} `json:"stop"`
	Messages       Messages    `json:"messages"`
	Stream         bool        `json:"stream"`
	Echo           bool        `json:"echo"`
	TopP           int         `json:"top_p"`
	TopK           int         `json:"top_k"`
	Temperature    float64     `json:"temperature"`
	MaxTokens      int         `json:"max_tokens"`
	N              int         `json:"n"`
	Batch          int         `json:"batch"`
	F16            bool        `json:"f16"`
	IgnoreEos      bool        `json:"ignore_eos"`
	RepeatPenalty  int         `json:"repeat_penalty"`
	NKeep          int         `json:"n_keep"`
	MirostatEta    int         `json:"mirostat_eta"`
	MirostatTau    int         `json:"mirostat_tau"`
	Mirostat       int         `json:"mirostat"`
	Seed           int         `json:"seed"`
	Mode           int         `json:"mode"`
	Step           int         `json:"step"`
}

type AiResponse struct {
	Object  string `json:"object"`
	Model   string `json:"model"`
	Choices []struct {
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

type Messages []struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatSummary struct {
	Role    string
	Content string
}
