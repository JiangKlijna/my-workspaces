#search image
docker search [key]
#download image :latest
docker pull [image-name]:[version]
#list all of image
docker images
#create container by image [command]=/bin/bash
sudo docker run -i -s --name [container-name] [image-name] [command]
#list all of container
docker ps -a
#start container
docker start [container-name]
#restart container
docker restart [container-name]
#connect container
docker attach [container-name]
#stop container
docker stop [container-name]
#remove container
docker rm [container-name]
#remove image
docker rmi [image-name]
#history image
docker histor image
#copy file form container
docker cp [container-name]:[path] [path]
#check the modification of the container file
docker diff [container-name]
#create image by container
docker commit -a "[author]" -m "[message]" [container-name] [image-name]:[version]
#show of all process
docker top [container-name]
#check port
docker port [container-name] [port]
#(pause or unpause) the container all of process
docker [pause|unpause] [container-name]
#save image to tar file
docker save -o [tar-file] [image]:[version]
#load tar image
docker load < [tar-file]
#set image version
docker tag -f [pre-image]:[pre-version] [new-image]:[new-version]
#detail information of docker
docker info
#detail information of (image or container)
docker inspect [container-name]

