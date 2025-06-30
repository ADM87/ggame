package exceptions

// =======================================================================
// Not Implemented Exception
// =======================================================================
type NotImplementedException string

func NotImplemented() NotImplementedException {
	return NotImplementedException("Not implemented")
}

func NotImplementedWith(message string) NotImplementedException {
	return NotImplementedException("Not implemented: " + message)
}

func IsNotImplemented(err error) bool {
	_, ok := err.(NotImplementedException)
	return ok
}

func (e NotImplementedException) Error() string {
	return string(e)
}

// =======================================================================
// Invalid Command Exception
// =======================================================================
type InvalidCommandException string

func InvalidCommand() InvalidCommandException {
	return InvalidCommandException("Invalid command")
}

func InvalidCommandWith(message string) InvalidCommandException {
	return InvalidCommandException("Invalid command: " + message)
}

func IsInvalidCommand(err error) bool {
	_, ok := err.(InvalidCommandException)
	return ok
}

func (e InvalidCommandException) Error() string {
	return string(e)
}

// =======================================================================
// Invalid Argument Exception
// =======================================================================
type InvalidArgumentException string

func InvalidArgument() InvalidArgumentException {
	return InvalidArgumentException("Invalid argument")
}

func InvalidArgumentWith(name, message string) InvalidArgumentException {
	return InvalidArgumentException("Invalid argument: " + name + ": " + message)
}

func IsInvalidArgument(err error) bool {
	_, ok := err.(InvalidArgumentException)
	return ok
}

func (e InvalidArgumentException) Error() string {
	return string(e)
}
