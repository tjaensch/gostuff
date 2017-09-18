package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/remeh/sizedwaitgroup"
)

type IsoFields struct {
	StationId        string
	Date             string
	Lat              string
	Lon              string
	MetadataKeywords []string
	BeginYear        string
	EndYear          string
	BeginMonth       string
	EndMonth         string
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

// Create file directories
func prepDirs() {
	os.Mkdir("./isolite", 0777)
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

func readInIndividualDataFileInfo(stationId string) ([]string, []string) {
	years := make([]string, 1)
	months := make([]string, 1)
	resp, err := http.Get("https://www1.ncdc.noaa.gov/pub/data/ghcn/daily/all/" + stationId + ".dly")
	checkError("getting individual data file failed", err)
	defer resp.Body.Close()
	out, err := os.Create(stationId + ".txt")
	checkError("write file failed", err)
	defer out.Close()
	io.Copy(out, resp.Body)
	f, _ := os.Open(stationId + ".txt")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		years = append(years, line[11:15])
		months = append(months, line[15:17])
	}
	os.Remove(stationId + ".txt")
    return years, months
}

func readInStationsFileInfo() ([]string, map[string]string, map[string]string) {
	stationIds := make([]string, 0)
	latMap := make(map[string]string, 1)
	lonMap := make(map[string]string, 1)
	f, _ := os.Open("ghcnd-stations.txt")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		stationIds = append(stationIds, line[0:11])
		latMap[line[0:11]] = line[12:20]
		lonMap[line[0:11]] = line[21:30]
	}
	os.Remove("ghcnd-stations.txt")
	return stationIds, latMap, lonMap
}

func getIndividualDataFileAsString(stationId string) string {
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
	stationData := getIndividualDataFileAsString(stationId)

	for key, value := range variablesGCMDDefinitionsMap {
		if strings.Contains(stationData, key) {
			metadataKeywords = append(metadataKeywords, value)
		}
	}
	return metadataKeywords
}

func processStationId(stationId string, latMap map[string]string, lonMap map[string]string) {
	years, months := readInIndividualDataFileInfo(stationId)
	data := IsoFields{
		stationId,
		time.Now().Local().Format("2006-01-02"),
		latMap[stationId],
		lonMap[stationId],
		getMetadataKeywordsForStationFile(stationId),
		getIndividualDataFileAsString(stationId)[0:4],
		years[len(years)-1],
		getIndividualDataFileAsString(stationId)[4:6],
		months[len(months)-1],
	}

	tmpl, err := template.ParseFiles("templates/isolite.tmpl")
	checkError("creating template failed", err)
	f, err := os.Create("isolite/ghcn-daily_v3.22." + time.Now().Local().Format("2006-01-02") + "_" + stationId + ".xml")
	checkError("create file failed", err)
	defer f.Close()
	err = tmpl.ExecuteTemplate(f, "isolite", data)
	checkError("executing template failed", err)
	fmt.Println(stationId + " successfully written to isolite directory")
}

func countOutputFiles() {
	files, err := ioutil.ReadDir("./isolite")
	checkError("counting output files failed, program exiting", err)
	log.Printf("%d files written to isolite directory\n", len(files))
}

func main() {
	log.Printf("Working digging up files...")
	t0 := time.Now()

	downloadStationsTextFile()
	prepDirs()
	stationIds, latMap, lonMap := readInStationsFileInfo()

	swg := sizedwaitgroup.New(200)
	for _, stationId := range stationIds {
		swg.Add()
		go func() {
			defer swg.Done()
			processStationId(stationId, latMap, lonMap)
		}()
	}

	swg.Wait()

	countOutputFiles()
	t1 := time.Now()
	log.Printf("The program took %v minutes to run.\n", t1.Sub(t0).Minutes())
}
