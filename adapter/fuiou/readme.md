### 富友支付技术对接群
有问题一定要艾特技术支持，因为群太多，望理解，艾特之后技术支持能快速响应！！！
前置接口测试地址：http://fundwx.fuiou.com/doc/#/aggregatePay/
前置接口测试参数：http://fundwx.fuiou.com/doc/#/aggregatePay/introduction?id=%e6%b5%8b%e8%af%95%e7%8e%af%e5%a2%83%e5%8f%82%e6%95%b0
前置接口验收流程：http://fundwx.fuiou.com/doc/#/aggregatePay/check

注意！注意！注意！
商户对接接口传值除了接口字段描述的之外，其他的尽量传英文或者数字，方便测试。如果需要用到2.6接口，测试环境可以使用富友测试APPID（富友的测试APPID   wxfa089da95020ba1a   wx5ac8eb4651fe544f），然后调用2.16接口获取openid即可。

1.请进群的合作方修改群名片，谢谢～合作方名称-对接的富友bd-某某某。例如，富友支付-李四-张三，以便群内交流
2.温馨提示：反馈问题请务必艾特@技术支持，以免漏掉消息，谢谢配合！
3.本群回复技术对接问题，运营及风控问题可咨询各自的富友业务对接人或富友运营人员。反馈问题尽量艾特对应的技术人员，以免漏掉信息，谢谢配合！
4.富友常用邮箱，请惠存~
商户组邮箱：fuiou_sh@fuioupay.com
终端组邮箱：fuiou_zd@fuioupay.com
风控邮箱：fu_fk@fuioupay.com
对接组邮箱：techsupport@fuioupay.com

推送优先级：商户>二级机构>一级机构，最多推送一种类型

接口文档：https://doc.weixin.qq.com/doc/w3_AAIAzgYZAE8SBEcd1ORTVGuB9Pke6?scode=AMoAIAeJAA03DEYQCoAD0AzgYZAE8
接口参数：https://doc.weixin.qq.com/sheet/e3_AD0AzgYZAE8cACLvEnUQKyGx0uLTN?scode=AMoAIAeJAA0sR0YTJAAD0AzgYZAE8&tab=mndixd

