package main
import (
  "fmt"
  "golang.org/x/net/html"
  "net/http"
  "github.com/moul/http2curl"
	"golang.org/x/net/proxy"

  // "os"
  // "strings"
)

func getImg(t html.Token) (ok bool, href string){
  for _, a := range t.Attr {
    if a.Key == "img" {
      href = a.Val
      ok = true
    }
  }
  return
}
func crawl(url string) {
  client := &http.Client{
	   CheckRedirect: redirectPolicyFunc,
  }
  resp, err := client.Get(url)
  if err != nil {
		fmt.Println("ERROR: Failed to crawl \"" + url + "\"")
		return
	}
  b := resp.Body
  defer b.Close()
  z := html.NewTokenizer(b)
  for {
    tt := z.Next()
    switch {
    case tt == html.ErrorToken:
      // End of the document, we're done
      return
    case tt == html.StartTagToken:
      t := z.Token()

      fmt.Println("Token: ", t)
      // Check if the token is an <a> tag
      isAnchor := t.Data == "img"
      if !isAnchor {
        continue
      }

      // Extract the href value, if there is one
      ok, url := getImg(t)
      if !ok {
        continue
      }
      fmt.Println("Img src: ", url)

      // Make sure the url begines in http**
      // hasProto := strings.Index(url, "http") == 0
      // if hasProto {
      //   ch <- url
      // }
    }
  }
  fmt.Print(b)
}
func main() {
  crawl("http://minitoon.net/bbs/board.php?bo_table=183&wr_id=166")
  // crawl("http://www.google.com")

}
