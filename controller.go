package main

import (
  "net/http"
  "log"
  "os"
  "strconv"
  "bytes"
  "image"
  "image/jpeg"
  "image/png"
)

func DoesFileExist(filename string) bool {
     _, err := os.Stat(filename)
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
     cachefile := file[0]
     file[0] = "images/"+file[0]
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

     infile, _ := os.Open(file[0])
     defer infile.Close()
     imagetype := getFormat(infile)
     finalfile := "images/cached/"+cachefile+"_final_"+widths[0]+"_"+heights[0]+"."+imagetype
   
     if DoesFileExist(finalfile) {
     	log.Println("File Cached Already")
    	DecodeImage(w, finalfile, imagetype)
    	return
     } else {
       log.Println("File Not Cached")
     }
     ResizeImage(w, finalfile, file[0], width, height, imagetype)
}

func DecodeImage(w http.ResponseWriter, finalfile string, imagetype string) {
  infile, _ := os.Open(finalfile)
  defer infile.Close()

  src, _, _ := image.Decode(infile)
  WriteImage(w, &src, imagetype)
}


func getFormat(file *os.File) (string) {
  bytes := make([]byte, 4)
  file.ReadAt(bytes, 0)
  if bytes[0] == 0x89 && bytes[1] == 0x50 && bytes[2] == 0x4E && bytes[3] == 0x47 { return "png" }
  if bytes[0] == 0xFF && bytes[1] == 0xD8 { return "jpg" }
  return ""
}

func WriteImage(w http.ResponseWriter, img *image.Image,imgtype string) {
  buffer := new(bytes.Buffer)
  switch imgtype {
    case "jpg" :
     	  jpeg.Encode(buffer, *img, nil)
     	   w.Header().Set("Content-Type", "image/jpeg")
    case "png" :
     	  png.Encode(buffer, *img)
       	  w.Header().Set("Content-Type", "image/png")
        }
  w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
   w.Write(buffer.Bytes())
}