```text
// 密钥相关
    商户号 "0002900F0370542";
    测试地址：https://richOperationFront-test.fuioupay.com
    商户公钥： "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCTK/DZ0Ve73u2ORRAYrpv07FNHyuTb87xabGMTwEIQV2PhDdAYtiRIO/dAtZ45PUN1N+rtiQQxwIgyKJpYIesFpCbBZ+3YIVf3wlkl9VVSfnUSDWcteN9n0WifBqrKbzJ3gaXi4wXveCMJViqTfgDkfgTV/EC/7h5nwj5VUF6LPQIDAQAB";
    商户私钥： "MIICdQIBADANBgkqhkiG9w0BAQEFAASCAl8wggJbAgEAAoGBAJMr8NnRV7ve7Y5FEBium/TsU0fK5NvzvFpsYxPAQhBXY+EN0Bi2JEg790C1njk9Q3U36u2JBDHAiDIomlgh6wWkJsFn7dghV/fCWSX1VVJ+dRINZy1432fRaJ8GqspvMneBpeLjBe94IwlWKpN+AOR+BNX8QL/uHmfCPlVQXos9AgMBAAECgYAzqbMs434m50UBMmFKKNF6kxNRGnpodBFktLO7FTybu/HF6TFp21a1PMe5IYhfk5AAsBZ6OCUOygWFhhdYZN+5W+dweF3kp1rLE4y5CjwqNlk/g22TAndf9znh/ltHFLvITToqu/eh/34tE1gyNxRbsi1olw/1wv8ZRjM3vtM9QQJBANvNwFq+CJHUyFzkXQB7+ycQFnY8wDq8Uw2Hv9ZMjgIntH7FSlJtdu5mAYPPo6f74slO5tFUMNP7EVppqsjYaNkCQQCraD6iKHo+OIlvvYIKiMXatJGD7N1GNhq5CrhUNPWLHwv/Ih2D3JJdF8IUZOPIJfUxTfM2fZYI+EVdsv6s4RcFAkAGjNYbnighOGcUJZYD6q3sVxVkRqEv3ubWs2HrH/Lna4l8caKqXCq8JfwLkod8/QugFiLYwBqIZqX4vMdjHtfZAkBsAl9dbWZCaPvpxp/4JWGPxDLhz9NLV/KU4bVvkoObq++yUHwKyGYOdVcd5MlIKOsNq5Hzp0Vw14lWVuF2bMxFAkBuNrZksvUULNIaWDKd4rQ6GVzUxXuIZW0ZE6atHYDiXPB4jVAjKRtLxZAV1qH9cr1zNJlcg+RbGYUdF9t4A9n5";
    富友私钥： "MIICdgIBADANBgkqhkiG9w0BAQEFAASCAmAwggJcAgEAAoGBAIaDw5Gb27ezyjrS1yAPtmwv/RZHm9tHSSuCY6SmbfK3WGAGhgAGhSGItNO3pXo3hEmD1graS70neFM0+aIZt9pShtxiz/1EszVMLD3HIkRKhBJFcq8p0gWH+l0Bpl0wGxlf4u2mTkf490YEiP+W/EyEnksQsoIdHpsn+/rZ20CnAgMBAAECgYB1r65ZJJ10+Y3DLVgdquGVgd7RsVEA5jt0H54CHcIwCoz9ZneyagHsNujOGuxiI1RP5VJNKHP/SBsT4VNOqWWaDoVn97thBr+JYCZX0SMdNznSEPMuZu6GKNSeuDl5alKZMB6CumIJmjl6JlDUeUturqbv5XwW1FaiNtrWE7x4sQJBAMEkNvj1hyKtopoLsOca5ydpR/vLOfQgPXGlKHRcOJfj45T3bSYCZ6fHc1t6IpvFUVqYElY1pLA8JYPc9PQaePMCQQCySv+mZfkAz3f2V0U8aw9nfbks4vnUi9XS3e+V593NaJ6IVX3M6sSXAHWJFgezEt7LdB+d54taMpKVeyUDpYZ9AkB/BjBZYDF2LzhHk/TOqbTpCKbdBPWihymh+ns2vAhEbQ6aRHg2jVJa2CQYP6VPSWCN8oHszO75MTWDGejIOjjdAkAKJb6bJ96eLzCyspDcOXOs/jjV1y1E7ZiD4eHK9GFpWXT8aXE5gnsh5QLLhJd3l7Fafwd1o0IJJiu1mkanCHq5AkEAwMfKw/GCAj8i+c7hHTccO0d69fC+fGDakb5zzTpzwNi/6VVIKYG6idOmnHqknGrj9h0UR3sKdE0iXvSBuHHBdg==";
    富友公钥： "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCGg8ORm9u3s8o60tcgD7ZsL/0WR5vbR0krgmOkpm3yt1hgBoYABoUhiLTTt6V6N4RJg9YK2ku9J3hTNPmiGbfaUobcYs/9RLM1TCw9xyJESoQSRXKvKdIFh/pdAaZdMBsZX+Ltpk5H+PdGBIj/lvxMhJ5LELKCHR6bJ/v62dtApwIDAQAB";

	


订单分账： 
20250301/030013436970
20250301/030013436973

可用入账方编号:
037054220230224000049871826
037054220230810000050292874
037054220230904000050379438


SFTP测试参数
机构名：0002900F0370542
机构代码：0002900F0370542
账户类型：FGJ
主机：ftp-1.fuioupay.com
端口：9022
协议：sftp
用户名：FGJ370542
密码：BnjiKm3CVI7rm1Dx


被归集商户号：
0002900F0096235
可归集订单：
20250730/500003860855
20250402/500003394652


正式：https://richfront.fuioupay.com
商户正式商户号
商户走对接流程下发正式环境参数，交换双方公钥
```
````text
富友公钥（分账类）：
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC4nlZ0Uk4/hPlmGUtrqdMCzrZ9b3Tt/zSaoH1Bc/9WN5N1C1eVDARtWYKAVbw1eLE9hCrT0GLzOKxK2YEVSfdUmhRo49gOz19jvAqpXCeTsylaxeXbaF+3ylziQ+XBtg4a8f9rLp1kmmNmumqBgENv2dJhvId+dpZjuGU1ZO/MGwIDAQAB
-----END PUBLIC KEY-----
平台私钥（分账类）：
-----BEGIN RSA PRIVATE KEY-----
MIICdwIBADANBgkqhkiG9w0BAQEFAASCAmEwggJdAgEAAoGBAL1XipPCUPVYHpVGDz1bLYsjdPv5Kd6teuU6y3Fk+82LNas1wQ1oxRgel8X2yNulKyG3e7m281hDOir5b+bQmSZNSYkK3tRFtSDLoAxRa6hdWTBT6USy0VeOwgLqySJNHX7VEThaKsB5OK9Q/HhSQsMcFiky499JX03plGqulgL3AgMBAAECgYBLxOZtrssbA0Jp1flvQjd9gJjl5JW+yIlvhhG3tQOXA2hctUwmA5Erz8JItDM4wmX6PiTC8tV6spxqahs/sKY4HuIP896BEMyeQPMAJaeoPkcIE21CvRNQMix7tBD1xT8PQv42lqXuG2ekcvDx6TQFQR8/sG7PbP7hM0tXXrIxQQJBAO4XHyIZXOT6dwEZTj0gPnZMj/Lb61RmnOykBWuIsBZ1Cu0Qxrrs3Fi61NGMtdvn7MpOmC06+PzdyjX0ZVe4vfECQQDLla2TRwJductB0RfqeqPOpzP8L28YHZAyNPgzel8PT9Wsy3arhDRGEzjxR2DoCSFuahCWAUyjZehWyi7nEAdnAkA/olZj2ruFR5v+4zCSDcgj/hqLIlQnXrWaWcxJDWNi3S1qZw12cFAENwsiZqVEfRxAhfkOPbDGhNDC0uszIuFBAkEAstLQ2pL/ExDF5xQhK8d552WbFiMcNFSTemZotd8BbNO1XsiBfnrr57muxNDr4CVVBkWIOBsAFG7JPKLvu+qqdQJBAJBPcmXHRlqsIPkRCA4rWHrPGdhnQEKAoiYA1ctnCWtJe+IhFyhFEDYzQV+AYQIylKq3mWzA8QNCcijqGUp8U84=
-----END RSA PRIVATE KEY-----

