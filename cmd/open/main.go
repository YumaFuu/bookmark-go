package main

import (
	"bookmarks/db"
	"bookmarks/models"
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	db := db.NewDB()
	defer db.Close()

	var u models.URL

	// fmt.Prinlt("id > ")
	i := getInput()
	inputs := strings.Fields(i)

	if len(inputs) == 0 {
		return
	}
	err := db.Where("id = ?", inputs[0]).First(&u).Error
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(u.Url)
	exec.Command("open", u.Url).Run()
}

func getInput() string {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	return stdin.Text()
}
