package enum

type RISK_FUND_TYPE string

const (
	// RiskFund
	// 【风险名称】
	//(1)、在先免模式下：只能传以下枚举值：
	//DEPOSIT：代表押金
	//ADVANCE：代表预付款
	//CASH_DEPOSIT：代表保证金
	//示例：若在充电宝场景传入"DEPOSIT"，当用户确认订单后，UI界面会显示“租借充电宝免押金XX元”。
	//(2)、先享模式：只能传“ESTIMATE_ORDER_COST”，代表订单预估费用。
	//详细说明参考先免模式和先享模式产品介绍
	RISK_FUND_TYPE_DEPOSIT             RISK_FUND_TYPE = "DEPOSIT"
	RISK_FUND_TYPE_ADVANCE             RISK_FUND_TYPE = "ADVANCE"
	RISK_FUND_TYPE_CASH_DEPOSIT        RISK_FUND_TYPE = "CASH_DEPOSIT"
	RISK_FUND_TYPE_ESTIMATE_ORDER_COST RISK_FUND_TYPE = "ESTIMATE_ORDER_COST"
)
