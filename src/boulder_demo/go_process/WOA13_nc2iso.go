package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type ncmlAdditions struct {
	ncFileName    string
	fileSize      int
	datapath      string
	browsegraphic string
}

//ncmlAdditions struct method that adds information missing from .ncml to same
func (p *ncmlAdditions) encode(w io.Writer) (int, error) {
	return fmt.Fprintf(w, "<title>%s</title>\n<filesize>%d</filesize>\n<path>%s</path>\n<browsegraphic>%s</browsegraphic>\n</netcdf>",
		p.ncFileName, p.fileSize, p.datapath, p.browsegraphic)
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
	os.Mkdir("./ncml", 0777)
	os.Mkdir("./xml_output", 0777)
}

var (
	ncFilePath string = "/nodc/web/data.nodc/htdocs/nodc/archive/data/0114815/public"
	xslFile    string = "/nodc/users/tjaensch/xsl.git/boulder_demo/go_process/XSL/ncml2iso_modified_from_UnidataDD2MI_demo_WOA_Thomas_edits.xsl"
	//WOA13 collection metadata template file
	isocofile     string   = "/nodc/web/data.nodc/htdocs/nodc/archive/metadata/approved/iso/0114815.xml"
	ncFiles       []string = findNcFiles("./netcdf")
	ncFileName    string
	fileSize      int
	dataPath      string
	browsegraphic string
	additions     = ncmlAdditions{ncFileName, fileSize, dataPath, browsegraphic}
)

func main() {
	log.Printf("Working digging up files...")
	t0 := time.Now()
	prepDirs()

	var wg sync.WaitGroup

	// Start goroutine for each files segment of ncFiles slice
	fileSegments := getFileSegments()
	for _, fileSegment := range fileSegments {
		wg.Add(1)
		go func(fileSegment []string) {
			defer wg.Done()
			nc2iso(fileSegment)
		}(fileSegment)
	}

	// Wait until all goroutines finish
	wg.Wait()

	countOutputFiles()
	t1 := time.Now()
	log.Printf("The program took %v seconds to run.\n", t1.Sub(t0).Seconds())
}

// Create fileSegments slice of slice for concurrent processing
func getFileSegments() [][]string {
	// Create a slice of ncFiles
	fileSegments := make([][]string, 0)
	// Determine the length of the subslices based on amount of files and how many files can be open at the same time in PuTTY
	increaseRate := 1
	// Add subslices to fileSegments slice
	for i := 0; i < len(ncFiles)-increaseRate; i += increaseRate {
		fileSegments = append(fileSegments, ncFiles[i:i+increaseRate])
	}
	fileSegments = append(fileSegments, ncFiles[len(ncFiles)-increaseRate:])
	return fileSegments
}

func nc2iso(ncFiles []string) {
	for _, ncFile := range ncFiles {
		var (
			ncFileName    string = getFileName(ncFile)
			fileSize      int    = getFileSize(ncFile)
			dataPath      string = getFilePath(ncFile)
			browsegraphic string = getBrowseGraphicLink(ncFile)
			additions            = ncmlAdditions{ncFileName, fileSize, dataPath, browsegraphic}
		)
		ncdump(ncFile)
		appendToNcml(ncFile, additions)
		xsltprocToISO(ncFile, xslFile)
		addCollectionMetadata(ncFile)
	}
}

func findNcFiles(string) []string {
	var ncFiles []string
	var files []byte
	var err error
	cmdName := "find"
	cmdArgs := []string{ncFilePath, "-type", "f", "-name", "*.nc"}
	if files, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Printf("Something went wrong with finding .nc files in source directory, program exiting.", err)
		os.Exit(1)
	}

	for _, rune := range bytes.Split(files, []byte{'\n'}) {
		ncFiles = append(ncFiles, string(rune))
	}
	log.Printf("%d files found in source directory", len(ncFiles)-1)
	//Last element is blank that's why -1
	return ncFiles[:len(ncFiles)-1]
}

