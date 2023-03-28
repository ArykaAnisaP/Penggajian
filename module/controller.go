package module

import (
	"context"
	"fmt"

	"github.com/ArykaAnisaP/Penggajian/model"
	"github.com/aiteung/atdb"
	"go.mongodb.org/mongo-driver/bson"

	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

var MongoString string = os.Getenv("MONGOSTRING")

var MongoInfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "Penggajian_db",
}

var MongoConn = atdb.MongoConnect(MongoInfo)

func InsertOneDoc(db *mongo.Database, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertPimpinan(db *mongo.Database, col string, nama string, phone_number string, jabatan string, gaji model.Uang, bonus model.Jamker, biodata model.Karyawan, absensi model.Presensi) (InsertedID interface{}) {
	var pimpinan model.Pimpinan
	pimpinan.Nama = nama
	pimpinan.Phone_number = phone_number
	pimpinan.Jabatan = jabatan
	pimpinan.Gaji = gaji
	pimpinan.Bonus = bonus
	pimpinan.Biodata = biodata
	pimpinan.Absensi = absensi
	return InsertOneDoc(db, col, pimpinan)
}

func InsertPresensi(db *mongo.Database, col string, location string, phone_number string, kehadiran string, biodata model.Karyawan, waktug model.Waktu) (InsertedID interface{}) {
	var presensi model.Presensi
	presensi.Location = location
	presensi.Phone_number = phone_number
	presensi.Biodata = biodata
	presensi.WaktuG = waktug
	return InsertOneDoc(db, col, presensi)
}

func InsertKaryawan(db *mongo.Database, col string, nama string, phone_number string, jabatan string, hari_gaji string) (InsertedID interface{}) {
	var karyawan model.Karyawan
	karyawan.Nama = nama
	karyawan.Phone_number = phone_number
	karyawan.Jabatan = jabatan
	karyawan.Hari_gaji = hari_gaji
	return InsertOneDoc(db, col, karyawan)
}

func InsertWaktu(db *mongo.Database, col string, hari string, jam string, tanggal string) (InsertedID interface{}) {
	var waktu model.Waktu
	waktu.Hari = hari
	waktu.Jam = jam
	waktu.Tanggal = tanggal
	return InsertOneDoc(db, col, waktu)
}

func InsertJamker(db *mongo.Database, col string, jam_masuk string, jam_keluar string, hari string, shift string, piket_tim string) (InsertedID interface{}) {
	var jamker model.Jamker
	jamker.Jam_masuk = jam_masuk
	jamker.Jam_keluar = jam_keluar
	jamker.Hari = hari
	jamker.Shift = shift
	jamker.Piket_tim = piket_tim
	return InsertOneDoc(db, col, jamker)
}

func InsertUang(db *mongo.Database, col string, tanggal string, gaji_pokok string, bonus model.Jamker, tunjangan_pph string, potongan string, biodata model.Karyawan, absensi model.Presensi) (InsertedID interface{}) {
	var uang model.Uang
	uang.Tanggal = tanggal
	uang.Gaji_pokok = gaji_pokok
	uang.Bonus = bonus
	uang.Tunjangan_pph = tunjangan_pph
	uang.Potongan = potongan
	uang.Biodata = biodata
	uang.Absensi = absensi
	return InsertOneDoc(db, col, uang)
}

func InsertBendahara(db *mongo.Database, col string, nama string, email string, phone_number string, hari string, biodata model.Karyawan, gaji model.Uang) (InsertedID interface{}) {
	var bendahara model.Bendahara
	bendahara.Nama = nama
	bendahara.Email = email
	bendahara.Phone_number = phone_number
	bendahara.Hari = hari
	bendahara.Biodata = biodata
	bendahara.Gaji = gaji
	return InsertOneDoc(db, col, bendahara)
}

func GetGajiFromNamaKaryawan(nama string, db *mongo.Database, col string) (staf model.Uang) {
	Nm_karyawan := db.Collection(col)
	filter := bson.M{"biodata.nama": nama}
	err := Nm_karyawan.FindOne(context.TODO(), filter).Decode(&staf)
	if err != nil {
		fmt.Printf("getgajiFromKaryawan: %v\n", err)
	}
	return staf
}

func GetPresensiromWaktu(tanggal string, db *mongo.Database, col string) (kehadiran model.Presensi) {
	waktu := db.Collection(col)
	filter := bson.M{"tanggal": tanggal}
	err := waktu.FindOne(context.TODO(), filter).Decode(&kehadiran)
	if err != nil {
		fmt.Printf("GetPresensiromWaktu: %v\n", err)
	}
	return kehadiran
}

func GetKaryawanromJamKerja(jam_masuk string, db *mongo.Database, col string) (masuk model.Jamker) {
	jammasuk := db.Collection(col)
	filter := bson.M{"jam_masuk": jam_masuk}
	err := jammasuk.FindOne(context.TODO(), filter).Decode(&masuk)
	if err != nil {
		fmt.Printf("GetKaryawanromJamKerja: %v\n", err)
	}
	return masuk
}
