# mcsm-instance-controller
## 功能
使用qq群控制mcsm的实例，群友把mc服务器搞崩后再也不用@狗管理上线开机啦，可以实现自助开机关机

## 配置文件
程序启动时会加载配置文件，可以使用`-config`参数指定配置文件，默认为config.json。
配置模板和注释如下：
```jsonc
{
    "mcsm_endpoint": "http://MCSM服务器地址:23333",
    "mcsm_api_key": "MCSM API KEY", //mcsm apikey
    "instance_list": [
        {
            "instance_name": "创造", // 服务器名称
            "instance_uuid": "远程/本地实例标识符",
            "node_uuid": "远程节点标识符"
        },
        {
            "instance_name": "生存", // 可以为一个qq群配置多个服务器
            "instance_uuid": "远程/本地实例标识符",
            "node_uuid": "远程节点标识符"
        }
    ],
    "cq_endpoint": "ws://go-cqhttp的ws地址+端口/", // 运行一个go-cqhttp，并创建一个ws接口
    "target_qq_group": 111222333, // qq群号，程序会忽略不是这个群的消息和私聊消息
    "admin_qq_list": [
        {  // 管理员的qq号，和权限等级。 此项可以通过 setuser 命令自助配置并保存。
            "id": 1122334455,
            "level": 999
        }
    ],
    "prefix": "/mcsm", // 命令前缀。只有在命令前缀为/mcsm 才会响应。程序也会相应/help命令，输出帮助信息。
    "function_level": [
        { 
            "function": "start", // 指令名称。  开启服务器。
            "level": 0           // 指令名称对应的权限。只有在用户权限大于等于此值时才有权限执行此命令。
        },
        {
            "function": "restart", // 重启服务器
            "level": 15
        },
        {
            "function": "stop", // 停止服务器
            "level": 20
        },
        {
            "function": "kill", // 强行停止服务器
            "level": 20
        },
        {
            "function": "list", // 列出已经配置的服务器
            "level": 0
        },
        {
            "function": "setuser", // 设置其它用户的权限
            "level": 999
        },
        {
            "function": "listuser", // 列出已设置权限的用户
            "level": 0
        }
    ]
}
```

## 快速开始
### 创建一个go-cqhttp ws服务器
在go-cqhttp的配置文件中，需要有ws的配置，示例如下。正常配置并启动go-cqhttp后，go-cqhttp会监听ws端口。在下面的示例中，会监听5790端口。
```
...

# 连接服务列表
servers:
  # 添加方式，同一连接方式可添加多个，具体配置说明请查看文档
  #- http: # http 通信
  #- ws:   # 正向 Websocket
  #- ws-reverse: # 反向 Websocket
  #- pprof: #性能分析服务器
  - ws:
      # 正向WS服务器监听地址
      address: 0.0.0.0:5790
      middlewares:
        <<: *default # 引用默认中间件

...
```
### 运行二进制文件
下载适合你平台的二进制文件，填好config.json，就可以运行了。当配置文件是当前目录下的config.json时，可以不加参数。

windows示例： 

`./mcsm-instance-controller-windows-amd64.exe -config config.json`

linux示例：

`./mcsm-instance-controller-linux-amd64 -config config.json`