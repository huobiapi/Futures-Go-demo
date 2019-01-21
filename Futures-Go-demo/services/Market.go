package services

import (
	"strconv"
	"Futures-Go-demo/config"
	"Futures-Go-demo/untils"
)

type Order struct {
	Symbol          string  `json:"symbol"`
	ContractType        string `json:"contractType"`
	ContractCode      string    `json:"contractCode"`
	ClientOrderId           string `json:"clientOrderId"`
	Price          string  `json:"price"`
	Volume          string  `json:"volume"`
	Direction string     `json:"direction"`
	Offset    string     `json:"offset"`
	LeverRate       string   `json:"leverRate"`
	OrderPriceType  string    `json:"orderPriceType"`

}
//------------------------------------------------------------------------------------------
//市场行情API
    /**
	 * 获取K线数据
	 *
	 * @param symbol
	 *            "BTC","ETH"...
	 * @param period
	 *            K线类型 1min, 5min, 15min, 30min, 60min,4hour,1day, 1mon
	 * @return
	 */

func FutureMarketHistoryKline(strSymbol, strPeriod string, nSize int) string {

	mapParams := make(map[string]string)
	mapParams["symbol"] = strSymbol
	mapParams["period"] = strPeriod
	mapParams["size"] = strconv.Itoa(nSize)

	strRequestUrl := "/market/history/kline"
	strUrl := config.MARKET_URL + strRequestUrl

	jsonKLineReturn := untils.HttpGetRequest(strUrl, mapParams)
	return jsonKLineReturn
}

    /**
	 * 获取聚合行情
	 *
	 * @param symbol
	 *            如"BTC_CW"表示BTC当周合约，"BTC_NW"表示BTC次周合约，"BTC_CQ"表示BTC季度合约
	 * @return
	 */
func FutureMarketDetailMerged(strSymbol string) string {


	mapParams := make(map[string]string)
	mapParams["symbol"] = strSymbol

	strRequestUrl := "/market/detail/merged"
	strUrl := config.MARKET_URL + strRequestUrl

	jsonTickReturn := untils.HttpGetRequest(strUrl, mapParams)


	return jsonTickReturn
}


	/**
	  * 获取行情深度数据
	  *
	  * @param symbol
	  *            "BTC","ETH"...
	  * @param type
	  * 	(150档数据)	step0, step1, step2, step3, step4, step5（合并深度1-5）；step0时，不合并深度
	  *     (20档数据)  step6, step7, step8, step9, step10, step11（合并深度7-11）；step6时，不合并深度
	  * @return
	  */
func FutureMarketDepth(strSymbol, strType string) string {
	mapParams := make(map[string]string)
	mapParams["symbol"] = strSymbol
	mapParams["type"] = strType
	strRequestUrl := "/market/depth"
	strUrl := config.MARKET_URL + strRequestUrl
	jsonMarketDepthReturn := untils.HttpGetRequest(strUrl, mapParams)
	return jsonMarketDepthReturn
}

    /**
	 * 获取市场最近成交记录
	 *
	 * @param symbol
	 *            如"BTC_CW"表示BTC当周合约，"BTC_NW"表示BTC次周合约，"BTC_CQ"表示BTC季度合约
	 * @return size 获取交易记录的数量 [1, 2000]
	 */
func FutureMarketDetailTrade(strSymbol string) string {

	mapParams := make(map[string]string)
	mapParams["symbol"] = strSymbol

	strRequestUrl := "/market/trade"
	strUrl := config.MARKET_URL + strRequestUrl

	jsonTradeDetailReturn := untils.HttpGetRequest(strUrl, mapParams)

	return jsonTradeDetailReturn
}

    /**
	 * 批量获取最近的交易记录
	 *
	 * @param symbol
	 *            如"BTC_CW"表示BTC当周合约，"BTC_NW"表示BTC次周合约，"BTC_CQ"表示BTC季度合约
	 * @return size 获取交易记录的数量 [1, 2000]
	 */
func FutureMarketHistoryTrade(strSymbol string, nSize int) string {
	mapParams := make(map[string]string)
	mapParams["symbol"] = strSymbol
	mapParams["size"] = strconv.Itoa(nSize)
	strRequestUrl := "/market/history/trade"
	strUrl := config.MARKET_URL + strRequestUrl
	jsonTradeReturn := untils.HttpGetRequest(strUrl, mapParams)
	return jsonTradeReturn
}


	/**
	  * 获取Market Detail 24小时成交量数据
	  *
	  * @param symbol
	  *            如"BTC_CW"表示BTC当周合约，"BTC_NW"表示BTC次周合约，"BTC_CQ"表示BTC季度合约
	  * @return
      */
func FutureMarketDetail(strSymbol string) string {

	mapParams := make(map[string]string)
	mapParams["symbol"] = strSymbol

	strRequestUrl := "/market/detail"
	strUrl := config.MARKET_URL + strRequestUrl
	jsonMarketDetailReturn := untils.HttpGetRequest(strUrl, mapParams)
	return jsonMarketDetailReturn
}

