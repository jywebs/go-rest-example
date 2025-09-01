# Go REST Command Executor API

A REST API server built in Go that allows remote command execution in a containerized environment. This example uses BusyBox as the base container but can be adapted for other environments.

## Purpose

This project demonstrates:
- Building a REST API server in Go
- Using Swagger for API documentation
- Containerized command execution
- Alternative to SSH/console access for remote command execution

## Features

- REST API endpoint for executing commands
- Swagger documentation
- Containerized with BusyBox
- Command output in JSON format
- Minimal attack surface (no SSH/console access needed)

## Quick Start

1. Build the Docker image:
```bash
docker build -t command-api .
```

2. Run the container:
```bash
docker run -p 8080:8080 command-api
```

## API Usage

The API server runs on port 8080 with the following endpoint:

- POST `/api/v1/execute`
  - Request body: `{"command": "your_command"}`
  - Returns: `{"output": "command_output", "error": "error_if_any"}`

Example commands:

```bash
# Get system information
curl -X POST http://localhost:8080/api/v1/execute \
  -H "Content-Type: application/json" \
  -d '{"command": "uname -a"}'

# List files
curl -X POST http://localhost:8080/api/v1/execute \
  -H "Content-Type: application/json" \
  -d '{"command": "ls -la"}'

# Check disk space
curl -X POST http://localhost:8080/api/v1/execute \
  -H "Content-Type: application/json" \
  -d '{"command": "df -h"}'
```

Access Swagger documentation at: `http://localhost:8080/swagger/index.html`

## Production Considerations

### Security

For production deployment, consider implementing:

1. Authentication/Authorization
   - JWT tokens
   - API keys
   - OAuth2

2. Command Validation
   - Whitelist allowed commands
   - Rate limiting
   - Input sanitization

3. TLS/HTTPS
   - Required for production
   - Secure communication

### Logging and Monitoring

Recommended sidecar setup:

1. Logging Sidecar
   - Log all commands executed
   - Record source IP/client information
   - Timestamp each request
   - Forward logs to central logging system

2. Monitoring
   - Track command execution times
   - Monitor resource usage
   - Alert on suspicious patterns

### Container Hardening

1. Minimal Base Image
   - Use distroless or minimal base images
   - Remove unnecessary tools
   - Limit available commands

2. Container Security
   - Run as non-root user
   - Read-only filesystem where possible
   - Resource limits
   - Network policy restrictions

## Use Cases

1. Remote System Management
   - Execute maintenance commands
   - System health checks
   - Configuration updates

2. Automated Operations
   - CI/CD pipelines
   - System automation
   - Scheduled tasks

3. Secure Command Execution
   - Alternative to SSH access
   - Auditable command execution
   - Controlled environment

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
