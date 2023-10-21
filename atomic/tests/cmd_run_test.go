package tests

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	_ "github.com/olegbespalov/xk6-atomic"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.k6.io/k6/cmd"
	k6Tests "go.k6.io/k6/cmd/tests"
)

func TestAtomicCounter_Inc(t *testing.T) {
	t.Parallel()

	ts := k6Tests.NewGlobalTestState(t)

	script, err := os.ReadFile("../../examples/script.js") //nolint:forbidigo
	require.NoError(t, err)

	require.NoError(t, afero.WriteFile(ts.FS, filepath.Join(ts.Cwd, "test.js"), []byte(script), 0o644))
	ts.CmdArgs = []string{"k6", "run", "-v", "--log-output=stdout", "--vus", "10", "-i", "20", filepath.Join(ts.Cwd, "test.js")}
	ts.ExpectedExitCode = int(0) // success

	cmd.ExecuteWithGlobalState(ts.GlobalState)

	stdout := ts.Stdout.String()

	// check that counter incremented 20 times
	for i := 1; i <= 20; i++ {
		assert.Contains(t, stdout, fmt.Sprintf("current value is: %d", i))
	}

	assert.Empty(t, ts.Stderr.String())
}

func TestAtomicCounter_Dec(t *testing.T) {
	t.Parallel()

	ts := k6Tests.NewGlobalTestState(t)

	script, err := os.ReadFile("../../examples/dec.js") //nolint:forbidigo
	require.NoError(t, err)

	require.NoError(t, afero.WriteFile(ts.FS, filepath.Join(ts.Cwd, "test.js"), []byte(script), 0o644))
	ts.CmdArgs = []string{"k6", "run", "-v", "--log-output=stdout", "--vus", "10", "-i", "20", filepath.Join(ts.Cwd, "test.js")}
	ts.ExpectedExitCode = int(0) // success

	cmd.ExecuteWithGlobalState(ts.GlobalState)

	stdout := ts.Stdout.String()

	// check that counter decremented 20 times
	for i := 1; i <= 20; i++ {
		assert.Contains(t, stdout, fmt.Sprintf("current value is: -%d", i))
	}

	assert.Empty(t, ts.Stderr.String())
}
