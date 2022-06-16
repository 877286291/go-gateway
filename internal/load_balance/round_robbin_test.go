package load_balance

import (
	"go-gateway/internal"
	"testing"
)

func TestRoundRobbin_AddServer(t *testing.T) {
	roundRobbin := RoundRobbin{}
	roundRobbin.AddServer(internal.NewServer("127.0.0.1", "1000"))
	if len(roundRobbin.servers) <= 0 {
		t.Fatal("add server failed")
	}
}
func TestRandom_NextServer(t *testing.T) {
	roundRobbin := RoundRobbin{}
	roundRobbin.AddServer(internal.NewServer("127.0.0.1", "1000"))
	roundRobbin.AddServer(internal.NewServer("127.0.0.1", "1001"))
	roundRobbin.AddServer(internal.NewServer("127.0.0.1", "1002"))
	roundRobbin.AddServer(internal.NewServer("127.0.0.1", "1003"))
	for i := 0; i < len(roundRobbin.servers)*2; i++ {
		roundRobbin.NextServer()
		if i%len(roundRobbin.servers)+1 != roundRobbin.currentIdx {
			t.Fatalf("expect nextIdx is:%d but got %d", i%len(roundRobbin.servers)+1, roundRobbin.currentIdx)
		}
	}
}
