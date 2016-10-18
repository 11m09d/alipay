# alipay
支付宝即时到帐，支付宝移动支付 Golang实现
++增加支付宝扫码支付
## 快速开始
### 获取安装
    go get -u github.com/11m09d/alipay

### 推荐使用localtunnel测试回调通知
可以先安装一个[localtunnel](https://localtunnel.github.io/www/)
可以方便快捷的实现你的本地web服务通过外网访问，无需修改DNS和防火墙设置

```console
$ npm install -g localtunnel
```

## 示例

#### 通过localtunnel获取外网地址:

```console
$ lt --port 9090
your url is: http://eygytquvvu.localtunnel.me
```

#### 修改示例代码中的配置:
记得修改示例中的对应的partner, key, email配置,
如果需要使用app支付记得添加public key path和private key path

```golang
var (
	partner = "your pid"
	key     = "your key"
	email   = "your email"

	publicKeyPath  = "your rsa pubKey path" // "支付宝的公共密钥xxx/rsa_public_key.pem"
	privateKeyPath = "your rsa priKey path" // "个人的私有密钥xxx/rsa_private_key.pem"

	a = alipay.NewPayment(partner, key, email)
	// app 支付需要加入rsa公钥密钥
	// a.InitRSA(publicKeyPath, privateKeyPath)

    // 示例监听的端口
	port = ":9090"

    // 通过 lt --port 9090 获取的外网地址
    localTunnel = "http://eygytquvvu.localtunnel.me"
    ...
)
```

#### 启动示例程序:

```console
$ go run example/main.go
```

#### 在浏览器中访问本地服务:
[http://localhost:9090/index](http://localhost:9090/index)

具体如何使用请查看[example/main.go](https://github.com/11m09d/alipay/blob/master/example/main.go)
