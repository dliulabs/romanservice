package main

import (
	"bytes"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os/exec"
)

func getCommandOutput(command string, args ...string) string {
	cmd := exec.Command(command, args...)
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Start()
	if err != nil {
		log.Fatal(fmt.Sprint(err) + ": " + stderr.String())
	}
	err = cmd.Wait()
	if err != nil {
		log.Fatal(fmt.Sprint(err) + ": " + stderr.String())
	}
	return stdout.String()
}

func goVersion(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprintf(w, getCommandOutput("/usr/local/go/bin/go", "version"))
}

func getFileContent(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprintf(w, getCommandOutput("/bin/cat", params.ByName("name")))
}

func main() {
	router := httprouter.New()
	log.Println("curl -X GET http://localhost:8000/api/v1/go-version")
	log.Println("curl -X GET http://localhost:8000/api/v1/show-file/latin.txt")
	router.GET("/api/v1/go-version", goVersion)
	router.GET("/api/v1/show-file/:name", getFileContent)
	log.Fatal(http.ListenAndServe(":8000", router))
}
