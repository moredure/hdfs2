package hdfs2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStatFs(t *testing.T) {
	client := getClient(t)

	_, err := client.StatFs()
	require.NoError(t, err)
}
