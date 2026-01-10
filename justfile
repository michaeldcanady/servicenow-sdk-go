# Variables
export DOC_IMAGE := "sdk-docs"
export DOC_PORT := "8000"
export DOC_PATH := "."

# Build the docs container image
build-docs:
    podman build -f doc.dockerfile -t {{DOC_IMAGE}}

# Serve docs locally with live reload
serve-docs: build-docs
    podman run --rm -p {{DOC_PORT}}:8000 {{DOC_IMAGE}}

# Build static site output (without serving)
generate-docs: build-docs
    podman run --rm {{DOC_IMAGE}} mkdocs build --clean

# Clean up local build artifacts
clean-docs:
    rm -rf site
