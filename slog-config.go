/*
Copyright 2024 Richard Kosegi

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package xlog

import (
	"flag"
	"log/slog"

	"github.com/prometheus/common/promslog"
	"github.com/spf13/pflag"
)

type LogFormat string

const (
	LogFormatJson   = LogFormat("json")
	LogFormatLogFmt = LogFormat("logfmt")
)

type SlogConfig struct {
	Level  Level
	Format Format
}

func (sc *SlogConfig) AddFlags(fs *flag.FlagSet) {
	fs.Var(&sc.Level, sc.Level.Type(), "Log level")
	fs.Var(&sc.Format, sc.Format.Type(), "Log format")
}

func (sc *SlogConfig) AddPFlags(fs *pflag.FlagSet) {
	fs.Var(&sc.Level, sc.Level.Type(), "Log level")
	fs.Var(&sc.Format, sc.Format.Type(), "Log format")
}

func (sc *SlogConfig) ToPromslogConfig() *promslog.Config {
	return &promslog.Config{
		Level:  &sc.Level.Level,
		Format: &sc.Format.Format,
		Style:  promslog.GoKitStyle,
	}
}

func (sc *SlogConfig) Logger() *slog.Logger {
	return promslog.New(sc.ToPromslogConfig())
}

func MustNew(level string, format LogFormat) *SlogConfig {
	xl, err := New(level, format)
	if err != nil {
		panic(err)
	}
	return xl
}

func New(level string, format LogFormat) (*SlogConfig, error) {
	var err error
	sc := &SlogConfig{}
	sc.Format.Format = *promslog.NewFormat()
	sc.Level.Level = *promslog.NewLevel()
	if err = sc.Format.Set(string(format)); err != nil {
		return nil, err
	}
	if err = sc.Level.Set(level); err != nil {
		return nil, err
	}
	return sc, nil
}

type Level struct {
	promslog.Level
}

func (l Level) Type() string {
	return "log-level"
}

type Format struct {
	promslog.Format
}

func (Y Format) Type() string {
	return "log-format"
}
