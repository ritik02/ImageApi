package main

import (
  "net/http"
  "log"
  "os"
  "strconv"
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
     heights, ok1 := r.URL.Query()["height"]
     widths, ok2 := r.URL.Query()["width"]

     if !ok || !ok1 || !ok2 || len(file) < 1 || len(widths) < 1 || len(heights) < 1 {
     	log.Println("Url Param 'file' is missing")
	    HandleError(w,"Invalid Parameters")
	    return
     }

     if !DoesFileExist(file[0]) {
     	HandleError(w,"No such file found!")
	return
     }
     width, errw := strconv.Atoi(widths[0])
     height, errh := strconv.Atoi(heights[0])
     if errh != nil || errw != nil {
     	HandleError(w,"Invalid Parameter Value")
	return
     }
     if width > 1000 || height > 1000 {
     	HandleError(w,"Too Big Size Parameters")
	return
     }
	
}