//------------------------------------------------------------------------------------------
// 公共API



    /**
	 * 获取当前可用合约总持仓量
	 *
	 * @param symbol
	 *            "BTC","ETH"...
	 * @param contractType
	 *            合约类型: this_week:当周 next_week:下周 quarter:季度
	 * @param contract_code
	 *            合约code
	 * @return
	 */
func FutureOpenInterest(symbol ,contractType,contractCode string) string {

	strRequest := "/api/v1/contract_open_interest"
	params :=make(map[string]string)
	params["symbol"] = symbol
	params["contract_type"] = contractType
	params["contractCode"] = contractCode
	jsonAccountsReturn := untils.ApiKeyGet(params, strRequest)

	return jsonAccountsReturn
}
	/**
	 * 获取当前可用合约总持仓量
	 *
	 * @param symbol
	 *            "BTC","ETH"...
	 * @param contractType
	 *            合约类型: this_week:当周 next_week:下周 quarter:季度
	 * @param contract_code
	 *            合约code
	 * @return
	 */
func FuturePriceLimit(symbol ,contractType,contractCode string) string {
	strRequest := "/api/v1/contract_price_limit"
	params :=make(map[string]string)
	params["symbol"] = symbol
	params["contract_type"] = contractType
	params["contractCode"] = contractCode
	jsonAccountsReturn := untils.ApiKeyGet(params, strRequest)
	return jsonAccountsReturn
}
	/**
	 * 获取合约指数
	 *
	 * @param symbol
	 *            "BTC","ETH"...
	 * @return
	 */
func FutureContractIndex(symbol string) string {

	strRequest := "/api/v1/contract_index"
	params :=make(map[string]string)
	params["symbol"] = symbol
	json := untils.ApiKeyGet(params, strRequest)
	return json
}
   /**
	 * 期货行情
	 *
	 * @param symbol
	 *            "BTC","ETH"...
	 * @param contractType
	 *            合约类型: this_week:当周 next_week:下周 quarter:季度
	 * @param contract_code
	 *            合约code
	 * @return
	 */
func FutureContractInfo(symbol ,contractType,contractCode string) string {
	strRequest := "/api/v1/contract_contract_info"
	params :=make(map[string]string)
	params["symbol"] = symbol
	params["contract_type"] = contractType
	params["contractCode"] = contractCode
	jsonAccountsReturn := untils.ApiKeyGet(params, strRequest)
	return jsonAccountsReturn
}
//------------------------------------------------------------------------------------------
// 用户资产API

    /**
	 * 获取订单明细信息
	 *
	 * @param symbol
	 *            "BTC","ETH"...
	 * @param orderId
	 *            订单id
	 * @param pageIndex
	 *            第几页,不填第一页
	 * @param pageSize
	 *            不填默认20，不得多于50
	 */
func FutureContractOrderDetail( symbol,  orderId,  pageIndex,  pageSize, createdAt, orderType string) string {


	strRequest := "/api/v1/contract_order_detail"
	params :=make(map[string]string)
	params["symbol"] = symbol
	params["order_id"] = orderId
	params["created_at"] = createdAt
	params["page_index"] = pageIndex
	params["page_size"] = pageSize
	params["order_type"] = orderType
	json := untils.ApiKeyPost(params, strRequest)


	return json
}
    /**
	 * 获取合约当前未成交委托
	 *
	 * @param symbol
	 *            "BTC","ETH"...
	 * @param pageIndex
	 *            第几页,不填第一页
	 * @param pageSize
	 *            不填默认20，不得多于50
	 */
func FutureContractOpenorders( symbol,  pageIndex,  pageSize string) string {


	strRequest := "/api/v1/contract_openorders"
	params :=make(map[string]string)
	params["symbol"] = symbol

	params["page_index"] = pageIndex
	params["page_size"] = pageSize
	jsonAccountsReturn := untils.ApiKeyPost(params, strRequest)


	return jsonAccountsReturn
}
	/**
	 * 获取合约历史委托
	 *
	 * @param symbol
	 *            "BTC","ETH"...
	 * @param tradeType
	 *            0:全部,1:买入开多,2: 卖出开空,3: 买入平空,4: 卖出平多,5: 卖出强平,6: 买入强平,7:交割平多,8: 交割平空
	 * @param type
	 *            1:所有订单，2：已结束订单
	 * @param status
	 *            0:全部,3:未成交, 4: 部分成交,5: 部分成交已撤单,6: 全部成交,7:已撤单 createDate
	 *            7，90（7天或者90天）
	 * @param pageIndex
	 *            第几页,不填第一页
	 * @param pageSize
	 *            不填默认20，不得多于50
	 */
