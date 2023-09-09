package keeper

import (
	"container/list"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

type Keeper struct {
	history *list.List
	current *list.Element
}

func New() *Keeper {
	return &Keeper{
		history: list.New(),
	}
}

func (k *Keeper) Add(expression string) {
	k.history.PushBack(expression)
	k.current = k.history.Back()
}

func (k *Keeper) Get() string {
	if k.current == nil {
		return ""
	}

	defer func() {
		k.current = k.current.Prev()
	}()

	return k.current.Value.(string)
}

func (k *Keeper) Save(path string) {
	if _, err := os.Stat(path); err == nil || errors.Is(err, os.ErrNotExist) {
		var expressions []string
		for item := k.history.Front(); item != nil; item = item.Next() {
			expressions = append(expressions, fmt.Sprintf("%s", item.Value.(string)))
		}

		jsonBody, err := json.MarshalIndent(struct {
			Expressions []string `json:"expressions"`
		}{
			Expressions: expressions,
		}, "", "\t")
		if err != nil {
			logrus.Warning(err)
		}

		if err = os.WriteFile(path, jsonBody, 0644); err != nil {
			logrus.Warning(err)
		}
	} else {
		logrus.Warning(err)
	}
}

func (k *Keeper) Load(path string) {
	jsonData, err := os.ReadFile(path)
	if err != nil {
		return
	}

	data := struct {
		Expressions []string `json:"expressions"`
	}{}
	if err := json.Unmarshal(jsonData, &data); err != nil {
		logrus.Warning(err)
		return
	}

	for _, item := range data.Expressions {
		k.Add(item)
	}

	return
}

func (k *Keeper) Clear() {
	k.history.Init()
	k.current = nil
}
