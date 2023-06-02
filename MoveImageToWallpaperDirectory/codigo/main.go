package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func main() {
	oldPath := {"path of images will be downloaded"}
	newPath := {"path of wallpaper directory/where i want images go"}

	for {
		files, err := ioutil.ReadDir(oldPath)
		if err != nil {
			fmt.Println(err)
		}

		for _, file := range files {

			time.Sleep(500 * time.Millisecond)

			fileName := file.Name()

			action := os.Rename(oldPath+fileName, newPath+fileName)
			if action != nil {
				fmt.Println(action)
			} else {
				log.Println("file: '" + fileName + "' successufully moved!")
			}
		}
	}
}
