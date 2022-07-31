package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

const scanNumber = 3
const delay = 5

func main() {
	printsIntro()
	for {
		printsOptions()
		selectedOption := readsOption()

		switch selectedOption {
		case 1:
			beginsTracking()
		case 2:
			printsLog()
		case 0:
			fmt.Println("Leaving...")
			os.Exit(0)
		default:
			fmt.Println("Option not recognized...")
			os.Exit(-1)
		}
	}
}

func printsIntro() {
	var userName string = "Matheus"
	var version float32 = 1.1
	fmt.Println("Hello,", userName)
	fmt.Println("Version:", version)
}

func printsOptions() {
	fmt.Println("1- Begin tracking;")
	fmt.Println("2- Show logs;")
	fmt.Println("0- Exit;")
}

func readsOption() int {
	var selectedOption int = 0
	fmt.Scan(&selectedOption)
	fmt.Println("Selected option is", selectedOption)
	return selectedOption
}

func beginsTracking() {
	sites := getSitesFromFile()
	fmt.Println()
	fmt.Println("Tracking...")

	for i := 0; i < scanNumber; i++ {
		for _, site := range sites {
			response, err := http.Get(site)
			if err != nil {
				fmt.Println("Error trying to access", site, ":", err)
			}
			fmt.Println("Site", site, "returned", response.StatusCode)
			savesOnLog(site, response.StatusCode)
		}
		time.Sleep(delay * time.Second)
		fmt.Println()
	}
	fmt.Println()
}

func getSitesFromFile() []string {
	var sites []string
	//file, err := ioutil.ReadFile("sites.txt")
	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Error trying to open the file:", err)
	}

	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error trying to read the file:", err)
		}
		sites = append(sites, string(line))
	}
	file.Close()
	return sites
}

func savesOnLog(site string, statusCode int) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if (err) != nil {
		fmt.Println(err)
	}

	file.WriteString("[" + time.Now().Format("2006-01-02 15:04:05") + "] Site " + site + " returned status code " + strconv.Itoa(statusCode) + "\n")

	file.Close()
}

func printsLog() {
	fmt.Println("Getting logs...")
	file, err := ioutil.ReadFile("log.txt")
	if (err) != nil {
		fmt.Println(err)
	}
	fmt.Println(string(file))
}
