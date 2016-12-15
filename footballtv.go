package main

import  (
	"fmt"
        "io"
	"strings"
        "os"
        "golang.org/x/net/html")


type Game struct {
  date, team1, team2, time, tv, league string
}

var games []Game
var schedules map[string]string
var currentDate string

func GetTvSchedule(country string, week string) {
  games = make([]Game, 0)
  schedules = make(map[string]string)
  schedules["uk"] = "http://www.footballonuktv.com/"
  schedules["it"] = "http://www.calciointv.com/"
  schedules["es"] = "http://www.futbolenlatele.com/"
  
  if week != "current" && week != "next" {
    fmt.Printf("The option -w=%s is not valid, will display this week fixtures\n", week)
  }

  weekOption := ""
  if week == "next" {
    weekOption = "/index2.php"
  }

  mainUrl := schedules[country]

  if mainUrl == "" {
    Pln("There's no service for TV listing for country " + country)
    os.Exit(1)
  }

  GetAndProcess(mainUrl+weekOption, parse) 
  for _, g := range games {
    formatGame(g)
  }
}

func formatGame(g Game) {
  fmt.Printf("%s\t%s\t %s\t\"%#s - %#s\" on %s\n", g.date, g.time, g.league, g.team1, g.team2, g.tv)	
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
    Pln(err.Error())
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
