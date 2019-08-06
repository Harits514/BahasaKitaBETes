# BahasaKitaBETes
Tes BackEnd from BahasaKita

Pada tes ini digunakan GoLang dan MySQL

Petunjuk Pemakaian:
1. Download GoLang https://golang.org/dl/
2. Ikuti petunjuk instalasi di : https://golang.org/doc/install
3. Import musik.sql pada phpmyadmin untuk mendapatkan data dan struktur tabel
4. Buka terminal/cmd pada folder yang memuat main.go
5. Untuk menjalankan, input: go run main.go
6. Buka postman

Untuk GET, POST dan PATCH dilakukan menggunakan link http://localhost:8080/ , bedakan method pada postman tergantung yang ingin digunakan.

Input menggunakan JSON, durasi lagu dalam detik, untuk PATCH diperlukan id dari instance yang ingin di-update.

Contoh input POST:
![alt text](https://github.com/Harits514/BahasaKitaBETes/blob/master/HowToUseImage/HowToPOST.JPG)

Contoh input PATCH:
![alt text](https://github.com/Harits514/BahasaKitaBETes/blob/master/HowToUseImage/HowToPATCH.JPG)

Contoh input GET:
![alt text](https://github.com/Harits514/BahasaKitaBETes/blob/master/HowToUseImage/HowToGET.JPG)
