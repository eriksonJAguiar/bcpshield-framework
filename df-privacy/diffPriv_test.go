package diffPriv

import (
	"testing"
)

func TestDiffPriv(t *testing.T) {
	database := randomGenerate(0, 1, 1000)
	qr := query(database)
	pdfNoise := diffPriv(qr, database, 1)
	t.Log("The noise query is: ")
	t.Log(pdfNoise)
}
