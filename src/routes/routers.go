package routes

import (
	"fmt"
	"io"
	"os"

	"github.com/gofiber/fiber/v2"
)

var fileName string
var path = ".txt"

func Router(tasknum int) *fiber.App {
	app := fiber.New() // fiber 인스턴스 생성
	switch tasknum {
	case 1:
		writest()
		app.Post("/file/create", createFile)
	case 2:
		// GET /dictionary.txt
		//app.PUT("/", func(c *fiber.Ctx) error {
		// 	file, err := c.FormFile("document")

		// 	if err == nil {
		// 		c.SaveFile(file, fmt.Sprintf("./%s", file.Filename))
		// 	}
		// 	return c.SendString("파일 읽기 종료")
		// })
		app.Put("/file/update", writeFile)
	case 3:
		app.Get("/file/Read", readFile)
	case 4:
		app.Delete("/file/delete", deleteFile)
	default:
		fmt.Println("잘못 입력하였습니다.")
	}
	// app.Get("/file/Read")    // 파일 안의 데이터 리턴
	// app.Put("/file/create")  // 파일 생성
	// app.Post("/file/update") // 파일에 데이터를 추가
	// app.Post("/file/delete") // 파일을 삭제

	// GET /api/list

	return app

}
func readFile(c *fiber.Ctx) error {
	var file, err2 = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err2) {
		return err2
	}
	defer file.Close()

	// Read file, line by line
	var text = make([]byte, 1024)
	for {
		//한줄 씩 읽기
		_, err2 = file.Read(text)

		// 파일 끝을 만나면 break
		if err2 == io.EOF {
			break
		}

		// Break if error occured
		if err2 != nil && err2 != io.EOF {
			isError(err2)
			break
		}
	}
	return c.SendString("파일 읽기 종료")
}
func deleteFile(c *fiber.Ctx) error {
	var err = os.Remove(path)
	if isError(err) {
		return err
	}
	return c.SendString("파일 삭제 종료")
}
func writest() string {
	fmt.Println("내용을 입력할 파일명 : ")
	fmt.Scanf("%s", &fileName) // 입력 받음
	fileName += ".txt"
	return fileName
}
func createFile(c *fiber.Ctx) error {
	// 파일 존재 확인 - Stat returns a FileInfo describing the named file
	fileName := "파일입니다.txt"
	var _, err = os.Stat(fileName)

	// create file if not exists
	if os.IsNotExist(err) { //os.IsNotExist(err) : 파일이 존재 하지 않을때 true 로 if 문 탐
		var file, err = os.Create(fileName) // 파일 생성
		if isError(err) {
			return err
		}
		defer file.Close()
		fmt.Println("파일 생성", fileName)
	} else {
		fmt.Println("해당 파일이 이미 존재합니다.")
		c.SendString("해당 파일이 이미 존재합니다.")
	}
	return c.SendString("파일 생성 종료")
}

func writeFile(c *fiber.Ctx) error {
	return c.SendString("파일 수정 종료")
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())

	}

	return (err != nil)
}
