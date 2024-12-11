# Installing `func` CLI on macOS and Ubuntu

The `func` CLI tool is essential for working with Knative functions. This guide provides installation instructions for macOS and Ubuntu.

## Installation Steps

### macOS
1. Install the Knative client:
   ```bash
   brew install knative/client
   brew install knative/client/kn
   brew install func


### ubuntu
1. Install the Knative client:
   ```bash
   wget https://github.com/knative/func/releases/download/knative-v1.16.1/func_linux_amd64
   chmod +x func_linux_amd64
   sudo mv func_linux_amd64 /usr/local/bin/func

