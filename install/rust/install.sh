#!/bin/bash

# https://gocode.cc/project/20/article/147
# https://learnku.com/rust/wikis/29017
# https://www.ixiqin.com/2019/02/macos-installation-rust-development-environment/

# 如果可以翻墙的话
curl https://sh.rustup.rs -sSf | sh
source $HOME/.cargo/env
# echo 'export PATH=$HOME/.cargo/bin:$PATH' >> ~/.bashrc
# source ~/.bashrc
# 验证
rustc --version

# 如果不能翻墙的话
# 使用中科大的代理
# Rust Toolchain 反向代理：https://mirrors.ustc.edu.cn/help/rust-static.html
# Rust Crates 源使用帮助：https://mirrors.ustc.edu.cn/help/crates.io-index.html
# Rust 官网：https://www.rust-lang.org/learn/get-started

# import USTC mirror
echo "export RUSTUP_DIST_SERVER=https://mirrors.ustc.edu.cn/rust-static" >> ~/.bashrc
echo "export RUSTUP_UPDATE_ROOT=https://mirrors.ustc.edu.cn/rust-static/rustup" >> ~/.bashrc
source ~/.bashrc

curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh

# 安装完毕后刷新环境变量
# echo 'export PATH=$HOME/.cargo/bin:$PATH' >> ~/.bashrc
# source ~/.bashrc
source $HOME/.cargo/env

# 安装rust后, 可以执行 rustc -V  和 cargo -V 看看是否正常输出版本

# 修改Rust Crates 源
cat >~/.cargo/config <<EOF
[source.crates-io]
registry = "https://github.com/rust-lang/crates.io-index"
replace-with = 'ustc'

[source.ustc]
registry = "git://mirrors.ustc.edu.cn/crates.io-index"
EOF

# 因为代码提示racer的一些#future功能不能在稳定版使用，无法安装
# 所以安装nightly版本的rust 并设置默认版本
rustup install nightly
rustup default nightly

# 安装RLS组件
rustup component add rls --toolchain nightly
rustup component add rust-analysis --toolchain nightly
rustup component add rust-src --toolchain nightly

# 安装racer 代码的自动补齐
cargo install racer
