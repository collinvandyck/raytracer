package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddNewlines(t *testing.T) {
	data := ``
	require.EqualValues(t, "", addNewlines(data))

	data = `|1| |2|`
	require.EqualValues(t, "|1|\n|2|", addNewlines(data))
}
