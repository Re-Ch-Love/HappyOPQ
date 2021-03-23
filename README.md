# HappyOPQ

![Badge](https://img.shields.io/badge/OneBot-v11-black)

<details>
<summary>目录</summary>
<p>

- [背景](#背景)

</p>
</details>

## 背景

[OneBot](https://github.com/howmanybots/onebot) 是一个有强大生态的 QQ 机器人协议。

[OPQBot](https://github.com/OPQBOT/OPQ) 是一个稳定性高、功能强大的 QQ 机器人框架。

然而令人可惜的是，OPQBot 并不使用 OneBot 协议。

## 动机

使得 OPQBot 用户也能使用到 OneBot 的强大生态，增强开发效率，尽量避免“重复造轮子”的问题。

## 预期

OneBot 的生态可以完美融入 OPQBot 。并在保证这一点的同时尽可能地使 OneBot 的用户能够无缝切换使用到 OPQBot。

## 使用方法

HappyOPQ 提倡开箱即用，只需运行可执行文件即可使用。

*注意：目前处于开发状态，暂不提供可执行文件，如有需要请自行编译*

在运行时使用参数`-c`可指定配置文件。

HappyOPQ 将会按以下逻辑加载配置文件

如果指定了配置文件，则使用指定的配置文件，如果指定的配置文件不存在，将会报错并结束程序。

如果没有指定，则使用工作目录下的`HappyOPQ.yml`文件，如果该文件不存在，将会使用默认配置。

*Q：为什么不设计成前一种失败自动尝试后一种？*

*A：因为这样很有可能导致用户想使用自定义配置（但不小心输入错误的路径），而 HappyOPQ 使用默认配置的非预期行为*

配置文件&默认配置：

```yaml
OPQBot:
  Host: "127.0.0.1"
  Port: 8080
OneBot:
  HTTP:
    Enabled: true
    URL: "http://127.0.0.1:8081"
  PositiveWebSocket:
    Enabled: false
    Host: "127.0.0.1"
    Port: 8082
  ReverseWebSocket:
    Enabled: false
    Host: "127.0.0.1"
    Port: 8083
```

*注意：HappyOPQ 仍在初期开发时期，任何改动都不会预先通知，**请勿用于生产场景！***

## 鸣谢

HappyOPQ的诞生离不开本项目以及下述依赖包的每一位贡献者！

<https://github.com/graarh/golang-socketio>

<https://gopkg.in/yaml.v2>