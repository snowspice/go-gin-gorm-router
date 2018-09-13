version=`git symbolic-ref --short -q HEAD`

image=`basename $PWD`:$version

# 上传本地私服
aliyunImage=harbornode.snowspice.com:5000/$image

docker rmi --force aliyunImage &> /dev/null

docker build -t $aliyunImage .
