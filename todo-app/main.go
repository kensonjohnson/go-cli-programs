package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Todos struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
}

func main() {
	counter := 3
	data := []Todos{
		{
			Id:          0,
			Description: "Create go app",
			Status:      true,
		},
		{
			Id:          1,
			Description: "Delete this todo with remove",
			Status:      false,
		},
		{
			Id:          2,
			Description: "Profit",
			Status:      false,
		},
	}
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Commands: show create remove done\n>")

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		switch split[0] {
		case "create":
			data = append(data, Todos{
				Id:          counter,
				Description: strings.Join(split[1:], " "),
				Status:      false,
			})
			counter++
			fmt.Println("Todo item created!")

		case "remove":
			index, err := strconv.Atoi(split[1])
			if err != nil {
				panic(err)
			}

			for i, todo := range data {
				if todo.Id == index {
					data = slices.Delete(data, i, i+1)
					break
				}
			}

		case "show":
			for _, todo := range data {
				statusLine := "[ ]"
				if todo.Status == true {
					statusLine = "[X]"
				}
				fmt.Printf("\n%d: %s %s\n", todo.Id, statusLine, todo.Description)
			}
			fmt.Print("\n")

		case "done":
			index, err := strconv.Atoi(split[1])
			if err != nil {
				panic(err)
			}
			for i, todo := range data {
				if todo.Id == index {
					data[i].Status = true
					break
				}
			}

		default:
			break

		}
		fmt.Print(">")
	}
}
