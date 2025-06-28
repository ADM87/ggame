package logger

import (
	"bytes"
	"fmt"
	"io"
	"sync"
	"testing"
)

const (
	debugTestMessage = "This is a debug message"
	infoTestMessage  = "This is an info message"
	warnTestMessage  = "This is a warning message"
	errorTestMessage = "This is an error message"

	debugfTestMessage = "This is a formatted debug message: %s"
	infofTestMessage  = "This is a formatted info message: %s"
	warnfTestMessage  = "This is a formatted warning message: %s"
	errorfTestMessage = "This is a formatted error message: %s"
)

const (
	debugPrefix = "[DEBUG] "
	infoPrefix  = "[INFO] "
	warnPrefix  = "[WARN] "
	errorPrefix = "[ERROR] "
)

const (
	debugStyle = StyleGray
	infoStyle  = StyleBlue
	warnStyle  = StyleYellow
	errorStyle = StyleRed
)

var (
	prefixMap = map[LogLevel]string{
		LevelDebug: debugPrefix,
		LevelInfo:  infoPrefix,
		LevelWarn:  warnPrefix,
		LevelError: errorPrefix,
	}
	styleMap = map[LogLevel]string{
		LevelDebug: debugStyle,
		LevelInfo:  infoStyle,
		LevelWarn:  warnStyle,
		LevelError: errorStyle,
	}
)

func mapBuffer(buffer *bytes.Buffer) map[LogLevel]io.Writer {
	return map[LogLevel]io.Writer{
		LevelDebug: buffer,
		LevelInfo:  buffer,
		LevelWarn:  buffer,
		LevelError: buffer,
	}
}

func TestSetVerboseLevels(t *testing.T) {
	testLogger := NewLogger()

	testLogger.SetVerboseLevel(LevelNone)
	if testLogger.GetVerboseLevel() != LevelNone {
		t.Errorf("Expected verbose level to be %d, got %d", LevelNone, testLogger.GetVerboseLevel())
	}
	testLogger.SetVerboseLevel(LevelDebug | LevelInfo)
	if testLogger.GetVerboseLevel() != (LevelDebug | LevelInfo) {
		t.Errorf("Expected verbose level to be %d, got %d", LevelDebug|LevelInfo, testLogger.GetVerboseLevel())
	}
	testLogger.SetVerboseLevel(LevelWarn | LevelError)
	if testLogger.GetVerboseLevel() != (LevelWarn | LevelError) {
		t.Errorf("Expected verbose level to be %d, got %d", LevelWarn|LevelError, testLogger.GetVerboseLevel())
	}
	testLogger.SetVerboseLevel(LevelDebug | LevelInfo | LevelWarn | LevelError)
	if testLogger.GetVerboseLevel() != (LevelDebug | LevelInfo | LevelWarn | LevelError) {
		t.Errorf("Expected verbose level to be %d, got %d", LevelDebug|LevelInfo|LevelWarn|LevelError, testLogger.GetVerboseLevel())
	}
}

func TestLogging(t *testing.T) {
	testLogger := NewLogger()
	testBuffer := &bytes.Buffer{}

	testLogger.MapWriters(mapBuffer(testBuffer))
	testLogger.SetVerboseLevel(LevelAll)

	testBuffer.Reset()
	testLogger.Debug(debugTestMessage)

	expectedOutput := debugTestMessage + "\n"
	if testBuffer.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, testBuffer.String())
	}

	testBuffer.Reset()
	testLogger.Info(infoTestMessage)

	expectedOutput = infoTestMessage + "\n"
	if testBuffer.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, testBuffer.String())
	}

	testBuffer.Reset()
	testLogger.Warn(warnTestMessage)

	expectedOutput = warnTestMessage + "\n"
	if testBuffer.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, testBuffer.String())
	}

	testBuffer.Reset()
	testLogger.Error(errorTestMessage)

	expectedOutput = errorTestMessage + "\n"
	if testBuffer.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, testBuffer.String())
	}

	testBuffer.Reset()
	testLogger.Debug(debugTestMessage)
	testLogger.Info(infoTestMessage)
	testLogger.Warn(warnTestMessage)
	testLogger.Error(errorTestMessage)

	expectedOutput = debugTestMessage + "\n" +
		infoTestMessage + "\n" +
		warnTestMessage + "\n" +
		errorTestMessage + "\n"
	if testBuffer.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, testBuffer.String())
	}
}

