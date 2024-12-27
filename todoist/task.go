package todoist

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"time"
)

type Task struct {
	ID          string `json:"id,omitempty"`
	Content     string `json:"content"`
	Description string `json:"description,omitempty"`
	Due         Due    `json:"due,omitempty"`
	ProjectID   string `json:"project_id,omitempty"`
	Project     string `json:"-"`
	Priority    int    `json:"priority"`
}

type Due struct {
	Date         string `json:"date"`
	text         string `json:"string"`
	is_recurring bool   `json:"is_recurring"`
}

func (c *Client) GetTasks(cacheFile string, showProjects bool, filter string) ([]Task, error) {
	var query map[string]string
	if filter != "" {
		query = map[string]string{
			"filter": filter,
		}
	} else {
		query = nil
	}
	data, err := c.makeRequest("GET", "tasks", query, nil)
	if err != nil {
		return nil, err
	}

	var tasks []Task

	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, fmt.Errorf("failed to decode tasks json: %v", err)
	}

	// Sort Tasks in Ascending order of duedates
	sort.Slice(tasks, func(i, j int) bool {
		if tasks[i].Due.Date == "" {
			return false
		} else if tasks[j].Due.Date != "" {
			i, err := time.Parse("2006-1-2", tasks[i].Due.Date)
			if err != nil {
				return false
			}
			j, err := time.Parse("2006-1-2", tasks[j].Due.Date)
			if err != nil {
				return true
			}
			return i.Compare(j) == -1
		} else {
			return true
		}
	})

	if showProjects {
		for i := range tasks {
			tasks[i].setProject(c)
		}
	}

	data, err = json.Marshal(tasks)
	if err != nil {
		log.Printf("Tasks could not be written to log file: %v\n", err)
	}
	err = os.WriteFile(cacheFile, data, 0644)
	if err != nil {
		log.Printf("Tasks could not be written to log file: %v\n", err)
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
		data, err := c.makeRequest("GET", "projects/"+t.ProjectID, nil, nil)
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

func (c *Client) CreateTask(content string, duedate string, priority int, description string) error {
	if content == "" {
		return errors.New("empty task content")
	}
	task := map[string]interface{}{
		"content":     content,
		"priority":    priority,
		"due_date":    duedate,
		"description": description,
	}
	_, err := c.makeRequest("POST", "tasks", nil, task)
	if err != nil {
		return err
	}
	return nil
}

func (t *Task) GetDate() string {
	if t.Due.Date == "" {
		return ""
	}
	date, err := time.Parse("2006-1-2", t.Due.Date)
	if err != nil {
		log.Fatalln(err)
		return t.Due.Date
	}
	y, m, d := time.Now().Date()
	if date.Year() == y && date.Month() == m && date.Day() == d {
		return "Today"
	}

	if date.Compare(time.Now()) == -1 {
		return "Overdue"
	}
	return date.Format("2.1.2006")
}

func (c *Client) CloseTask(index int, cacheFile string) (string, error) {
	data, err := os.ReadFile(cacheFile)
	if err != nil {
		return "", err
	}
	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return "", err
	}
	if len(tasks) < index {
		return "", errors.New("task not found in cache-file")
	}
	id := tasks[index-1].ID
	_, err = c.makeRequest("POST", "tasks/"+id+"/close", nil, nil)
	if err != nil {
		return "", err
	}
	return tasks[index-1].Content, nil
}

func (c *Client) DeleteTask(index int, cacheFile string) (string, error) {
	data, err := os.ReadFile(cacheFile)
	if err != nil {
		return "", err
	}
	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return "", err
	}
	if len(tasks) < index {
		return "", errors.New("task not found in cache-file")
	}
	scanner := bufio.NewScanner(os.Stdin)
	task := tasks[index-1]
	fmt.Printf("Delete Task: '%s'? (y/n)\n", task.Content)
	scanner.Scan()
	del := scanner.Text() == "y"
	if err := scanner.Err(); err != nil {
		return "", err
	}
	if !del {
		return fmt.Sprintf("Task: '%s' was not deleted\n", task.Content), nil
	}
	_, err = c.makeRequest("DELETE", "tasks/"+task.ID, nil, nil)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Task: '%s' was deleted\n", task.Content), nil

}

func (t *Task) GetPriority() string {
	switch t.Priority {
	case 2:
		return "P3"
	case 3:
		return "P2"
	case 4:
		return "P1"
	default:
		return "P4"
	}
}
