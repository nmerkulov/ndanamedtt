package portDomain

//go:generate bash -c "protolock status --lockdir=$PROTO_DIR --protoroot=$PROTO_DIR && protoc $PROTO_DIR/PortDomainService.proto --go_out=. --go-grpc_out=.  --proto_path $PROTO_DIR"
