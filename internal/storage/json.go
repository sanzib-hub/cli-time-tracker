package storage

import (
	"cli/internal/task"
	"encoding/json"
	"os"
	"path/filepath"
)



var dataFile = filepath.Join("data","tasks.json")

func LoadTask() ([]task.Task, error){
	file, err := os.Open(dataFile)
	if os.IsNotExist(err) {
		return []task.Task{}, nil
	}else if err !=nil{
		return nil,err
	}

	defer file.Close()

	var tasks []task.Task
	err = json.NewDecoder(file).Decode(&tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func saveTask(tasks []task.Task)error{

	err := os.MkdirAll(filepath.Dir(dataFile), 0755)
	if err != nil{
		return err
	}

	file, err := os.Create(dataFile)
	if err != nil{
		return nil
	}

	defer file.Close()

	enc := json.NewEncoder(file)

	enc.SetIndent("", "  ")
	err = enc.Encode(tasks)
	return enc.Encode(tasks)
}
