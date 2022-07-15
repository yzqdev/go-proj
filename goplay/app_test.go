package main

import (
	"fmt"
	goquery "github.com/PuerkitoBio/goquery"
	"github.com/google/go-querystring/query"
	"github.com/gookit/color"
	jsoniter "github.com/json-iterator/go"
	"golang.org/x/sys/windows/registry"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/user"
	"regexp"
	"strings"
	"testing"
)

func TestWalk(t *testing.T) {
	file, err := ioutil.ReadFile("a.json")
	gache := &Gacha{}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	json.UnmarshalFromString(string(file), gache)
	if err != nil {
		fmt.Println(err)
		return
	}

}

type ExportGacha struct {
	Info struct {
		Uid              string `json:"uid"`
		Lang             string `json:"lang"`
		ExportTime       string `json:"export_time"`
		ExportTimestamp  string `json:"export_timestamp"`
		ExportApp        string `json:"export_app"`
		ExportAppVersion string `json:"export_app_version"`
		UigfVersion      string `json:"uigf_version"`
	} `json:"info"`
	List []Gacha `json:"list"`
}
type Gacha struct {
	Uid           string `json:"uid"`
	GachaType     string `json:"gacha_type"`
	ItemId        string `json:"item_id"`
	Count         string `json:"count"`
	Time          string `json:"time"`
	Name          string `json:"name"`
	Lang          string `json:"lang"`
	ItemType      string `json:"item_type"`
	RankType      string `json:"rank_type"`
	Id            string `json:"id"`
	UigfGachaType string `json:"uigf_gacha_type"`
}

func TestGen(t *testing.T) {
	var reg = "SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Uninstall\\原神\\"
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, reg, registry.ALL_ACCESS)
	if err != nil {

	}
	var ud, _, err2 = k.GetStringValue("InstallPath")
	if err2 != nil {
		color.Redln(err2)
	}
	color.Cyanln(ud)
}
func TestIcon(t *testing.T) {
	res, err := http.Get("https://github.com")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	icon := doc.Find("head link[rel*='icon']")
	href, _ := icon.Attr("href")
	color.Redln(href)
	title := doc.Find("head title")

	color.Redln(title.Text())
}
func TestExampleScrape(t *testing.T) {
	// Request the HTML page.
	res, err := http.Get("http://metalsucks.net")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".left-content article .post-title").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		title := s.Find("a").Text()
		fmt.Printf("Review %d: %s\n", i, title)
	})
}

type GachaQuery struct {
	GachaType string `url:"gacha_type"`
	Page      string `url:"page"`

	Size  string `url:"size"`
	EndId string `url:"end_id"`
	//gacha_type={(int)WishType}&page={Page}&size={Size}&end_id={EndId}
}

func TestGetGenshinWish(t *testing.T) {
	u, _ := user.Current()
	logPath := u.HomeDir + "\\AppData\\LocalLow\\miHoYo\\原神\\output_log.txt"
	file, err := os.ReadFile(logPath)
	if err != nil {
		return
	}
	var regX = regexp.MustCompile("OnGetWebViewPageFinish:http.*#/log")

	res := strings.ReplaceAll(regX.FindAllStringSubmatch(string(file), -1)[0][0], "OnGetWebViewPageFinish:", "")
	var param = strings.ReplaceAll(res[strings.Index(res, "?"):], "#/log", "")
	color.Redln(res)
	//url?.substring(url.indexOf('?')).replaceAll("#/log", "");
	var queryParam = &GachaQuery{
		GachaType: "301",
		Page:      "16",
		Size:      "6",
		EndId:     "1654229160001129743",
	}
	color.Redln(param)
	v, _ := query.Values(queryParam)

	var gacharUrl = "https://hk4e-api.mihoyo.com/event/gacha_info/api/getGachaLog" + param + "&" + v.Encode()
	color.Redln(gacharUrl)
}
