package main

import (
	"flag"
	"log"
	"os"

	"github.com/jasontconnell/sqlcsv/conf"
	"github.com/jasontconnell/sqlcsv/process"
)

func main() {
	configFilename := flag.String("c", "config.json", "config filename")
	query := flag.String("q", "", "query to run")
	out := flag.String("out", "out.csv", "output filename")
	headers := flag.Bool("headers", false, "include headers in output")
	flag.Parse()

	cfg, err := conf.LoadConfig(*configFilename)
	if err != nil {
		log.Fatal(err)
	}
	if cfg.ConnectionString == "" {
		log.Fatal("no connection string defined")
	}

	tbl, err := process.Read(cfg.ConnectionString, *query)
	if err != nil {
		log.Fatal(err)
	}

	output := os.Stdout
	if *out != "" {
		log.Println("writing output to ", *out)
		output, err = os.OpenFile(*out, os.O_CREATE|os.O_TRUNC, os.ModePerm)
		if err != nil {
			log.Fatalf("can't open file for writing, %s. %v", *out, err)
		}
		defer output.Close()
	}

	process.Write(output, *headers, tbl)
}
