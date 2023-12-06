<h1 align="center">Tech Tweakers - Polaris AI API v0.0.1 </h1>
<p align="center"><i>API Interface to deal with GGML AI Models, based on go-llama.cpp/llama.cpp</i></p>

<div align="center">
  <a href="https://github.com/Tech-Tweakers/ecatrom2000/stargazers"><img src="https://img.shields.io/github/stars/andreh1982/ecaterminal" alt="Stars Badge"/></a>
<a href="https://github.com/Tech-Tweakers/ecatrom2000/network/members"><img src="https://img.shields.io/github/forks/andreh1982/ecaterminal" alt="Forks Badge"/></a>
<a href="https://github.com/Tech-Tweakers/ecatrom2000/pulls"><img src="https://img.shields.io/github/issues-pr/andreh1982/ecaterminal" alt="Pull Requests Badge"/></a>
<a href="https://github.com/Tech-Tweakers/ecatrom2000/issues"><img src="https://img.shields.io/github/issues/andreh1982/ecaterminal" alt="Issues Badge"/></a>
<a href="https://github.com/Tech-Tweakers/ecatrom2000/graphs/contributors"><img alt="GitHub contributors" src="https://img.shields.io/github/contributors/andreh1982/ecaterminal?color=2b9348"></a>
<a href="https://github.com/Tech-Tweakers/ecatrom2000/blob/master/LICENSE"><img src="https://img.shields.io/github/license/andreh1982/ecaterminal?color=2b9348" alt="License Badge"/></a>
</div>

<br>
<p align="center"><i>Have some time to help? Please open an <a href="https://github.com/Tech-Tweakers/ecatrom2000/issues/new">Issue</a>.</i></p>

## TODO

**A lot!** :sweat_smile:

## Usage (Not tested yet)

```bash
# Clone this repository
$ git clone https://github.com/Tech-Tweakers/ecatrom2000.git

# Go into the repository
$ cd ecatrom2000

# Create a folder called "models"
$ mkdir models

# Copy the LLM file to the models folder
$ cp <path to LLM file> models/

# Install dependencies
$ go mod tidy

# Run the app
$ LIBRARY_PATH=$PWD C_INCLUDE_PATH=$PWD go run cmd/ecatrom2000/main.go
```

## Credits

Such awesome projects that made this possible:
| Tool | Link |
|------|------|
| **Go 1.21** | https://golang.org/ |
| **Go-LLama.cpp** | https://github.com/go-skynet/go-llama.cpp |
| **LLama.cpp** | https://github.com/ggerganov/llama.cpp |
| **The Bloke** | https://huggingface.co/TheBloke/Llama-2-7B-Chat-GGML/tree/main |

Love you all! Thank you so much for your hard work! :blue_heart: