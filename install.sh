#!/bin/bash

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Default install directory
INSTALL_DIR="/usr/local/bin"

# ASCII Art
echo -e "${BLUE}"
cat << "EOF"
   ____       __  __
  / ___| ___ |  \/  | ___ _ __
 | |  _ / _ \| |\/| |/ _ \ '_ \
 | |_| | (_) | |  | |  __/ | | |
  \____|\___/|_|  |_|\___|_| |_|

  GoMen CLI Installer
EOF
echo -e "${NC}"

# Check if Go is installed
echo -e "${BLUE}[1/3]${NC} Checking Go installation..."
if ! command -v go &> /dev/null; then
    echo -e "${RED}Error: Go is not installed.${NC}"
    echo "Please install Go 1.18+ from https://golang.org/dl/"
    exit 1
fi

GO_VERSION=$(go version | awk '{print $3}' | sed 's/go//')
echo -e "${GREEN}✓ Go $GO_VERSION detected${NC}\n"

# Build CLI tool
echo -e "${BLUE}[2/3]${NC} Building GoMen CLI..."
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR"

# Get dependencies
go mod tidy 2>/dev/null

if go build -o gomen ./cmd/gomen; then
    echo -e "${GREEN}✓ CLI built successfully${NC}\n"
else
    echo -e "${RED}Error: Failed to build CLI${NC}"
    exit 1
fi

# Install to /usr/local/bin
echo -e "${BLUE}[3/3]${NC} Installing to ${INSTALL_DIR}..."

if [ -w "$INSTALL_DIR" ]; then
    mv gomen "$INSTALL_DIR/gomen"
    chmod +x "$INSTALL_DIR/gomen"
else
    echo -e "${YELLOW}Requires sudo to install to ${INSTALL_DIR}${NC}"
    sudo mv gomen "$INSTALL_DIR/gomen"
    sudo chmod +x "$INSTALL_DIR/gomen"
fi

if [ $? -eq 0 ]; then
    echo -e "${GREEN}✓ Installed to ${INSTALL_DIR}/gomen${NC}\n"
else
    echo -e "${RED}Error: Failed to install to ${INSTALL_DIR}${NC}"
    echo -e "You can manually copy: ${YELLOW}sudo cp gomen ${INSTALL_DIR}/${NC}"
    exit 1
fi

# Verify installation
if command -v gomen &> /dev/null; then
    echo -e "${GREEN}============================================${NC}"
    echo -e "${GREEN}  GoMen CLI installed successfully!${NC}"
    echo -e "${GREEN}============================================${NC}"
    echo ""
    echo -e "Installed version: ${YELLOW}$(gomen version)${NC}"
    echo ""
    echo -e "Usage:"
    echo -e "  ${YELLOW}gomen help${NC}                    Show all commands"
    echo -e "  ${YELLOW}gomen make:controller User${NC}    Create a controller"
    echo -e "  ${YELLOW}gomen make:model User${NC}         Create a model"
    echo -e "  ${YELLOW}gomen make:resource Product${NC}   Create full resource"
    echo ""
    echo -e "${BLUE}Happy coding! 🚀${NC}"
else
    echo -e "${YELLOW}Warning: gomen not found in PATH${NC}"
    echo -e "Make sure ${INSTALL_DIR} is in your PATH"
fi
