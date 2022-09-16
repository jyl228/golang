package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"routes"
)

// Login
type Login struct {
	password string
	username string
}

var path = "test.txt"
var fileName string
var tasknum int

func init() {

	fmt.Println("1 ) 파일 생성 ")
	fmt.Println("2 ) 파일 쓰기 ")
	fmt.Println("3 ) 파일 읽기 ")
	fmt.Println("4 ) 파일 지우기 ")
	fmt.Println("원하는 작업의 번호를 입력해 주세요. : ")
	fmt.Scanf("%d", &tasknum)
}

func main() {
	//  http.Post
	// login := Login{"VMware1!", "stoneuser"}
	// lbytes, _ := json.Marshal(login)
	// buff := bytes.NewBuffer(lbytes)
	// resp, err := http.Post("http://192.168.204.158:3000/swagger/index.html#/default/post_api_user_auth", "application/json", buff)
	// if err != nil {
	// 	panic(err)
	// }

	// defer resp.Body.Close()

	// // Response 체크.
	// respBody, err := ioutil.ReadAll(resp.Body)
	// if err == nil {
	// 	str := string(respBody)

	// 	println(str)
	// }

	app := routes.Router(tasknum)
	log.Fatal(app.Listen(":3000"))

}

func createFile() {
	fmt.Println("생성할 파일 명 : ")
	fmt.Scanf("%s", &fileName) // 입력 받음
	fileName += ".txt"
	// 파일 존재 확인 - Stat returns a FileInfo describing the named file
	var _, err = os.Stat(fileName)

	// create file if not exists
	if os.IsNotExist(err) { //os.IsNotExist(err) : 파일이 존재 하지 않을때 true 로 if 문 탐
		var file, err = os.Create(fileName) // 파일 생성
		if isError(err) {
			return
		}
		defer file.Close()
		fmt.Println("파일 생성", fileName)
	} else {
		fmt.Println("해당 파일이 이미 존재합니다.")
	}

}
func writeFile() {
	fmt.Println("내용을 입력할 파일명 : ")
	fmt.Scanf("%s", &fileName) // 입력 받음
	fileName += ".txt"
	// 파일 플래그, 파일 모드를 지정하여 파일 열기가 가능해집니다.
	var file, err = os.OpenFile(fileName, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// Write some text line-by-line to file.
	_, err = file.WriteString("Hello \n")
	if isError(err) {
		return
	}
	_, err = file.WriteString("World \n")
	if isError(err) {
		return
	}

	// Save file changes.
	err = file.Sync()
	if isError(err) {
		return
	}

	fmt.Println("파일 업데이트")
}
func readFile() {
	// Open file for reading.
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// Read file, line by line
	var text = make([]byte, 1024)
	for {
		//한줄 씩 읽기
		_, err = file.Read(text)

		// 파일 끝을 만나면 break
		if err == io.EOF {
			break
		}

		// Break if error occured
		if err != nil && err != io.EOF {
			isError(err)
			break
		}
	}

	fmt.Println("Reading from file.")
	fmt.Println(string(text))
}

func deleteFile() {
	var err = os.Remove(path)
	if isError(err) {
		return
	}

	fmt.Println("File Deleted")
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}
