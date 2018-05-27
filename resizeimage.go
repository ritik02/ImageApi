package main

import (
       "gopkg.in/gographics/imagick.v2/imagick"
       "net/http"
)

func ResizeImage(w http.ResponseWriter,finalfile string, inputfile string, width int, height int , imagetype string) {

  imagick.Initialize()
  defer imagick.Terminate()
  mw := imagick.NewMagickWand()
  err := mw.ReadImage(inputfile)
  if err != nil {
    HandleError(w,"Invalid Image Type")
    return
  }

  err = mw.ResizeImage(uint(width),uint(height), imagick.FILTER_LANCZOS, 1)
  
  if err == nil {
  err = mw.SetImageCompressionQuality(100)
  }
  if err == nil {
  err = mw.WriteImage(finalfile)
  DecodeImage(w, finalfile, imagetype)
  }
}
