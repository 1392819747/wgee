#!/bin/bash

# 检查输入参数是否为空
if [ -z "$1" ]; then
    echo "Usage: $0 <os>"
    echo "  <os> - Specify the target operating system (linux, darwin, windows), or 'all' to compile for all platforms"
    exit 1
fi

# 设置目标操作系统
OS=$1

# 确定输出目录
OUTPUT_DIR="./bin"

# 创建输出目录（如果不存在）
mkdir -p $OUTPUT_DIR

# 检查是否编译全部平台
if [ "$OS" == "all" ]; then
    echo "Compiling for all platforms..."
    
    # 编译 Linux 平台
    echo "Compiling for linux and arm64..."
    export GOOS=linux
    export GOARCH=arm64
    go build -o $OUTPUT_DIR/wgee-linux-arm64 wgee.go

    echo "Compiling for linux and amd64..."
    export GOARCH=amd64
    go build -o $OUTPUT_DIR/wgee-linux-amd64 wgee.go

    # 编译 macOS 平台
    echo "Compiling for darwin and arm64..."
    export GOOS=darwin
    go build -o $OUTPUT_DIR/wgee-darwin-arm64 wgee.go

    echo "Compiling for darwin and amd64..."
    export GOARCH=amd64
    go build -o $OUTPUT_DIR/wgee-darwin-amd64 wgee.go

    # 编译 Windows 平台
    echo "Compiling for windows and amd64..."
    export GOOS=windows
    export GOARCH=amd64
    go build -o $OUTPUT_DIR/wgee-windows-amd64.exe wgee.go

    echo "Compilation complete for all platforms."
    exit 0
fi

# 检查操作系统参数是否有效
case "$OS" in
    linux)
        export GOOS=linux
        ;;
    darwin)
        export GOOS=darwin
        ;;
    windows)
        export GOOS=windows
        ;;
    *)
        echo "Unsupported operating system: $OS"
        exit 1
        ;;
esac

# 根据操作系统确定输出文件后缀
OUTPUT_EXT=""
if [ "$GOOS" == "windows" ]; then
    OUTPUT_EXT=".exe"
fi

# 编译 ARM 架构
echo "Compiling for $OS and arm64..."
export GOARCH=arm64
go build -o $OUTPUT_DIR/wgee-$OS-arm64$OUTPUT_EXT wgee.go

# 编译 AMD64 架构
echo "Compiling for $OS and amd64..."
export GOARCH=amd64
go build -o $OUTPUT_DIR/wgee-$OS-amd64$OUTPUT_EXT wgee.go

echo "Compilation complete."

# 移动文件到输出目录
echo "Moving binaries to $OUTPUT_DIR..."
mv wgee-$OS-* $OUTPUT_DIR/

echo "All binaries moved to $OUTPUT_DIR."