func TestLoggingWithPrefixes(t *testing.T) {
	testLogger := NewLogger()
	testBuffer := &bytes.Buffer{}

	testLogger.MapWriters(mapBuffer(testBuffer))
	testLogger.MapPrefixes(prefixMap)
	testLogger.SetVerboseLevel(LevelAll)

	testBuffer.Reset()
	testLogger.Debug(debugTestMessage)

	expectedOutput := debugPrefix + debugTestMessage + "\n"
	if testBuffer.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, testBuffer.String())
	}
	testBuffer.Reset()
	testLogger.Info(infoTestMessage)

	expectedOutput = infoPrefix + infoTestMessage + "\n"
	if testBuffer.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, testBuffer.String())
	}
	testBuffer.Reset()
	testLogger.Warn(warnTestMessage)

	expectedOutput = warnPrefix + warnTestMessage + "\n"
	if testBuffer.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, testBuffer.String())
	}
	testBuffer.Reset()
	testLogger.Error(errorTestMessage)

	expectedOutput = errorPrefix + errorTestMessage + "\n"
	if testBuffer.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, testBuffer.String())
	}

	testBuffer.Reset()
	testLogger.Debug(debugTestMessage)
	testLogger.Info(infoTestMessage)
	testLogger.Warn(warnTestMessage)
	testLogger.Error(errorTestMessage)

	expectedOutput = debugPrefix + debugTestMessage + "\n" +
		infoPrefix + infoTestMessage + "\n" +
		warnPrefix + warnTestMessage + "\n" +
		errorPrefix + errorTestMessage + "\n"
	if testBuffer.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, testBuffer.String())
	}
}

func TestFormatLogging(t *testing.T) {
	testLogger := NewLogger()
	testBuffer := &bytes.Buffer{}

	testLogger.MapWriters(mapBuffer(testBuffer))
	testLogger.SetVerboseLevel(LevelAll)

	testBuffer.Reset()
	testLogger.Debugf(debugfTestMessage, "formatted debug")

	expectedOutput := fmt.Sprintf(debugfTestMessage, "formatted debug") + "\n"
	if testBuffer.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, testBuffer.String())
	}

	testBuffer.Reset()
	testLogger.Infof(infofTestMessage, "formatted info")

	expectedOutput = fmt.Sprintf(infofTestMessage, "formatted info") + "\n"
	if testBuffer.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, testBuffer.String())
	}

	testBuffer.Reset()
	testLogger.Warnf(warnfTestMessage, "formatted warning")

	expectedOutput = fmt.Sprintf(warnfTestMessage, "formatted warning") + "\n"
	if testBuffer.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, testBuffer.String())
	}

	testBuffer.Reset()
	testLogger.Errorf(errorfTestMessage, "formatted error")

	expectedOutput = fmt.Sprintf(errorfTestMessage, "formatted error") + "\n"
	if testBuffer.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, testBuffer.String())
	}

	testBuffer.Reset()
	testLogger.Debugf(debugfTestMessage, "formatted debug")
	testLogger.Infof(infofTestMessage, "formatted info")
	testLogger.Warnf(warnfTestMessage, "formatted warning")
	testLogger.Errorf(errorfTestMessage, "formatted error")

	expectedOutput = fmt.Sprintf(debugfTestMessage, "formatted debug") + "\n" +
		fmt.Sprintf(infofTestMessage, "formatted info") + "\n" +
		fmt.Sprintf(warnfTestMessage, "formatted warning") + "\n" +
		fmt.Sprintf(errorfTestMessage, "formatted error") + "\n"
	if testBuffer.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, testBuffer.String())
	}
}

func TestFormatLoggingWithPrefix(t *testing.T) {
	testLogger := NewLogger()
	testBuffer := &bytes.Buffer{}

	testLogger.MapWriters(mapBuffer(testBuffer))
	testLogger.MapPrefixes(prefixMap)
	testLogger.SetVerboseLevel(LevelAll)

	testBuffer.Reset()
	testLogger.Debugf(debugfTestMessage, "formatted debug")

	expectedOutput := debugPrefix + fmt.Sprintf(debugfTestMessage, "formatted debug") + "\n"
	if testBuffer.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, testBuffer.String())
	}

	testBuffer.Reset()
	testLogger.Infof(infofTestMessage, "formatted info")

	expectedOutput = infoPrefix + fmt.Sprintf(infofTestMessage, "formatted info") + "\n"
	if testBuffer.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, testBuffer.String())
	}

	testBuffer.Reset()
	testLogger.Warnf(warnfTestMessage, "formatted warning")

	expectedOutput = warnPrefix + fmt.Sprintf(warnfTestMessage, "formatted warning") + "\n"
	if testBuffer.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, testBuffer.String())
	}

	testBuffer.Reset()
	testLogger.Errorf(errorfTestMessage, "formatted error")

	expectedOutput = errorPrefix + fmt.Sprintf(errorfTestMessage, "formatted error") + "\n"
	if testBuffer.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, testBuffer.String())
	}

	testBuffer.Reset()
	testLogger.Debugf(debugfTestMessage, "formatted debug")
	testLogger.Infof(infofTestMessage, "formatted info")
	testLogger.Warnf(warnfTestMessage, "formatted warning")
	testLogger.Errorf(errorfTestMessage, "formatted error")

	expectedOutput = debugPrefix + fmt.Sprintf(debugfTestMessage, "formatted debug") + "\n" +
		infoPrefix + fmt.Sprintf(infofTestMessage, "formatted info") + "\n" +
		warnPrefix + fmt.Sprintf(warnfTestMessage, "formatted warning") + "\n" +
		errorPrefix + fmt.Sprintf(errorfTestMessage, "formatted error") + "\n"
	if testBuffer.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, testBuffer.String())
	}
}

