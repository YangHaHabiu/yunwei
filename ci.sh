#!/bin/bash
MY_PATH=$(cd $(dirname $0) && pwd)
VERSION=""
REMARK=""
OLD_VERSION=""
##最大数值
MAX_NUM=100

get_tag(){
    OLD_VERSION=$(git tag |grep ^${1}|sort -t . -k1,1r -k2,2nr -k3,3nr|sed -n 1p)
    FIRST=$(echo $OLD_VERSION|awk -F '[.t|v]' '{print $2}')
    SECOND=$(echo $OLD_VERSION|awk -F '[.t|v]' '{print $3}')
    THIRD=$(echo $OLD_VERSION|awk -F '[.t|v]' '{print $4}')
    if [[ "$THIRD" -ge  "${MAX_NUM}" ]];then
        if [[ "$SECOND" -ge  "${MAX_NUM}" ]];then
            FIRST=$(($FIRST+1))
            SECOND=0
        else
            SECOND=$(($SECOND+1))
        fi
        THIRD=0
    else
        THIRD=$(($THIRD+1))
    fi
    VERSION="${1}${FIRST}.${SECOND}.${THIRD}"
    REMARK="${REMARK}${VERSION}"
	echo $VERSION
	echo $REMARK
	git add .
	git commit -m "$REMARK"
	git push origin main
	git tag -a $VERSION -m "$REMARK"
	git push origin $VERSION
}

crete_swagger_json(){
  for i in admin yunwei intranet monitor sendMsg;do
    json_file="$i"
    file_path="./service/$i/api/"
    goctl api plugin -plugin goctl-swagger="swagger -filename ${json_file}.json" -api ${file_path}${json_file}.api -dir ${file_path}
  done
}

if [[ $1 == "" ]];then
	echo "请输入参数t[测试]或者v[正式]"
	exit 0
fi

if [[ $1 == "t" ]];then
	REMARK="test_"
	crete_swagger_json
	get_tag $1
elif [[ $1 == "v" ]];then
	REMARK="stable_"
	crete_swagger_json
	get_tag $1
else
	echo "参数错误"
	exit 1
fi