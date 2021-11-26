package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

//defer -> It actually executes any function passed into it, after the function finishes its final statement but before it actually returns
// primary uses of a defer statement is for cleaning up resources, such as open files, network connections, and database handles.

//Follows LIFO order : as if stored in stack.

func main() {
	defer fmt.Println("Bye")
	fmt.Println("Hi")

	a := "start"
	defer fmt.Println(a) //it will print start, even if a's value is changed.
	a = "end"


	if err:= write("game.txt", "This is a readme file."); err != nil{
		log.Fatal("Failed to write file: ", err)
	}

	res, err := http.Get("https://www.google.com/robots.txt")
	if err != nil{
		log.Fatal("Error in network call: ", err)
	}
	// network calls.
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}

func write(fileName string, text string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal("Error occurred while closing file: ",err)
		}
	}(file)
	_, err = io.WriteString(file, text)
	if err != nil {
		return err
	}
	return nil
}