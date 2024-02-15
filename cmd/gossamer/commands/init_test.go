// Copyright 2023 ChainSafe Systems (ON)
// SPDX-License-Identifier: LGPL-3.0-only

package commands

import (
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const testChainSpec = "./test_inputs/test-chain-spec-raw.json"

// TestInitFromChainSpec test "gossamer init --chain=./test_inputs/test-chain-spec-raw.json"
func TestInitFromChainSpec(t *testing.T) {
	basepath := t.TempDir()

	pubIp := "123.456.789.123"
	rootCmd, err := NewRootCommand()
	require.NoError(t, err)
	rootCmd.AddCommand(InitCmd)

	rootCmd.SetArgs([]string{InitCmd.Name(), "--base-path", basepath, "--chain", testChainSpec, "--public-ip", pubIp})
	err = rootCmd.Execute()
	require.NoError(t, err)

	out, err := exec.Command("ping", pubIp).Output()
	require.NoError(t, err)
	if strings.Contains(string(out), "Destination Host Unreachable") {
		t.Logf("ping %s failed", pubIp)
	} else {
		t.Log("public ip is reachable")
	}
}
