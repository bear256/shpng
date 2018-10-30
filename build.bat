:: Azure Storgae Base
protoc --go_out=. src/azsb/blob.proto

:: Agent of Azure Storage Base
go build -o bin/azsb-agent.exe src/azsb/agent

:: Server of Azure Storage Base
go build -o bin/azsb-server.exe src/azsb/server

:: Raw Data
protoc --go_out=. src/model/raw.proto