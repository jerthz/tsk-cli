package utils

import (
    "fmt"
    "os"
    "errors"
    "encoding/json"
    "strings"
)

import . "go-todo-cli/models"

func GetTasksDirectoryPath() string {
    home, err := os.UserHomeDir()
    if err != nil {
        fmt.Println("Error : can't find the user home repository")
        os.Exit(1)
    }

    return home + "/.tskcli"
}

func GetMasterConfigPath() string {
    return GetTasksDirectoryPath() + "/task-cli-master.json"
}

func checkDirectoryExistence() bool {
    return checkPathExistence(GetTasksDirectoryPath())
}

func checkMasterConfigExistence() bool {
    return checkPathExistence(GetMasterConfigPath())
}

func checkPathExistence(path string) bool { 
    _, err := os.OpenFile(path, os.O_RDONLY, 644)
    res := errors.Is(err, os.ErrNotExist)
    return !res
}

func SaveMaster(config MasterConfig) {
    data,_ := json.MarshalIndent(config, "", "  ")
    file, err := os.OpenFile(GetMasterConfigPath(), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
    if err != nil {
        fmt.Println("Error while opening the master config file")
    }
    defer file.Close()

    _, err = file.Write(data)
    if err != nil {
        fmt.Println("Error while saving master config")
    }
}   

func LoadMaster() MasterConfig {
    var config MasterConfig
    file, err := os.ReadFile(GetMasterConfigPath())
    if err != nil {
        fmt.Println("Error while opening master configuration")
        os.Exit(1)
    }
    json.Unmarshal(file, &config)
    return config
}

func MasterInit() {
    if !checkDirectoryExistence() {
        fmt.Println("$home/.tskcli folder does not exist, creating...");
        err := os.Mkdir(GetTasksDirectoryPath(), 0755)
        if err != nil{
            fmt.Println("Error :", err)
            os.Exit(1)
        }
        fmt.Println("Repository created")
    }
    if !checkMasterConfigExistence() {
        fmt.Println("$home/.tskcli file does not exist, creating...")
        masterConfig := MasterConfig {
            LastId: 0,
        }
        SaveMaster(masterConfig)
    }
}

func LoadTasks() []Task{
    var tasks []Task
    file, err := os.ReadFile(GetTaskFilePath())
    if err == nil {
        json.Unmarshal(file, &tasks)
    }
    return tasks
}

func GetTaskFilePath() string {
    return GetTasksDirectoryPath() + "/task-cli-tasks.json"
}

func SaveTasks(tasks []Task) {
    data, _ := json.MarshalIndent(tasks, "", "  ")
    file, err := os.OpenFile(GetTaskFilePath(), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
    if err != nil {
        fmt.Println("Erreur lors de l'ouverture du fichier :", err)
        return
    }
    defer file.Close()

    _, err = file.Write(data)
    if err != nil {
        fmt.Println("Error while writing in the tasks file", err)
    }
}



var Filter string
var Category string
var Detail bool
var All bool
var Status string

func FilterTasks(tasks []Task) []Task{
    var filteredCategory []Task
    for _, task := range tasks {
        if Category == "" || strings.Contains(task.Category,Category) {
            filteredCategory = append(filteredCategory, task)
        }
    }
    var filteredText []Task
    for _, task := range filteredCategory{
        if Filter == "" || strings.Contains(task.Description,Filter) {
            filteredText = append(filteredText, task)
        }
    }  
    var filteredStatus []Task
    for _, task := range filteredText {
        if All || strings.ToLower(Status) == "completed" || task.Status != Completed {
            filteredStatus = append(filteredStatus, task)
        }
    }

    var filteredStatus2 []Task
    for _, task := range filteredStatus {
        if Status != "" {
            taskStatus, err := StringToTaskStatus(Status)
            if err != nil{
                fmt.Printf("\nThe given status does not exist, please use Pending|InProgress|Completed|Stashed\n\n")
            }

            if taskStatus == task.Status {
                filteredStatus2 = append(filteredStatus2, task)
            }
        } else{
            filteredStatus2 = append(filteredStatus2, task)
        }
    }

    return filteredStatus2
}


func GetTaskById(taskId int) (Task, error) {
    tasks := LoadTasks()
    
        for i := range tasks{
            if tasks[i].Id == taskId {
                return tasks[i], nil
            }
        }
		return Task{}, fmt.Errorf("Unknown task : %d", taskId)
}

func UpdateTask(task Task) error {
    tasks := LoadTasks()
    
        for i := range tasks{
            if tasks[i].Id == task.Id {
                tasks[i].Comment = task.Comment
                SaveTasks(tasks)
                return nil
            }
        }
        return fmt.Errorf("Error while saving the task")
}
