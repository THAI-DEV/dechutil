package dechutil

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"math/rand"
	"path/filepath"
	"syscall"
	"time"

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

func ReadFileAndDirName(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		// fmt.Printf("failed opening directory: %s\n", err)
		return nil, err
	}
	defer file.Close()

	list, _ := file.Readdirnames(0) // 0 to read all files and folders
	// fmt.Printf("%#v\n", list)
	// for _, name := range list {
	// 	fmt.Println(name)
	// }

	return list, nil
}

func ReadFileName(path string) ([]string, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		// fmt.Printf("%s\n", err)
		return nil, err
	}

	result := []string{}

	for _, file := range files {
		// fmt.Println(file.Name(), file.IsDir())
		if !file.IsDir() {
			result = append(result, file.Name())
		}
	}

	return result, nil
}

func ReadDirName(path string) ([]string, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		// fmt.Printf("%s\n", err)
		return nil, err
	}

	result := []string{}

	for _, file := range files {
		// fmt.Println(file.Name(), file.IsDir())
		if file.IsDir() {
			result = append(result, file.Name())
		}
	}

	return result, nil
}

func CreateFolderPath(Path string) error {
	err := os.MkdirAll(Path, os.ModePerm)
	return err
}

func WriteJsonFile(createFolderName string, fileName string, dataJson any) error {
	if createFolderName != "" {
		CreateFolderPath(createFolderName)
	}

	data, err := json.MarshalIndent(dataJson, "", "  ")
	if err != nil {
		// fmt.Printf("%s\n", err)
		return err
	}

	err = os.WriteFile(fileName, data, 0644)
	if err != nil {
		// fmt.Printf("%s\n", err)
		return err
	}

	return nil
}

func ReadJsonFile(fileName string, refData any) error {
	plan, err := os.ReadFile(fileName)
	if err != nil {
		// fmt.Printf("%s\n", err)
		return err
	}

	err = json.Unmarshal(plan, &refData)
	if err != nil {
		// fmt.Printf("%s\n", err)
		return err
	}

	return nil
}

func WriteDataFile(createFolderName string, fileName string, data []byte) error {
	if createFolderName != "" {
		CreateFolderPath(createFolderName)
	}

	err := os.WriteFile(fileName, data, 0644)
	if err != nil {
		// fmt.Printf("%s\n", err)
		return err
	}

	return nil
}

func CreateSingleDir(filePathName string) error {
	//* Can not create existing directory and create only single directory
	if err := os.Mkdir(filePathName, os.ModePerm); err != nil {
		return err
	}

	return nil
}

func CreateHierarchyDir(filePathName string) error {
	//* Can create existing directory and hierarchy directory
	if err := os.MkdirAll(filePathName, os.ModePerm); err != nil {
		return err
	}

	return nil
}

func CreateNewFile(filePathName string) error {
	f, err := os.Create(filePathName)
	if err != nil {
		return err
	}

	defer f.Close()

	return nil
}

func DeleteFile(filePathName string) error {
	err := os.Remove(filePathName)
	if err != nil {
		return err
	}

	return nil
}

func MoveFile(fileName string, oldLocation string, newLocation string) error {
	from := oldLocation + "/" + fileName
	to := newLocation + "/" + fileName

	err := os.Rename(from, to)
	if err != nil {
		return err
	}

	return nil
}

func ReadDirAndFileInfoOneLevel(folderName string) ([]fs.FileInfo, error) {
	result := []fs.FileInfo{}

	files, err := os.ReadDir(folderName)
	if err != nil {
		return nil, err
	}

	for _, files := range files {

		info, _ := files.Info()
		result = append(result, info)
	}

	return result, nil
}

func Walk(folderName string) ([]fs.FileInfo, []string, error) {
	result := []fs.FileInfo{}
	pathList := []string{}

	err := filepath.Walk(folderName,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			pathList = append(pathList, path)
			result = append(result, info)
			return nil
		})
	if err != nil {
		// log.Fatal(err)
		return nil, nil, err
	}

	return result, pathList, nil
}

func ShowFileInfo(fileInfoList []fs.FileInfo) {
	for i, info := range fileInfoList {
		fmt.Printf("No : %d | Name : %s | IsDir : %v | Size : %d | Mod Time : %v \n", i+1, info.Name(), info.IsDir(), info.Size(), info.ModTime())
	}
}

func ShowWalkInfo(fileInfoList []fs.FileInfo, pathList []string) {
	for i, info := range fileInfoList {
		fmt.Printf("No : %d | Name : %s | Path : %s | IsDir : %v | Size : %d | Mod Time : %v \n", i+1, info.Name(), pathList[i], info.IsDir(), info.Size(), info.ModTime())
	}
}

// * Ex : ModifyFileTime(`/Temp/t/a.txt`, GetFileTime("2022-09-20 13:21:56"))
func ModifyFileTime(filePath string, ftime syscall.Filetime) error {
	fd, err := syscall.Open(filePath, os.O_RDWR, 0755)
	if err != nil {
		log.Println(filePath, err)
	}

	err = syscall.SetFileTime(fd, &ftime, &ftime, &ftime)
	if err != nil {
		log.Println(filePath, err)
	}
	syscall.Close(fd)

	return err
}

func GetFileTime(dt string) syscall.Filetime {
	longTimeFormat := "2006-01-02 15:04:05"

	loc, _ := time.LoadLocation("Local")
	st, err := time.ParseInLocation(longTimeFormat,
		dt, loc)
	if err != nil {
		log.Println(err)
	}

	t, err := time.Parse(longTimeFormat, dt)
	if err != nil {
		log.Println(err)
	}

	t = t.Add(time.Second * 1)
	dt2 := t.Format(longTimeFormat)

	et, err := time.ParseInLocation(longTimeFormat,
		dt2, loc)
	if err != nil {
		log.Println(err)
	}

	gapTimestamp := et.Unix() - st.Unix()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// rand.Seed(time.Now().UnixNano())

	randGap := r.Int63n(gapTimestamp)

	distTime := uniformTime(time.Unix(st.Unix()+randGap, 0))

	ftime := syscall.NsecToFiletime(distTime.UnixNano())

	return ftime
}

func uniformTime(t time.Time) time.Time {
	abbrevTimeFormat := "2006-1-2 15:4:5"

	loc, _ := time.LoadLocation("Local")
	// rand.Seed(time.Now().UnixNano())
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	startYmdStr := fmt.Sprintf("%d-%d-%d 09:30:00", t.Year(), t.Month(), t.Day())
	endYmdStr := fmt.Sprintf("%d-%d-%d 18:30:00", t.Year(), t.Month(), t.Day())

	daySt, _ := time.ParseInLocation(abbrevTimeFormat, startYmdStr, loc)
	dayEt, _ := time.ParseInLocation(abbrevTimeFormat, endYmdStr, loc)
	dayGap := dayEt.Unix() - daySt.Unix()
	randGap := r.Int63n(dayGap)

	//fmt.Println(t, startYmdStr, daySt, dayEt, dayGap, randGap)

	if t.Unix() < daySt.Unix() || t.Unix() > dayEt.Unix() {
		return time.Unix(daySt.Unix()+randGap, 0)
	}
	return t
}