func ncdump(ncFile string) {
	var ncml []byte
	var err error
	cmdName := "ncdump"
	cmdArgs := []string{"-x", ncFile}
	if ncml, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Printf("Something went wrong with ncdump, program exiting.", err)
		os.Exit(1)
	}
	// Write ncdump conversion to file
	err = ioutil.WriteFile("./ncml/"+getFileName(ncFile)+".ncml", ncml, 0644)
	checkError("write ncml file failed, program exiting", err)
}

// Get just the file name without file extension
func getFileName(ncFile string) (ncFileName string) {
	ncFileBasePath := filepath.Base(ncFile)
	ncFileName = strings.TrimSuffix(ncFileBasePath, ".nc")
	return
}

// Get .nc file size in MB
func getFileSize(ncFile string) (fileSize int) {
	file, err := os.Open(ncFile)
	checkError("Open .nc file failed", err)
	defer file.Close()
	stat, err := file.Stat()
	checkError("Get file size failed", err)
	return (int)(stat.Size() / 1024 / 1024)
}

// Get .nc file path on WAF
func getFilePath(ncFile string) string {
	trimmedPath := ncFile[27:len(ncFile)]
	return strings.Replace(trimmedPath, filepath.Base(ncFile), "", -1)
}

// Append struct fields to .ncml file so XSLT can grab the added values
func appendToNcml(ncFile string, additions ncmlAdditions) {
	// Mess with <netcdf> tag for XSL wo work later on
	input, err := ioutil.ReadFile("./ncml/" + getFileName(ncFile) + ".ncml")
	checkError("read ncml file failed", err)
	lines := strings.Split(string(input), "\n")
	for i, line := range lines {
		if strings.Contains(line, "<netcdf") {
			lines[i] = "<netcdf>"
		}
		if strings.Contains(line, "</netcdf>") {
			lines[i] = ""
		}
		output := strings.Join(lines, "\n")
		err = ioutil.WriteFile("./ncml/" + getFileName(ncFile) + ".ncml", []byte(output), 0644)
		checkError("write ncml file failed", err)
	}
	// Append extra lines to .ncml
	file, err := os.OpenFile("./ncml/" + getFileName(ncFile) + ".ncml", os.O_RDWR|os.O_APPEND, 0)
	checkError("open ncml file failed", err)
	if _, err := additions.encode(file); err != nil {
		fmt.Printf("append to ncml file failed: %s\n", err)
		os.Exit(1)
	}
	defer file.Close()
}

func xsltprocToISO(ncFile string, xslFile string) {
	var isoXML []byte
	var err error
	cmdName := "xsltproc"
	//Convert to ISO
	cmdArgs := []string{xslFile, "./ncml/" + getFileName(ncFile) + ".ncml"}
	if isoXML, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Printf("Something went wrong with the XSLT conversion, program exiting.", err)
		os.Exit(1)
	}
	// Write xsltproc conversion to file
	err = ioutil.WriteFile("./xml_output/" + getFileName(ncFile) + ".xml", isoXML, 0644)
	checkError("write xml file failed", err)
}

func addCollectionMetadata(ncFile string) {
	var isoXML []byte
	var err error
	cmdName := "xsltproc"
	cmdArgs := []string{"--stringparam", "collFile", isocofile, "/nodc/users/tjaensch/xsl.git/boulder_demo/go_process/XSL/granule.xsl", "./xml_output/" + getFileName(ncFile) + ".xml"}
	if isoXML, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Printf("Something went wrong with the collection metadata addition, program exiting.", err)
		os.Exit(1)
	}
	err = ioutil.WriteFile("./xml_output/" + getFileName(ncFile) + ".xml", isoXML, 0644)
	if err == nil {
		fmt.Printf("%s.xml successfully written to xml output directory\n", getFileName(ncFile))
	}
	checkError("write xml file failed, program exiting", err)
}

