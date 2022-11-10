package ratio_test

import (
	ratio "katas/05_ratio"
	"testing"
)

func TestRatio1920x1080(t *testing.T) {
	xRatio, yRatio := ratio.CalcRatio(1920, 1080)
	if xRatio != 16 || yRatio != 9 {
		t.Fail()
	}
}

func TestRatio1350x900(t *testing.T) {
	xRatio, yRatio := ratio.CalcRatio(1350, 900)
	if xRatio != 3 || yRatio != 2 {
		t.Fail()
	}
}

func TestRatio1280x720(t *testing.T) {
	xRatio, yRatio := ratio.CalcRatio(1280, 720)
	if xRatio != 16 || yRatio != 9 {
		t.Fail()
	}
}

func TestRatio2100x900(t *testing.T) {
	xRatio, yRatio := ratio.CalcRatio(2100, 900)
	if xRatio != 7 || yRatio != 3 {
		t.Fail()
	}
}

func TestRatio1440x900(t *testing.T) {
	xRatio, yRatio := ratio.CalcRatio(1440, 900)
	if xRatio != 8 || yRatio != 5 {
		t.Fail()
	}
}
