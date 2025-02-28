[![Go Report Card](https://goreportcard.com/badge/github.com/dev-shimada/urlencode)](https://goreportcard.com/report/github.com/dev-shimada/urlencode)
[![CI](https://github.com/dev-shimada/urlencode/actions/workflows/CI.yaml/badge.svg)](https://github.com/dev-shimada/urlencode/actions/workflows/CI.yaml)
[![License](https://img.shields.io/badge/license-MIT-blue)](https://github.com/dev-shimada/urlencode/blob/master/LICENSE)

# urlencode

## Usage
To encode a string:
```
echo "Hello World" | urlencode encode
```
To encode spaces as %20:
```
echo "Hello World" | urlencode encode --use-percent-20
```
To decode a string:
```
echo "Hello%20World" | urlencode decode
```

## Installation
Install via Homebrew:
```
brew install dev-shimada/urlencode/urlencode
```

## Generating Shell Completions
For bash:
```
urlencode completion bash > /etc/bash_completion.d/urlencode
```
For zsh:
```
urlencode completion zsh > ~/.zshrc.d/_urlencode
```
For fish:
```
urlencode completion fish | source
```
