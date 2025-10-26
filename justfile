# Task to build docs
serve-docs:
	# Run container with mkdocs build steps
	podman run --rm -p "8000:8000" -v .:/docs $(podman build -f doc.dockerfile -q)

build-docs:
	podman build -f doc.dockerfile
