package dto

type NasabahCreateDto struct {
	Nama             string `json:"nama" binding:"required"`
	NoKtp            string `json:"no_ktp" binding:"required"`
	TempatLahir      string `json:"tempat_lahir" binding:"required"`
	TanggalLahir     string `json:"tanggal_lahir" binding:"required" time_format:"2006-01-02"`
	StatusKawin      string `json:"status_kawin" binding:"required"`
	Alamat           string `json:"alamat" binding:"required"`
	Pekerjaan        string `json:"pekerjaan" binding:"required"`
	NamaPerusahaan   string `json:"nama_perusahaan" binding:"required"`
	AlamatPerusahaan string `json:"alamat_perusahaan" binding:"required"`
	NoHp             string `json:"no_hp" binding:"required"`
	Email            string `json:"email" binding:"required,email"`
	NamaIbuKandung   string `json:"nama_ibu_kandung" binding:"required"`
	JenisKelamin     string `json:"jenis_kelamin" binding:"required"`
}

type NasabahUpdateDto struct {
	Nama             string `json:"nama" binding:"required"`
	NoKtp            string `json:"no_ktp" binding:"required"`
	TempatLahir      string `json:"tempat_lahir" binding:"required"`
	TanggalLahir     string `json:"tanggal_lahir" binding:"required" time_format:"2006-01-02"`
	StatusKawin      string `json:"status_kawin" binding:"required"`
	Alamat           string `json:"alamat" binding:"required"`
	Pekerjaan        string `json:"pekerjaan" binding:"required"`
	NamaPerusahaan   string `json:"nama_perusahaan" binding:"required"`
	AlamatPerusahaan string `json:"alamat_perusahaan" binding:"required"`
	NoHp             string `json:"no_hp" binding:"required"`
	Email            string `json:"email" binding:"required,email"`
	NamaIbuKandung   string `json:"nama_ibu_kandung" binding:"required"`
	JenisKelamin     string `json:"jenis_kelamin" binding:"required"`
}