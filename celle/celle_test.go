package cell

import "testing"

func TestInitCell(t *testing.T) {
    wanted := false; // Bruker semantikken i Golangs konseptuelle modell
    state := InitCell();
    if state != wanted {
        t.Errorf("Feil, fikk %t, ønsket %t.", state, wanted)
    }
}


func TestSetCellValue(t *testing.T) {
    wanted := true; // Bruker semantikken i Golangs konseptuelle modell
    state := SetCellValue(true);
    if state != wanted {
        t.Errorf("Feil, fikk %t, ønsket %t.", state, wanted)
    }
}


func TestGetCellValue(t *testing.T) {
    wanted := true; // Bruker semantikken i Golangs konseptuelle modell
    SetCellValue(true)
    state := GetCellValue();
    if state != wanted {
        t.Errorf("Feil, fikk %t, ønsket %t.", state, wanted)
    }
}