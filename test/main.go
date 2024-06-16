package main

import "fmt"

func main() {
    var a int = 42
    var p *int = &a  // p adalah pointer yang menyimpan alamat dari variabel a
    
    fmt.Println("Nilai a:", a)
    fmt.Println("Alamat a:", &a)
    fmt.Println("Pointer p:", p)
    fmt.Println("Nilai yang ditunjuk oleh p:", *p)  // Menggunakan dereferensi untuk mendapatkan nilai dari variabel a melalui pointer p
    
    *p = 21  // Mengubah nilai variabel a melalui pointer p
    fmt.Println("Nilai a setelah perubahan:", a)
}
