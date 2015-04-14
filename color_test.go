package colorlog

import "testing"

func TestColor(t *testing.T) {
	log := NewColorLog()
	log.Error("this is Error  test %s", "success")
	log.Info("this i Info test %s", "success")
	log.Warning("this i Warning test %s", "success")
	log.Debug("this i Debug test %s", "success")
}
