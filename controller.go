package main

import (
  "net/http"
  "log"
  "os"
)

func DoesFileExist(filename string) bool {
     _, err := os.Stat("images/"+filename)
     if !os.IsNotExist(err) {
       return  true
     } else {
       return false
     }
}

func ImageHandler(w http.ResponseWriter, r *http.Request) {
     file, ok := r.URL.Query()["file"]
     height, ok1 := r.URL.Query()["height"]
     width, ok2 := r.URL.Query()["width"]

     if !ok || !ok1 || !ok2 || len(file) < 1 || len(width) < 1 || len(height) < 1 {
     	log.Println("Url Param 'file' is missing")
	    HandleError(w,"Invalid Parameters")
	    return
     }

     if !DoesFileExist(file[0]) {
     	HandleError(w,"No such file found!")
	return
     }

     infile, _ := os.Open(file[0])
     defer infile.Close()
	
}
