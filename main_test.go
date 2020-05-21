package main

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
  total := Sum(5, 5)

  assert.Equal(t, 10, total, "The total should be 10")
}
