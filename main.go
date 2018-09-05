package main

import (
	"flag"
	"net/http"
	"log"
	"fmt"
)

func main()  {
	addr:=flag.String("addr",":6789","listen address")
	dir:=flag.String("dir",".","directory")
	help:=flag.Bool("help",false,"usage")

	flag.Parse()

	if *help{
		flag.Usage()
	}

	h := http.FileServer(http.Dir(*dir))
	fmt.Println("listening on "+*addr)
	err := http.ListenAndServe(*addr, h)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
