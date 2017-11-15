package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mock-test-revel/app/configs"
	"net/http"

	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

type ViewResult []SiteData

type SiteData struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func (c App) Index() revel.Result {
	conf := configs.Initialize()
	SeastallionOwndURL := conf.SeastallionOwnd
	resp, err := http.Get(SeastallionOwndURL)
	if err != nil {
		fmt.Printf("Error has occured on get request")
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	v := ViewResult{}
	json.Unmarshal(body, &v)
	return c.RenderJson(v)
}
