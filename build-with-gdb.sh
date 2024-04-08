go build -gcflags="all=-N -l " -o bin/orchestrator go/cmd/orchestrator/main.go

# 创建输出目录
rm -rf orchestrator
rm -rf orchestrator-3.2.6.tar.gz

mkdir orchestrator
cp -rf resources orchestrator/
mkdir orchestrator/bin/
cp bin/orchestrator orchestrator/bin/


tar -czvf orchestrator-3.2.6.tar.gz orchestrator


