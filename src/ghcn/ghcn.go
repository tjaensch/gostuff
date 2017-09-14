package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"text/template"
	"time"
)

type IsoFields struct {
	StationId        string
	Date             string
	Lat              string
	Lon              string
	MetadataKeywords []string
}

var variablesGCMDDefinitionsMap = map[string]string{
	"PRCP": "Earth Science > Atmosphere > Precipitation > Precipitation Amount > 24 Hour Precipitation Amount",
	"SNOW": "Earth Science > Terrestrial Hydrosphere > Snow/Ice > Snow Depth",
	"SNWD": "Earth Science > Terrestrial Hydrosphere > Snow/Ice > Snow Depth",
	"TM":   "Earth Science > Atmosphere > Atmospheric Temperature > Surface Temperature > Maximum/Minimum Temperature > 24 Hour Maximum Temperature",
	"AC":   "Earth Science > Atmosphere > Clouds > Cloud Properties > Cloud Fraction",
	"AWDR": "Earth Science > Atmosphere > Atmospheric Winds > Surface Winds > Wind Direction",
	"AWND": "Earth Science > Atmosphere > Atmospheric Winds > Surface Winds > Wind Speed",
	"EVAP": "Earth Science > Atmosphere > Atmospheric Water Vapor > Water Vapor Processes > Evaporation",
	"FR":   "Earth Science > Land Surface > Frozen Ground > Seasonally Frozen Ground",
	"MDEV": "Earth Science > Atmosphere > Atmospheric Water Vapor > Water Vapor Processes > Evaporation",
	"MDPR": "Earth Science > Atmosphere > Precipitation > Precipitation Amount > 24 Hour Precipitation Amount",
	"MDSF": "Earth Science > Terrestrial Hydrosphere > Snow/Ice > Snow Depth",
	"MDT":  "Earth Science > Atmosphere > Atmospheric Temperature > Surface Temperature > Maximum/Minimum Temperature > 24 Hour Maximum Temperature",
	"PSUN": "Earth Science > Atmosphere > Atmospheric Radiation > Sunshine",
	"SN":   "Earth Science > Land Surface > Soils > Soil Temperature",
	"SX":   "Earth Science > Land Surface > Soils > Soil Temperature",
	"TAVG": "Earth Science > Atmosphere > Atmospheric Temperature > Surface Temperature > Air Temperature",
	"TOBS": "Earth Science > Atmosphere > Atmospheric Temperature > Surface Temperature > Air Temperature",
	"TSUN": "Earth Science > Atmosphere > Atmospheric Radiation > Sunshine",
	"WDF":  "Earth Science > Atmosphere > Atmospheric Winds > Surface Winds > Wind Direction",
	"WES":  "Earth Science > Atmosphere > Precipitation > Snow Water Equivalent",
	"WSF":  "Earth Science > Atmosphere > Atmospheric Winds > Surface Winds > Wind Speed",
	"WT":   "Earth Science > Atmosphere > Weather Events",
	"WV":   "Earth Science > Atmosphere > Weather Events",
}

// Generic error checking function
func checkError(reason string, err error) {
	if err != nil {
		fmt.Printf("%s: %s\n", reason, err)
		os.Exit(1)
	}
}

func downloadStationsTextFile() {
	resp, err := http.Get("https://www1.ncdc.noaa.gov/pub/data/ghcn/daily/ghcnd-stations.txt")
	checkError("getting stations file failed", err)
	defer resp.Body.Close()
	out, err := os.Create("ghcnd-stations.txt")
	checkError("write file failed", err)
	defer out.Close()
	io.Copy(out, resp.Body)
}

func readInStationsFileInfo() ([]string, map[string]string, map[string]string, map[string]string) {
	stationIds := make([]string, 0)
	latMap := make(map[string]string)
	lonMap := make(map[string]string)
	stationLongNameMap := make(map[string]string)
	f, _ := os.Open("ghcnd-stations.txt")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		stationIds = append(stationIds, line[0:11])
		latMap[line[0:11]] = line[12:20]
		lonMap[line[0:11]] = line[21:30]
		stationLongNameMap[line[0:11]] = line[38:71]
	}
	return stationIds, latMap, lonMap, stationLongNameMap
}

func getIndividualDataFile(stationId string) string {
	resp, err := http.Get("https://www1.ncdc.noaa.gov/pub/data/ghcn/daily/all/" + stationId + ".dly")
	checkError("http.Get failed", err)
	defer resp.Body.Close()
	stationData, err := ioutil.ReadAll(resp.Body)
	checkError("ioutil.ReadAll(resp.Body) failed", err)
	// Remove stationIds from data string before returning
	return strings.Replace(string(stationData), stationId, "", -1)
}

func getMetadataKeywordsForStationFile(stationId string) []string {
	metadataKeywords := make([]string, 0)
	stationData := getIndividualDataFile(stationId)

	for key, value := range variablesGCMDDefinitionsMap {
		if strings.Contains(stationData, key) {
			metadataKeywords = append(metadataKeywords, value)
		}
	}
	return metadataKeywords
}

func main() {

	downloadStationsTextFile()
	stationIds, latMap, lonMap, _ := readInStationsFileInfo()

	data := IsoFields{
		stationIds[0],
		time.Now().Local().Format("2006-01-02"),
		latMap[stationIds[0]],
		lonMap[stationIds[0]],
		getMetadataKeywordsForStationFile(stationIds[0]),
	}

	tmpl, err := template.New("test").Parse("stationId: {{.StationId}}, date: {{.Date}}, lat: {{.Lat}}, lon: {{.Lon}}, metadataKeywords: {{.MetadataKeywords}}")
	checkError("creating template failed", err)
	err = tmpl.Execute(os.Stdout, data)
	checkError("executing template failed", err)

}
