package websocket

import (
	"fmt"
	"log"

	"Futures-Go-demo/untils"
	"golang.org/x/net/websocket"

	"Futures-Go-demo/config"
)

func sendOrder(message []byte, ws *websocket.Conn) {
	_, err := ws.Write(message)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("Send: %s\n", message)
}
func WSWithOrder() {

	params :=make(map[string]string)
	resultParams := untils.ApiKeyGetOrder(params, "/notification")

	run(resultParams)

}
func run(mapParams  map[string]string) {




	ws, err := websocket.Dial(config.WS_ORDER_URL, "", origin)
	if err != nil {
		log.Fatal(err)
	}

    /************************************************/
	/**
	* 发送账户订单请求鉴权
	* Send account order authentication request
	*/
	str := fmt.Sprintf("{\"op\":\"%s\",\"type\":\"%s\",\"AccessKeyId\":\"%s\",\"SignatureMethod\":\"%s\"," +
		"\"SignatureVersion\":\"%s\",\"Timestamp\":\"%s\",\"Signature\":\"%s\"}",mapParams["op"],mapParams["type"],
		mapParams["AccessKeyId"],mapParams["SignatureMethod"],mapParams["SignatureVersion"],mapParams["Timestamp"],mapParams["Signature"])

	message := []byte(str)
	sendOrder(message, ws)

    //sub
	/**
		* 发送账户订单sub请求
		* Send account order sub request
		* @param topic 订阅topic
		* @param cid   标识
		*/
     // orders.btc 这样的格式，可以填写orders.eth orders.ltc ...
	subStr := fmt.Sprintf("{\"op\":\"%s\",\"topic\":\"%s\",\"cid\":\"%s\"}","sub","orders.btc","btc")
	subMessage := []byte(subStr)
	sendOrder(subMessage, ws)
	fmt.Printf("Send: %s\n", subMessage)

    //unsub
	/**
		* 发送账户订单取消订阅发送unsub请求
		* Send account order unsub request
		* @param topic unsub请求topic、 topic unsub request topic
		* @param cid   标识、 cid identifying
		*/

/*
	unsubStr := fmt.Sprintf("{\"op\":\"%s\",\"topic\":\"%s\",\"cid\":\"%s\"}","unsub","orders.btc","btc")
	unsubMessage := []byte(unsubStr)
	sendOrder(unsubMessage, ws)
    fmt.Printf("Send: %s\n", unsubMessage)

*/



	var msg = make([]byte, 512000)

	for {
		m, err := ws.Read(msg)
		if err != nil {
			log.Fatal(err)
		}

		newmsg := msg[:m]

		unzipmsg, _ := ParseGzip(newmsg, true)
		fmt.Printf("Receive: %s\n", unzipmsg)


		if len(unzipmsg) > 33 {
			pingcmd := string(unzipmsg[7:11])
			if "ping" == pingcmd {

				pingtime := string(unzipmsg[19:33])
				pongsStr := fmt.Sprintf("{\"op\":\"%s\",\"ts\":\"%s}","pong",pingtime)
				pongMessage := []byte(pongsStr)
				sendOrder(pongMessage, ws)
				fmt.Printf("Send: %s\n", pongMessage)
			}
		}
	}

	ws.Close() //关闭连接


}


