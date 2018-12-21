package main

import (
	"fmt"
	"WS-REST-GO-demo-master/services"
	"WS-REST-GO-demo-master/websocket"
	"time"
)


func main() {
	//------------------------------------------------------------------------------------------
	// 公共API
    // 获取合约信息
	contractInfo := services.FutureContractInfo("BTC","","")
	fmt.Println(time.Now().Format("2006-01-02 15:04:05.000000"),"  获取合约信息: ", contractInfo)
    time.Sleep(time.Second)


	// 获取合约指数信息
	contract_index := services.FutureContractIndex("BTC")
	fmt.Println(time.Now().Format("2006-01-02 15:04:05.000000"),"  获取合约指数信息: ", contract_index)
	time.Sleep(time.Second)

	// 获取合约最高限价和最低限价
	contract_price_limit := services.FuturePriceLimit("BTC","","")
	fmt.Println(time.Now().Format("2006-01-02 15:04:05.000000"),"  获取合约最高限价和最低限价: ", contract_price_limit)
	time.Sleep(time.Second)

	//获取当前可用合约总持仓量
	contract_open_interest := services.FutureOpenInterest("BTC", "", "")
	fmt.Println(time.Now().Format("2006-01-02 15:04:05.000000"),"  获取当前可用合约总持仓量: ", contract_open_interest)
	time.Sleep(time.Second)
	//------------------------------------------------------------------------------------------
	//市场行情API
	////获取交易深度信息
	marketDepthReturn :=services.FutureMarketDepth("BTC_CW","step0")
	fmt.Println(time.Now().Format("2006-01-02 15:04:05.000000"),"  获取交易深度信息:",marketDepthReturn)
	time.Sleep(time.Second)

	// 获取K线数据
	kline :=services.FutureMarketHistoryKline("BTC_CW","1min",10)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05.000000"),"  获取K线数据:",kline)
	time.Sleep(time.Second)

	// 获取聚合行情
	tickerReturn :=services.FutureMarketDetailMerged("BTC_CW")
	fmt.Println(time.Now().Format("2006-01-02 15:04:05.000000"),"  获取聚合行情:",tickerReturn)
	time.Sleep(time.Second)


	//获取市场最近成交记录
	tradeDetailReturn :=services.FutureMarketDetailTrade("BTC_CW")
	fmt.Println(time.Now().Format("2006-01-02 15:04:05.000000"),"  获取市场最近成交记录:",tradeDetailReturn)
	time.Sleep(time.Second)

	//批量获取最近的交易记录
	tradeReturn :=services.FutureMarketHistoryTrade("BTC_CW",10)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05.000000"),"  批量获取最近的交易记录:",tradeReturn)
	time.Sleep(time.Second)

	//获取Market Detail 24小时成交量数据
	marketDetailReturn :=services.FutureMarketDetail("BTC_CW")
	fmt.Println(time.Now().Format("2006-01-02 15:04:05.000000"),"  获取Market Detail 24小时成交量数据:",marketDetailReturn)
	time.Sleep(time.Second)

	//------------------------------------------------------------------------------------------
	// 用户相关API

	//获取用户账户信息
	accountInfo := services.FutureContractAccountInfo("BTC")
	fmt.Println(time.Now().Format("2006-01-02 15:04:05.000000"),"  获取用户账户信息: ", accountInfo)
	time.Sleep(time.Second)

	//获取用户持仓信息
	contract_position_info := services.FutureContractPositionInfo("BTC")
	fmt.Println(time.Now().Format("2006-01-02 15:04:05.000000"),"  获取用户持仓信息: ", contract_position_info)
	time.Sleep(time.Second)


	//合约下单
	contract_order := services.FutureContractOrder("BTC", "this_week", "BTC181214", "", "6188", "12",
		"buy", "open", "10", "limit")
	fmt.Println(time.Now().Format("2006-01-02 15:04:05.000000"),"  合约下单: ", contract_order)
	time.Sleep(time.Second)

	//合约批量下单
	orders := make([]*services.Order,0)
	order1 := &services.Order{
		Symbol :           "BTC",
		ContractType :     "this_week",
		ContractCode :     "BTC181214",
		ClientOrderId :    "10",
		Price    :         "6188",
		Volume   :         "1",
		Cirection :        "buy",
		Offset :           "open",
		LeverRate :        "10",
		OrderPriceType :   "limit",
	}
	order2 := &services.Order{
		Symbol :           "BTC",
		ContractType :     "this_week",
		ContractCode :     "BTC181214",
		ClientOrderId :    "12",
		Price    :         "6188",
		Volume   :         "2",
		Cirection :        "buy",
		Offset :           "open",
		LeverRate :        "10",
		OrderPriceType :   "limit",
	}
	orders = append(orders ,order1)
	orders = append(orders ,order2)

	contract_batchorder := services.FutureContractBatchorder(orders)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05.000000"),"  合约批量下单: ", contract_batchorder)
	time.Sleep(time.Second)

	//合约取消订单
	contract_cancel := services.FutureContractCancel("123456","BTC", "123456")
	fmt.Println(time.Now().Format("2006-01-02 15:04:05.000000"),"  合约取消订单: ", contract_cancel)
	time.Sleep(time.Second)


	//合约全部撤单
	contract_cancelall := services.FutureContractCancelall("BTC")
	fmt.Println(time.Now().Format("2006-01-02 15:04:05.000000"),"  合约全部撤单: ", contract_cancelall)
	time.Sleep(time.Second)

	//获取合约订单信息
	contract_order_info := services.FutureContractOrderInfo("4", "","BTC","1")
	fmt.Println(time.Now().Format("2006-01-02 15:04:05.000000"),"  获取合约订单信息: ", contract_order_info)
	time.Sleep(time.Second)

	//获取订单明细信息
	contract_order_detail := services.FutureContractOrderDetail("BTC", "123456", "1", "100","1539345271124","1")
	fmt.Println(time.Now().Format("2006-01-02 15:04:05.000000"),"  获取订单明细信息: ", contract_order_detail)
	time.Sleep(time.Second)


	//获取合约当前未成交委托
	contract_openorders := services.FutureContractOpenorders("BTC", "1", "100")
	fmt.Println(time.Now().Format("2006-01-02 15:04:05.000000"),"  获取合约当前未成交委托: ", contract_openorders)
	time.Sleep(time.Second)

	//获取合约历史委托
	contract_hisorders := services.FutureContractHisorders("BTC", "0", "1", "0", "90", "1", "20")
	fmt.Println(time.Now().Format("2006-01-02 15:04:05.000000"),"  获取合约历史委托: ", contract_hisorders)
	time.Sleep(time.Second)

    // WebSocket 行情,交易 API
	fmt.Println()
	fmt.Println("********************websocket  Run******************************")
	websocket.Run()
}




