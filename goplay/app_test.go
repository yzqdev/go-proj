package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
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
