package todoist

import (
	"encoding/json"
	"fmt"
)

type Task struct {
	ID          string `json:"id"`
	Content     string `json:"content"`
	Description string `json:"description,omitempty"`
	DueDate     string `json:"due_date,omitempty"`
}

func (c *Client) GetTasks() ([]Task, error) {
	data, err := c.makeRequest("GET", "tasks", nil)
	if err != nil {
		return nil, err
	}

	var tasks []Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, fmt.Errorf("failed to decode tasks json: %v", err)
	}

	return tasks, nil
}

func (t *Task) ShortContent(maxLength int) string {
	if len(t.Content) > maxLength {
		return t.Content[:maxLength-3] + "..."
	}
	return t.Content
}
