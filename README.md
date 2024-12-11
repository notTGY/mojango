# mojango

Mojang api golang client

This packages primary use is in
[https://foragingupdate.com](https://foragingupdate.com).

So, if you need more data from mojang api,
file in pull request. No NITs, typos, tests, etc.

## Example Usage

```go
package main

import (
  "fmt"
  "github.com/nottgy/mojango"
)

func main() {
  player := "<YOUR MINECRAFT NICKNAME>"
  playerResponse, _ := mojango.QueryPlayerApi(player)
  fmt.Printf("%v\n", playerResponse)

  reversePlayerResponse, _ := mojango.
    QueryReversePlayerApi(playerResponse.UUID)
  fmt.Printf("%v", reversePlayerResponse)
}
```
