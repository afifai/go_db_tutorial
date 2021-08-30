# MySQL tutorial Golang

*Afif A. Iskandar*

- email: <afif@ngodingpython.com>
- youtube: [NgodingPython](https://youtube.com/NgodingPython)
- github: [afifai](http://github.com/afifai)

Repositori ini merupakan material untuk training MySQL menggunakan golang

## Catatan Instalasi

### MySQL on Docker

Pastikan MySQL sudah berjalan di docker agar lebih mudah, apabila Anda belum menginstall docker, silahkan download docker [disini](https://www.docker.com/products/docker-desktop)

Apabila Anda sudah menginstall docker, pertama-tama anda perlu pull MySQL docker image dengan cara : 

```
docker pull mysql
```

Sebagai catatan, anda hanya perlu melakukan satu kali pull docker.

Selanjutnya, buat container yang menjalankan image docker diatas dengan cara :

```
docker run --name mysql_docker -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password mysql
```

Apabila port `3306` sudah dipakai oleh komputer anda, maka anda dapat mengganti bagian  `xxxx` pada `xxxx:3306` ke port yang tersedia. Anda juga dapat mengubah `password` sesuai yang anda inginkan.

Setelah container dijalankan, silahkan tes koneksi MySQL di terminal WSL pada VSCode. Apabila anda belum menginstal MySQL client di WSL, anda dapat menjalankan command berikut :

```
sudo apt update

sudo apt install mysql-client-core-8.0

```

setelah selesai instalasi silahkan coba command berikut (pastikan container mysql berjalan di docker)

```
mysql -u root -h 127.0.0.1 -P 3306 -p
```

Apabila anda terkoneksi dengan MySQL console, maka anda sudah siap

### GO

Saya asumsikan bahwa golang sudah terinstall dengan benar di komputer anda. Pertama-tama silahkan clone repository ini :

```
git clone https://github.com/afifai/go_db_tutorial.git

cd go_db_tutorial
```

selanjutnya inisiasi project go dengan cara :

```
go mod init

go mod tidy

```

setelah selesai, silahkan masuk ke salah satu folder, misalnya folder `basic_sql_with_go/` dengan cara

```
cd basic_sql_with_go
```

lalu coba jalankan program di folder tersebut dengan cara

```
go run .
```

## Authors

* **Afif A. Iskandar** - [afifai](https://github.com/afifai) [email](mailto:afifai@sci.ui.ac.id)
