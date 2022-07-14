package main

import (
  "encoding/base64"
  "fmt"
)

func main() {
  encoded := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"
  decoded, _ := base64.StdEncoding.DecodeString(encoded)
  fmt.Println(string(decoded))
}