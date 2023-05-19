# NexLog

NexLog is a customizable logging package for Go (golang) applications. It provides capabilities to customize log levels, formats, filters, and hooks to enhance your logging system.

## Features

- Multiple log levels: DBG, INF, ERR, WRN, FTL
- Log formatting using `LogFormatter` interface.
- Log filtering using `LogFilter` interface.
- Support for adding hooks to perform additional actions using `LogHook` interface.
- Thread-safe logging using `sync.Pool`.

## Installation

To install NexLog, run:

```bash
go get github.com/nex-gen-tech/nexlog
```


## Usage

Here's a basic usage example:

```go
package main

import (
	"github.com/nex-gen-tech/nexlog"
)

func main() {
	logger := nexlog.NewLogger("myapp")

	logger.Debug("This is a debug log")
	logger.Info("This is an info log")
	logger.Warn("This is a warning log")
	logger.Error("This is an error log")
	logger.Fatal("This is a fatal log")
}
```

## Interfaces

*LogFormatter* - Implement the LogFormatter interface to customize the format of your log messages.

*LogFilter* - Implement the LogFilter interface to decide which log messages to print based on the log level and/or other conditions.

*LogHook* - Implement the LogHook interface to perform additional actions when a log message is printed. This could be useful for sending alerts, storing log messages in a database, etc.

## Contributing
Contributions are welcome! Please submit a Pull Request with any enhancements, fixes, or features.

## License
This project is licensed under the MIT License.
