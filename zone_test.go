package mtga

import "testing"

func TestChangeZoneNormal(t *testing.T) {
	l := RawLog{
		Body: []string{`<<<<<<<<<< ZoneChange of type Destroy for ["Law-Rune Enforcer" InstanceId:294, GrpId:67740] ("Law-Rune Enforcer") had Instigator 291 ("Tyrant's Scorn").`},
	}

	parser := Parser{}
	parser.OnZoneChange(func(change ZoneChange) {
		if change.Type != Destroy || change.Target != "Law-Rune Enforcer" || change.InstanceId != 294 ||
			change.GrpId != 67740 || change.Instigator != 291 || change.Source != "Tyrant's Scorn" {
			t.Error()
		}
	})
	parser.Parse(l)
}

func TestChangeZoneNull(t *testing.T) {
	l := RawLog{
		Body: []string{`<<<<<<<<<< ZoneChange of type ZeroToughness for 338 ("[NULL]") had Instigator 334 ("Cry of the Carnarium").`},
	}

	parser := Parser{}
	parser.OnZoneChange(func(change ZoneChange) {
		if change.Type != ZeroToughness || change.Target != "NULL" || change.InstanceId != 338 ||
			change.GrpId != 0 || change.Instigator != 334 || change.Source != "Cry of the Carnarium" {
			t.Error()
		}
	})
	parser.Parse(l)
}
