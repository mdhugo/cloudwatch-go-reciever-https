package main

import (
	"log"
	"os"
	"encoding/json"
	"fmt"
	"net/http"
//	"io"
	"io/ioutil"
//	"html/template"
)

func writeoutput() {
	// temp testing with local file
	
	file, err := os.Open("alert.json")
	if err != nil {
		log.Fatal(err)
	}
	// end test

	decoder := json.NewDecoder(file)
//	encoder := json.NewEncoder(os.Stdout)

	for {
	    var v map[string]interface{}
	    if err := decoder.Decode(&v); err != nil {
		    log.Println(err)
		    return
	    }

	    
	    for k := range v {
		    if k == "MessageId" {
		        delete (v, k)
		    }
	    }
	    //if err := encoder.Encode(&v); err != nil {
	    //    log.Println(err)
	   // }
	   fmt.Println(v)
	}
}

func check(err error) { if err != nil { panic(err) } }

func upload(w http.ResponseWriter, r *http.Request) {
	//if r.Method != "POST" {
	//    uploadTemplate.Execute(w, nil)
        //    return
//	}
	    f, _, err := r.FormFile("json")
	log.Println(f)

//check(err)
	t, err := ioutil.TempFile("./", "json-")
	//check(err)
	defer t.Close()
//	io.Copy(t, f)
//	_, err = io.Copy(t, f)
//	check(err)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}	
}

func main() {
	http.HandleFunc("/cloudwatch2", upload) //errorHandler(upload))
	http.ListenAndServe(":8082", nil)

}
