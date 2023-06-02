package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoringValue = 3
const timeDelay = 2

func main() {
	intro()
	// registerLog("fake site", false)

	for {
		showMenu()

		command := inputCommand()

		switch command {
		case 1:
			monitoring()
		case 2:
			logs()
			printLogs()
		case 0:
			programmeExit()
			os.Exit(0)
		default:
			notRecognised()
			os.Exit(-1)
		}
	}
}

func intro() {
	name := "Tom"
	version := "1.20.4"
	age := 14
	fmt.Println("Hello", name, ". Welcome to the Go language!")
	fmt.Println("The current Go version is", version)
	fmt.Println("Go language exists for", age, "years")
}

func inputCommand() int {
	var command int
	fmt.Scan(&command)
	fmt.Println("The var address is", &command)
	fmt.Println("The command chosen was", command)

	return command
}

func showMenu() {
	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - Show system logs")
	fmt.Println("0 - Exit programme")
}

func monitoring() {
	fmt.Println("Monitoring...")
	sites := readFile()

	for x := 0; x < monitoringValue; x++ {
		for x, site := range sites {
			fmt.Println("Testing site", x, ":", site)
			testSite(site)
		}
		time.Sleep(timeDelay * time.Second)
		fmt.Println("")
	}

	site := "https://alura.com.br"
	http.Get(site)
}

func logs() {
	fmt.Println("Showing logs ...")
}

func programmeExit() {
	fmt.Println("Leaving programme")
}

func notRecognised() {
	fmt.Println("Command not recognised!")
}

func testSite(site string) {
	response, error := http.Get(site)

	if error != nil {
		fmt.Println("An error ocurred. See further:", error)
	}

	if response.StatusCode == 200 {
		fmt.Println("Site", site, "has been loaded sucessfully!")
		registerLog(site, true)
	} else {
		fmt.Println("Site:", site, "it has problems. Status Code:", response.StatusCode)
	}
}

func readFile() []string {
	var sites []string

	file, error := os.Open("sites.txt")

	if error != nil {
		fmt.Println("An error ocurred. See further:", error)
	}

	reader := bufio.NewReader(file)

	for {
		line, error := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		fmt.Println(line)

		sites = append(sites, line)

		if error == io.EOF {
			break
		}

	}

	file.Close()
	return sites
}

func registerLog(site string, status bool) {
	file, error := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if error != nil {
		fmt.Println(error)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + "-" + site + "- online: " + strconv.FormatBool(status) + "\n")
	file.Close()
}

func printLogs() {
	file, error := ioutil.ReadFile("log.txt")

	if error != nil {
		fmt.Println(error)
	}

	fmt.Println(string(file))
}
