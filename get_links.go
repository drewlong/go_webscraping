package main

import (
  "fmt"
  "net/http"
  "golang.org/x/net/html"
  "os"
)

func getHTML(url string) string {
  resp, err := http.Get(url)
  if err != nil {
  }
  defer resp.Body.Close()
  z := html.NewTokenizer(resp.Body)
  for {
      tt := z.Next()

      switch {
      case tt == html.ErrorToken:
          return "Done."
      case tt == html.StartTagToken:
          t := z.Token()
          if t.Data == "a"{
            for _, a := range t.Attr{
              if a.Key == "href" {
                fmt.Println(a.Val)
              }
            }
          }
      }
  }
}

func main(){
  url   := os.Args[1]
  getHTML(url)
}

