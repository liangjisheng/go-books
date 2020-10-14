#!/bin/bash

# 首先安装 Command Line Tools 
# 选择自己系统版本相应的Command Line Tools，下载安装就ok了
# https://developer.apple.com/download/more/

# 然后安装 HomeBrew
# https://zhuanlan.zhihu.com/p/111014448

# 安装脚本
/bin/zsh -c "$(curl -fsSL https://gitee.com/cunkai/HomebrewCN/raw/master/Homebrew.sh)"

# 卸载脚本
/bin/zsh -c "$(curl -fsSL https://gitee.com/cunkai/HomebrewCN/raw/master/HomebrewUninstall.sh)"


