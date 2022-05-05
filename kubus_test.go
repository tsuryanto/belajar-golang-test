package main

import "testing"

var (
	kubus              Kubus   = Kubus{4}
	volumeSeharusnya   float64 = 64
	luasSeharusnya     float64 = 96
	kelilingSeharusnya float64 = 48
)

func TestVolume(t *testing.T) {
	t.Log("Volume : ", volumeSeharusnya)
	if kubus.Volume() != volumeSeharusnya {
		t.Error("Salah, seharusnya ", volumeSeharusnya)
	}
}

func TestKeliling(t *testing.T) {
	t.Log("Keliling : ", kelilingSeharusnya)
	if kubus.Keliling() != kelilingSeharusnya {
		t.Error("Salah, Keliling seharusnya ", kelilingSeharusnya)
	}
}

func TestLuas(t *testing.T) {
	t.Log("Luas : ", luasSeharusnya)
	if kubus.Luas() != luasSeharusnya {
		t.Error("Salah, Luas seharusnya ", luasSeharusnya)
	}
}
