// 예제 7-13. 인코더를 사용해 구조체를 JSON으로 생성하기

package main

import (
	"encoding/json"
	"fmt"
	_ "io"
	"os"
)

type Post struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {

	post := Post{
		Id:      1,
		Content: "Hello World!",
		Author:  Author{
			Id:   2,
			Name: "Sau Sheong",
		},
		Comments: []Comment{
			Comment{
				Id:      3,
				Content: "Have a great day!",
				Author:  "Adam",
			},
			Comment{
				Id:      4,
				Content: "How are you today?",
				Author:  "Betty",
			},
		},
	}
	jsonFile, err := os.Create("post.json")
	if err != nil {
		fmt.Println("Error creating JSON file:", err)
		return
	}
	encoder := json.NewEncoder(jsonFile)
	err = encoder.Encode(&post)
	if err != nil {
		fmt.Println("error encoding JSON to file:", err)
		return
	}
}
