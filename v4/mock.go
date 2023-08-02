package logger

type MockLogger struct {
	LoggedData    Event
	Format        string
	LogType       []LogType
	wasCalledLog  bool
	wasCalledType bool
}

func (l *MockLogger) Log(e Event, timeFormat string) error {
	l.LoggedData = e
	l.Format = timeFormat
	l.wasCalledLog = true

	return nil
}

func (l *MockLogger) Type() []LogType {
	l.wasCalledType = true
	return l.LogType
}

type MockFile struct {
	Text []string
}

func (l *MockFile) Write(b []byte) (n int, err error) {
	l.Text = append(l.Text, string(b))

	return 0, nil
}

func (l *MockFile) WriteString(s string) (n int, err error) {
	l.Text = append(l.Text, s)

	return 0, nil
}

func (l *MockFile) Close() error {
	return nil
}
