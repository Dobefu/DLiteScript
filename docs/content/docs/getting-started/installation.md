+++
title = 'Installation'
description = 'Install DLiteScript on Windows, macOS, and Linux. Download prebuilt binaries or build it from source with step-by-step instructions.'
weight = 0
draft = false
+++

## Binary Downloads

Download the latest release for your platform from [GitHub releases](https://github.com/Dobefu/DLiteScript/releases/latest):

- **macOS (Apple Silicon)**: `dlitescript-darwin-arm64`
- **macOS (Intel x64)**: `dlitescript-darwin-x64`
- **Linux (ARM32)**: `dlitescript-linux-arm`
- **Linux (ARM64)**: `dlitescript-linux-arm64`
- **Linux (x64)**: `dlitescript-linux-x64`
- **Windows (ARM32)**: `dlitescript-win32-arm`
- **Windows (ARM64)**: `dlitescript-win32-arm64`
- **Windows (x64)**: `dlitescript-win32-x64`

### Installation Steps

1. Download the binary for your platform

2. Make the file executable (Not needed for Windows):

   ```bash {linenos=false}
   chmod +x dlitescript-linux-x64
   ```

3. Move the file to your PATH, to call the command from any directory (optional):

   For Unix-like systems:

   ```bash {linenos=false}
   sudo mv dlitescript-linux-x64 /usr/local/bin/dlitescript
   ```

   For Windows:

   ```powershell {linenos=false}
   Move-Item dlitescript-win32-x64 $env:USERPROFILE\bin\dlitescript
   ```

4. Run DLiteScript:

   If not moved to PATH:

   ```bash {linenos=false}
   ./dlitescript script.dl
   ```

   If moved to PATH:

   ```bash {linenos=false}
   dlitescript script.dl
   ```

## Building from Source

### Prerequisites

- Go 1.24.6 or later
- Git

### Build Steps

1. **Clone the repository:**

   ```bash {linenos=false}
   git clone https://github.com/Dobefu/DLiteScript.git
   cd DLiteScript
   ```

2. **Download dependencies:**

   ```bash {linenos=false}
   go mod download
   ```

3. **Build the binary:**

   ```bash {linenos=false}
   go build -buildvcs
   ```

4. **Run DLiteScript:**

   ```bash {linenos=false}
   ./dlitescript your-script.dl
   ```
