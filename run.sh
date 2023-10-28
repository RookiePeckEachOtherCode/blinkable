cd ./cmd/api
go build -o api && ./api

cd ..

cd ./user
sh build && sh ./output/bootstrap.sh

cd ..