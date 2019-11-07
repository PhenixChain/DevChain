# 客户端命令

## 创建钱包账户并开启自动挖矿
```
# 生成随机数种子，建议可以手工生成并保存好，后继可以使用此种子恢复钱包
./cli seed generate -l 0

# 保存种子，并设置钱包密码为test1234 (密码可以自定义，并且牢牢记住，后面解锁钱包时会用到)
./cli seed save -p test1234 -s "上一步中生成的seed"

# 使用我们刚刚设置的密码解锁钱包 (钱包新创建后默认为锁定状态)
./cli wallet unlock -p test1234

# 开启自动挖矿 (确保区块自动增长)
./cli wallet auto_mine -f 1

# 导入私钥 (挖矿地址的私钥)
./cli account import_key -k 3990969DF92A5914F7B71EEB9A4E58D6E255F32BF042FEA5318FC8B3D50EE6E8 -l miner

# 查看当前钱包中的账户列表
./cli account list
```

## 交易相关

### 转账 
```
./cli send coins transfer -a 100 -t <接收地址> -k <发送地址私钥>
```

### 查询
```
./cli account balance -a <钱包地址>
```