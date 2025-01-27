package ports

type Logger interface {
	// Debug logs a message at level Debug with fields.
	Debug(msg string, fields map[string]interface{})
	// Info logs a message at level Info with fields.
	Info(msg string, fields map[string]interface{})
	// Warn logs a message at level Warn with fields.
	Warn(msg string, fields map[string]interface{})
	// Error logs a message at level Error with fields.
	Error(msg string, fields map[string]interface{})
	// Fatal logs a message at level Fatal with fields and then calls os.Exit(1).
	Fatal(msg string, fields map[string]interface{})

	// SetLogLevel sets the log level.
	SetLogLevel(level string)
}
