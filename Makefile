# Makefile

# 设置交叉编译的目标操作系统和体系结构
TARGET_OS_WINDOWS = windows
TARGET_ARCH_WINDOWS = amd64
TARGET_OS_MACOS = darwin
TARGET_ARCH_MACOS = amd64

# 设置输出目录和可执行文件名
OUTPUT_DIR = bin
EXECUTABLE_WINDOWS = $(OUTPUT_DIR)/laike.exe
EXECUTABLE_MACOS = $(OUTPUT_DIR)/laike

# 默认目标：编译 Windows 平台可执行文件
all: windows

# 编译 Windows 平台可执行文件
windows:
	@echo "=== 编译 Windows 平台可执行文件 ==="
	@echo "设置环境变量..."
	@SET GOOS=$(TARGET_OS_WINDOWS)
	@SET GOARCH=$(TARGET_ARCH_WINDOWS)
	@echo "开始编译..."
	go build -o $(EXECUTABLE_WINDOWS)

# 编译 macOS 平台可执行文件
macos:
	@echo "=== 编译 macOS 平台可执行文件 ==="
	@echo "设置环境变量..."
	@set GOOS=$(TARGET_OS_MACOS)
	@set GOARCH=$(TARGET_ARCH_MACOS)
	@echo "开始编译..."
	go build -o $(EXECUTABLE_MACOS)

# 清理生成的可执行文件和目录
clean:
	@echo "=== 清理 ==="
	@echo "删除输出目录和可执行文件..."
	del /F /Q $(EXECUTABLE_WINDOWS)
	del /F /Q $(EXECUTABLE_MACOS)
	rmdir /Q /S $(OUTPUT_DIR)
