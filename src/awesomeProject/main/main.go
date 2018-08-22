package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"bufio"
	"awesomeProject/io"
	"strings"
	"awesomeProject/repositories"
	"awesomeProject/minerals"
)

//Constants
const ResourcesFolder = "resources/"
const TxtResourcesFolder = ResourcesFolder + "txt/"

func readFile(fileName string) {
	fmt.Println("asd")
	fileContent, err := ioutil.ReadFile( TxtResourcesFolder + fileName)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("File Content: %s", fileContent)
}

func writeFile(fileName string, content string) {
	ioutil.WriteFile(fileName, []byte(content), os.ModeAppend)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	budget := 100.0

	miners := repositories.MinersRepository{}
	var minedMinerals []minerals.Mineral

	inputLine := io.ReadLine(*reader)

	for inputLine != "Quit" {
		if(strings.HasPrefix(inputLine, "Miner")) {
			budget -= 25
			miners.AddMiner();
		} else if(strings.HasPrefix(inputLine, "Status")) {
			fmt.Printf("Current Budget: %.2f\r\n", budget);
			fmt.Printf("Current Miners: %d\r\n", miners.Count());
			fmt.Printf("Mined Minerals: %s\r\n", minedMinerals);
		} else if(strings.HasPrefix(inputLine, "Harvest")) {
			budget -= float64(miners.Count()) * 5;

			for _, mineral := range miners.Harvest() {
				minedMinerals =
				append(minedMinerals, mineral)
			}
		} else if(strings.HasPrefix(inputLine, "Sell")) {
			for _, mineral := range minedMinerals {
				budget += mineral.GetPrice()
			}

			minedMinerals = minedMinerals[:0]
		}

		inputLine = strings.Trim(io.ReadLine(*reader), "\n")
	}
}