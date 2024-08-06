#! /bin/bash

tar -zcf /home/john/Code/untruth-www/backup/$(date +%F-%H_%M_%S).tar.gz /srv/http /home/john/Code/untruth-www/backend

echo "Backup created"

cd frontend 
yarn build
sudo cp -r ./dist/* /srv/http/untruth
rm -rf ./dist
sudo chown -R http:http /srv/http
# sudo chmod -R 777 /home/john/Code/untruth-www/www     
echo "Fronend deployed"

pwd
cd ../backend/
go build -o /home/john/Code/untruth-www/backend
sudo systemctl restart untruth-backend
echo "Backend deployed"

cd ..
