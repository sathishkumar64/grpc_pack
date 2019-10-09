#### API Gateway Implementation.
***
-   Define API's & Grpc
-   Implement REST & Swagger
-   Docker & Kubernetes.

##### Add Plugins:
-   go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
-   go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

EX:
-    https://medium.com/@amsokol.com/tutorial-how-to-develop-go-grpc-microservice-with-http-rest-endpoint-middleware-kubernetes-daebb36a97e9
-   https://github.com/grpc-ecosystem/grpc-gateway

-   https://grpc-ecosystem.github.io/grpc-gateway/docs/grpcapiconfiguration.html (Nice)


##### Validate YMAL File.
-   http://www.yamllint.com/

##### Generated commands

protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:. \
  notes.proto


protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --grpc-gateway_out=logtostderr=true,grpc_api_configuration=notes.yaml:. \
  notes.proto


protoc -I/usr/local/include -I. \
     -I$GOPATH/src \
     -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
     --swagger_out=logtostderr=true,grpc_api_configuration=notes.yaml:. \
     notes.proto


#### Watch This
- https://www.youtube.com/watch?v=KyuFeiG3Y60

#### Logger 
- https://dev-journal.in/2019/05/27/adding-uber-go-zap-logger-to-golang-project/ (NICE)