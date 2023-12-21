<h1 align="center">Tech Tweakers - Polaris Chat API v0.0.1 </h1>
<p align="center"><i>API Interface for GGUF Models, based on go-llama.cpp / llama.cpp</i></p>

<div align="center">
  <a href="https://github.com/Tech-Tweakers/ecatrom2000/stargazers"><img src="https://img.shields.io/github/stars/andreh1982/ecaterminal" alt="Stars Badge"/></a>
<a href="https://github.com/Tech-Tweakers/ecatrom2000/network/members"><img src="https://img.shields.io/github/forks/andreh1982/ecaterminal" alt="Forks Badge"/></a>
<a href="https://github.com/Tech-Tweakers/ecatrom2000/pulls"><img src="https://img.shields.io/github/issues-pr/andreh1982/ecaterminal" alt="Pull Requests Badge"/></a>
<a href="https://github.com/Tech-Tweakers/ecatrom2000/issues"><img src="https://img.shields.io/github/issues/andreh1982/ecaterminal" alt="Issues Badge"/></a>
<a href="https://github.com/Tech-Tweakers/ecatrom2000/graphs/contributors"><img alt="GitHub contributors" src="https://img.shields.io/github/contributors/andreh1982/ecaterminal?color=2b9348"></a>
<a href="https://github.com/Tech-Tweakers/ecatrom2000/blob/master/LICENSE"><img src="https://img.shields.io/github/license/andreh1982/ecaterminal?color=2b9348" alt="License Badge"/></a>
</div>

<br>
<p align="center"><i>Have some time to help? Please open an <a href="https://github.com/Tech-Tweakers/ecatrom2000/issues/new">Issue</a> to say hello!</i></p>

## About

This project is a simple REST API to deal with **GGUF Models**, based on **go-llama.cpp** / **llama.cpp**. It' can be used to create a chatbot, or even a simple chat interface.

## To Do

**VectorDB** to chat persistance;
Work directly with **llama.cpp**;

## Install (Not tested yet)

```bash
# Clone this repository:
git clone https://github.com/Tech-Tweakers/polaris-chatbot.git --recurse-submodules

# Enter in the folder:
cd polaris-chatbot

# Create a new folder called "models":
mkdir models

# Copy the LLM file to the models folder:
cp <path to LLM file> models/

# Install dependencies:
go get all
go mod tidy

```
## Setup

Edit the **.env.local** file with your own settings:

```bash
ENVIRONMENT=development
APP_VERSION=v1.0.0
APP_PORT=9001
APP_URL=http://localhost:9001
LOG_LEVEL=debug
#
# Setup to use DynamoDB
#
DEFAULT_PERSISTENT=false # False use MemoryDB, True use DynamoDB

AWS_REGION=us-east-1 
AWS_ENDPOINT=http://localhost:4566
AWS_PROFILE=default # AWS Profile to be used even if you use localstack

DYNAMO_AWS_ENDPOINT=http://localhost:4566
DYNAMO_TABLE_NAME=ecatrom2000

# Path to the model - GGUF Models ONLY
MODEL_PATH="./models/llama-2-7b-chat.Q2_K.gguf"
```
After that, goes inside the folder go-llama.cpp and run the following command:

```bash
make clean
make libbinding.a
```
## Run the API:

Run in one line:

```bash
LIBRARY_PATH=$PWD C_INCLUDE_PATH=$PWD go run cmd/ecatrom2000/main.go
```
Or just run the script:

```bash
./run-api.sh
```

## API Usage

Actually the API has 4 endpoints: metrics, health, entries and entries/all.

```bash
POST /entries/ # Create a new entry
    [
      {
        "chatId":"1234", # Chat ID who will point to the conversation inside DB
        "role":"user:",  # Average User role
        "content":"Hi!"  # Message to be sent to the model
      }
    ]

GET /entries/all # Get all entries in DB

GET /health # Check if the API is up and running

GET /metrics # Get some metrics about the API
```

## Credits

Such awesome projects that made this possible:
| Tool | Link |
|------|------|
| **Go 1.21** | https://golang.org/ |
| **Go-LLama.cpp** | https://github.com/go-skynet/go-llama.cpp |
| **LLama.cpp** | https://github.com/ggerganov/llama.cpp |
| **The Bloke** | https://huggingface.co/TheBloke/Llama-2-7B-Chat-GGUF/tree/main |

Love you all! Thank you so much for your hard work! :blue_heart: