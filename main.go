package main

import (
	"encoding/csv"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	// Create csv file to write results to
	csvFile, err := os.Create("status.csv")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer csvFile.Close()

	w := csv.NewWriter(csvFile)
	defer w.Flush()

	// Open the csv file to read our links
	csvfile, err := os.Open("links.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the csv file
	r := csv.NewReader(csvfile)

	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		link := record[0]
		status := checkLink(link)

		//Write the lniks which are down
		if status == "404" {
			row := []string{link, status}
			//fmt.Println(link + " - " + status)
			if err := w.Write(row); err != nil {
				log.Fatalln("error writing record to file", err)
			}
		}
	}

}
func checkLink(link string) string {

	resp, err := http.Get(link)
	if err != nil {
		message := "link is down"
		return message
	}

	if resp.StatusCode == 404 {
		message := "404"
		return message
	}

	message := "link is up"
	return message

}
