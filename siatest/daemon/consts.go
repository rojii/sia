package daemon

import (
	"os"

	"go.sia.tech/siad/persist"
	"go.sia.tech/siad/siatest"
)

// daemonTestDir creates a temporary testing directory for daemon tests. This
// should only every be called once per test. Otherwise it will delete the
// directory again.
func daemonTestDir(testName string) string {
	path := siatest.TestDir("daemon", testName)
	if err := os.MkdirAll(path, persist.DefaultDiskPermissionsTest); err != nil {
		panic(err)
	}
	return path
}
