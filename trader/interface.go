package trader

// Trade 结构体，用于标准化表示一笔已成交的交易
type Trade struct {
	Symbol    string
	ID        int64
	OrderID   int64
	Price     float64
	Quantity  float64
	Fee       float64
	FeeAsset  string
	Time      int64 // Unix apoch in milliseconds
	IsBuyer   bool
	IsMaker   bool
	Side      string // "BUY" or "SELL"
	IsReduce  bool   // 是否为只减仓
}

// Trader 交易接口
type Trader interface {
	// GetBalance 获取账户余额
	GetBalance() (map[string]interface{}, error)

	// GetPositions 获取所有持仓
	GetPositions() ([]map[string]interface{}, error)

	// OpenLong 开多仓
	OpenLong(symbol string, quantity float64, leverage int) (map[string]interface{}, error)

	// OpenShort 开空仓
	OpenShort(symbol string, quantity float64, leverage int) (map[string]interface{}, error)

	// CloseLong 平多仓（quantity=0表示全部平仓）
	CloseLong(symbol string, quantity float64) (map[string]interface{}, error)

	// CloseShort 平空仓（quantity=0表示全部平仓）
	CloseShort(symbol string, quantity float64) (map[string]interface{}, error)

	// SetLeverage 设置杠杆
	SetLeverage(symbol string, leverage int) error

	// SetMarginMode 设置仓位模式 (true=全仓, false=逐仓)
	SetMarginMode(symbol string, isCrossMargin bool) error

	// GetMarketPrice 获取市场价格
	GetMarketPrice(symbol string) (float64, error)

	// SetStopLoss 设置止损单
	SetStopLoss(symbol string, positionSide string, quantity, stopPrice float64) error

	// SetTakeProfit 设置止盈单
	SetTakeProfit(symbol string, positionSide string, quantity, takeProfitPrice float64) error

	// CancelAllOrders 取消该币种的所有挂单
	CancelAllOrders(symbol string) error

	// FormatQuantity 格式化数量到正确的精度
	FormatQuantity(symbol string, quantity float64) (string, error)

	// GetTradeHistory 方法，用于获取交易所的真实成交历史
	GetTradeHistory(symbol string, startTime int64) ([]Trade, error)
}
