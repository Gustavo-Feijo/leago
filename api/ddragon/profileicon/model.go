package profileicon

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type (
	ProfileIconResponse struct {
		Type    string                 `json:"type"`
		Version string                 `json:"version"`
		Data    map[string]ProfileIcon `json:"data"`
	}

	ProfileIcon struct {
		ID    IntOrString `json:"id"`
		Image Image       `json:"image"`
	}

	Image struct {
		Full   string `json:"full"`
		Sprite string `json:"sprite"`
		Group  string `json:"group"`
		X      int    `json:"x"`
		Y      int    `json:"y"`
		W      int    `json:"w"`
		H      int    `json:"h"`
	}
)

// For some reason the profile iconIDs can come as normal integers or as integer strings.
type IntOrString int

func (i *IntOrString) UnmarshalJSON(b []byte) error {
	var num int
	if err := json.Unmarshal(b, &num); err == nil {
		*i = IntOrString(num)
		return nil
	}

	var str string
	if err := json.Unmarshal(b, &str); err == nil {
		n, err := strconv.Atoi(str)
		if err != nil {
			return err
		}
		*i = IntOrString(n)
		return nil
	}

	return fmt.Errorf("invalid id")
}
