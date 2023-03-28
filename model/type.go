package model

type Pimpinan struct {
	Nama         string   `bson:"nama,omitempty" json:"nama,omitempty"`
	Phone_number string   `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Jabatan      string   `bson:"jabatan,omitempty" json:"jabatan,omitempty"`
	Gaji         Uang     `bson:"gaji,omitempty" json:"gaji,omitempty"`
	Bonus        Jamker   `bson:"bonus,omitempty" json:"bonus,omitempty"`
	Biodata      Karyawan `bson:"biodata,omitempty" json:"biodata,omitempty"`
	Absensi      Presensi `bson:"absensi,omitempty" json:"absensi,omitempty"`
}

type Presensi struct {
	Location     string   `bson:"location,omitempty" json:"location,omitempty"`
	Phone_number string   `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Kehadiran    string   `bson:"kehadiran,omitempty" json:"kehadiran,omitempty"`
	Biodata      Karyawan `bson:"biodata,omitempty" json:"biodata,omitempty"`
	WaktuG       Waktu    `bson:"waktug,omitempty" json:"waktug,omitempty"`
}

type Karyawan struct {
	Nama         string `bson:"nama,omitempty" json:"nama,omitempty"`
	Phone_number string `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Jabatan      string `bson:"jabatan,omitempty" json:"jabatan,omitempty"`
	Hari_gaji    string `bson:"hari_gaji,omitempty" json:"hari_gaji,omitempty"`
}

type Waktu struct {
	Hari    string `bson:"hari,omitempty" json:"hari,omitempty"`
	Jam     string `bson:"jam,omitempty" json:"jam,omitempty"`
	Tanggal string `bson:"tanggal,omitempty" json:"tanggal,omitempty"`
}

type Jamker struct {
	Jam_masuk  string `bson:"jam_masuk,omitempty" json:"jam_masuk,omitempty"`
	Jam_keluar string `bson:"jam_keluar,omitempty" json:"jam_keluar,omitempty"`
	Hari       string `bson:"hari,omitempty" json:"hari,omitempty"`
	Shift      string `bson:"shift,omitempty" json:"shift,omitempty"`
	Piket_tim  string `bson:"piket_tim,omitempty" json:"piket_tim,omitempty"`
}

type Uang struct {
	Tanggal       string   `bson:"tanggal,omitempty" json:"tanggal,omitempty"`
	Gaji_pokok    string   `bson:"gaji_pokok,omitempty" json:"gaji_pokok,omitempty"`
	Bonus         Jamker   `bson:"bonus,omitempty" json:"bonus,omitempty"`
	Tunjangan_pph string   `bson:"tunjangan_pph,omitempty" json:"Tunjangan_pph,omitempty"`
	Potongan      string   `bson:"potongan,omitempty" json:"potongan,omitempty"`
	Biodata       Karyawan `bson:"biodata,omitempty" json:"biodata,omitempty"`
	Absensi       Presensi `bson:"absensi,omitempty" json:"absensi,omitempty"`
}

type Bendahara struct {
	Nama         string   `bson:"nama,omitempty" json:"nama,omitempty"`
	Email        string   `bson:"email,omitempty" json:"email,omitempty"`
	Phone_number string   `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Hari         string   `bson:"hari,omitempty" json:"hari,omitempty"`
	Biodata      Karyawan `bson:"biodata,omitempty" json:"biodata,omitempty"`
	Gaji         Uang     `bson:"gaji,omitempty" json:"gaji,omitempty"`
}
