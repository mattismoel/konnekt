#!/bin/sh

CWD=$(pwd)
echo "Building from $CWD..."

SRV_USERNAME="mattiskristensen"
SRV_HOSTNAME="knnkt.dk"
SRV_SRC_DIR="/home/$SRV_USERNAME/website"

ORIGIN="https://$SRV_HOSTNAME"

echo "Change diretory to $CWD/frontend..."
cd $CWD/frontend

echo "Building frontend with base $ORIGIN..."
vite build --base=$ORIGIN

DST=$SRV_USERNAME@$SRV_HOSTNAME:$SRV_SRC_DIR
echo "Copying build output to remote server at $DST"
scp -r dist $DST


echo "Executing Docker compose at remote..."
ssh -t $SRV_USERNAME@$SRV_HOSTNAME "\
	cd ./website &&\
	sudo docker compose down &&\
	sudo docker compose up --remove-orphans --pull always -d"

echo "Frontend sucessfully deployed at remote."
