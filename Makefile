.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=. proto/config/healthy.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/config/shared.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/config/text.proto