func TestStyledLogging(t *testing.T) {
	testLogger := NewLogger()
	testBuffer := &bytes.Buffer{}

	testLogger.MapWriters(mapBuffer(testBuffer))
	testLogger.MapStyles(styleMap)
	testLogger.SetVerboseLevel(LevelAll)

	testBuffer.Reset()
	testLogger.Debug(debugTestMessage)

	expectedOutput := Gray(debugTestMessage) + "\n"
	if testBuffer.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, testBuffer.String())
	}

	testBuffer.Reset()
	testLogger.Info(infoTestMessage)

	expectedOutput = Blue(infoTestMessage) + "\n"
	if testBuffer.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, testBuffer.String())
	}

	testBuffer.Reset()
	testLogger.Warn(warnTestMessage)

	expectedOutput = Yellow(warnTestMessage) + "\n"
	if testBuffer.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, testBuffer.String())
	}

	testBuffer.Reset()
	testLogger.Error(errorTestMessage)

	expectedOutput = Red(errorTestMessage) + "\n"
	if testBuffer.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, testBuffer.String())
	}

	testBuffer.Reset()
	testLogger.Debug(debugTestMessage)
	testLogger.Info(infoTestMessage)
	testLogger.Warn(warnTestMessage)
	testLogger.Error(errorTestMessage)

	expectedOutput = Gray(debugTestMessage) + "\n" +
		Blue(infoTestMessage) + "\n" +
		Yellow(warnTestMessage) + "\n" +
		Red(errorTestMessage) + "\n"
	if testBuffer.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, testBuffer.String())
	}
}

func TestStyledLoggingWithPrefixes(t *testing.T) {
	testLogger := NewLogger()
	testBuffer := &bytes.Buffer{}

	testLogger.MapWriters(mapBuffer(testBuffer))
	testLogger.MapPrefixes(prefixMap)
	testLogger.MapStyles(styleMap)
	testLogger.SetVerboseLevel(LevelAll)

	testBuffer.Reset()
	testLogger.Debug(debugTestMessage)

	expectedOutput := debugPrefix + Gray(debugTestMessage) + "\n"
	if testBuffer.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, testBuffer.String())
	}

	testBuffer.Reset()
	testLogger.Info(infoTestMessage)

	expectedOutput = infoPrefix + Blue(infoTestMessage) + "\n"
	if testBuffer.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, testBuffer.String())
	}

	testBuffer.Reset()
	testLogger.Warn(warnTestMessage)

	expectedOutput = warnPrefix + Yellow(warnTestMessage) + "\n"
	if testBuffer.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, testBuffer.String())
	}

	testBuffer.Reset()
	testLogger.Error(errorTestMessage)

	expectedOutput = errorPrefix + Red(errorTestMessage) + "\n"
	if testBuffer.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, testBuffer.String())
	}

	testBuffer.Reset()
	testLogger.Debug(debugTestMessage)
	testLogger.Info(infoTestMessage)
	testLogger.Warn(warnTestMessage)
	testLogger.Error(errorTestMessage)

	expectedOutput = debugPrefix + Gray(debugTestMessage) + "\n" +
		infoPrefix + Blue(infoTestMessage) + "\n" +
		warnPrefix + Yellow(warnTestMessage) + "\n" +
		errorPrefix + Red(errorTestMessage) + "\n"
	if testBuffer.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, testBuffer.String())
	}
}

