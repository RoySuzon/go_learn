package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"runtime"
	"time"
)

func openBrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = exec.Command("xdg-open", url).Start()
	}
	if err != nil {
		fmt.Println("ℹ️ Please open your browser and navigate to:", url)
	}
}

func main() {
	port := ":8080"
	url := "http://localhost" + port

	// Serve current directory static files
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)

	fmt.Println("==================================================")
	fmt.Println(" 🌐 Go Master Documentation Web Server Started!")
	fmt.Println(" 🔗 Address:", url)
	fmt.Println("==================================================")

	go func() {
		time.Sleep(500 * time.Millisecond)
		openBrowser(url)
	}()

	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Printf("❌ Failed to start web server: %v\n", err)
	}
}
