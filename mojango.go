package mojango

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
  "errors"
)

func QueryReversePlayerApi(uuid string) (ReversePlayerResponse, error) {
	var reversePlayerResponse ReversePlayerResponse
	url := fmt.Sprintf("https://sessionserver.mojang.com/session/minecraft/profile/%s", uuid)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return reversePlayerResponse, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return reversePlayerResponse, err
	}
  if resp.StatusCode != 200 {
    err = errors.New(fmt.Sprintf(
      "Code %d", resp.StatusCode,
    ))
    return reversePlayerResponse, err
  }
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return reversePlayerResponse, err
	}

	if err := json.Unmarshal(body, &reversePlayerResponse); err != nil {
		return reversePlayerResponse, err
	}
	return reversePlayerResponse, nil
}

func QueryPlayerApi(player string) (
  PlayerResponse, error,
) {
	var playerResponse PlayerResponse
	url := fmt.Sprintf(
    "https://api.mojang.com/users/profiles/minecraft/%s?at=0",
    player,
  )
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return playerResponse, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return playerResponse, err
	}
  if resp.StatusCode != 200 {
    err = errors.New(fmt.Sprintf(
      "Code %d", resp.StatusCode,
    ))
    return playerResponse, err
  }
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return playerResponse, err
	}

	if err := json.Unmarshal(body, &playerResponse); err != nil {
		return playerResponse, err
	}
	return playerResponse, nil
}
