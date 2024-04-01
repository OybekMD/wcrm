package main

import (
	"fmt"
	"os"
)

func main() {
	// Get the current working directory
	req := "0af153d9-37f3-4376-aa07-27f052adf4fd.png"
	fileDir := "/home/oybek/projects/go/src/Asrlan/backend/api-gateway/media/"+req
	err := os.Remove(fileDir)
	if err != nil {
		// Handle the error if getting the working directory fails
		fmt.Println("Error:", err)
		return
	}
	
	fmt.Println("File deleted successfully:", fileDir)
}