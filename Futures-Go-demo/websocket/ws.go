package websocket

import (
	"Futures-Go-demo/config"
	"bytes"
	"compress/gzip"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/net/websocket"
)

var origin = "http://www.baidu.com"

//var url = "wss://www.hbdm.com/ws"
var buffer bytes.Buffer

func ParseGzip(data []byte, handleErr bool) ([]byte, error) {
	b := new(bytes.Buffer)
	binary.Write(b, binary.LittleEndian, data)
	r, err := gzip.NewReader(b)
	if err != nil {
		//with error
		fmt.Printf("[ParseGzip %t %d] NewReader error: %v, maybe data is ungzip: [%s]\n", handleErr, len(data), err)
		if handleErr {
			errHandler(data)
		}
		return nil, err
	} else {
		defer r.Close()
		undatas, err := ioutil.ReadAll(r)
		if err != nil {
			//with error
			fmt.Printf("[ParseGzip %t %d] ioutil.ReadAll error: %v: [%s]\n", handleErr, len(data), err)
			if handleErr {
				errHandler(data)
			}
			return nil, err
		} else {
			//buffer.Reset()
			return undatas, nil
		}
	}
}

func errHandler(data []byte) {
	buffer.Write(data)
	msg, err := ParseGzip(buffer.Bytes(), false)
	if err == nil {
		fmt.Println("!!!!!!", string(msg[:]))
	}
}

func send(message []byte, ws *websocket.Conn) {
	_, err := ws.Write(message)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Send: %s\n", message)
}

func WSRun() {
	ws, err := websocket.Dial(config.WS_URL, "", origin)
	if err != nil {
		log.Fatal(err)
	}

	//===============================================================================

	//订阅websocket kline
	/*
		"market.$symbol.kline.$period"
		symbol	true	string	交易对		如"BTC_CW"表示BTC当周合约，"BTC_NW"表示BTC次周合约，"BTC_CQ"表示BTC季度合约
		period	true	string	K线周期		1min, 5min, 15min, 30min, 60min,4hour,1day, 1mon
	*/
	message := []byte("{\"Sub\":\"market.BTC_CW.kline.1min\"}")
	send(message, ws)


	//订阅websocket Market Detail 数据
	/*
	  "market.$symbol.detail"
	  symbol	true	string	交易对		如"BTC_CW"表示BTC当周合约，"BTC_NW"表示BTC次周合约，"BTC_CQ"表示BTC季度合约
	*/
	message = []byte("{\"Sub\":\"market.BTC_CW.detail\"}")
	send(message, ws)

	//订阅websocket Trade Detail 数据
	/*
	  "market.$symbol.trade.detail"
	  symbol	true	string	交易对		如"BTC_CW"表示BTC当周合约，"BTC_NW"表示BTC次周合约，"BTC_CQ"表示BTC季度合约
	*/
	message = []byte("{\"Sub\":\"market.BTC_CW.trade.detail\"}")
	send(message, ws)


	//订阅websocket Market Depth 数据
	/*
			  "market.$symbol.depth.$type"
			 symbol	true	string	交易对		如"BTC_CW"表示BTC当周合约，"BTC_NW"表示BTC次周合约，"BTC_CQ"表示BTC季度合约.
		     type	true	string	Depth 类型	(150档数据)	step0, step1, step2, step3, step4, step5（合并深度1-5）；step0时，不合并深度
	                                            (20档数据)  step6, step7, step8, step9, step10, step11（合并深度7-11）；step6时，不合并深度
	*/
	message = []byte("{\"Sub\":\"market.BTC_CW.depth.step0\"}")
	send(message, ws)


	//请求websocket KLine 数据
	/*
			  "market.$symbol.kline.$period"
			 symbol	true	string	交易对		如"BTC_CW"表示BTC当周合约，"BTC_NW"表示BTC次周合约，"BTC_CQ"表示BTC季度合约
		     period	true	string	K线周期		1min, 5min, 15min, 30min, 60min,4hour,1day, 1mon

			"from": "optional, type: long, 2017-07-28T00:00:00+08:00 至2050-01-01T00:00:00+08:00 之间的时间点，单位：秒",
		  	"to": "optional, type: long, 2017-07-28T00:00:00+08:00 至2050-01-01T00:00:00+08:00 之间的时间点，单位：秒，必须比 from 大"}

			[t1, t5] 假设有 t1  ~ t5 的K线：

			from: t1, to: t5, return [t1, t5].
			from: t5, to: t1, which t5  > t1, return [].
			from: t5, return [t5].
			from: t3, return [t3, t5].
			to: t5, return [t1, t5].
			from: t which t3  < t  <t4, return [t4, t5].
			to: t which t3  < t  <t4, return [t1, t3].
	*/
	message = []byte("{\"req\":\"market.BTC_CQ.kline.1day\",\"from\":1544170607,\"to\":1544602608}")
	send(message, ws)


	//请求websocket Trade Detail 数据
	/*
	  "market.$symbol.trade.detail"
	  symbol	true	string	交易对		如"BTC_CW"表示BTC当周合约，"BTC_NW"表示BTC次周合约，"BTC_CQ"表示BTC季度合约
	*/
	message = []byte("{\"req\":\"market.BTC_CW.trade.detail\"}")
	send(message, ws)



	//===============================================================================

	var msg = make([]byte, 512000)

	for {
		m, err := ws.Read(msg)
		if err != nil {
			log.Fatal(err)
		}
		newmsg := msg[:m]

		unzipmsg, _ := ParseGzip(newmsg, true)

		fmt.Printf("Receive[UNZIP]: [%d:%d] %s\n", m, len(unzipmsg), unzipmsg[:])

		if len(unzipmsg) > 21 {
			pingcmd := string(unzipmsg[2:6])
			if "ping" == pingcmd {
				pingtime := string(unzipmsg[8:21])
				pongstr := fmt.Sprintf("{\"pong\":%s}", pingtime)
				message := []byte(pongstr)

				send(message, ws)
			}
		}
	}

	ws.Close() //关闭连接

}
