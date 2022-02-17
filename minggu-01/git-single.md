#PRAKTIKUM TEKNOLOGI CLOUD

### Nama: NURUL HALIMAH

### NIM: 215611106

### Prodi: Sistem Informasi

##A. Instalasi GIT di Windows

1. Setelah selesai download Git, jalankan program GIT. Akan dimunculkan lisensi. Klik Next untuk lanjut.

![01](image/Installgit/1.png)

2. Setelah itu, pilih lokasi instalasi. Secara default akan terisi C:\Program Files\Git. Ganti lokasi jika memang anda menginginkan lokasi lain. Klik Next untuk lanjut.

![02](image/Installgit/2.png)

3. Pilih komponen. Tidak perlu diubah-ubah, sesuai dengan default saja. Klik pada Next.

![03](image/Installgit/3.png)

4. Mengisi shortcut untuk menu Start. Gunakan default (Git), ganti jika ingin mengganti - misalnya Git VCS.

![04](image/Installgit/4.png)

5. Pilih editor yang akan digunakan bersama dengan Git. Pada pilihan ini, digunakan Visual Studio Code.

![05](image/Installgit/5.png)

6. Setelah itu proses instalasi akan dilakukan.

![15](image/Installgit/15.png)

7. Jika selesai akan muncul dialog pemberitahuan. Klik pada Finish.

![16](image/Installgit/16.png)

8. Untuk menjalankan, dari Start, ketikkan "Git", akan muncul beberapa pilihan. Pilih "Git Bash" atau "Git GUI".

![17](image/Installgit/17.png)

9. Tampilan jika akan menggunakan Git Bash

![18](image/Installgit/18.png)

10. Tampilan jika akan menggunakan Git GUI

![19](image/Installgit/19.png)

11. Untuk mencoba dari command prompt, masuk ke command prompt, setelah itu eksekusi "git --version" untuk melihat apakah sudah terinstall atau belum. Jika sudah terinstall dengan benar, makan akan muncul hasil berikut:

![20](image/Installgit/20.png)

##B. Konfigurasi GIT
Untuk melakukan konfigurasi GIT bisa melakukan melalui Command prompt dengan memakai perintah sebagai berikut.
Konfigurasi Username dan Email Isian di bawah harus disesuaikan dengan nama serta email yang digunakan untuk mendaftar di GitHub.

1. Konfigurasi Username dan Email Isian di bawah harus disesuaikan dengan nama serta email yang digunakan untuk mendaftar di GitHub.

```sh
$ git config --global user.name "Nama Anda di GitHub"
$ git config --global user.email email@domain.tld
```

2. untuk melihat hasil konfigurasi dengan perintah

```sh
$ git config --list
```

untuk hasilnya seperti ini
![01](image/konfig/1.png)

##C. Mengelola Repo Sendiri di Account Sendiri

Untuk membuat repo, gunakan langkah-langkan berikut:

1. Klik tanda + pada bagian atas setelah login, pilih New repository
   ![01](image/git/1.png)

2. Isikan nama, keterangan, serta lisensi. Jika dikehendaki, bisa membuat repo Private
   ![02](image/git/2.png)

##D. Mengelola Repo Sendiri di Organisasi
Untuk membuat repo, gunakan langkah-langkan berikut:

1. Masuk pada Organisasi Kalian lalu Klik tanda + pada bagian atas setelah login, pilih New repository

![04](image/git/4.png)

2. Isikan nama, keterangan, serta lisensi. Jika dikehendaki, bisa membuat repo Private

![05](image/git/5.png)
