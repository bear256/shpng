# Azure Storgae Base
protoc --go_out=. src/azsb/blob.proto

# Agent of Azure Storage Base
go build -o bin/linux/azsb-agent ./src/azsb/agent

# Server of Azure Storage Base
go build -o bin/linux/azsb-server ./src/azsb/server

# Raw Data
protoc --go_out=. src/model/raw.proto