# Function Setup Guide

This guide walks you through installing the `func` CLI tool and creating a new Knative function using the Ucode.io Knative template.

## Prerequisites

- **Homebrew**: Ensure that [Homebrew](https://brew.sh/) is installed on your machine.

## Steps to Create a New Function

1. Install the `func` CLI tool using Homebrew:

   ```bash
   brew install func

2. Create your Knative function by running the following command. Replace {FUNCTION-NAME} with your desired function name:
   ```bash
   func create {FUNCTION-NAME} -l go -t function --repository https://github.com/Ucode-io/knative-template