func TestStyledFormatLogging(t *testing.T) {
	testLogger := NewLogger()
	testBuffer := &bytes.Buffer{}

	testLogger.MapWriters(mapBuffer(testBuffer))
	testLogger.MapStyles(styleMap)
	testLogger.SetVerboseLevel(LevelAll)

	testBuffer.Reset()
	testLogger.Debugf(debugfTestMessage, "styled debug")

	expectedOutput := Gray(fmt.Sprintf(debugfTestMessage, "styled debug")) + "\n"
	if testBuffer.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, testBuffer.String())
	}

	testBuffer.Reset()
	testLogger.Infof(infofTestMessage, "styled info")

	expectedOutput = Blue(fmt.Sprintf(infofTestMessage, "styled info")) + "\n"
	if testBuffer.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, testBuffer.String())
	}

	testBuffer.Reset()
	testLogger.Warnf(warnfTestMessage, "styled warning")

	expectedOutput = Yellow(fmt.Sprintf(warnfTestMessage, "styled warning")) + "\n"
	if testBuffer.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, testBuffer.String())
	}

	testBuffer.Reset()
	testLogger.Errorf(errorfTestMessage, "styled error")

	expectedOutput = Red(fmt.Sprintf(errorfTestMessage, "styled error")) + "\n"
	if testBuffer.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, testBuffer.String())
	}

	testBuffer.Reset()
	testLogger.Debugf(debugfTestMessage, "styled debug")
	testLogger.Infof(infofTestMessage, "styled info")
	testLogger.Warnf(warnfTestMessage, "styled warning")
	testLogger.Errorf(errorfTestMessage, "styled error")

	expectedOutput = Gray(fmt.Sprintf(debugfTestMessage, "styled debug")) + "\n" +
		Blue(fmt.Sprintf(infofTestMessage, "styled info")) + "\n" +
		Yellow(fmt.Sprintf(warnfTestMessage, "styled warning")) + "\n" +
		Red(fmt.Sprintf(errorfTestMessage, "styled error")) + "\n"
	if testBuffer.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, testBuffer.String())
	}
}

func TestStyledFormatLoggingWithPrefixes(t *testing.T) {
	testLogger := NewLogger()
	testBuffer := &bytes.Buffer{}

	testLogger.MapWriters(mapBuffer(testBuffer))
	testLogger.MapPrefixes(prefixMap)
	testLogger.MapStyles(styleMap)
	testLogger.SetVerboseLevel(LevelAll)

	testBuffer.Reset()
	testLogger.Debugf(debugfTestMessage, "styled debug")

	expectedOutput := debugPrefix + Gray(fmt.Sprintf(debugfTestMessage, "styled debug")) + "\n"
	if testBuffer.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, testBuffer.String())
	}

	testBuffer.Reset()
	testLogger.Infof(infofTestMessage, "styled info")

	expectedOutput = infoPrefix + Blue(fmt.Sprintf(infofTestMessage, "styled info")) + "\n"
	if testBuffer.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, testBuffer.String())
	}

	testBuffer.Reset()
	testLogger.Warnf(warnfTestMessage, "styled warning")

	expectedOutput = warnPrefix + Yellow(fmt.Sprintf(warnfTestMessage, "styled warning")) + "\n"
	if testBuffer.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, testBuffer.String())
	}

	testBuffer.Reset()
	testLogger.Errorf(errorfTestMessage, "styled error")

	expectedOutput = errorPrefix + Red(fmt.Sprintf(errorfTestMessage, "styled error")) + "\n"
	if testBuffer.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, testBuffer.String())
	}

	testBuffer.Reset()
	testLogger.Debugf(debugfTestMessage, "styled debug")
	testLogger.Infof(infofTestMessage, "styled info")
	testLogger.Warnf(warnfTestMessage, "styled warning")
	testLogger.Errorf(errorfTestMessage, "styled error")

	expectedOutput = debugPrefix + Gray(fmt.Sprintf(debugfTestMessage, "styled debug")) + "\n" +
		infoPrefix + Blue(fmt.Sprintf(infofTestMessage, "styled info")) + "\n" +
		warnPrefix + Yellow(fmt.Sprintf(warnfTestMessage, "styled warning")) + "\n" +
		errorPrefix + Red(fmt.Sprintf(errorfTestMessage, "styled error")) + "\n"
	if testBuffer.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, testBuffer.String())
	}
}

func TestConcurrentLogging(t *testing.T) {
	const numGoroutines = 1000
	const messagesPerGoroutine = 1000

	buffer := &bytes.Buffer{}
	logger := NewLogger()
	logger.MapWriters(map[LogLevel]io.Writer{
		LevelInfo: buffer,
	})
	logger.SetVerboseLevel(LevelInfo)

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	for i := range numGoroutines {
		go func(id int) {
			defer wg.Done()
			for j := range messagesPerGoroutine {
				logger.Infof("Goroutine %d - message %d", id, j)
			}
		}(i)
	}

	wg.Wait()

	lines := bytes.Count(buffer.Bytes(), []byte("\n"))
	expectedLines := numGoroutines * messagesPerGoroutine
	if lines != expectedLines {
		t.Errorf("Expected %d log lines, got %d", expectedLines, lines)
	}
}
