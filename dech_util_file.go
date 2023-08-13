package dechutil

import (
	"encoding/json"
	"fmt"

	"os"
	"strings"
)

func CurrentCallDir() string {
	path, err := os.Getwd()
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	// fmt.Println(path) // for example /home/user
	return path
}

func CurrentExecDir() string {
	path, err := os.Executable()
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	str := "/"
	ind := strings.LastIndex(path, str)
	if ind == -1 { // os is windows
		str = `\`
		ind = strings.LastIndex(path, str)
	}

	path = path[:ind]

	return path
}

func ReadFileAndDirName(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("failed opening directory: %s\n", err)
	}
	defer file.Close()

	list, _ := file.Readdirnames(0) // 0 to read all files and folders
	// fmt.Printf("%#v\n", list)
	// for _, name := range list {
	// 	fmt.Println(name)
	// }

	return list
}

func ReadFileName(path string) []string {
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	result := []string{}

	for _, file := range files {
		// fmt.Println(file.Name(), file.IsDir())
		if !file.IsDir() {
			result = append(result, file.Name())
		}
	}

	return result
}

func ReadDirName(path string) []string {
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	result := []string{}

	for _, file := range files {
		// fmt.Println(file.Name(), file.IsDir())
		if file.IsDir() {
			result = append(result, file.Name())
		}
	}

	return result
}

func MoveFile(fileName string, oldLocation string, newLocation string) {
	from := oldLocation + "/" + fileName
	to := newLocation + "/" + fileName

	err := os.Rename(from, to)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
}

func CreateFolderPath(Path string) error {
	err := os.MkdirAll(Path, os.ModePerm)
	return err
}

func WriteJsonFile(createFolderName string, fileName string, dataJson any) {
	if createFolderName != "" {
		CreateFolderPath(createFolderName)
	}

	data, err := json.MarshalIndent(dataJson, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = os.WriteFile(fileName, data, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

func ReadJsonFile(fileName string, refData any) {
	plan, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = json.Unmarshal(plan, &refData)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func WriteBinaryFile(createFolderName string, fileName string, data []byte) {
	if createFolderName != "" {
		CreateFolderPath(createFolderName)
	}

	err := os.WriteFile(fileName, data, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