func getBrowseGraphicLink(ncFile string) string {
	var (
		graphictype   string
		graphictime   string
		graphicmonth  string
		graphicdegree string
		graphicid     string
	)

	//graphictype
	if strings.Contains(getFileName(ncFile), "_s") {
		graphictype = "salinity"
	} else if strings.Contains(getFileName(ncFile), "_t") {
		graphictype = "temperature"
	} else if strings.Contains(getFileName(ncFile), "_A") {
		graphictype = "AOU"
	} else if strings.Contains(getFileName(ncFile), "_i") {
		graphictype = "silicate"
	} else if strings.Contains(getFileName(ncFile), "_n") {
		graphictype = "nitrate"
	} else if strings.Contains(getFileName(ncFile), "_O") {
		graphictype = "o2sat"
	} else if strings.Contains(getFileName(ncFile), "_o") {
		graphictype = "oxygen"
	} else if strings.Contains(getFileName(ncFile), "_p") {
		graphictype = "phosphate"
	} else {
		graphictype = ""
	}

	//graphictime
	if getFileName(ncFile)[len(getFileName(ncFile))-5:len(getFileName(ncFile))-3] == "00" {
		graphictime = "annual"
	} else if getFileName(ncFile)[len(getFileName(ncFile))-5:len(getFileName(ncFile))-3] > "12" {
		graphictime = "seasonal"
	} else {
		graphictime = "monthly"
	}

	//graphicmonth
	if getFileName(ncFile)[len(getFileName(ncFile))-5:len(getFileName(ncFile))-3] == "01" {
		graphicmonth = "0.5"
	} else if getFileName(ncFile)[len(getFileName(ncFile))-5:len(getFileName(ncFile))-3] == "02" {
		graphicmonth = "1.5"
	} else if getFileName(ncFile)[len(getFileName(ncFile))-5:len(getFileName(ncFile))-3] == "03" {
		graphicmonth = "2.5"
	} else if getFileName(ncFile)[len(getFileName(ncFile))-5:len(getFileName(ncFile))-3] == "04" {
		graphicmonth = "3.5"
	} else if getFileName(ncFile)[len(getFileName(ncFile))-5:len(getFileName(ncFile))-3] == "05" {
		graphicmonth = "4.5"
	} else if getFileName(ncFile)[len(getFileName(ncFile))-5:len(getFileName(ncFile))-3] == "06" {
		graphicmonth = "5.5"
	} else if getFileName(ncFile)[len(getFileName(ncFile))-5:len(getFileName(ncFile))-3] == "07" {
		graphicmonth = "6.5"
	} else if getFileName(ncFile)[len(getFileName(ncFile))-5:len(getFileName(ncFile))-3] == "08" {
		graphicmonth = "7.5"
	} else if getFileName(ncFile)[len(getFileName(ncFile))-5:len(getFileName(ncFile))-3] == "09" {
		graphicmonth = "8.5"
	} else if getFileName(ncFile)[len(getFileName(ncFile))-5:len(getFileName(ncFile))-3] == "10" {
		graphicmonth = "9.5"
	} else if getFileName(ncFile)[len(getFileName(ncFile))-5:len(getFileName(ncFile))-3] == "11" {
		graphicmonth = "10.5"
	} else {
		graphicmonth = "11.5"
	}

	//graphicdegree
	if strings.Contains(getFileName(ncFile), "_01") {
		graphicdegree = "1degree"
	} else if strings.Contains(getFileName(ncFile), "_04") {
		graphicdegree = "quarterdegree"
	} else {
		graphicdegree = "5degree"
	}

	//graphicid
	switch {
	case graphictype == "salinity":
		s_ids := []string{"s_mn", "s_dd", "s_sd", "s_se"}
		graphicid = s_ids[rand.Intn(len(s_ids))]
	case graphictype == "temperature":
		t_ids := []string{"t_mn", "t_dd", "t_sd", "t_se"}
		graphicid = t_ids[rand.Intn(len(t_ids))]
	case graphictype == "AOU":
		A_ids := []string{"A_mn", "A_dd", "A_sd", "A_se"}
		graphicid = A_ids[rand.Intn(len(A_ids))]
	case graphictype == "silicate":
		i_ids := []string{"i_mn", "i_dd", "i_sd", "i_se"}
		graphicid = i_ids[rand.Intn(len(i_ids))]
	case graphictype == "nitrate":
		n_ids := []string{"n_mn", "n_dd", "n_sd", "n_se"}
		graphicid = n_ids[rand.Intn(len(n_ids))]
	case graphictype == "o2sat":
		O2sat_ids := []string{"O_mn", "O_dd", "O_sd", "O_se"}
		graphicid = O2sat_ids[rand.Intn(len(O2sat_ids))]
	case graphictype == "oxygen":
		o_ids := []string{"o_mn", "o_dd", "o_sd", "o_se"}
		graphicid = o_ids[rand.Intn(len(o_ids))]
	case graphictype == "phosphate":
		p_ids := []string{"p_mn", "p_dd", "p_sd", "p_se"}
		graphicid = p_ids[rand.Intn(len(p_ids))]
	}

	if graphictype == "" {
		return ""
	} else {
		return "http://data.nodc.noaa.gov/las/ProductServer.do?xml=%3C%3Fxml+version%3D%221.0%22%3F%3E%3ClasRequest+href%3D%22file%3Alas.xml%22%3E%3Clink+match%3D%22%2Flasdata%2Foperations%2Foperation%5B%40ID%3D%27Plot_2D_XY_zoom%27%5D%22%3E%3C%2Flink%3E%3Cproperties%3E%3Cferret%3E%3Cview%3Exy%3C%2Fview%3E%3Cland_type%3Edefault%3C%2Fland_type%3E%3Cset_aspect%3Edefault%3C%2Fset_aspect%3E%3Cmark_grid%3Eno%3C%2Fmark_grid%3E%3Ccontour_levels%3E%3C%2Fcontour_levels%3E%3Cfill_levels%3E%3C%2Ffill_levels%3E%3Ccontour_style%3Edefault%3C%2Fcontour_style%3E%3Cpalette%3Edefault%3C%2Fpalette%3E%3Cdeg_min_sec%3Edefault%3C%2Fdeg_min_sec%3E%3Cmargins%3Edefault%3C%2Fmargins%3E%3Cuse_graticules%3Edefault%3C%2Fuse_graticules%3E%3Csize%3E0.5%3C%2Fsize%3E%3Cimage_format%3Edefault%3C%2Fimage_format%3E%3Cinterpolate_data%3Efalse%3C%2Finterpolate_data%3E%3Cexpression%3E%3C%2Fexpression%3E%3C%2Fferret%3E%3C%2Fproperties%3E%3Cargs%3E%3Clink+match%3D%22%2Flasdata%2Fdatasets%2Fid-woa13-" + graphictype + "-" + graphictime + "-" + graphicdegree + "%2Fvariables%2F" + graphicid + "-id-woa13-" + graphictype + "-" + graphictime + "-" + graphicdegree + "%22%3E%3C%2Flink%3E%3Cregion%3E%3Cpoint+type%3D%22t%22+v%3D%22" + graphicmonth + "%22%3E%3C%2Fpoint%3E%3Cpoint+type%3D%22z%22+v%3D%220%22%3E%3C%2Fpoint%3E%3Crange+type%3D%22y%22+low%3D%22-87.5%22+high%3D%2287.5%22%3E%3C%2Frange%3E%3Crange+type%3D%22x%22+low%3D%22-177.5%22+high%3D%22177.5%22%3E%3C%2Frange%3E%3C%2Fregion%3E%3C%2Fargs%3E%3C%2FlasRequest%3E&amp;stream=true&amp;stream_ID=plot_image"
	}
}

func countOutputFiles() {
	files, err := ioutil.ReadDir("./xml_output")
	checkError("counting output files failed, program exiting", err)
	log.Printf("%d files written to xml_output directory\n", len(files))
}
