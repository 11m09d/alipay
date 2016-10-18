package alipay

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

var initQRCodeParamMap = map[string]bool{
	"app_id":         true,  //  支付宝分配给开发者的应用ID. 不可空
	"method":         true,  //  接口名称。	不可空	alipay.trade.precreate
	"format":         false, //  仅支持JSON。	可空 JSON
	"charset":        true,  //  请求使用的编码格式，如utf-8,gbk,gb2312等。	utf-8
	"sign_type":      false, //  商户生成签名字符串所使用的签名算法类型，目前支持RSA。	不可空 RSA
	"sign":           false, //  商户请求参数的签名串，详见签名。	不可空	https://doc.open.alipay.com/doc2/detail.htm?treeId=200&articleId=105351&docType=1
	"timestamp":      true,  //  发送请求的时间，格式"yyyy-MM-dd HH:mm:ss"。	不可空	2014-07-24 03:07:50
	"version":        true,  //  调用的接口版本，固定为：1.0	不可空	1.0
	"notify_url":     false, //  支付宝服务器主动通知商户服务器里指定的页面http/https路径。	可空	http://api.test.alipay.net/atinterface/receive_notify.htm
	"app_auth_token": false, //
	"biz_content":    true,  //  请求参数的集合，最大长度不限，除公共参数外所有请求参数都必须放在这个参数中传递，具体参照各产品快速接入文档。	不可空
}

type QRCodePaymentNotify struct {
	NotifyTime       string  //  notify_time	通知时间	Date	通知的发送时间。格式为yyyy-MM-dd HH:mm:ss。	不可空	2013-08-22 14:45:24
	NotifyType       string  //  notify_type	通知类型	String	通知的类型。	不可空	trade_status_sync
	NotifyID         string  //  notify_id	通知校验ID	String	通知校验ID。	不可空	64ce1b6ab92d00ede0ee56ade98fdf2f4c
	SignType         string  //  sign_type	签名方式	String	固定取值为RSA。	不可空	RSA
	Sign             string  //  sign	签名	String	请参见签名机制。	不可空	lBBK%2F0w5LOajrMrji7DUgEqNjIhQbidR13GovA5r3TgIbNqv231yC1NksLdw%2Ba3JnfHXoXuet6XNNHtn7VE%2BeCoRO1O%2BR1KugLrQEZMtG5jmJI
	OutTradeNo       string  //  out_trade_no	商户网站唯一订单号	String(64)	对应商户网站的订单系统中的唯一订单号，非支付宝交易号。需保证在商户网站中的唯一性。是请求时对应的参数，原样返回。	可空	082215222612710
	Subject          string  //  subject	商品名称	String(128)	商品的标题/交易标题/订单标题/订单关键字等。它在支付宝的交易明细中排在第一列，对于财务对账尤为重要。是请求时对应的参数，原样通知回来。	可空	测试
	PaymentType      string  //  payment_type	支付类型	String(4)	支付类型。默认值为：1（商品购买）。	可空	1
	TradeNo          string  //  trade_no	支付宝交易号	String(64)	该交易在支付宝系统中的交易流水号。最短16位，最长64位。	不可空	2013082244524842
	TradeStatus      string  //  trade_status	交易状态	String	交易状态，取值范围请参见“交易状态”。	不可空	TRADE_SUCCESS
	SellerID         string  //  seller_id	卖家支付宝用户号	String(30)	卖家支付宝账号对应的支付宝唯一用户号。以2088开头的纯16位数字。	不可空	2088501624816263
	SellerEmail      string  //  seller_email	卖家支付宝账号	String(100)	卖家支付宝账号，可以是email和手机号码。	不可空	xxx@alipay.com
	BuyerID          string  //  buyer_id	买家支付宝用户号	String(30)	买家支付宝账号对应的支付宝唯一用户号。以2088开头的纯16位数字。	不可空	2088602315385429
	BuyerEmail       string  //  buyer_email	买家支付宝账号	String(100)	买家支付宝账号，可以是Email或手机号码。	不可空	dlwdgl@gmail.com
	TotalFee         float64 //  total_fee	交易金额	Number	该笔订单的总金额。请求时对应的参数，原样通知回来。	不可空	1.00
	Quantity         int64   //  quantity	购买数量	Number	购买数量，固定取值为1（请求时使用的是total_fee）。	可空	1
	Price            float64 //  price	商品单价	Number	price等于total_fee（请求时使用的是total_fee）。	可空	1.00
	Body             string  //  body	商品描述	String(512)	该笔订单的备注、描述、明细等。对应请求时的body参数，原样通知回来。	可空	测试测试
	GMTCreate        string  //  gmt_create	交易创建时间	Date	该笔交易创建的时间。格式为yyyy-MM-dd HH:mm:ss。	可空	2013-08-22 14:45:23
	GMTPayment       string  //  gmt_payment	交易付款时间	Date	该笔交易的买家付款时间。格式为yyyy-MM-dd HH:mm:ss。	可空	2013-08-22 14:45:24
	IsTotalFeeAdjust string  //  is_total_fee_adjust	是否调整总价	String(1)	该交易是否调整过价格。	可空	N
	UseCoupon        string  //  use_coupon	是否使用红包买家	String(1)	是否在交易过程中使用了红包。	可空	N
	Discount         float64 //  discount	折扣	String	支付宝系统会把discount的值加到交易金额上，如果有折扣，本参数为负数，单位为元。	可空	0.00
	RefundStatus     string  //  refund_status	退款状态	String	取值范围请参见“退款状态”。	可空	REFUND_SUCCESS
	GMTRefund        string  //  gmt_refund	退款时间	Date	卖家退款的时间，退款通知时会发送。格式为yyyy-MM-dd HH:mm:ss。	可空	2008-10-29 19:38:25
}

