package test

import (
	"testing"

	"github.com/mgechev/revive/rule"
)

func TestDeepExit(t *testing.T) {
	testRule(t, "deep-exit", &rule.DeepExitRule{})
}
