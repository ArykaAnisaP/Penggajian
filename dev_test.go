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
			Nama:         "anisa",
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
		Nama:         "anisa",
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

func TestInsertJamker(t *testing.T) {
	jam_masuk := "08.00"
	jam_keluar := "17.00"
	hari := "senin"
	shift := "pagi"
	piket_tim := "a"
	hasil := module.InsertJamker(module.MongoConn, "jamker", jam_masuk, jam_keluar, hari, shift, piket_tim)
	fmt.Println(hasil)
}

func TestInsertUang(t *testing.T) {
	tanggal := "01-09-2023"
	gaji_pokok := "5300000"
	bonus := model.Jamker{
		Jam_masuk:  "08.00",
		Jam_keluar: "17.00",
		Hari:       "Selasa",
		Shift:      "pagi",
		Piket_tim:  "a",
	}
	tunjangan_pph := "250000"
	potongan := "40000"
	biodata := model.Karyawan{
		Nama:         "winda",
		Phone_number: "08963253689",
		Jabatan:      " Computer Hardware Engineer",
		Hari_gaji:    "Senin",
	}
	absensi := model.Presensi{
		Location:     "Bali",
		Phone_number: "08963253689",
		Kehadiran:    "Hadir",
		Biodata: model.Karyawan{
			Nama:         "winda",
			Phone_number: "08963253689",
			Jabatan:      " Computer Hardware Engineer",
			Hari_gaji:    "Senin",
		},
		WaktuG: model.Waktu{
			Hari:    "Senin",
			Jam:     "09.00",
			Tanggal: "20-11-2023",
		},
	}
	hasil := module.InsertUang(module.MongoConn, "uang", tanggal, gaji_pokok, bonus, tunjangan_pph, potongan, biodata, absensi)
	fmt.Println(hasil)
}

func TestInsertBendahara(t *testing.T) {
	nama := "anisa"
	email := "aryka@gmail.com"
	phone_number := "086558747457"
	hari := "Kamis"
	biodata := model.Karyawan{
		Nama:         "antonio",
		Phone_number: "086558747457",
		Jabatan:      "programmer",
		Hari_gaji:    "Sabtu",
	}
	gaji := model.Uang{
		Tanggal:    "08-04-2023",
		Gaji_pokok: "4500000",
		Bonus: model.Jamker{
			Jam_masuk:  "15.00",
			Jam_keluar: "20.30",
			Hari:       "Sabtu",
			Shift:      "sore",
			Piket_tim:  "c",
		},
		Tunjangan_pph: "125000",
		Potongan:      "130000",
		Biodata: model.Karyawan{
			Nama:         "antonio",
			Phone_number: "08542136987",
			Hari_gaji:    "Kamis",
		},
		Absensi: model.Presensi{
			Location:     "Yogyakarta",
			Phone_number: "08542136987",
			Biodata: model.Karyawan{
				Nama:         "antonio",
				Phone_number: "08542136987",
				Jabatan:      "programmer",
				Hari_gaji:    "selasa",
			},
			WaktuG: model.Waktu{
				Hari:    "Kamis",
				Jam:     "09.00",
				Tanggal: "28-12-2023",
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

func TestGetPresensiFromWaktu(t *testing.T) {
	tanggal := "12-03-2023"
	data := module.GetPresensiFromWaktu(tanggal, module.MongoConn, "presensi")
	fmt.Println(data)
}

func TestGetKaryawanFromNama(t *testing.T) {
	nama := "nia"
	data := module.GetKaryawanFromNama(nama, module.MongoConn, "karyawan")
	fmt.Println(data)
}

func TestGetGajiFromPresensi(t *testing.T) {
	phone_number := "0854632178"
	data := module.GetGajiFromPresensi(phone_number, module.MongoConn, "uang")
	fmt.Println(data)
}

func TestGetAll(t *testing.T) {
	data := module.GetAllUang(module.MongoConn, "presensi")
	fmt.Println(data)
}
