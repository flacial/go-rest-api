package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

type PageData struct {
  PageTitle string
  PageSubtitle string
}

type IncrementNumber struct {
  NewNumber string
}

// States
var currentNum = 0

func main() {
  PORT := 8123

  http.HandleFunc("/", handleHomePagefunc) 
  http.HandleFunc("/increment-number", handleIncrementNumber)

  fmt.Printf("Server is running on port %s", strconv.Itoa(PORT))

  tscPort := ":" + strconv.Itoa(PORT)
  err := http.ListenAndServe(tscPort, nil)

  log.Fatal(err)
}

// GET
func handleHomePagefunc(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path != "/" {
    return
  }

  if r.Method != "GET" {
    fmt.Fprintf(w, "We only support GET, get, and gEt big guy")
    return
  }

  tmpl := template.Must(template.ParseFiles("layout.html"))

  data := PageData{
    PageTitle: "Just got my hands on Golang",
    PageSubtitle: `Ever since I touched Goland, I no longer think straight,<br/>
    Gotta keep the *facts straight, no room for debate,<br/>
    I code with precision, my flow so tight,<br/>
    I'm the master of the keys, I shine so bright.<br/><br/>
    
    I navigate the syntax like a lyrical wand,<br/>
    Coding rhythms that's smooth, never offbeat or wrong,<br/>
    From functions to loops, I bring the heat,<br/>
    Every line I drop, a masterpiece so sweet.<br/><br/>
    
    Goland's my playground, where I create,<br/>
    Lines of code that innovate, never hesitate,<br/>
    I stay on my grind, never miss a beat,<br/>
    In this rap game of code, I'm on the elite.<br/><br/>
    
    So step aside, let me show you my skill,<br/>
    Goland's the tool, and I'm the one with the will,<br/>
    I'll rap and code, break the boundaries and mold,<br/>
    This rhythm and logic, a tale yet untold.<br/>
    `,
  }

  tmpl.Execute(w, data)
}

func handleIncrementNumber(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path != "/increment-number" {
    return
  }

  if r.Method != "POST" {
    fmt.Fprintf(w, "We only support POST, post, and pOst big guy")
    return
  }

  currentNum += 1

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusCreated)

  json.NewEncoder(w).Encode(IncrementNumber{
    // newNumber with lowercase "n" doesn't work and Golang doesn't throw an error
    // To construct a new struct, the keys should be in PascalCase
    // If the struct key are invalid, it'll output an empty object/struct "{}"
    NewNumber: strconv.Itoa(currentNum),
  })
}