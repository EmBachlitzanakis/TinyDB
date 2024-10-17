package main

import (
	"TinyDB/TinyDB"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	db, err := TinyDB.NewTinyDB("data.json")
	if err != nil {
		fmt.Println("Error loading database:", err)
		return
	}
	defer db.Save("data.json")

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to TinyDB! Type 'help' for commands.")

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		parts := strings.SplitN(input, " ", 3)

		switch parts[0] {
		case "put":
			if len(parts) != 3 {
				fmt.Println("Usage: put <key> <value>")
				continue
			}
			db.Put(parts[1], parts[2])
			fmt.Println("Added:", parts[1])
		case "get":
			if len(parts) != 2 {
				fmt.Println("Usage: get <key>")
				continue
			}
			if value, exists := db.Get(parts[1]); exists {
				fmt.Println("Value:", value)
			} else {
				fmt.Println("Key not found.")
			}
		case "delete":
			if len(parts) != 2 {
				fmt.Println("Usage: delete <key>")
				continue
			}
			db.Delete(parts[1])
			fmt.Println("Deleted:", parts[1])
		case "help":
			fmt.Println("Commands: put <key> <value>, get <key>, delete <key>, exit")
		case "exit":
			return
		case "Search":
			if len(parts) != 2 {
				fmt.Println("Usage: getbyvalue <value>")
				continue
			}
			keys := db.Search(parts[1])
			if len(keys) > 0 {
				fmt.Println("Keys for value", parts[1], ":", keys)
			} else {
				fmt.Println("No keys found for value:", parts[1])
			}
		default:
			fmt.Println("Unknown command:", parts[0])
		}
	}
}
