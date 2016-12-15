package main

import (
  	"fmt"
	"io"
	"net/http"
	"os"
	"io/ioutil"
      )

func Pln(some string) {
  fmt.Println(some)
}

type processBody func(someBody io.Reader)

func GetAndProcess(url string, fn processBody)  {
  resp, err := http.Get(url)
  if err != nil {
    Pln(err.Error())
    os.Exit(1)
  }
  fn(resp.Body)
  defer resp.Body.Close()
}

func Dump(some io.Reader) {
  s, _ := ioutil.ReadAll(some)
  Pln(string(s))
}