func (a *Alipay) QRCodePayment(outTradeNo, subject string, totalFee float64, notifyURL string, extraParams map[string]string) (s string, err error) {
	if outTradeNo == "" {
		err = fmt.Errorf("%s out_trade_no : Required parameter missing", LogPrefix)
		return
	}

	if subject == "" {
		err = fmt.Errorf("%s subject is required parameter", LogPrefix)
		return
	}

	if notifyURL == "" {
		err = fmt.Errorf("%s notify_url is required parameter", LogPrefix)
		return
	}

	if totalFee == 0 {
		err = fmt.Errorf("%s total_fee is required parameter", LogPrefix)
		return
	}

	if a.privateKey == nil {
		err = fmt.Errorf("%s rsa private key is not init", LogPrefix)
		return
	}

	params := a.initQRCodeParams(outTradeNo, subject, notifyURL, totalFee, extraParams)
	kvs, err := GenKVpairs(initQRCodeParamMap, params, "sign", "sign_type")
	if err != nil {
		return
	}

	for i, kv := range kvs {
		kvs[i] = KVpair{K: kv.K, V: fmt.Sprintf(`%s`, kv.V)}
	}

	var sig string
	sig, err = a.rsaSign(kvs)
	if err != nil {
		return
	}

	kvs = append(kvs, KVpair{K: "sign", V: fmt.Sprintf(`%s`, url.QueryEscape(sig))})
	//kvs = append(kvs, KVpair{K: "sign_type", V: `RSA`})

	s = kvs.Join("&")
	return
}

func (a *Alipay) initQRCodeParams(outTradeNo, subject, notifyURL string, totalFee float64) (params map[string]string) {
	var t int64 = time.Now().Unix()
	params = make(map[string]string)
	biz_content := "{'out_trade_no':'%s','total_amount':'%f','subject':'%s',}"
	params["app_id"] = a.partner
	params["method"] = "alipay.trade.precreate"
	params["format"] = "json"

	params["charset"] = "utf-8"
	params["sign_type"] = "RSA"

	params["timestamp"] = time.Unix(t, 0).Format("2006-01-02 15:04:05")

	//params["notify_url"] = notifyURL
	params["version"] = "1.0"
	params["biz_content"] = fmt.Sprintf(biz_content, outTradeNo, totalFee, "Test")
	return
}

func (a *Alipay) QRCodePaymentNotify(req *http.Request) (result *QRCodePaymentNotify, err error) {
	vals, err := parsePostData(req)
	if err != nil {
		return
	}

	if len(vals) == 0 {
		err = ErrNotifyDataIsEmpty
		return
	}

	var fields = []string{
		"notify_time",
		"notify_type",
		"notify_id",
		"sign_type",
		"sign",
		"out_trade_no",
		"subject",
		"payment_type",
		"trade_no",
		"trade_status",
		"seller_id",
		"seller_email",
		"buyer_id",
		"buyer_email",
		"total_fee",
		"quantity",
		"price",
		"body",
		"gmt_create",
		"gmt_payment",
		"is_total_fee_adjust",
		"use_coupon",
		"discount",
		"refund_status",
		"gmt_refund",
		"gmt_close",
	}

	err = a.rsaVerify(vals, fields)
	if err != nil {
		return
	}

	var price, totalFee, discount float64
	price, _ = strconv.ParseFloat(vals.Get("price"), 64)
	totalFee, _ = strconv.ParseFloat(vals.Get("total_fee"), 64)
	discount, _ = strconv.ParseFloat(vals.Get("discount"), 64)

	var quantity int64
	quantity, _ = strconv.ParseInt(vals.Get("quantity"), 10, 64)

	result = &QRCodePaymentNotify{
		NotifyTime:       vals.Get("notify_time"),
		NotifyType:       vals.Get("notify_type"),
		NotifyID:         vals.Get("notify_id"),
		SignType:         vals.Get("sign_type"),
		Sign:             vals.Get("sign"),
		OutTradeNo:       vals.Get("out_trade_no"),
		Subject:          vals.Get("subject"),
		PaymentType:      vals.Get("payment_type"),
		TradeNo:          vals.Get("trade_no"),
		TradeStatus:      vals.Get("trade_status"),
		SellerID:         vals.Get("seller_id"),
		SellerEmail:      vals.Get("seller_email"),
		BuyerID:          vals.Get("buyer_id"),
		BuyerEmail:       vals.Get("buyer_email"),
		TotalFee:         totalFee,
		Quantity:         quantity,
		Price:            price,
		Body:             vals.Get("body"),
		GMTCreate:        vals.Get("gmt_create"),
		GMTPayment:       vals.Get("gmt_payment"),
		IsTotalFeeAdjust: vals.Get("is_total_fee_adjust"),
		UseCoupon:        vals.Get("use_coupon"),
		Discount:         discount,
		RefundStatus:     vals.Get("refund_status"),
		GMTRefund:        vals.Get("gmt_refund"),
	}

	return
}
