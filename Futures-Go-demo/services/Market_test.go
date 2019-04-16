package services

import (
	"Futures-Go-demo/config"
	"Futures-Go-demo/websocket"
	"fmt"
	"testing"
	"time"
)

//测试获取合约信息接口
func Test_FutureContractInfo(t *testing.T) {
	contractInfo := FutureContractInfo("BTC", "", "")
	fmt.Println("获取合约信息: ", contractInfo)

}

//测试获取合约指数信息接口
func Test_FutureContractIndex(t *testing.T) {
	//contract_index := services.FutureContractIndex("BTC")
	contract_index := FutureContractIndex("BTC")
	fmt.Println("获取合约指数信息: ", contract_index)

}

//获取订单明细信息
func Test_FutureContractOrderDetail(t *testing.T) {

	contract_order_detail := FutureContractOrderDetail("BTC", "123556", "1", "100", "1539345271124", "1")
	fmt.Println("获取订单明细信息: ", contract_order_detail)

}

//合约取消订单
func Test_FutureContractCancel(t *testing.T) {
	contract_cancel := FutureContractCancel("123456", "BTC", "123456")
	fmt.Println("合约取消订单: ", contract_cancel)

}

//合约全部撤单
func Test_FutureContractCancelall(t *testing.T) {
	contract_cancelall := FutureContractCancelall("BTC")
	fmt.Println("合约全部撤单: ", contract_cancelall)

}

//获取合约当前未成交委托
func Test_FutureContractOpenorders(t *testing.T) {
	contract_openorders := FutureContractOpenorders("BTC", "1", "100")
	fmt.Println("获取合约当前未成交委托: ", contract_openorders)

}

//获取合约历史委托
func Test_FutureContractHisorders(t *testing.T) {
	contract_hisorders := FutureContractHisorders("BTC", "0", "1", "0", "90", "1", "50")
	fmt.Println("获取合约历史委托: ", contract_hisorders)
	time.Sleep(time.Second)

}

//测试合约下单接口
func Test_FutureContractOrder(t *testing.T) {
	//合约下单
	contract_order := FutureContractOrder("BTC", "this_week", "BTC181214", "", "6188", "12",
		"buy", "open", "10", "limit")
	fmt.Println("合约下单: ", contract_order)

}

//测试批量下单接口
func Test_FutureContractBatchorder(t *testing.T) {
	//合约批量下单
	ordersData := make([]*Order, 0)
	order1 := &Order{

		Symbol:         "BTC",
		ContractType:   "quarter",
		ContractCode:   "BTC181228",
		ClientOrderId:  "10",
		Price:          "6188",
		Volume:         "1",
		Direction:      "buy",
		Offset:         "open",
		LeverRate:      "10",
		OrderPriceType: "limit",
	}

	ordersData = append(ordersData, order1)
	order2 := &Order{

		Symbol:         "BTC",
		ContractType:   "quarter",
		ContractCode:   "BTC181228",
		ClientOrderId:  "11",
		Price:          "6189",
		Volume:         "2",
		Direction:      "buy",
		Offset:         "open",
		LeverRate:      "10",
		OrderPriceType: "limit",
	}
	ordersData = append(ordersData, order2)
	fmt.Println("ordersData:", ordersData)

	contract_batchorder := FutureContractBatchorder(ordersData)
	fmt.Println("合约批量下单ordersDataResult: ", contract_batchorder)

}

//测试 WebSocket 行情,交易 API
func Test_Websocket(t *testing.T) {
	//websocket.WSRun()   //无需本地IP地址，直接运行
	websocket.WSRunWithIP(config.Local_IP) //配置文件须填写本地IP地址，WS运行太久，外部原因可能断开，支持自动重连
}

//测试 WebSocket 订单推送 API
func Test_Websocket_order(t *testing.T) {

	websocket.WSWithOrder()
}
