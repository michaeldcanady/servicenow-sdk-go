# Justfile equivalent of the Makefile

OS := if env('OS') == "Windows_NT" { "Windows_NT" } else { `uname -s` }

set windows-shell := ["C:\\windows\\System32\\WindowsPowerShell\\v1.0\\powershell.exe"]

# Detect OS and set current directory
#if [ "$OS" = "Windows_NT" ]; then
#	set windows-shell := ["C:\\windows\\System32\\WindowsPowerShell\\v1.0\\powershell.exe"]
#	CURRENT_DIR=$(powershell -Command "(Get-Location).Path");
#	CONTAINER_ENGINE="unknown"
#	if (Get-Command podman -ErrorAction SilentlyContinue) {
#		$CONTAINER_ENGINE="podman"
#	} elseif (Get-Command docker -ErrorAction SilentlyContinue) {
#		$CONTAINER_ENGINE="docker"
#	} else {
#		Write-Host "❌ Neither Docker nor Podman found."
#		exit 1
#	}
#else
#	CURRENT_DIR=$(pwd);
#fi;
DOCKER_ARGS := ""
CURRENT_DIR := if OS == "Windows_NT" { `powershell -Command "(Get-Location).Path"` } else { "Shell pwd" }
CONTAINER_ENGINE := "podman"
CONTAINER_ROOT := "/docs"
PORT := "8000"
CONTAINER_PORT := "8000"

# Default image and options
MKDOCS_DOCKER_IMAGE := "squidfunk/mkdocs-material"
MKDOCS_RUN_ARGS := ""


# Determine additional engine args
ADDITIONAL_ARGS := ""

# Task to build docs
serve-docs:
	# Run container with mkdocs build steps
	{{CONTAINER_ENGINE}} run {{MKDOCS_RUN_ARGS}} --rm -p "{{PORT}}:{{CONTAINER_PORT}}" {{ADDITIONAL_ARGS}} \
	    -v ${PWD}:{{CONTAINER_ROOT}} {{MKDOCS_DOCKER_IMAGE}} \
	    serve --dirtyreload --dev-addr=0.0.0.0:{{CONTAINER_PORT}}

build-docs:
	# Run container with mkdocs build steps
	{{CONTAINER_ENGINE}} run {{MKDOCS_RUN_ARGS}} --rm -p "{{PORT}}:{{CONTAINER_PORT}}" {{ADDITIONAL_ARGS}} \
	    -v ${PWD}:{{CONTAINER_ROOT}} {{MKDOCS_DOCKER_IMAGE}} \
