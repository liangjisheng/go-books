package main

// 国密即国家密码局认定的国产密码算法。主要有SM1，SM2，SM3，SM4。密钥长度和分组长度均为128位。
// SM1 为对称加密。其加密强度与AES相当。该算法不公开，调用该算法时，需要通过加密芯片的接口进行调用。
// SM2为非对称加密，基于ECC。该算法已公开。由于该算法基于ECC，故其签名速度与秘钥生成速度都快于RSA。ECC 256位（SM2采用的就是ECC 256位的一种）安全强度比RSA 2048位高，但运算速度快于RSA。
// 国家密码管理局公布的公钥算法，其加密强度为256位
// SM3 消息摘要。可以用MD5作为对比理解。该算法已公开。校验结果为256位。
// SM4 无线局域网标准的分组数据算法。对称加密，密钥长度和分组长度均为128位。

// 由于SM1、SM4加解密的分组大小为128bit，故对消息进行加解密时，若消息长度过长，需要进行分组，要消息长度不足，则要进行填充。
// 国际算法与国密算法分类

// 分组密码算法(DES和SM4)、将明文数据按固定长度进行分组，然后在同一密钥控制下逐组进行加密，
// 公钥密码算法(RSA和SM2)、公开加密算法本身和公开公钥，保存私钥

// 摘要算法(SM3 md5) 这个都比较熟悉，用于数字签名，消息认证，数据完整性，但是sm3安全度比md5高

// 总得来说国密算法的安全度比较高，2010年12月推出，也是国家安全战略，现在银行都要要求国际算法改造，要把国际算法都给去掉
