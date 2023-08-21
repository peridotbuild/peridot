// Copyright 2023 Peridot Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package base

import (
	"fmt"
	"log"
	"strings"
)

type LogLevel string

const (
	LogLevelDebug LogLevel = "DEBUG"
	LogLevelInfo  LogLevel = "INFO"
	LogLevelWarn  LogLevel = "WARN"
	LogLevelError LogLevel = "ERROR"
	LogLevelFatal LogLevel = "FATAL"
)

func Logf(level LogLevel, format string, args ...interface{}) {
	if !strings.HasPrefix(format, fmt.Sprintf("[%s] ", level)) {
		format = fmt.Sprintf("[%s]: %s", level, format)
	}

	if level == LogLevelFatal {
		log.Fatalf(format, args...)
	} else {
		log.Printf(format, args...)
	}
}

func LogErrorf(format string, args ...interface{}) {
	Logf(LogLevelError, format, args...)
}

func LogWarnf(format string, args ...interface{}) {
	Logf(LogLevelWarn, format, args...)
}

func LogInfof(format string, args ...interface{}) {
	Logf(LogLevelInfo, format, args...)
}

func LogDebugf(format string, args ...interface{}) {
	Logf(LogLevelDebug, format, args...)
}

func LogFatalf(format string, args ...interface{}) {
	Logf(LogLevelFatal, format, args...)
}
