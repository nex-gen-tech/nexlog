# **NexLog**

NexLog is a customizable logging package for Go (Golang) applications. It provides capabilities to customize log levels, formats, filters, and hooks to enhance your logging system.

## **Features**

- Support for different log levels: Debug, Info, Warn, Error, and Fatal.
- Customizable log formatters to modify the output format of log messages.
- Customizable log filters to decide which log messages to print based on the log level and/or other conditions.
- Support for adding hooks to perform additional actions using the LogHook interface.

## **Installation**

Use the following command to install the NexLog package:

`go get github.com/nex-gen-tech/nexlog`

## **Usage**

Here's a simple example of how to use NexLog in your Go application:

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

## **Customization**

NexLog provides several interfaces which you can implement to customize the behavior of the logger:

- **LogFormatter**: Implement the LogFormatter interface to customize the format of your log messages.
- **LogFilter**: Implement the LogFilter interface to decide which log messages to print based on the log level and/or other conditions.
- **LogHook**: Implement the LogHook interface to perform additional actions when a log message is printed. This could be useful for sending alerts, storing log messages in a database, etc.

****Custom Log Formatter****

To create a custom log formatter, implement the `LogFormatter` interface by defining the `Format` method. The `Format` method takes a `LogEntry` object and returns a formatted log message as a string. Here's an example:

```go
type CustomFormatter struct{}

func (f *CustomFormatter) Format(entry *nexlog.LogEntry) string {
	return fmt.Sprintf("Custom format: [%s] - %s: %s", entry.Timestamp.Format("2006-01-02 15:04:05"), entry.Level, entry.Message)
}

func main() {
	logger := nexlog.NewLogger("myapp")
	logger.SetFormatter(&CustomFormatter{})

	logger.Info("This is a custom formatted log message")
}
```

****Custom Log Filter****

To create a custom log filter, implement the `LogFilter` interface by defining the `Filter` method. The `Filter` method takes a `LogEntry` object as input and returns a boolean value. If the value is `true`, the log message will be printed; if `false`, the log message will be ignored. Here's an example:

```go
type CustomFilter struct{}

func (f *CustomFilter) Filter(entry *nexlog.LogEntry) bool {
	// Filter out log messages containing the word "ignore"
	return !strings.Contains(entry.Message, "ignore")
}

func main() {
	logger := nexlog.NewLogger("myapp")
	logger.SetFilter(&CustomFilter{})

	logger.Info("This log message will be printed")
	logger.Info("This log message will be ignored because it contains the word 'ignore'")
}
```

****Custom Log Hook****

To create a custom log hook, implement the `LogHook` interface by defining the `Fire` method. The `Fire` method takes a `LogEntry` object as input and performs additional actions when a log message is printed. Here's an example:

```go
type CustomHook struct{}

func (h *CustomHook) Fire(entry *nexlog.LogEntry) {
	fmt.Printf("Custom hook fired for log message: %s\n", entry.Message)
}

func main() {
	logger := nexlog.NewLogger("myapp")
	logger.AddHook(&CustomHook{})

	logger.Info("This is a log message with a custom hook")
}
```

In this example, the `CustomHook` struct implements the `LogHook` interface by defining the `Fire` method. The `Fire` method takes a `LogEntry` object and performs an additional action (in this case, printing a message) when a log message is printed. We then create a new logger and add the custom hook to it. When logging messages with this logger, the custom hook will be executed.

## **Contributing**

If you want to contribute to the project, you can submit a Pull Request with any enhancements, fixes, or features.