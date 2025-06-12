package dhan

import "testing"

func TestShouldExit_TargetHit(t *testing.T) {
	dc := &DhanClient{Target: 10}
	if exit, reason := dc.ShouldExit(12); !exit || reason != "🎯 Target hit" {
		t.Fatalf("expected target hit, got exit=%v reason=%q", exit, reason)
	}
}

func TestShouldExit_StoplossHit(t *testing.T) {
	dc := &DhanClient{Target: 10, Stoploss: -5}
	if exit, reason := dc.ShouldExit(-5.1); !exit || reason != "🛑 Stop-loss hit" {
		t.Fatalf("expected stop-loss hit, got exit=%v reason=%q", exit, reason)
	}
}

func TestShouldExit_Trailing(t *testing.T) {
	dc := &DhanClient{Target: 100, Stoploss: -10, TrailStart: 10, TrailStep: 5, CurrentTrail: -10}
	if exit, _ := dc.ShouldExit(5); exit {
		t.Fatalf("expected no exit before trail start")
	}
	exit, _ := dc.ShouldExit(15)
	if exit {
		t.Fatalf("did not expect exit when moving trail")
	}
	if dc.CurrentTrail != -5 {
		t.Fatalf("expected CurrentTrail -5, got %.2f", dc.CurrentTrail)
	}
	if exit, reason := dc.ShouldExit(-6); !exit || reason != "🛑 Trailing SL hit (₹-5.00)" {
		t.Fatalf("expected trailing SL hit, got exit=%v reason=%q", exit, reason)
	}
}

func TestPlaceOrder(t *testing.T) {
	dc := &DhanClient{}
	if res := dc.PlaceOrder("buy"); res != "✅ Order placed: buy (simulated)" {
		t.Fatalf("unexpected result: %s", res)
	}
	if res := dc.PlaceOrder("invalid"); res != "❌ Invalid order type!" {
		t.Fatalf("unexpected invalid order result: %s", res)
	}
}
