package services

import (
    "encoding/json"
    "fmt"
    "os"
    "time"
    "yask-tracker/internal/enums"
    "yask-tracker/internal/models"
)

// TaskService provides methods to manage tasks
type TaskService struct {
    filePath string
}

// NewTaskService creates a new TaskService
func NewTaskService(filePath string) *TaskService {
    return &TaskService{filePath: filePath}
}

func (s *TaskService) LoadTasks() ([]models.Task, error) {

    // Try open the file
    file, err := os.Open("/app/data/tasks.json")
    if err != nil {

        // If the error is "no such file or directory", create the file
        if os.IsNotExist(err) {
            
            // Create the file and initialize it as an empty array
            emptyTasks := []models.Task{}
            if err := s.SaveTasks(emptyTasks); err != nil {
                return nil, fmt.Errorf("Error creating tasks.json: %v", err)
            }

            return emptyTasks, nil
        }
        return nil, fmt.Errorf("Error opening tasks.json: %v", err)
    }
    defer file.Close()

    // Read the file and deserialize it
    tasks := []models.Task{}
    decoder := json.NewDecoder(file)
    err = decoder.Decode(&tasks)
    if err != nil {
        return nil, fmt.Errorf("Error decoding tasks.json: %v", err)
    }

    return tasks, nil
}

// SaveTasks saves tasks to the JSON file
func (s *TaskService) SaveTasks(tasks []models.Task) error {
    file, err := os.Create(s.filePath)
    if err != nil {
        return err
    }
    defer file.Close()

    return json.NewEncoder(file).Encode(tasks)
}

// CreateTask adds a new task
func (s *TaskService) CreateTask(description string) (models.Task, error) {
    tasks, err := s.LoadTasks()
    if err != nil {
        return models.Task{}, err
    }

    newTask := models.Task{
        Id:          len(tasks) + 1, // Simple ID generation
        Description: description,
        Status:      enums.Todo,
        CreatedAt:   time.Now(),
        UpdatedAt:   time.Now(),
    }

    tasks = append(tasks, newTask)
    err = s.SaveTasks(tasks)
    return newTask, err
}

// UpdateTask updates an existing task
func (s *TaskService) UpdateTask(id int, description string, status enums.Status) (models.Task, error) {
    tasks, err := s.LoadTasks()
    if err != nil {
        return models.Task{}, err
    }

    for i, task := range tasks {
        if task.Id == id {
            tasks[i].Description = description
            tasks[i].Status = status
            tasks[i].UpdatedAt = time.Now()
            err = s.SaveTasks(tasks)
            return tasks[i], err
        }
    }
    return models.Task{}, nil // Task not found
}

// DeleteTask removes a task
func (s *TaskService) DeleteTask(id int) error {
    tasks, err := s.LoadTasks()
    if err != nil {
        return err
    }

    for i, task := range tasks {
        if task.Id == id {
            tasks = append(tasks[:i], tasks[i+1:]...) // Remove the task
            return s.SaveTasks(tasks)
        }
    }
    return nil // Task not found
}

// ListTasks lists all tasks or filters them by status if provided
func (s *TaskService) ListTasks(status *enums.Status) ([]models.Task, error) {
    tasks, err := s.LoadTasks()
    if err != nil {
        return nil, err
    }

    // If status is nil, return all tasks
    if status == nil {
        return tasks, nil
    }

    // Filter tasks by the provided status
    var filteredTasks []models.Task
    for _, task := range tasks {
        if task.Status == *status {
            filteredTasks = append(filteredTasks, task)
        }
    }
    return filteredTasks, nil
}

// ChangeTaskStatus updates the status of a task
func (s *TaskService) ChangeTaskStatus(id int, newStatus enums.Status) (models.Task, error) {
    tasks, err := s.LoadTasks()
    if err != nil {
        return models.Task{}, err
    }

    for i, task := range tasks {
        if task.Id == id {
            tasks[i].Status = newStatus
            tasks[i].UpdatedAt = time.Now()
            err = s.SaveTasks(tasks)
            return tasks[i], err
        }
    }
    return models.Task{}, nil // Task not found
}
