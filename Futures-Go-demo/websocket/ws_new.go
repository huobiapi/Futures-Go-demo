package websocket

import (
	"Futures-Go-demo/config"
	"bytes"
	"compress/gzip"
	"encoding/binary" //"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"time"

	"github.com/gorilla/websocket"
	"github.com/spf13/cast"
)

type Message struct {
	Ts      int    `json:"ts"`
	Status  string `json:"status"`
	ErrCode string `json:"err-code"`
	ErrMsg  string `json:"err-msg"`
	Ping    int    `json:"ping"`
}

type Client struct {
	Name string
	Addr string
	Ws   *websocket.Conn
}

var (
	localIP string
)

type Moniter struct {
	clientNum  int
	addChan    chan int
	subChan    chan int
	lastUseSec int
}

var mon *Moniter

func InitMoniter() {
	mon = &Moniter{}
	mon.addChan = make(chan int, 1000)
	mon.subChan = make(chan int, 1000)
	go func() {

		for {
			select {
			case <-mon.addChan:
				mon.clientNum++
			case <-mon.subChan:
				mon.clientNum--

			}
		}
	}()
}

func AddClientNum() {
	mon.addChan <- 1
}

func SubClientNum() {
	mon.subChan <- 1
}

func NewClient(addr string, name string) *Client {
	return &Client{Name: name, Addr: addr}
}

func NowSec() int {
	return int(time.Now().UnixNano() / 1000000000)
}

func (cli *Client) RunClient() {

	AddClientNum()
	dialer := websocket.DefaultDialer
	dialer.NetDial = func(network, addr string) (net.Conn, error) {

		addrs := []string{string(localIP)}
		fmt.Println("addrs=", addrs)
		localAddr := &net.TCPAddr{IP: net.ParseIP(addrs[rand.Int()%len(addrs)]), Port: 0}
		d := net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			LocalAddr: localAddr,
			DualStack: true,
		}
		c, err := d.Dial(network, addr)
		return c, err
	}
	c, _, err := dialer.Dial(cli.Addr, nil)
	if err != nil {
		log.Println("Dial Erro:", err)
		SubClientNum()
		return
	}
	log.Println(c.LocalAddr().String())
	//	defer c.Close()
	defer func() {
		c.Close()
		SubClientNum()
	}()

	//===============================================================================

	//订阅websocket kline
	/*
		"market.$symbol.kline.$period"
		symbol	true	string	交易对		如"BTC_CW"表示BTC当周合约，"BTC_NW"表示BTC次周合约，"BTC_CQ"表示BTC季度合约
		period	true	string	K线周期		1min, 5min, 15min, 30min, 60min,4hour,1day, 1mon
	*/
	message := []byte("{\"Sub\":\"market.BTC_CW.kline.60min\"}")
	err = c.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		log.Println("write err :", err)
	}

	//订阅websocket Market Detail 数据
	/*
	  "market.$symbol.detail"
	  symbol	true	string	交易对		如"BTC_CW"表示BTC当周合约，"BTC_NW"表示BTC次周合约，"BTC_CQ"表示BTC季度合约
	*/
	message = []byte("{\"Sub\":\"market.BTC_CW.detail\"}")
	err = c.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		log.Println("write err :", err)
	}

	//订阅websocket Trade Detail 数据
	/*
	  "market.$symbol.trade.detail"
	  symbol	true	string	交易对		如"BTC_CW"表示BTC当周合约，"BTC_NW"表示BTC次周合约，"BTC_CQ"表示BTC季度合约
	*/
	message = []byte("{\"Sub\":\"market.BTC_CW.trade.detail\"}")
	err = c.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		log.Println("write err :", err)
	}

	//订阅websocket Market Depth 数据
	/*
				  "market.$symbol.depth.$type"
				 symbol	true	string	交易对		如"BTC_CW"表示BTC当周合约，"BTC_NW"表示BTC次周合约，"BTC_CQ"表示BTC季度合约.
		        type	true	string	Depth 类型	(150档数据)	step0, step1, step2, step3, step4, step5（合并深度1-5）；step0时，不合并深度
		                                            (20档数据)  step6, step7, step8, step9, step10, step11（合并深度7-11）；step6时，不合并深度
	*/

	message = []byte("{\"Sub\":\"market.BTC_CW.depth.step0\"}")
	err = c.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		log.Println("write err :", err)
	}

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
	err = c.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		log.Println("write err :", err)
	}

	go func() {

		pangTicker := time.NewTicker(time.Second * 5)

		for {
			select {
			case <-pangTicker.C:
				message = []byte(fmt.Sprintf("{\"pong\":%d}", time.Now().Unix()))
				err = c.WriteMessage(websocket.TextMessage, message)
				if err != nil {
					log.Println("send msg err:", err)
					return
				}
			}
		}
	}()

	for {
		_, zipmsg, err := c.ReadMessage()
		if err != nil {
			log.Println("Read Error : ", err, cli.Name)
			c.Close()
			return
		}

		msg, err := parseGzip(zipmsg)
		if err != nil {
			log.Println("gzip Error : ", err)
		}
		log.Println(string(msg[:]))

	}
}

var address string

func WSRunWithIP(ip string) {

	address = config.WS_URL

	localIP = ip

	InitMoniter()

	createClient(address)
	time.Sleep(1 * time.Millisecond)

	reCreateClient()

	for {
		time.Sleep(10 * time.Second)
	}
}

func reCreateClient() {
	go func() {
		time.Sleep(time.Second * 100)
		checkTicker := time.NewTicker(time.Second * 20)
		for {
			select {
			case <-checkTicker.C:

				if mon.clientNum <= 0 {

					createClient(address)

				}
			}
		}
	}()

}

var clientNameNum int

func createClient(addr string) {
	clientNameNum++
	c := NewClient(addr, cast.ToString(clientNameNum))
	go c.RunClient()
}

func parseGzip(data []byte) ([]byte, error) {
	b := new(bytes.Buffer)
	binary.Write(b, binary.LittleEndian, data)
	r, err := gzip.NewReader(b)
	if err != nil {
		fmt.Println("[ParseGzip] NewReader error: , maybe data is ungzip: ", err, string(data))
		fmt.Println("[ParseGzip] NewReader error: , maybe data is ungzip: ", err, string(b.Bytes()))
		return nil, err
	} else {
		defer r.Close()
		undatas, err := ioutil.ReadAll(r)
		if err != nil {
			fmt.Println("[ParseGzip]  ioutil.ReadAll error: :", err, string(data))
			return nil, err
		}
		return undatas, nil
	}
}
