package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Airport struct {
	CSVRow      string  `json:"_csv_row"`
	AirportName string  `json:"airport_name"`
	City        string  `json:"city"`
	Country     string  `json:"country"`
	IATA        string  `json:"iata"`
	ICAO        string  `json:"icao"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Elevation   int     `json:"elevation"`
	UTCOffset   float64 `json:"utc_offset"`
	Class       string  `json:"_class"`
	Timezone    string  `json:"timezone"`
}

var (
	commit  = "0000000"
	version = "0.0.0-pre0+0000000"
	date    = "0000-00-00T00:00:00+0000"
)

func showVersion() {
	fmt.Printf("airports-csv-to-json v%s (%s) %s\n", version, commit[:7], date)
}

func showHelp() {
	showVersion()
	fmt.Println()
	fmt.Println("USAGE")
	fmt.Println("    airports-csv-to-json --csv <csv>")
	fmt.Println()
	fmt.Println("EXAMPLE")
	fmt.Println("    airports-csv-to-json --csv ./airport-data/airports.csv")
	fmt.Println()
}

func main() {
	if len(os.Args) > 1 {
		if "version" == strings.TrimLeft(os.Args[1], "-") || "-V" == os.Args[1] {
			showVersion()
			os.Exit(0)
			return
		}
		if "help" == strings.TrimLeft(os.Args[1], "-") {
			showHelp()
			os.Exit(0)
			return
		}
	}

	if len(os.Args) == 1 {
		showHelp()
		os.Exit(1)
		return
	}

	var csvFile string
	flag.StringVar(&csvFile, "csv", "", "path to CSV file")
	flag.Parse()

	if csvFile == "" {
		showHelp()
		os.Exit(1)
		return
	}

	file, err := os.Open(csvFile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	errlog, err := os.OpenFile("errors.log", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	os.MkdirAll("icao", 0o750)

	var airports []Airport
	recordsLen := len(records)
	for i, record := range records {
		icao := record[5]

		latitude := 0.0
		{
			latitudeStr := record[6]
			latitude, err = strconv.ParseFloat(latitudeStr, 64)
			if err != nil {
				fmt.Printf("parse error on line %d (%s): latitude: %s\n", i, icao, latitudeStr)
			}
		}

		longitude := 0.0
		{
			longitudeStr := record[7]
			longitude, err = strconv.ParseFloat(longitudeStr, 64)
			if err != nil {
				fmt.Printf("parse error on line %d (%s): longitude: %s\n", i, icao, longitudeStr)
			}
		}

		elevation := 0
		{
			elevationStr := record[8]
			elevation, err = strconv.Atoi(elevationStr)
			if err != nil {
				fmt.Printf("parse error on line %d (%s): elevation: %s\n", i, icao, elevationStr)
			}
		}

		utcOffset := 0.0
		{
			utcOffsetStr := record[9]
			utcOffset, err = strconv.ParseFloat(utcOffsetStr, 64)
			if err != nil {
				fmt.Printf("parse error on line %d (%s): utc_offset: %s\n", i, icao, utcOffsetStr)
			}
		}

		airport := Airport{
			CSVRow:      record[0],
			AirportName: record[1],
			City:        record[2],
			Country:     record[3],
			IATA:        record[4],
			ICAO:        icao,
			Latitude:    latitude,
			Longitude:   longitude,
			Elevation:   elevation,
			UTCOffset:   utcOffset,
			Class:       record[10],
			Timezone:    record[11],
		}
		if len(icao) == 0 {
			line := fmt.Sprintf("bad input for line %d: %#v\n", i, airport)
			errlog.WriteString(line)
			fmt.Fprintf(os.Stderr, line)
			continue
		}

		jsonData, err := json.MarshalIndent(airport, "", "    ")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		filename := fmt.Sprintf("icao/%s.json", airport.ICAO)
		filename = strings.ToLower(filename)
		if err := os.WriteFile(filename, jsonData, 0o640); err != nil {
			panic(err)
		}
		fmt.Printf("    wrote %d of %d: %s\n", i, recordsLen, filename)

		airports = append(airports, airport)
	}

	{
		jsonData, err := json.MarshalIndent(airports, "", "    ")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		filename := "icao.json"
		if err := os.WriteFile(filename, jsonData, 0o640); err != nil {
			panic(err)
		}
		fmt.Printf("wrote %s\n", filename)
	}
}
