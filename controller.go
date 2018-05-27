package main

import (
  "net/http"
  "log"
)

func ImageHandler(w http.ResponseWriter, r *http.Request) {
     keys, ok := r.URL.Query()["file"]
     height, ok1 := r.URL.Query()["height"]
     width, ok2 := r.URL.Query()["width"]

     if !ok || !ok1 || !ok2 || len(keys) < 1 || len(width) < 1 || len(height) < 1 {
     	log.Println("Url Param 'file' is missing")
	    HandleError(w,"Invalid Parameters")
	    return
	}
}
