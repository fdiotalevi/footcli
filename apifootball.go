package main

import (
	"github.com/icedream/go-footballdata"
	"net/http"
      )

func FootballApi() {
  // Create client
  client := footballdata.NewClient(http.DefaultClient)

//  client.AuthToken = "<insert your api token here>"

  seasons, err := client.SoccerSeasons().Do()
  if err != nil {
  }
  for _, season := range seasons {
    Pln(season.Caption)
  }
}


