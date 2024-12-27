# **Godoist: A Todoist CLI Tool**

[![Go Version](https://img.shields.io/github/go-mod/go-version/DukicDev/godoist)](https://golang.org/)  
[![License](https://img.shields.io/github/license/DukicDev/godoist)](LICENSE)

Godoist is a powerful command-line interface (CLI) tool for interacting with your [Todoist](https://todoist.com/) account. With Godoist, you can manage tasks, mark them as done, delete them, and much more—all from the comfort of your terminal.

---

## **Features**

- Add tasks with due dates, priorities, and descriptions.
- List tasks with filters like `today` or `overdue`.
- Mark tasks as completed.
- Delete tasks with confirmation.
- Optional project display for enhanced context.
- Fully customizable and easy to use.

---

## **Installation**

### **Using Go**
If you have Go installed, you can easily install Godoist:
```bash
go install github.com/DukicDev/godoist@latest
```

### **From Source**
Clone the repository and build the binary:
```bash
git clone https://github.com/DukicDev/godoist.git
cd godoist
go build -o godoist
```

---

## **Setup**

### **Todoist API Token**
1. Obtain your Todoist API token from [Todoist Settings](https://todoist.com/prefs/integrations).
2. Export the token as an environment variable:
   ```bash
   export TODOIST_API_TOKEN=your_api_token
   ```

---

## **Usage**

### **Available Commands**
| Command          | Description                                 | Example                                            |
|-------------------|---------------------------------------------|----------------------------------------------------|
| `godoist add`     | Add a new task                             | `godoist add "Buy groceries" -d 2024-12-31 -p 2`  |
| `godoist list`    | List tasks                                 | `godoist list --filter today`                     |
| `godoist done`    | Mark a task as completed                   | `godoist done 1`                                  |
| `godoist delete`  | Delete a task                              | `godoist delete 1`                                |

### **Flags**
- **Add Command (`add`)**:
  - `-d, --duedate`: Set the task's due date (e.g., `-d 2024-12-31`).
  - `-p, --priority`: Set the task's priority (1=highest, 4=lowest).
  - `--desc`: Add a description to the task.
- **List Command (`list`)**:
  - `-f, --filter`: Filter tasks (e.g., `today`, `overdue`).
  - `--show-projects`: Include project names in the output.

---

## **Examples**

### Add a Task
```bash
godoist add "Write documentation" -d 31.12.2024 --desc "Prepare for release"
```

### List Tasks Due Today
```bash
godoist list --filter today
```

### Mark a Task as Done
```bash
godoist done 1
```

### Delete a Task
```bash
godoist delete 2
```

---

## **License**

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

