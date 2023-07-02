# 说明

官方文档： https://vercel.com/docs/concepts/functions/serverless-functions

这是一个部署在 vercel 上的 serverless，使用 go，js，或者其他语言来写。

所有的接口都要在 api 目录下。

| 接口 | 说明 |
|---------|---------|
| / | 官网直登获取 Token 功能页面 |
| /third | 第三方账号获取 Token 页面 |
| /getOpenAItoken | 官网直登获取 Token 接口 |
| /getOpenAItokenByCode | 通过 Code 获取 Token 接口 |
| /revokeOpenAIToken | 撤销 refresh_token 接口 |
| /refreshOpenAIToken | 刷新 Token 接口 |


部署地区：regions ：sin1 新加坡


如果要部署官网直登的api接口，需要在 vercel 上设置环境变量，key 是 LoginURL,值就是自己部署的 pandora 项目域名，注意不要用 cf 的安全模式，不然没法直接调用，因为安全问题需要手动验证。


Vercel
[![Deploy with Vercel](https://vercel.com/button)](https://vercel.com/new/clone?repository-url=https%3A%2F%2Fgithub.com%2Fayuayue%2Fchatgpt_token_vercel)