func FutureContractHisorders( symbol,  tradeType, ordertype,  status,
 createDate,  pageIndex,  pageSize string) string {

	strRequest := "/api/v1/contract_hisorders"
	params :=make(map[string]string)
	params["symbol"] = symbol
	params["trade_type"] =tradeType
	params["type"] = ordertype
	params["status"] =status
	params["create_date"] = createDate
	params["page_index"] = pageIndex
	params["page_size"] = pageSize

	jsonAccountsReturn := untils.ApiKeyPost(params, strRequest)


	return jsonAccountsReturn
}
	/**
	 * 获取合约订单信息
	 *
	 * @param orderId
	 *            订单ID（ 多个订单ID中间以","分隔,一次最多允许撤消50个订单 ）
	 * @param clientOrderId
	 *            客户订单ID(多个订单ID中间以","分隔,一次最多允许撤消50个订单)
	 */
func FutureContractOrderInfo( orderId,  clientOrderId, symbol,order_type string) string {


	strRequest := "/api/v1/contract_order_info"
	params :=make(map[string]string)

	params["order_id"] = orderId
	params["client_order_id"] = clientOrderId
	params["symbol"] = symbol
	params["order_type"] = order_type

	jsonAccountsReturn := untils.ApiKeyPost(params, strRequest)


	return jsonAccountsReturn
}
	/**
	 * 全部撤单
	 *
	 * @param symbol
	 *            品种代码，如"BTC","ETH"...
	 */
func FutureContractCancelall(symbol string) string {


	strRequest := "/api/v1/contract_cancelall"
	params :=make(map[string]string)

	params["symbol"] = symbol
	json := untils.ApiKeyPost(params, strRequest)


	return json
}

     /**
	 * 撤销订单
	 *
	 * @param orderId
	 *            订单ID（ 多个订单ID中间以","分隔,一次最多允许撤消50个订单 ）
	 * @return clientOrderId 客户订单ID(多个订单ID中间以","分隔,一次最多允许撤消50个订单)
	 */
func FutureContractCancel(orderId,symbol ,client_order_id string) string {

	strRequest := "/api/v1/contract_cancel"
	params :=make(map[string]string)
	params["symbol"] = symbol
	params["order_id"] = orderId
	params["client_order_id "] = client_order_id
	jsonAccountsReturn := untils.ApiKeyPost(params, strRequest)


	return jsonAccountsReturn
}
    /**
	 * 批量下单
	 *
	 * @param orders_data   订单数组切片

	 * @return
	 */
func FutureContractBatchorder(orders_data [] *Order) string {

	strRequest := "/api/v1/contract_batchorder"
	params :=make(map[string]interface{})
	params["orders_data"] = orders_data
	json := untils.ApiKeyPostBatchorder(params, strRequest)

	return json
}
    /**
	 * 合约下单
	 *
	 * @param symbol
	 *            "BTC","ETH"...
	 * @param contractType
	 *            合约类型: this_week:当周 next_week:下周 month:当月 quarter:季度
	 * @param contractCode
	 *            BTC1403
	 * @param client_order_id
	 *            客户自己填写和维护，这次一定要大于上一次
	 * @param price
	 *            价格
	 * @param volume
	 *            委托数量(张)
	 * @param direction
	 *            "buy":买 "sell":卖
	 * @param offset
	 *            "open":开 "close":平
	 * @param leverRate
	 *            杠杆倍数[“开仓”若有10倍多单，就不能再下20倍多单]
	 * @param orderPriceType
	 *            "limit":限价 "opponent":对手价
	 * @return
	 */
func FutureContractOrder(symbol ,contractType,contractCode,clientOrderId,price,volume,direction,offset,leverRate,orderPriceType string) string {

	///v1/contract_contract_info
	strRequest := "/api/v1/contract_order"
	params :=make(map[string]string)
	params["symbol"] = symbol
	params["contract_type"] = contractType
	params["contract_code"] = contractCode
	params["client_order_id"] = clientOrderId
	params["price"] = price
	params["volume"] = volume
	params["direction"] = direction
	params["offset"] = offset
	params["lever_rate"] = leverRate
	params["order_price_type"] = orderPriceType
	json := untils.ApiKeyPost(params, strRequest)


	return json
}

    /**
	 * 获取用户持仓信息
	 *
	 * @param symbol
	 *            "BTC","ETH"...如果缺省，默认返回所有品种
	 * @return size 获取交易记录的数量 [1, 2000]
	 */
func FutureContractPositionInfo(symbol string) string {
	strRequest := "/api/v1/contract_position_info"
	params :=make(map[string]string)
	params["symbol"] = symbol
	json := untils.ApiKeyPost(params, strRequest)

	return json
}

    /**
	 * 获取用户账户信息
	 *
	 * @param symbol
	 *            "BTC","ETH"...如果缺省，默认返回所有品种
	 * @return size 获取交易记录的数量 [1, 2000]
	 */
func FutureContractAccountInfo(symbol string) string {

	strRequest := "/api/v1/contract_account_info"
	params :=make(map[string]string)
	params["symbol"] = symbol
	json := untils.ApiKeyPost(params, strRequest)
	return json
}




