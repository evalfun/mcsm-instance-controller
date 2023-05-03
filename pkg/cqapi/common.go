package cqapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/gorilla/websocket"
)

type CQAPIClient struct {
	APIEndpoint string
	RecvChan    chan interface{}
	Status      string
	wsConn      *websocket.Conn
}

func CreateCQAPIClient(serverEndpoint string) *CQAPIClient {
	cqAPI := &CQAPIClient{
		APIEndpoint: serverEndpoint,
		RecvChan:    make(chan interface{}, 100),
		Status:      "connecting",
	}
	go cqAPI.cqAPIWSProc()
	return cqAPI
}

func (cq *CQAPIClient) SendGroupMessage(groupID int64, message string) {
	params := WebsocketJsonRequestParams{
		Action: "send_msg",
		Params: SendMessageParams{
			MessageType: "group",
			GroupID:     groupID,
			Message:     message,
			AutoEscape:  true,
		},
		Echo: fmt.Sprint(rand.Int63()),
	}
	log.Printf("send_msg to %d echo %s: %s", groupID, params.Echo, message)
	cq.InvokeAPI(params)
}

func (cq *CQAPIClient) InvokeAPI(params interface{}) error {
	if cq.Status != "connected" {
		return errors.New("websocket not connected")
	}
	jsonData, err := json.Marshal(params)
	if err != nil {
		return err
	}
	//log.Println(string(jsonData))
	err = cq.wsConn.WriteMessage(websocket.TextMessage, jsonData)
	return err
}

func (cq *CQAPIClient) cqAPIWSProc() {
	for {
		var err error
		cq.wsConn, _, err = websocket.DefaultDialer.Dial(cq.APIEndpoint, nil)
		if err != nil {
			log.Fatalf("Connect to go-cqhttp ws server failed: %v\n Reconnect in 3s.\n", err.Error())
		} else {
			cq.Status = "connected"
			log.Printf("connect to go-cqhttp ws server %s success\n", cq.APIEndpoint)
			for {
				_, data, err := cq.wsConn.ReadMessage()
				if err != nil {
					log.Fatalf("Disconect from go-cqhttp ws server: %v\n Reconnect in 3s.\n", err.Error())
					cq.Status = "connecting"
					break
				}
				commonReport := &CommonReport{}
				err = json.Unmarshal(data, commonReport)
				if err != nil {
					log.Fatalf("Can not unmarshal json from cq api server:%v\n", err.Error())
					continue
				}
				if commonReport.PostType == "message" {
					report := &MessageReport{}
					json.Unmarshal(data, report)
					cq.RecvChan <- report
				} else if commonReport.PostType == "request" {
					report := &RequestReport{}
					json.Unmarshal(data, report)
					cq.RecvChan <- report
				} else if commonReport.PostType == "notice" {
					report := &NoticeReport{}
					json.Unmarshal(data, report)
					cq.RecvChan <- report
				} else if commonReport.PostType == "meta_event" {
					report := &MetaEventReport{}
					json.Unmarshal(data, report)
					cq.RecvChan <- report
				} else {
					log.Println("receive ws data:", string(data))
				}
			}
		}
		time.Sleep(3 * time.Second)
	}
}
