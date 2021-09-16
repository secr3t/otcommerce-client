package client

import "testing"

func TestResetCacheClient_ResetCaches(t *testing.T) {
	t.Log(NewResetCacheClient(apiKey).ResetCaches())
}
