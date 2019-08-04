package stats

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Stats struct holds the server stats
type Stats struct {
	Count         int     `json:"total"`
	AvgRespTime   float64 `json:"average"`
	TotalRespTime float64 `json:"-"`
}

var CurrentStats Stats

// Init initializes stats with either past data or a fresh start
/*
	Input:
		- none
	Output:
		- none
*/
func Init() {
	CurrentStats = Stats{}
	GetStatsFromStore()
}

// UpdateStats updates the stats for the current server session
/*
	Input:
		- respTime - float64 - response time for the current request
	Output:
		- none
*/
func UpdateStats(respTime float64) {
	CurrentStats.Count++
	CurrentStats.TotalRespTime += respTime

	if CurrentStats.Count == 0 {
		CurrentStats.AvgRespTime = 0.0
	} else {
		CurrentStats.AvgRespTime = CurrentStats.TotalRespTime / float64(CurrentStats.Count)
	}
}

// GetCurrentStats gets the current stats to return
/*
	Input:
		- none
	Output:
		- data - map[string]interface{} - marshalled byte array of the current stats
		- err - error - error encountered when marshalling
*/
func GetCurrentStats() (data map[string]interface{}, err error) {
	d, err := json.Marshal(CurrentStats)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(d, &data)
	return
}

// GetStatsFromStore retrieves previously stored stats from a file store to build on that with a fresh instance of the server
/*
	Input:
		- none
	Output:
		- none
*/
func GetStatsFromStore() {
	file, err := ioutil.ReadFile("stats/stats.json")
	err = json.Unmarshal([]byte(file), &CurrentStats)
	if err != nil {
		log.Println("Could not get past stats...resetting stats")
		return
	}
	CurrentStats.TotalRespTime = CurrentStats.AvgRespTime * float64(CurrentStats.Count)
}

// UpdateStatsInStore store current stats toa file store
/*
	Input:
		- none
	Output:
		- none
*/
func UpdateStatsInStore() {
	file, err := json.MarshalIndent(CurrentStats, "", " ")
	if err != nil {
		log.Fatalf("COuld not update stats in store: %s", err.Error())
		return
	}
	err = ioutil.WriteFile("stats/stats.json", file, 0644)
	if err != nil {
		log.Fatalf("Could not update stats in store: %s", err.Error())
	}
}
