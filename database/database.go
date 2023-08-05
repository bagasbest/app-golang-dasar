package database

var connection string

// / KETIKA PACKAGE DATABASE DIAKSES DARRI PACKAGE LAIN
// / MAKA FUNCTION INIT AKAN DIJALANKAN LANGSUNG
func init() {
	connection = "MySQL"
}

// / SEHINGGA KETIKA GETDATABASE DI PANGGIL
// / KONEKSI AKAN MENJADI MYSQL
func GetDatabase() string {
	return connection
}
