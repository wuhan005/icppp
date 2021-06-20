# icppp

一个帮我应对 ICP 备案页面检查的小玩意儿

## 为何要写这个项目

前阵子备案的时候，被告知自己之前备案的好几个域名处于无法访问状态。实际原因是自己只用了那些域名的部分二级域名，（例如 `cdn.xxx.xxx`），并没有去管 `www` 的解析。

同时之前也有接到过服务商打来的电话向我确认备案域名网站的访问情况。

一般遇到这些问题，我都是匆匆忙忙找个 Bootstrap 的模板站放上去，想着不如一步到位，所以便有了这么个项目哈哈。

## 效果展示

![效果展示](https://s2.ax1x.com/2019/12/21/QjCW11.md.png)

## 开始使用

### Step 1:

在 `/home/app/conf` 路径下创建配置文件 `icppp.toml`（路径也可自定义）。

输入如下配置，其中 ICP 为你网站的域名以及备案信息。

```toml
[[icp]]
url="abc.xyz"
no="粤ICP备88888888号-1"

[[icp]]
url="qaz.wsx"
no="粤ICP备88888888号-2"
```

### Step 2:

```sh
docker run -dt --name icppp -p 9315:9315 -v /home/app/conf:/home/app/conf wuhan005/icppp
```

**注意 `-v` 文件夹的映射路径**

Enjoy it!

## Caddyfile 配置

若您是使用 Caddy 进行反向代理，可以参考使用如下配置：

```caddyfile
# Caddy 1
www.abc.xyz {
    gzip
    proxy / http://127.0.0.1:9315 {
    	header_upstream Host abc.xyz
    }
    tls abc@xyz.com
}
```
