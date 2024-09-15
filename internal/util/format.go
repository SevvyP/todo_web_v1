package util

import (
	"encoding/json"

	"github.com/SevvyP/tasks_v1/pkg/model"
)

type Node struct {
	ID        string  `json:"id"`
	Body      string  `json:"body"`
	Completed bool    `json:"completed"`
	Reminder  *string `json:"reminder"`
	Children  *[]Node `json:"children"`
}

func FormatTaskJson(tasks []model.Task) (json.RawMessage, error) {
	if tasks == nil {
		return nil, nil
	}

	nodeMap := make(map[string]*Node)
	var roots []Node

	// Populate the map with node references
	for _, task := range tasks {
		node := Node{
			ID:        task.ID,
			Body:      task.Body,
			Completed: task.Completed,
			Reminder:  task.Reminder,
			Children:  &[]Node{},
		}
		nodeMap[task.ID] = &node
	}

	// Build the tree
	for _, task := range tasks {
		node := nodeMap[task.ID]
		if task.Parent == nil {
			roots = append(roots, *node)
		} else {
			parent := nodeMap[*task.Parent]
			*parent.Children = append(*parent.Children, *node)
		}
	}

	// Marshal the root nodes into JSON
	jsonData, err := json.MarshalIndent(roots, "", "  ")
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}
