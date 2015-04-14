package colorlog

import "testing"

func TestColor(t *testing.T){
    Error("this is Error  test %s", "success")
    Info("this i Info test %s", "success")
    Warning("this i Warning test %s", "success")
    Debug("this i Debug test %s", "success")
}