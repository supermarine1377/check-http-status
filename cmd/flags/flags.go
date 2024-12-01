package flags

import (
	"github.com/spf13/cobra"
)

type Flags struct {
	intervalSeconds int
	timeoutSeconds  int
	createLogFile   bool
}

func (f *Flags) IntervalSeconds() int {
	return f.intervalSeconds
}

func (f *Flags) TimeoutSeconds() int {
	return f.timeoutSeconds
}

func (f *Flags) CreateLogFile() bool {
	return f.createLogFile
}

const INTERVAL_SECONDS = "interval-seconds"
const INTERVAL_SECONDS_SHORTHAND = "i"
const DEFAULT_INTERVAL_SECONDS = 10

const CREATE_LOG_FILE = "create-log-file"
const CREATE_LOG_FILE_SHORTHAND = "c"
const DEFAULT_CREATE_LOG_FILE = false

const TIMEOUT_SECONDS = "timeout-seconds"
const TIMEOUT_SECONDS_SHORTHAND = "t"
const DEFAULT_TIMEOUT_SECONDS = 30

func Parse(cmd *cobra.Command) (*Flags, error) {
	is, err := cmd.Flags().GetInt(INTERVAL_SECONDS)
	if err != nil {
		return nil, err
	}

	clf, err := cmd.Flags().GetBool(CREATE_LOG_FILE)
	if err != nil {
		return nil, err
	}

	ts, err := cmd.Flags().GetInt(TIMEOUT_SECONDS)
	if err != nil {
		return nil, err
	}

	return &Flags{
		intervalSeconds: is,
		timeoutSeconds:  ts,
		createLogFile:   clf,
	}, nil
}
