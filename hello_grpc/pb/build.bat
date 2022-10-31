protoc --proto_path=./api \
       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./api \
 	       --go-http_out=paths=source_relative:./api \
 	       --go-grpc_out=paths=source_relative:./api \
       --openapi_out=fq_schema_naming=true,default_response=false:. \
       api/realworld/v1/realworld.proto api/realworld/v1/error_reason.proto
