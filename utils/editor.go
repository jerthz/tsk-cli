package utils

import (
    "os"
    "fmt"
    "os/exec"
    "strings"
    "regexp"
)

type EditorTask struct {
    Title string
    Category string
    Description string
}


func OpenTaskEditor(title string, category string, comment string) (string, error){
    template := `%s

# Category
%s

# Description
%s
    `

    tmpFile, err := os.CreateTemp("", "task_add_*.md")
    if err != nil {
        return "", err
    }
    defer tmpFile.Close()

    variableTemplate := fmt.Sprintf(template, title, category, comment)

    tmpFile.WriteString(variableTemplate)

    editor := "nvim"
    cmd := exec.Command(editor, tmpFile.Name())

    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    err = cmd.Run()

    if err != nil {
        return "", err
    }

    content, err := os.ReadFile(tmpFile.Name())
    if err != nil {
        return "", err
    }

    return string(content), nil 
}

func ParseTask(content string) EditorTask {
    var editedTask EditorTask

    titleRegex := regexp.MustCompile(`(?m)^(.*)$`)
	categoryRegex := regexp.MustCompile(`(?m)^# Category\n(.+)\n`)
	descriptionRegex := regexp.MustCompile(`(?m)^# Description\n((?:.|\n)*)`)

    if match := titleRegex.FindStringSubmatch(content); len(match) > 1 {
		editedTask.Title = strings.TrimSpace(match[1])
	}

	if match := categoryRegex.FindStringSubmatch(content); len(match) > 1 {
		editedTask.Category = strings.TrimSpace(match[1])
	}

	if match := descriptionRegex.FindStringSubmatch(content); len(match) > 1 {
		editedTask.Description = strings.TrimSpace(match[1])
	}


    return editedTask
}

