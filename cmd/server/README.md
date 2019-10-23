# 启动主链节点

## 一、多节点配置(单节点无需更改，直接启动主链)
[1、singleMode](https://github.com/PhenixChain/devchain/blob/master/cmd/server/main.toml#L35)
修改为：singleMode=false

[2、seeds](https://github.com/PhenixChain/devchain/blob/master/cmd/server/main.toml#L44)
修改为实际ip，例如：
seeds=["192.168.0.1:13802","192.168.0.2:13802","192.168.0.3:13802"]

## 二、启动主链
```
./server -f main.toml
```
如要部署在docker里运行：[参考](https://github.com/PhenixChain/devchain/tree/master/docker)

## 三、导入挖矿账户私钥，开启自动挖矿
[开启挖矿](https://github.com/PhenixChain/devchain/blob/master/cmd/cli/README.md#L3)


备注：main.toml为主链配置文件；para.toml为平行链配置文件