````




````text
互联网扫码接口：

----------生产环境----------
获取openid     			  https://hlwnets.fuioupay.com/aggpos/getOpenid.fuiou  
订单支付接口   			  https://hlwnets.fuioupay.com/aggpos/order.fuiou  
订单支付被扫接口          https://hlwnets.fuioupay.com/aggpos/orderScan.fuiou
支付查询接口   			  https://hlwnets.fuioupay.com/aggpos/orderQuery.fuiou  
订单关闭接口   			  https://hlwnets.fuioupay.com/close.fuiou
订单手续费查询接口   	  https://hlwnets.fuioupay.com/aggposFeeQry.fuiou 
银联小程序签约接口        https://hlwnets.fuioupay.com/union/sign.fuiou  
银联APP签约接口       	  https://hlwnets.fuioupay.com/union/appSign.fuiou
签约查询接口    		  https://hlwnets.fuioupay.com/union/query.fuiou
银联APP签约支付接口       https://hlwnets.fuioupay.com/union/appSignPay.fuiou
银联APP签约支付查询接口   https://hlwnets.fuioupay.com/union/appSignPayQuery.fuiou
银联解约接口   			  https://hlwnets.fuioupay.com/union/cancelSign.fuiou
订单退款接口   			  https://refund-transfer.fuioupay.com/refund_transfer/aggposRefund.fuiou
退款查询接口   			  https://refund-transfer.fuioupay.com/refund_transfer/aggposRefundQuery.fuiou
分佣回退接口              https://hlwnets.fuioupay.com/allocate/allocateRefund.fuiou
分佣回退查询接口          https://hlwnets.fuioupay.com/allocate/refundQuery.fuiou
分佣接口                  https://hlwnets.fuioupay.com/allocate/allocateQuery.fuiou
分佣协议电签接口	      https://hlwsign.fuioupay.com/account/create.fuiou

