package timecop_test

import (
	"github.com/bluele/go-timecop"
	"testing"
	"time"
)

func TestFreeze(t *testing.T) {
	now := timecop.Now()
	timecop.Freeze(now)

	if timecop.Now() != now {
		t.Errorf("Expected time is not %v.", now)
	}

	timecop.Return()

	if !timecop.Now().Before(time.Now()) {
		t.Error("timecop should be reolve freezing.")
	}
}

func TestTravel(t *testing.T) {
	now := timecop.Now()
	future := now.AddDate(1, 0, 0)
	timecop.Travel(future)

	if timecop.Now() == now {
		t.Errorf("Expected time is not %v.", now)
	}

	if !timecop.Now().After(future) {
		t.Errorf("Expected time should be greater than %v.", future)
	}
}
