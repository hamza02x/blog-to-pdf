package main

import (
	"fmt"
	"unicode/utf8"
	"io/ioutil"
	"strings"
	"os"
	"net/http"
	"encoding/json"
	"regexp"
	"log"
)

func getURLContent(urlStr string) []byte {

	// fmt.Printf("HTML code of %s ...\n", urlStr)

	// Create HTTP client with timeout
	client := &http.Client{}

	// Create and modify HTTP request before sending
	request, err := http.NewRequest("GET", urlStr, nil)

	if err != nil {
		pp("Error request, err := http.NewRequest: " + err.Error())
	}

	request.Header.Set("User-Agent", ConstUserAgent)

	// Make request
	response, err := client.Do(request)

	if err != nil {
		p("Error response, err := client.Do: " + err.Error())
		pp("Please run the program again!")
	}

	htmlBytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		pp("Error htmlBytes, err := ioutil.ReadAll: " + err.Error())
	}

	response.Body.Close()
	client.CloseIdleConnections()

	return htmlBytes
}

func getFileContents(filePath string) []byte {

	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Error reading file: " + filePath)
	}

	b, err := ioutil.ReadAll(file)

	if err != nil {
		fmt.Println("Error ioutil.ReadAll: " + filePath)
	}
	file.Close()
	return b
}
func fileExists(filename string) bool {

	info, err := os.Stat(filename)

	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
func pathExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func AZ_AND_NUMBER_ONLY(urlStr string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	return reg.ReplaceAllString(urlStr, "")
}

func checkDomain(name string) error {

	switch {
	case len(name) == 0:
		return nil // an empty domain name will result in a cookie without a domain restriction
	case len(name) > 255:
		return fmt.Errorf("cookie domain: name length is %d, can't exceed 255", len(name))
	}
	var l int
	for i := 0; i < len(name); i++ {
		b := name[i]
		if b == '.' {
			// check domain labels validity
			switch {
			case i == l:
				return fmt.Errorf("cookie domain: invalid character '%c' at offset %d: label can't begin with a period", b, i)
			case i-l > 63:
				return fmt.Errorf("cookie domain: byte length of label '%s' is %d, can't exceed 63", name[l:i], i-l)
			case name[l] == '-':
				return fmt.Errorf("cookie domain: label '%s' at offset %d begins with a hyphen", name[l:i], l)
			case name[i-1] == '-':
				return fmt.Errorf("cookie domain: label '%s' at offset %d ends with a hyphen", name[l:i], l)
			}
			l = i + 1
			continue
		}
		// test label character validity, note: tests are ordered by decreasing validity frequency
		if !(b >= 'a' && b <= 'z' || b >= '0' && b <= '9' || b == '-' || b >= 'A' && b <= 'Z') {
			// show the printable unicode character starting at byte offset i
			c, _ := utf8.DecodeRuneInString(name[i:])
			if c == utf8.RuneError {
				return fmt.Errorf("cookie domain: invalid rune at offset %d", i)
			}
			return fmt.Errorf("cookie domain: invalid character '%c' at offset %d", c, i)
		}
	}
	// check top level domain validity
	switch {
	case l == len(name):
		return fmt.Errorf("cookie domain: missing top level domain, domain can't end with a period")
	case len(name)-l > 63:
		return fmt.Errorf("cookie domain: byte length of top level domain '%s' is %d, can't exceed 63", name[l:], len(name)-l)
	case name[l] == '-':
		return fmt.Errorf("cookie domain: top level domain '%s' at offset %d begins with a hyphen", name[l:], l)
	case name[len(name)-1] == '-':
		return fmt.Errorf("cookie domain: top level domain '%s' at offset %d ends with a hyphen", name[l:], l)
	case name[l] >= '0' && name[l] <= '9':
		return fmt.Errorf("cookie domain: top level domain '%s' at offset %d begins with a digit", name[l:], l)
	}
	return nil
}

func ps(str string) {
	fmt.Println("-------------------------------------------------------")
	fmt.Println("+ " + str)
}
func pm(str string) {
	fmt.Println("+ " + str)
}
func pe(str string) {
	fmt.Println("+ " + str)
	fmt.Println("-------------------------------------------------------")
}
func p(str string) {
	fmt.Println("-------------------------------------------------------")
	fmt.Println("+ " + str)
	fmt.Println("-------------------------------------------------------")
}
func pf(a ...interface{}) (n int, err error) {
	fmt.Println("-------------------------------------------------------")
	return fmt.Fprintln(os.Stdout, a...)
}
func pp(str string) {
	ps(str)
	pe("I Quit :'(")
	os.Exit(0)
}
func ContainsStr(array []string, value string) bool {
	for _, a := range array {
		if a == value {
			return true
		}
	}
	return false
}

func tempWrite(path string, str string) {
	f, _ := os.Create(path)
	f.WriteString(str)
	f.Close()
}
func strToArr(str string, sep string) []string {
	return strings.Split(
		strings.TrimSpace(str),
		sep,
	)
}
func arrToStr(strs []string, sep string) string {
	var str = ""
	t := len(strs) - 1
	for i, v := range strs {
		str += v
		if t != i {
			str += sep
		}
	}
	return str
}

// file data to []byte
func FileDataToByte(filePath string) (byteValue []byte) {
	// Open our jsonFile
	jsonFile, err := os.Open(filePath)
	// if we os.Open returns an error then handle it
	if err != nil {
		pp("Error opening file: " + err.Error())
	}
	// log.NoticeF("Successfully opened: %s", filePath)
	// defer the closing of our file so that we can parse it later on
	defer jsonFile.Close()
	// get bytes
	byteValue, err = ioutil.ReadAll(jsonFile)

	if err != nil {
		pp("Error ioutil.ReadAll: " + err.Error())
	}

	return
}

func PrettyPrint(data interface{}) {
	var p []byte
	//    var err := error
	p, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s \n", p)
}

func hashifyDollar(str string) string {
	return strings.ReplaceAll(str, "$", "#")
}
