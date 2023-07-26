# v2ray

[clashX](https://github.com/yichengchen/clashX/releases)

mac 下载 dmg 文件

[github](https://github.com/yanue/v2rayu/releases)

```sh
brew install cask v2rayu
```

MacOS 打开软件出现 ‘xxx‘ “将对您的电脑造成伤害。 您应该将它移到废纸篓。“的解决方式
但尝试过没用

```shell
sudo spctl –master-disable
sudo codesign --force --deep --sign - /Applications/V2rayU.app

sudo xattr -r -d com.apple.quarantine /Applications/V2rayU.app
sudo xattr -cr /Applications/V2rayU.app
```
