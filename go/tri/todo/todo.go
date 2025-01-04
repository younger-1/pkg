package todo

import (
	"encoding/json"
	"os"
)

type Item struct {
	Text     string
	Priority int
}

func SaveItems(filename string, items []Item) error {
	b, err := json.Marshal(items)
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, b, 0644)
	if err != nil {
		return err
	}

	return nil
}

func ReadItems(filename string) ([]Item, error) {
	items := []Item{}

	b, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &items)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (i *Item) SetPriority(p int) {
	switch p {
	case 1:
		i.Priority = 1
	case 3:
		i.Priority = 3
	default:
		i.Priority = 2
	}
}

func (i *Item) PrettyP() string {
	switch i.Priority {
	case 1:
		return "[p1]"
	case 3:
		return "[p3]"
	}
	return ""
}