----------测试环境----------
获取openid     			  https://hlwnets-test.fuioupay.com/aggpos/getOpenid.fuiou  
订单支付接口   			  https://hlwnets-test.fuioupay.com/aggpos/order.fuiou  
订单支付被扫接口          https://hlwnets-test.fuioupay.com/aggpos/orderScan.fuiou
支付查询接口   			  https://hlwnets-test.fuioupay.com/aggpos/orderQuery.fuiou  
订单关闭接口   			  https://hlwnets-test.fuioupay.com/close.fuiou
订单手续费查询接口   	  https://hlwnets-test.fuioupay.com/aggposFeeQry.fuiou
银联小程序签约接口        https://hlwnets-test.fuioupay.com/union/sign.fuiou  
银联APP签约接口       	  https://hlwnets-test.fuioupay.com/union/appSign.fuiou
银联签约查询接口          https://hlwnets-test.fuioupay.com/union/query.fuiou
银联APP签约支付接口       https://hlwnets-test.fuioupay.com/union/appSignPay.fuiou
银联APP签约支付查询接口   https://hlwnets-test.fuioupay.com/union/appSignPayQuery.fuiou
银联解约接口   			  https://hlwnets-test.fuioupay.com/union/cancelSign.fuiou
订单退款接口   			  https://refund-transfer-test.fuioupay.com/refund_transfer/aggposRefund.fuiou
退款查询接口   			  https://refund-transfer-test.fuioupay.com/refund_transfer/aggposRefundQuery.fuiou 
分佣回退接口              https://hlwnets-test.fuioupay.com/allocate/allocateRefund.fuiou
分佣回退查询接口          https://hlwnets-test.fuioupay.com/allocate/refundQuery.fuiou
分佣接口                  https://hlwnets-test.fuioupay.com/allocate/allocateQuery.fuiou
分佣协议电签接口	      https://hlwsign-test.fuioupay.com/account/create.fuiou

测试商户代码
0001000F0040992

测试商户私钥
MIICdQIBADANBgkqhkiG9w0BAQEFAASCAl8wggJbAgEAAoGBAJMr8NnRV7ve7Y5FEBium/TsU0fK5NvzvFpsYxPAQhBXY+EN0Bi2JEg790C1njk9Q3U36u2JBDHAiDIomlgh6wWkJsFn7dghV/fCWSX1VVJ+dRINZy1432fRaJ8GqspvMneBpeLjBe94IwlWKpN+AOR+BNX8QL/uHmfCPlVQXos9AgMBAAECgYAzqbMs434m50UBMmFKKNF6kxNRGnpodBFktLO7FTybu/HF6TFp21a1PMe5IYhfk5AAsBZ6OCUOygWFhhdYZN+5W+dweF3kp1rLE4y5CjwqNlk/g22TAndf9znh/ltHFLvITToqu/eh/34tE1gyNxRbsi1olw/1wv8ZRjM3vtM9QQJBANvNwFq+CJHUyFzkXQB7+ycQFnY8wDq8Uw2Hv9ZMjgIntH7FSlJtdu5mAYPPo6f74slO5tFUMNP7EVppqsjYaNkCQQCraD6iKHo+OIlvvYIKiMXatJGD7N1GNhq5CrhUNPWLHwv/Ih2D3JJdF8IUZOPIJfUxTfM2fZYI+EVdsv6s4RcFAkAGjNYbnighOGcUJZYD6q3sVxVkRqEv3ubWs2HrH/Lna4l8caKqXCq8JfwLkod8/QugFiLYwBqIZqX4vMdjHtfZAkBsAl9dbWZCaPvpxp/4JWGPxDLhz9NLV/KU4bVvkoObq++yUHwKyGYOdVcd5MlIKOsNq5Hzp0Vw14lWVuF2bMxFAkBuNrZksvUULNIaWDKd4rQ6GVzUxXuIZW0ZE6atHYDiXPB4jVAjKRtLxZAV1qH9cr1zNJlcg+RbGYUdF9t4A9n5
测试商户公钥
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCTK/DZ0Ve73u2ORRAYrpv07FNHyuTb87xabGMTwEIQV2PhDdAYtiRIO/dAtZ45PUN1N+rtiQQxwIgyKJpYIesFpCbBZ+3YIVf3wlkl9VVSfnUSDWcteN9n0WifBqrKbzJ3gaXi4wXveCMJViqTfgDkfgTV/EC/7h5nwj5VUF6LPQIDAQAB
测试富友公钥
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCKPD4f/3xMjPuHcQSGxPIYZNgF2i0sJhzmKpN3dmzhbKH/1nG+aXUJDhswyitzrI+U0ic/GL/IWB1wQ3noWuLFr8jDSGafenTFiY9c0H9ZAEfqW/oywx95G5JWu5y/ffp4yCmlt8k5kHO/4kE1qnJcGaQlb6/+7t3MPSV5ybmBZwIDAQAB

秘钥说明：双方交换公钥，用富友公钥加密报文发送给富友。收到富友响应报文后，用商户私钥解密响应报文。

备注：测试环境只能用支付宝1分钱测试流程(支付。退款。查询。回调) , 微信要生产环境测试。
生产环境参数：看接口文档的最下面生产参数。


小程序/公众号/扫码支付接口文档：	http://180.168.100.158:13318/fuiouWposApipay/
````