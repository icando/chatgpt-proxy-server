RED_COLOR='\E[1;31m'  #红
GREEN_COLOR='\E[1;32m' #绿
YELOW_COLOR='\E[1;33m' #黄
BLUE_COLOR='\E[1;34m'  #蓝
PINK='\E[1;35m'      #粉红
RES='\E[0m'

RUN_NAME="chatgpt-server"

echo -e  "${RED_COLOR}****生产环境****${RES}\n\n"
echo -e  "${GREEN_COLOR}****开始执行自动化部署****${RES}\n\n"

echo -e "${YELOW_COLOR}---step1:环境准备---${RES}"
CURDIR=$(cd $(dirname $0); pwd)
mkdir -p output/bin output/conf
cp conf/* output/conf/

RUNTIME_ROOT=${CURDIR}
export KITEX_RUNTIME_ROOT=$RUNTIME_ROOT
export KITEX_LOG_DIR="$RUNTIME_ROOT/logs"
if [ ! -d "$KITEX_LOG_DIR/app" ]; then
    mkdir -p "$KITEX_LOG_DIR/app"
fi
if [ ! -d "$KITEX_LOG_DIR/rpc" ]; then
    mkdir -p "$KITEX_LOG_DIR/rpc"
fi

git fetch origin master
git reset --hard FETCH_HEAD
git clean -df
echo -e "${BLUE_COLOR}合并代码成功${RES}\n"

echo -e "${BLUE_COLOR}环境准备完成${RES}\n"

echo -e "${YELOW_COLOR}---step2:编译---${RES}"
go build -o output/bin/${RUN_NAME}
echo -e "${BLUE_COLOR}编译完成${RES}\n"

echo -e "${YELOW_COLOR}---step3:杀掉进程并且运行---${RES}"
i1=`ps -ef|grep -E ${RUN_NAME}|grep -v grep|awk '{print $2}'`
if [ -z $i1 ];then
  nohup output/bin/${RUN_NAME} >/dev/null 2>&1 &
else
  echo -e "${BLUE_COLOR}杀掉进程$i1${RES}\n"
  kill -9 $i1 && nohup output/bin/${RUN_NAME} >/dev/null 2>&1 &
fi

i2=`ps -ef|grep -E ${RUN_NAME}|grep -v grep|awk '{print $2}'`
echo -e "${GREEN_COLOR}****部署成功,部署的进程ID为:$i2${RES}****"