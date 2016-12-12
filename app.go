package main

import  (
	"fmt"
        "net/http"
        "io"
	"strings"
        "os"
        "golang.org/x/net/html")


type Game struct {
  date, team1, team2, time, tv, league string
}

var games []Game
var currentDate string

func main() {
  games = make([]Game, 0)
  const mainUrl = "http://www.footballonuktv.com/"
  getGames(mainUrl)
  for _, g := range games {
    formatGame(g)
  }
}

func formatGame(g Game) {
  fmt.Printf("%s\t%s\t %s\t\"%#s - %#s\" on %s\n", g.date, g.time, g.league, g.team1, g.team2, g.tv)	
}

func getGames(url string)  {
  resp, err := http.Get(url)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  parse(resp.Body)
  defer resp.Body.Close()
}

func visit(node *html.Node) {
  if hasAttr(node, "dia.php?fecha") {
    currentDate = node.FirstChild.FirstChild.Data
  }
  if hasAttr(node, "div_partido") {
    extractGame(node)
  } else {
    for c := node.FirstChild; c != nil; c = c.NextSibling {
      visit(c)
    }
  }
}

func extractGame(node *html.Node) {
  g := Game{}
  g.date = currentDate
  for c := node.FirstChild; c != nil; c = c.NextSibling {
    if hasAttr(c, "div_cadena") {
      g.tv = c.FirstChild.FirstChild.Data
    }
    if hasAttr(c, "div_equipo1") {
      g.team1 = c.FirstChild.FirstChild.Data
    }
    if hasAttr(c, "div_equipo2") {
      g.team2 = c.FirstChild.FirstChild.Data
    }
    if hasAttr(c, "div_hora") {
      g.time = c.FirstChild.Data
    }
    if hasAttr(c, "div_campeonato") {
       g.league = c.FirstChild.FirstChild.Data
    }
  }
  games = append(games, g)
}

func parse(body io.Reader) {
  doc, err := html.Parse(body)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  visit(doc)
}

func hasAttr(node *html.Node, attrName string) bool {
  for _, a := range node.Attr {
    if strings.Contains(a.Val, attrName) {
      return true
    }
  }
  return false
}
