package npm

import (
	"fmt"
	"testing"

	"github.com/ArykaAnisaP/Penggajian/model"
	"github.com/ArykaAnisaP/Penggajian/module"
)

func TestInsertPimpinan(t *testing.T) {
	nama := "aryka"
	phone_number := "080123412340"
	jabatan := "programmer"
	gaji := model.Uang{
		Tanggal:    "08-04-2023",
		Gaji_pokok: "100000000",
		Bonus: model.Jamker{
			Jam_masuk:  "08.00",
			Jam_keluar: "17.00",
			Hari:       "senin",
			Shift:      "pagi",
			Piket_tim:  "a",
		},
		Tunjangan_pph: "3200000",
		Potongan:      "5670000",
		Biodata: model.Karyawan{
			Nama:         "aryka",
			Phone_number: "09897656565",
			Jabatan:      "ui/ux",
			Hari_gaji:    "selasa",
		},
		Absensi: model.Presensi{
			Location:     "Bandung",
			Phone_number: "085214235678",
			Biodata: model.Karyawan{
				Nama:         "aryka",
				Phone_number: "09897656565",
				Jabatan:      "ui/ux",
				Hari_gaji:    "selasa",
			},
			WaktuG: model.Waktu{
				Hari:    "Senin",
				Jam:     "08.00",
				Tanggal: "12-03-2023",
			},
		},
	}
	bonus := model.Jamker{
		Jam_masuk:  "08.00",
		Jam_keluar: "17.00",
		Hari:       "senin",
		Shift:      "pagi",
		Piket_tim:  "a",
	}
	biodata := model.Karyawan{
		Nama:         "aryka",
		Phone_number: "09897656565",
		Jabatan:      "ui/ux",
		Hari_gaji:    "selasa",
	}
	absensi := model.Presensi{
		Location:     "Bandung",
		Phone_number: "085214235678",
		Biodata: model.Karyawan{
			Nama:         "aryka",
			Phone_number: "09897656565",
			Jabatan:      "ui/ux",
			Hari_gaji:    "selasa",
		},
		WaktuG: model.Waktu{
			Hari:    "Senin",
			Jam:     "08.00",
			Tanggal: "12-03-2023",
		},
	}
	hasil := module.InsertPimpinan(module.MongoConn, "pimpinan", nama, phone_number, jabatan, gaji, bonus, biodata, absensi)
	fmt.Println(hasil)
}

func TestInsertPresensi(t *testing.T) {
	location := "Bandung"
	phone_number := "0854632178"
	kehadiran := "Hadir"
	biodata := model.Karyawan{
		Nama:         "aryka",
		Phone_number: "09897656565",
		Jabatan:      "ui/ux",
		Hari_gaji:    "selasa",
	}
	waktug := model.Waktu{
		Hari:    "Senin",
		Jam:     "08.00",
		Tanggal: "12-03-2023",
	}
	hasil := module.InsertPresensi(module.MongoConn, "presensi", location, phone_number, kehadiran, biodata, waktug)
	fmt.Println(hasil)
}

func TestInsertKaryawan(t *testing.T) {
	nama := "nia"
	phone_number := "06945322332353"
	jabatan := "ui/ux"
	hari_gaji := "senin"
	hasil := module.InsertKaryawan(module.MongoConn, "karyawan", nama, phone_number, jabatan, hari_gaji)
	fmt.Println(hasil)
}

func TestInsertWaktu(t *testing.T) {
	hari := "Senin"
	jam := "08.00"
	tanggal := "12-03-2023"
	hasil := module.InsertWaktu(module.MongoConn, "waktu", hari, jam, tanggal)
	fmt.Println(hasil)
}

func TestInserJamker(t *testing.T) {
	jam_masuk := "08.00"
	jam_keluar := "17.00"
	hari := "senin"
	shift := "pagi"
	piket_tim := "a"
	hasil := module.InsertJamker(module.MongoConn, "jamker", jam_masuk, jam_keluar, hari, shift, piket_tim)
	fmt.Println(hasil)
}

func TestInsertUang(t *testing.T) {
	tanggal := "29-03-2023"
	gaji_pokok := "100000000"
	bonus := model.Jamker{
		Jam_masuk:  "08.00",
		Jam_keluar: "17.00",
		Hari:       "senin",
		Shift:      "pagi",
		Piket_tim:  "a",
	}
	tunjangan_pph := "12000000"
	potongan := "5000000"
	biodata := model.Karyawan{
		Nama:         "aryka",
		Phone_number: "09897656565",
		Jabatan:      "ui/ux",
		Hari_gaji:    "selasa",
	}
	absensi := model.Presensi{
		Location:     "Bandung",
		Phone_number: "0854632178",
		Kehadiran:    "Hadir",
		Biodata: model.Karyawan{
			Nama:         "aryka",
			Phone_number: "09897656565",
			Jabatan:      "ui/ux",
			Hari_gaji:    "selasa",
		},
		WaktuG: model.Waktu{
			Hari:    "Senin",
			Jam:     "08.00",
			Tanggal: "12-03-2023",
		},
	}
	hasil := module.InsertUang(module.MongoConn, "uang", tanggal, gaji_pokok, bonus, tunjangan_pph, potongan, biodata, absensi)
	fmt.Println(hasil)
}

func TestInserBendahara(t *testing.T) {
	nama := "anisa"
	email := "aryka@gm,ail.com"
	phone_number := "085842138"
	hari := "selasa"
	biodata := model.Karyawan{
		Nama:         "aryka",
		Phone_number: "09897656565",
		Jabatan:      "ui/ux",
		Hari_gaji:    "selasa",
	}
	gaji := model.Uang{
		Tanggal:    "08-04-2023",
		Gaji_pokok: "100000000",
		Bonus: model.Jamker{
			Jam_masuk:  "08.00",
			Jam_keluar: "17.00",
			Hari:       "senin",
			Shift:      "pagi",
			Piket_tim:  "a",
		},
		Tunjangan_pph: "3200000",
		Potongan:      "5670000",
		Biodata: model.Karyawan{
			Nama:         "aryka",
			Phone_number: "09897656565",
			Hari_gaji:    "selasa",
		},
		Absensi: model.Presensi{
			Location:     "Bandung",
			Phone_number: "085214235678",
			Biodata: model.Karyawan{
				Nama:         "aryka",
				Phone_number: "09897656565",
				Jabatan:      "ui/ux",
				Hari_gaji:    "selasa",
			},
			WaktuG: model.Waktu{
				Hari:    "Senin",
				Jam:     "08.00",
				Tanggal: "12-03-2023",
			},
		},
	}
	hasil := module.InsertBendahara(module.MongoConn, "bendahara", nama, email, phone_number, hari, biodata, gaji)
	fmt.Println(hasil)
}

func TestGetGajiFromNamaKaryawan(t *testing.T) {
	nama := "aryka"
	biodata := module.GetGajiFromNamaKaryawan(nama, module.MongoConn, "uang")
	fmt.Println(biodata)
}
