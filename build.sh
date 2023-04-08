# --push -t fjiabinc/chat-proxy:latest 这里换成自己的镜像名
#错误处理:ERROR: multiple platforms feature is currently not supported for docker driver. Please switch to a different driver (eg. "docker buildx create --use")
#解决方法:docker buildx create --name mybuilder --use
docker buildx build --platform linux/amd64,linux/arm64 --push -t fjiabinc/chat-proxy:latest .