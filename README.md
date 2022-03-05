# Avatar Proxy

使用Express框架和Axios代理各大站点的头像

## Docker 部署

```bash
docker pull ghcr.io/sourceglobalcdn/gravatar-proxy:master
docker run -p 9000:9000 -d ghcr.io/sourceglobalcdn/gravatar-proxy:master
```

你可以通过更改环境变量`PORT`的方法来改变监听端口，但我更推荐你直接改变Docker映射的端口。
