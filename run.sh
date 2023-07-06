#!/bin/bash

MY_PATH=$(cd $(dirname $0) && pwd)

do_exit(){
    echo -e "\e[0;31;1m$1\e[0m\n"
    exit $2
}

do_p(){
        [ $# -eq 2 ] && echo -e "\e[0;$2;1m$1\e[0m" || echo -e "\e[0;32;1m$1\e[0m"
}

update_server(){
    do_p "开始更新"
    rsync -avzP --exclude-from='./exclude.list' /data/tmp_data/ywadminv3/ /data/ywadminv3/ 
    do_p "更新完成"
}

processlist(){
    ps -ef | egrep "${1}" | egrep -v "${EXCLUDE}" | awk '{printf("%6s %6s %6s %4s %6s %6s %10s\t",$1,$2,$3,$4,$5,$6,$7)}{for (i=8;i<=NF;i++) {printf(" %s",$i)} printf("\n")}'

}


operate(){
    case $1 in
    start)
        [ $(ps -ef |grep "$SERVER_FILE" |egrep -v "${EXCLUDE}" |wc -l) -eq 1 ] && { do_p "$SERVER_FILE 已启动,无需重复操作";exit 0; }
        
        cd $SERVER_PATH 
        if [ "${DEV_MODE}" == "1" ];then
            nohup go run $SERVER_FILE -f $SERVER_CONF  >/dev/null 2>&1 &
        else
            nohup $SERVER_FILE -f $SERVER_CONF  >/dev/null 2>&1 &
        fi
        processlist "${SERVER_FILE}|${SERVER_CONF}"
        cd $MY_PATH
    ;;
    stop)
        ps aux |egrep "$SERVER_FILE|${SERVER_CONF}" |egrep -v "${EXCLUDE}"|awk '{print $2}' |xargs --no-run-if-empty kill || do_exit "停止${a}-${m}服务失败" 1
    ;;
    status)
        processlist "${SERVER_FILE}|${SERVER_CONF}"
    ;;
    esac


}

action_server(){
    ACTIONS=$1
    
    [ ${ACTIONS} == "status" ] && printf "%6s %6s %6s %4s %6s %6s %10s %16s" "UID" "PID" "PPID" "C" "STIME" "TTY" "TIME" "CMD"
    echo
    for TMP in $SERVICES
    do
        [ ${ACTIONS} == "status" ] || do_p "开始执行[$TMP] ${ACTIONS}操作"
        SERVER_NAME=$(echo $TMP|awk -F - '{print $1}')
        SERVER_OPER=$(echo $TMP|awk -F - '{print $2}')
        if [ "${DEV_MODE}" == "1" ];then
            SERVER_PATH="$MY_PATH/service/$SERVER_NAME/$SERVER_OPER"
            SERVER_FILE="$SERVER_PATH/${SERVER_NAME}.go"
        else
            SERVER_PATH="$MY_PATH/$SERVER_NAME/$SERVER_OPER"
            SERVER_FILE="$SERVER_PATH/${SERVER_NAME}-${SERVER_OPER}"
        fi

        SERVER_CONF="$SERVER_PATH/etc/${SERVER_NAME}.yaml"
        operate $ACTIONS
        [ ${ACTIONS} == "status" ] || do_p "执行[$TMP] ${ACTIONS}操作完成"
    done

}

handle_services(){
for var in ${ARRAY[@]};do
    START_SERVER=(${var//:/ })
    for m in ${START_SERVER[@]};do
        if [ "$m" != "${START_SERVER}" ];then
            SER=${START_SERVER}-${m}
            ALL_SERVICES+=" $SER"
        fi
    done
done
}

HELP(){
do_p "USAGE: sh $0 start|status|stop|restart|update [OPTS] [DEV_MODE]
OPTS:
$(echo -e "${ALL_SERVICES}"|tr '[ ]+' ' \n' |awk '{print "\t"$1}')
DEV_MODE:
1:开发者启动
空:正式服启动
" 32
exit 1;
}


STRING="identity:rpc,admin:rpc:api,yunwei:rpc:api,ws:api,qqGroupv3:api,intranet:rpc:api"
EXCLUDE="grep|/root/go/bin"
ARRAY=(${STRING//,/ })
ALL_SERVICES=""
SERVICES=$2
DEV_MODE=$3
handle_services
if [ -n "$SERVICES" ] ;then
    [ $(echo "$ALL_SERVICES"|grep $SERVICES|wc -l) -eq 0 ] && HELP
else
    SERVICES=$ALL_SERVICES
fi

##主程序开始
case $1 in
    start)
        action_server start
    ;;
    stop)
        action_server stop
    ;;
    status)
        action_server status
    ;;
    restart)
        action_server stop
        action_server start
    ;;
    update)
        action_server stop
        update_server
        action_server start
    ;;
    *)
        HELP
    ;;
esac
