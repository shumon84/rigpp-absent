PROJECT   = $(shell basename $(CURDIR))
BUILD_DIR = build
BUILD_CMD = go build -o $(BUILD_DIR)/$(PROJECT)

WINDOWS = windows
LINUX   = linux
MAC     = darwin

X86 = 386
X64 = amd64

$(BUILD_DIR)/$(PROJECT):
	$(BUILD_CMD)

.PHONY: all
all: windows-amd64 windows-386 mac-amd64 mac-386 linux-amd64 linux-386

$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)

.PHONY: windows-386 windows-amd64
windows-386:
	GOOS=$(WINDOWS) GOARCH=$(X86) $(BUILD_CMD)-$(WINDOWS)-$(X86).exe
windows-amd64:
	GOOS=$(WINDOWS) GOARCH=$(X64) $(BUILD_CMD)-$(WINDOWS)-$(X64).exe

.PHONY: mac-386 mac-amd64
mac-386:
	GOOS=$(MAC) GOARCH=$(X86) $(BUILD_CMD)-$(MAC)-$(X86)
mac-amd64:
	GOOS=$(MAC) GOARCH=$(X64) $(BUILD_CMD)-$(MAC)-$(X64)

.PHONY: linux-386 linux-amd64
linux-386:
	GOOS=$(LINUX) GOARCH=$(X86) $(BUILD_CMD)-$(LINUX)-$(X86)
linux-amd64:
	GOOS=$(LINUX) GOARCH=$(X64) $(BUILD_CMD)-$(LINUX)-$(X64)
