package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Response struct {
	XMLName xml.Name `xml:"FindPlaceFromTextResponse"`
	Status  string   `xml:"status"`
	Candidates Candidates  `xml:"candidates"`
}
type Candidates struct {
	XMLName  xml.Name `xml:"candidates"`
	Geometry Geometry `xml:"geometry"`
}
type Geometry struct {
	XMLName  xml.Name `xml:"geometry"`
	Location Location `xml:"location"`
}
type Location struct {
	XMLName   xml.Name `xml:"location"`
	Latitude  float64 `xml:"lat"`
	Longitude float64 `xml:"lng"`
}
// XML follows: result > geometry > location > lat,lng

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	// creates scanner variable

	fmt.Println("What is the name of the park you would like to visit?")
	scanner.Scan()
	var parkname = scanner.Text()
	fmt.Println("The park name is: " + parkname)
	var parknameAdj = strings.ReplaceAll(parkname," ","%20")
	//prompts user for park name

	urlEmpty := "https://maps.googleapis.com/maps/api/place/findplacefromtext/xml?input=[]&inputtype=textquery&fields=geometry(location)&key=[]"
	url := strings.ReplaceAll(urlEmpty,"[]",parknameAdj)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method,url,nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var LatLong = Response{}
	xml.Unmarshal(body, &LatLong)

	fmt.Println(string(body))
	lati := strconv.FormatFloat(LatLong.Candidates.Geometry.Location.Latitude,'E',-1,64)
	long :=  strconv.FormatFloat(LatLong.Candidates.Geometry.Location.Longitude,'E',-1,64)
	fmt.Println(lati)

	weatherUrlEmpty := "https://api.weather.gov/points/{lat},{lon}"
	weatherUrl := strings.ReplaceAll(weatherUrlEmpty,"{lat}",lati)
	weatherUrl = strings.ReplaceAll(weatherUrlEmpty,"{lat}",long)
}
