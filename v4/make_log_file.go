package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/lazybark/go-helpers/fsw"
)

// makeLogFile creates new log file with addition of current date (Y_M_D_H_M_S)
func makeLogFile(path string, truncate bool, ext string) (*os.File, error) {
	now := time.Now()
	f, err := fsw.MakePathToFile(
		fmt.Sprintf("%s-%d_%d_%d_%d_%d_%d.%s",
			path,
			now.Year(), now.Month(), now.Day(),
			now.Hour(), now.Minute(), now.Second(),
			ext,
		),
		truncate,
	)
	if err != nil {
		return nil, fmt.Errorf("[makeLogFile] error making log path: %w", err)
	}

	return f, err
}
