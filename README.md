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
