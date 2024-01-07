package queryapi

import (
	"ecatrom/internal/infrastructure/environment"
	"fmt"
	"os"

	"github.com/go-skynet/go-llama.cpp"
)

func SendMessage(ecatromEntity string, m *llama.LLama) (replyMessage string) {

	replyMessage = WorkerLlama(m, ecatromEntity)

	return replyMessage
}

func WorkerLlama(l *llama.LLama, question string) (replyMessage string) {

	send2AI := " user: " + question
	replyMessage, err := l.Predict(send2AI, llama.SetTokenCallback(func(token string) bool { return true }),
		llama.SetTokens(2048),
		llama.SetThreads(6),
		llama.SetTopK(20),
		llama.SetTopP(0.50),
		llama.SetTemperature(0.1),
		llama.SetNKeep(0),
		llama.SetSeed(0),
		llama.SetPresencePenalty(0),
		llama.SetFrequencyPenalty(2),
		// llama.SetPathPromptCache("./cache"),
		llama.SetStopWords("user:", "User:", "system:", "System:"),
	)
	if err != nil {
		panic(err)
	}

	l.Free()

	return replyMessage
}

func LoadAiModel(kind string) (l *llama.LLama) {

	var err error
	var loadModel string

	env := environment.GetInstance()

	if kind == "chat" {
		loadModel = env.CHAT_MODEL_PATH
	} else if kind == "code" {
		loadModel = env.CODE_MODEL_PATH
	}

	l, err = llama.New(loadModel, llama.EnableF16Memory, llama.SetContext(2048), llama.SetGPULayers(0))
	if err != nil {
		fmt.Println("Loading the model failed:", err.Error())
		os.Exit(1)
	}

	fmt.Println("Model " + kind + " loaded successfully.")

	return l
}
