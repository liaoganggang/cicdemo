#!/bin/bash
#author:mrl

#$1git下载源码目录，$2源码地址，$3源码版本tag，$4对应的是checkout_sha,即commit ID
#$5项目名称 $6项目名称空间即用户组
Time=`date '+%Y%m%d%H%M%S'`
CloneDir=$1
CodeAddr=$2
CodeTag=$3
CommitID=$4
ProjectName=$5
ProjectNamespace=$6

PDIR=`cd $(dirname "$0");pwd`
cd $PDIR


#下载代码
download() {
    cd $CloneDir && git clone $CodeAddr $CodeTag
    cd ${CloneDir}/${CodeTag} && git checkout -b $CodeTag && git reset --hard $CommitID 
}

#build
#构建tar包，还要考虑包配置文件
build_package(){
    BuildDir=${CloneDir}/${CodeTag}
    AppName=${ProjectName}
    NameSpace=${ProjectNamespace}
    cd ${BuildDir}
    go build .  #需要安装go
    mv ${CodeTag} ${AppName}  #go程序名为项目名
    cp ${AppName} ../dockerfile/web/
    cd $PDIR
    docker build --build-arg AppName=${AppName} -t ${AppName}:${CodeTag} ../dockerfile/web/
    docker tag ${AppName}:${CodeTag} 192.168.2.4/devel/${AppName}:${CodeTag}

    #远程登录验证,${url}为docker仓库地址，需写到配置文件
    #如果没有expect命令则需要安装
    (/usr/bin/expect -c  "
        set timeout 300
        spawn docker login ${url}   

        expect {
            \"Username:\" { send \"$username\r\";exp_continue }
            \"Password:\" { send \"$password\r\"}
        }
        expect eof
    ") &> /dev/null

    docker push 192.168.2.4/devel/${AppName}:${CodeTag} #上传至harbor仓库

    if [ $? -eq 0 ];then
        exit 0
    else
        echo "faild"
    fi

}

download && build_package



