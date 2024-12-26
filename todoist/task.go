package todoist

import (
	"encoding/json"
	"fmt"
	"log"
)

type Task struct {
	ID          string `json:"id"`
	Content     string `json:"content"`
	Priority    int    `json:"priority,omitempty"`
	Description string `json:"description,omitempty"`
	DueDate     string `json:"due_date,omitempty"`
	ProjectID   string `json:"project_id,omitempty"`
	Project     string `json:"-"`
}

func (c *Client) GetTasks(showProjects bool) ([]Task, error) {
	data, err := c.makeRequest("GET", "tasks", nil)
	if err != nil {
		return nil, err
	}

	var tasks []Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, fmt.Errorf("failed to decode tasks json: %v", err)
	}

	if showProjects {
		for i := range tasks {
			tasks[i].setProject(c)
		}

	}
	return tasks, nil
}

func (t *Task) ShortContent(maxLength int) string {
	if len(t.Content) > maxLength {
		return t.Content[:maxLength-3] + "..."
	}
	return t.Content
}

func (t *Task) setProject(c *Client) {
	if t.ProjectID != "" {
		data, err := c.makeRequest("GET", "projects/"+t.ProjectID, nil)
		if err != nil {
			log.Printf("error while setting project for task: %v\n error: %v", t, err)
		}
		var project map[string]interface{}
		if err := json.Unmarshal(data, &project); err != nil {
			log.Printf("error while setting project for task: %v\n error: %v", t, err)
		}
		t.Project = project["name"].(string)
	}
}
