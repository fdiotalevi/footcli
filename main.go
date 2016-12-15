package main

import  (
	"flag"
        )


func main() {
  country := flag.String("c", "uk", "Country: \"uk\", \"it\", \"es\" are supported")
  week := flag.String("w", "current", "\"current\" for current week, \"next\" for next week")
  flag.Parse()

  Pln("\n   footcli - Football on the Command Line")
  Pln("-----------------------------------------------")
  GetTvSchedule(*country, *week)
}

