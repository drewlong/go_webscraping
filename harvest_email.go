package main

import (
  "fmt"
  "strings"
  "net/http"
  "io/ioutil"
  "os"
  "regexp"
)

func getHTML(url string) string {
  resp, err := http.Get(url)
  if err != nil {
  }
  // Close http stream after body is read and html is parsed
  defer resp.Body.Close()
  // Get bytes of http response body
  body_bytes, err := ioutil.ReadAll(resp.Body)
  // Convert bytes to string
  body := string(body_bytes)
  return body
}

func harvestEmail(word string) []string{
  re := regexp.MustCompile(`['.'\w|-]*@+[a-z]+[.]+\w{2,}`)
  match := re.FindStringSubmatch(word)
  return match
}

func main(){
  url   := os.Args[1]
  html  := getHTML(url)
  words := strings.Fields(html)
  for i := 0; i < len(words); i++ {
    match := harvestEmail(words[i])
    l := len(match)
    if l > 0 {
      fmt.Println(match)
    }
  }
}

// Usage: ./harvest http://example.com
