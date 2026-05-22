# Ollama Server Feature

Adds an Ollama server service to the devcontainer setup via Docker Compose.

## GPU Support

The default configuration attempts to reserve all available NVIDIA GPUs.

```yaml
deploy:
  resources:
    reservations:
      devices:
        - driver: nvidia
          count: all
          capabilities: [gpu]
```

If your host does not have an NVIDIA GPU or the NVIDIA Container Toolkit installed, you may need to comment out these lines in `.devcontainer/ollama-server/docker-compose.yml` to prevent startup errors.
