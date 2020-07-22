package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetSignNameById(signId int) string {
	for name, id := range SignMap {
		if id == signId {
			return name
		}
	}
	return ""
}

func GetHoroscope(signName string) (string, error) {
	signName = strings.ToUpper(signName)
	signId, ok := SignMap[signName]
	if !ok {
		fmt.Errorf("The sign name is not existed")
	}

	url := fmt.Sprintf("%s?sign=%d", endpoint, signId)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	content := string(body)
	return content, nil

}

var endpoint string = "http://www.horoscope.com/us/horoscopes/general/horoscope-general-daily-today.aspx"

var SignMap = map[string]int{
	"ARIES":       1,
	"TAURUS":      2,
	"GEMINI":      3,
	"CANCER":      4,
	"LEO":         5,
	"VIRGO":       6,
	"LIBRA":       7,
	"SCORPIO":     8,
	"SAGITTARIUS": 9,
	"CAPRICORN":   10,
	"AQUARIUS":    11,
	"PISCES":      12,
}

func main() {
	var signName string
	flag.StringVar(&signName, "sign", "", "Sign name")
	flag.Parse()
	if signName == "" {
		var signId int = -1
		for k, v := range SignMap {
			fmt.Println(k, "--", v)
		}
		isFirst := true
		for signName == "" {
			if !isFirst {
				fmt.Print("This sign is not exist, please type a number from the list below :")
			} else {
				isFirst = false
				fmt.Print("Please choose a horoscope (number):")
			}
			fmt.Scanf("%d", &signId)
			if signId == 0 {
				return
			}
			signName = GetSignNameById(signId)
			fmt.Println(signName)
		}
	}
	fmt.Println("Please wait ...")
	content, err := GetHoroscope(signName)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(content)
}
