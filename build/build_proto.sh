# https://github.com/gogo/protobuf#speed
# go get github.com/gogo/protobuf/protoc-gen-gofast
# cd .. && protoc --gofast_out=. spore.proto

cd ../protocol && protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    spore.proto
