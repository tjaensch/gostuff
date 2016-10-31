package main

 import (
         "fmt"
         "net/http"
 )

 func Home(w http.ResponseWriter, r *http.Request) {
         html := `<html>
  <title>Go upload</title>
  <body>

  <form action="http://10.90.235.15:1313/getcheckbox" method="post" enctype="multipart/form-data">
  <label for="file">Filenames:</label>
  <input type="file" name="multiplefiles" id="multiplefiles" multiple>
  <input type="checkbox" name="compress" value="compressboxchecked">
  <input type="submit" name="submit" value="Submit">
  </form>

  </body>
  </html>`

         //output to web browser(client)

         w.Write([]byte(fmt.Sprintf(html)))

 }

 func GetCheckBox(w http.ResponseWriter, r *http.Request) {

         err := r.ParseMultipartForm(200000) // grab the multipart form

         if err != nil {
                 fmt.Fprintln(w, err)
                 return
         }

         // get the value of checkbox

         // if the checkbox is not ticked, the value will be empty

         fmt.Fprintln(w, r.FormValue("compress"))
         fmt.Println("Compress the file ? : ", r.FormValue("compress"))

 }

 func main() {

         mux := http.NewServeMux()

         mux.HandleFunc("/", Home)
         mux.HandleFunc("/getcheckbox", GetCheckBox)

         http.ListenAndServe("10.90.235.15:1313", mux)
 }
