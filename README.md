This is a Go implemented REST API for Resizing images.

STEPS TO FOLLOW -

  1) Download or clone the repository using "git clone https://github.com/ritik02/ImageApi.git" .
  2) Now, cd into the directory where the project has been downloaded.
  3) Now Run the Server using "go build && ./imageapi" .
  4) Open the Browser and type "http://localhost:8000/api/resize?file=test.jpg&width=400&height=300"  to view the resized image


---------------------------

BASIC FUNCTIONING -

    The Code uses "ImageMagick" for Resizing Images based upon given width and height in URL parameters.
    The Code has all necessary validations for checking parameter and image types .
    There is also implementation of caching using file system if the image is already saved in "images/cached/" directory.

  
