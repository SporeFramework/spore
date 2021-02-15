# https://github.com/gogo/protobuf#speed
# go get github.com/gogo/protobuf/protoc-gen-gofast
cd .. && protoc --gofast_out=. chat.proto
