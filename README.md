# check-http-status

`check-http-status` is a CLI tool written in Go that monitors the HTTP status code of a specified website at regular intervals. It can log the results to a file and supports configurable timeouts and intervals.

## Features
- Monitors HTTP status codes of a target URL.
- Configurable interval for sending requests.
- Optional logging of results to a file.
- Supports request timeouts to avoid hanging on unresponsive servers.
- Gracefully handles termination via `Ctrl+C`.

## Installation

### Install via `go install`
You can install the command directly using `go install`:
```bash
go install github.com/supermarine1377/check-http-status@latest
```
Ensure your `GOPATH/bin` is in your system's `PATH` to run the command directly.

### Build from Source

1. Clone the repository:
   ```bash
   git clone https://github.com/supermarine1377/monitoring-scripts.git
   cd monitoring-scripts/go/check-http-status
   ```

2. Build the CLI:
   ```bash
   go build -o check-http-status
   ```

3. Add the built binary to your `PATH` or execute it directly.

## Usage

```bash
check-http-status <URL> [flags]
```

### Example

Monitor the HTTP status code of `https://example.com` every 30 seconds and log results to a file:
```bash
check-http-status https://example.com -i 30 -c
```

### Flags

| Flag                          | Shorthand | Description                                                                                         | Default |
|-------------------------------|-----------|-----------------------------------------------------------------------------------------------------|---------|
| `--interval-seconds`          | `-i`      | Interval in seconds between HTTP requests.                                                         | `10`    |
| `--create-log-file`           | `-c`      | Create a log file to save the results. The log file name format is `check-http-status_<timestamp>.log`. | `false` |
| `--timeout-seconds`           | `-t`      | Timeout in seconds for each HTTP request. If no response is received within this time, the request is considered failed. | `30`    |

### Handling Interruptions

The tool handles interruptions gracefully. Press `Ctrl+C` to terminate the monitoring process.

## Output

- On the console, the tool prints the HTTP status code and the timestamp for each request.
- If logging is enabled (`--create-log-file`), the results are written to a log file in the working directory.

## Development

### Prerequisites
- [Go](https://golang.org/) 1.20 or later.

### Directory Structure

```plaintext
├── cmd
│   ├── root.go       # Main command logic.
├── internal
│   ├── http_status   # Package for monitoring HTTP status.
│   ├── log_files     # Package for managing log files.
```

### Testing
To run tests:
```bash
go test ./...
```

## Contributing
Contributions are welcome! Feel free to open an issue or submit a pull request.

## License
This project is licensed under the MIT License. See the `LICENSE` file for details.

## Contact
For questions or feedback, please contact `NAME HERE` at `<EMAIL ADDRESS>`. Replace placeholders in this README as necessary.

