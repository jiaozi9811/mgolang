#JWT(JSON Web Token)

## token
token的意思是“令牌”，是用户身份的验证方式，最简单的token组成:uid(用户唯一的身份标识)、time(当前时间的时间戳)、sign(签名，由token的前几位+盐以哈希算法压缩成一定长的十六进制字符串，可以防止恶意第三方拼接token请求服务器)。还可以把不变的参数也放进token，避免多次查库

## jwt
jwt包含3个部分
- header:使用的算法和token类型
- payload:使用sub key指定用户id,还可包含其他email,username等信息
- signature:保证jwt的真实性
