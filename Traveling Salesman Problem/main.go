package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"time"
)

type Matrix struct {
	Matrix [][]struct {
		DistanceInMeters    float64 `json:"distance_in_meters"`
		TravelTimeInMinutes float64 `json:"travel_time_in_minutes"`
	} `json:"matrix"`
}

type Node struct {
	id        int
	distances []float64
}

type Cycle struct {
	node Node
}

//go run main.go --file 10_cm.json
func main() {

	start := time.Now()

	if len(os.Args) < 3 {
		fmt.Println("Missing Filename")
		return
	}

	var data Matrix
	//file, _ := ioutil.ReadFile("10_cm.json")
	file, _ := ioutil.ReadFile(os.Args[2])
	_ = json.Unmarshal([]byte(file), &data)
	// Assign all nodes (each row) to a new object called 'Node'
	nodeList := assignNodes(data)

	cycle := make([]Node, 0)
	//fmt.Println("Nodes", len(nodeList))
	//last t
	cycle, totalcost := generateAlgorithm(nodeList, cycle, true)

	fmt.Println("Lenght of cycle", len(cycle))

	fmt.Println("Results:")
	for i := range cycle {
		fmt.Print("->", cycle[i].id)
	}

	fmt.Println("\nTotal Cost:", totalcost, " mts")
	duration := time.Since(start)
	fmt.Println("Took:", duration.Seconds(), " seconds")
}

// if nearestIncersion is false, it will generate a Nearest Neighbor solution; True will throw the nearest Insertion
func generateAlgorithm(nodeList []Node, cycle []Node, nearestIncersion bool) ([]Node, float64) {
	nodeListSize := len(nodeList)
	var lastNode Node
	totalCost := 0.0
	for len(cycle) < nodeListSize { //while cycle has not taken all the nodes, it will continue

		if len(cycle) == 0 { //first iteration
			var tmp Node
			tmp = nodeList[0]
			cycle = append(cycle, tmp)

		} else if len(cycle) == 1 { //second iteration
			min := math.MaxFloat64
			id := 0
			for i := 0; i < len(cycle[0].distances); i++ {
				if cycle[0].distances[i] < min && cycle[0].distances[i] != 0 {
					min = cycle[0].distances[i]
					id = i
				}
			}
			totalCost += min
			var tmp Node
			tmp = nodeList[id]
			cycle = append(cycle, tmp)
			lastNode = tmp

		} else { //i+2 iterations

			takenNode := lastNode
			min := math.MaxFloat64
			id := 0

			for i := 0; i < len(nodeList[takenNode.id].distances); i++ {

				if nodeList[takenNode.id].distances[i] < min && nodeList[takenNode.id].distances[i] != 0 && checkIfExists(cycle, i) == false {
					min = nodeList[takenNode.id].distances[i]
					id = i

				}
			}
			totalCost += min
			var tmp Node
			tmp = nodeList[id]
			cycle = append(cycle, tmp)
			lastNode = tmp

			if len(nodeList) == len(cycle) {

				break
			}

			if nearestIncersion { //if false, it will break the nearest insertion algorithim and go with the nearest neighbor

				nextNode := id
				distance := 0.0
				min = math.MaxFloat64
				dik := nodeList[takenNode.id].distances[nextNode]
				//fmt.Println("Numero tomado:", takenNode.id)
				//fmt.Println("No de veces:", len(nodeList[takenNode.id].distances))
				for j := 0; j < len(nodeList); j++ {
					if checkIfExists(cycle, j) == false {
						dkj := nodeList[nextNode].distances[j]
						dij := nodeList[takenNode.id].distances[j]
						distance = dik + dkj - dij
						if distance < min && distance != 0 {
							min = distance
							id = j
							//fmt.Println("id:", id)
						}
					}

				}
				totalCost += min
				tmp = nodeList[id]
				cycle = append(cycle, tmp)
				lastNode = tmp
			}

		}

	}
	return cycle, totalCost
}

//checkIfExists receives the cycle list of nodes, and the id of the node that needs to check. Returns false if the id has not been added
func checkIfExists(cycle []Node, id int) bool {
	result := false
	for _, node := range cycle {
		if node.id == id {
			result = true
		}
	}
	return result
}

//assignNodes receives an object of type Matrix, returns an array of elements of type Node
func assignNodes(data Matrix) []Node {
	nodeList := make([]Node, len(data.Matrix))
	for i := 0; i < len(data.Matrix); i++ { //for each didn't work, so had to use i,j
		distances := make([]float64, len(data.Matrix[i])) //takes the lenght of the distances obj of each node
		for j := 0; j < len(distances); j++ {
			distances[j] = data.Matrix[i][j].DistanceInMeters
		}
		//assign elements to new format
		var tmp Node
		tmp.id = i
		tmp.distances = distances
		nodeList[i] = tmp

	}
	//prints all the nodes received
	/*
		for i := 0; i < len(nodeList); i++ {
			fmt.Println("Nodo:", nodeList[i].id, "Distancias:", nodeList[i].distances)

		}
	*/
	return nodeList

}
