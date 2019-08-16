package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

type testConfig struct {
	Count       int32  `json:"count"`
	URL         string `json:"url"`
	CommodityID int32  `json:"commodityID"`
}

type requestData struct {
	ID int32 `json:"id"`
}

type rspData struct {
	Success bool `json:"success"`
}

func buyRequest(lock *sync.RWMutex, ret chan<- bool, url string, reqData []byte) {
	body := bytes.NewBuffer(reqData)
	lock.RLock()
	defer func() {
		lock.RUnlock()
	}()

	resp, err := http.Post(url, "application/json;charset=utf-8", body)
	if err != nil {
		log.Println("send post request failed:", err)
		ret <- false
		return
	}
	if resp.StatusCode != 200 {
		ret <- false
		return
	}

	rsp := &rspData{}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("read request content failed:", err)
		ret <- false
		return
	}
	err = json.Unmarshal(content, rsp)
	if err != nil {
		log.Panicf("parse config json error: %v\n", err)
		ret <- false
		return
	}
	if rsp.Success {
		ret <- true
	} else {
		ret <- false
	}
}

func main() {
	data, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Panicf("read ./config.json file error: %v\n", err)
	}
	config := &testConfig{}
	err = json.Unmarshal(data, config)
	if err != nil {
		log.Panicf("parse config json error: %+v\n", err)
	}

	reqData, _ := json.Marshal(requestData{ID: config.CommodityID})
	ret := make(chan bool)
	rwlock := &sync.RWMutex{}
	rwlock.Lock()

	var num int32
	for num = 0; num < config.Count; num++ {
		go buyRequest(rwlock, ret, config.URL, reqData)
	}
	log.Printf("create %d goroutine complete", num)
	time.Sleep(1 * time.Second)
	rwlock.Unlock()

	var reqRet bool
	var success, fail int
	for num = 0; num < config.Count; num++ {
		reqRet = <-ret
		if reqRet {
			success++
		} else {
			fail++
		}
	}

	log.Printf("test complete, success: %d fail: %d\n", success, fail)
}
