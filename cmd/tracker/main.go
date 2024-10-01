package main

import (
    "fmt"
    "os"
    "strings"
    "strconv"
    "yask-tracker/internal/enums"
    "yask-tracker/internal/services"
)

func main() {
    service := services.NewTaskService("/app/data/tasks.json")

    if len(os.Args) < 2 {
        fmt.Println("No command provided. Run 'yask-tracker help' for usage.")
        return
    }

    args := os.Args
    command := args[1]

    switch command {
    case "addTask":
        if len(args) < 3 {
            fmt.Println("Usage: addTask <description>")
            return
        }
        description := strings.Join(args[2:], " ")
        task, err := service.CreateTask(description)
        if err != nil {
            fmt.Println("Error adding task:", err)
            return
        }
        fmt.Printf("Task added: %+v\n", task)

    case "deleteTask":
        if len(args) < 3 {
            fmt.Println("Usage: deleteTask <id>")
            return
        }
        id, err := strconv.Atoi(args[2])
        if err != nil {
            fmt.Println("Invalid ID:", args[2])
            return
        }
        err = service.DeleteTask(id)
        if err != nil {
            fmt.Println("Error deleting task:", err)
            return
        }
        fmt.Println("Task deleted.")

    case "listTasks":
        var status *enums.Status
        if len(args) >= 3 {
            stat := args[2]
            if stat != "all" {
                statusVal := enums.Status(strings.ToLower(stat))
                status = &statusVal
            }
        }
        tasks, err := service.ListTasks(status)
        if err != nil {
            fmt.Println("Error listing tasks:", err)
            return
        }
        for _, task := range tasks {
            fmt.Printf("ID: %d, Description: %s, Status: %s, CreatedAt: %s, UpdatedAt: %s\n",
                task.Id, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)
        }

    case "updateTask":
        if len(args) < 4 {
            fmt.Println("Usage: updateTask <id> <description>")
            return
        }
        id, err := strconv.Atoi(args[2])
        if err != nil {
            fmt.Println("Invalid ID:", args[2])
            return
        }
        description := strings.Join(args[3:], " ")
        task, err := service.UpdateTask(id, description, enums.Todo) // Status is set to Todo by default
        if err != nil {
            fmt.Println("Error updating task:", err)
            return
        }
        fmt.Printf("Task updated: %+v\n", task)

    case "updateStatusTask":
        if len(args) < 4 {
            fmt.Println("Usage: updateStatusTask <id> <status>")
            return
        }
        id, err := strconv.Atoi(args[2])
        if err != nil {
            fmt.Println("Invalid ID:", args[2])
            return
        }
        status := enums.Status(strings.ToLower(args[3]))
        task, err := service.ChangeTaskStatus(id, status)
        if err != nil {
            fmt.Println("Error updating task status:", err)
            return
        }
        fmt.Printf("Task status updated: %+v\n", task)

    case "help":
        fmt.Println("Yask Tracker CLI Help")
        fmt.Println("Usage:")
        fmt.Println("  addTask <description>        - Add a new task with the provided description.")
        fmt.Println("  deleteTask <id>              - Delete the task with the provided ID.")
        fmt.Println("  listTasks <status>           - List tasks by status. Status can be: all, todo, in-progress, done.")
        fmt.Println("  updateTask <id> <description> - Update the description of the task with the provided ID.")
        fmt.Println("  updateStatusTask <id> <status> - Update the status of the task with the provided ID. Status can be: todo, in-progress, done.")
        fmt.Println("  exit                        - Exit the program.")
        fmt.Println()
        fmt.Println("Example usage:")
        fmt.Println("  docker run -it --rm -v $(pwd)/data:/app/data yask-tracker addTask \"Finish Go project\"")
        fmt.Println("  docker run -it --rm -v $(pwd)/data:/app/data yask-tracker deleteTask 1")
        fmt.Println("  docker run -it --rm -v $(pwd)/data:/app/data yask-tracker listTasks all")
        fmt.Println("  docker run -it --rm -v $(pwd)/data:/app/data yask-tracker updateTask 2 \"Update Go project description\"")
        fmt.Println("  docker run -it --rm -v $(pwd)/data:/app/data yask-tracker updateStatusTask 2 in-progress")

    default:
        fmt.Println("Unknown command:", command)
        fmt.Println("If you need help you can run: yask-tracker help")
    }
}