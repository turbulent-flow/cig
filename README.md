## 项目简介
该项目的全称是 Customer Is God，旨在维持或者提高高端模玩藏品的稀缺性。它是由 web 3.0 技术实现，目前只做了一个简单的原型。合约文件的地址是：`contract/CIG.sol`。


## 测试
### 1. 安装 Ganache
Ganache 是个人版的以太坊区块链，可以在本地开发、调试智能合约项目。官方的地址为：[https://archive.trufflesuite.com/ganache/](https://archive.trufflesuite.com/ganache/)。

### 2. 配置环境变量
在项目根目录放置`.env.test`，文件内容如下：
```text
ETHEREUM_ENDPOINT=http://127.0.0.1:[YOUR_PORT]
PRIVATE_KEY=[YOUR_PRIVATE_KEY]
WALLET_ADDRESS=[YOUR_WALLET_ADDRESS]
GAS_LIMIT=3000000
CONTRACT_ADDRESS=[YOUR_CONTRACT_ADDRESS]
```
启动 Ganache UI 后，可以找到`ETHEREUM_ENDPOINT`、`PRIVATE_KEY`以及`WALLET_ADDRESS`。`CONTRACT_ADDRESS`可以先不填，因为合约地址是在合约部署后生成的，等到测试具体的行为时，再把部署时得到的合约地址填进环境变量中。

### 3. 编译合约文件
使用 solc 编译合约文件，得到 abi 和字节码。solc 是 Solidity 的源码编译器。如果没有安装，可以使用`brew`安装（如果使用的是 Mac 开发的话）；安装后，可以运行如下命令：
```shell
solc --abi --bin ./contract/CIG.sol --evm-version paris -o ./build/contract
```
注意：这里使用了参数`--evm-version paris`，特意指定了 evm 的版本为 paris，这是因为 Ganache GUI 2.7.1 不支持最新的 Solidity 的编译器，需要降级兼容。

### 4. 生成 go 版的合约文件
使用 abigen 工具并结合上一步得到的`build/contract/CIG.bin`与`build/contract/CIG.abi`，生成 go 版的合约文件。
```shell
abigen --bin=build/contract/CIG.bin --abi=build/contract/CIG.abi --pkg=contract --out=./internal/pkg/contract/cig.go
```

### 5. 运行测试
```shell
go test -v ./...
```
注意：运行`test/internal/pkg/deploy`，可以得到合约地址，需要把它填进环境变量。
