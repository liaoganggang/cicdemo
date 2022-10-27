#!/bin/bash
#author:mrl

#$1git下载源码目录，$2源码地址，$3源码版本tag，$4对应的是checkout_sha,即commit ID
#$5项目名称 $6项目名称空间即用户组
Time=`date '+%Y/%m/%d%H%M%S'`
CloneDir=$1
CodeAddr=$2
CodeTag=$3
CommitID=$4
ProjectName=$5
ProjectNamespace=$6

PDIR=`cd $(dirname "$0");pwd`
cd $PDIR

Deploy() {
    
}