#!/bin/bash

# https://www.cnblogs.com/vactor/p/9286425.html
# http://joey771.cn/2019/01/18/ubuntu%E5%AE%89%E8%A3%85OpenCL%E8%BF%90%E8%A1%8C%E5%8F%8A%E7%BC%96%E8%AF%91%E7%8E%AF%E5%A2%83/
# 下载 skd, 但没有用这个
# https://registrationcenter.intel.com/en/products/postregistration/?sn=CD38-399ZZFHK&EmailID=716600596%40qq.com&Sequence=2518794&dnld=t
# https://blog.csdn.net/sinat_23084397/article/details/98871090

sudo apt-get install build-essential libgtk2.0-dev libjpeg-dev libtiff5-dev libjasper-dev libopenexr-dev cmake libeigen3-dev yasm libfaac-dev libtheora-dev libx264-dev libv4l-dev libavcodec-dev libavformat-dev libswscale-dev libv4l-dev ffmpeg

sudo apt-get install clinfo

sudo apt install dkms xz-utils openssl libnuma1 libpciaccess0 bc curl libssl-dev lsb-core libicu-dev
sudo apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys 3FA7E0328081BFF6A14DA29AA6A19B38D3D831EF
echo "deb http://download.mono-project.com/repo/debian wheezy main" | sudo tee /etc/apt/sources.list.d/mono-xamarin.list
sudo apt-get update
sudo apt-get install mono-complete

wget http://registrationcenter-download.intel.com/akdlm/irc_nas/vcp/12526/opencl_runtime_16.1.2_x64_rh_6.4.0.37.tgz
cd opencl_runtime_16.1.2_x64_rh_6.4.0.37/
sudo sh install_GUI.sh

# 上面是安装运行时的一种方式, 下面是安装openCL应用的一种方式, 实际编译 filecoin 的时候, 下面的这种安装方式是管用的

wget http://registrationcenter-download.intel.com/akdlm/irc_nas/vcp/16624/intel_sdk_for_opencl_applications_2020.1.395.tar.gz
tar xzf intel_sdk_for_opencl_applications_2020.1.395.tar.gz
cd intel_sdk_for_opencl_applications_2020.1.395/
sudo ./install.sh
