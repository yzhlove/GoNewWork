#!/bin/zsh

# 创建私钥
# openssl genrsa -out server.key 2048

# 创建公钥/证书
# 一旦我们有了私钥，我们就需要证书。

#生成自签式公共证书

# openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650

#You are about to be asked to enter information that will be incorporated
#into your certificate request.
#What you are about to enter is what is called a Distinguished Name or a DN.
#There are quite a few fields but you can leave some blank
#For some fields there will be a default value,
#If you enter '.', the field will be left blank.
#-----
#Country Name (2 letter code) [AU]:CN
#State or Province Name (full name) [Some-State]:shanghai
#Locality Name (eg, city) []:shanghai
#Organization Name (eg, company) [Internet Widgits Pty Ltd]:
#Organizational Unit Name (eg, section) []:
#Common Name (e.g. server FQDN or YOUR name) []:localhost
#Email Address []:
