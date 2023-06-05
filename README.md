# 说明

官方文档： https://vercel.com/docs/concepts/functions/serverless-functions

这是一个部署在 vercel 上的 serverless，使用 go 进行编写的，可以获取 chatgpt 上 oauth 认证的一些数据。

功能：

1. 通过登录跳转地址，获取到 code
2. 通过 code，获取到 access_token | refresh_token
3. 使用 refresh_token 刷新 access_token 的过期时间
4. 撤销 refresh_token, 使 refresh_token 失效


所有的接口都在 api 目录下。

部署地区：regions ：sin1 新加坡。

## 关于部署

自己使用 vercel 部署，不用担心数据泄露问题。

1. fork 项目
2. 在 vercel 中导入项目即可。

## 访问

vercel 部署完成后会生成一个域名，或者可以自定义一个域名，直接访问就可以看到操作 token 的页面了。

## 致谢

感谢 [pengzhile](https://github.com/pengzhile) 大佬开源的 [pandora](https://github.com/pengzhile/pandora) 项目，以及 [该博客](https://zhile.io/2023/05/19/how-to-get-chatgpt-access-token-via-pkce.html) 中提及到的操作 token 的方式。
