-- phpMyAdmin SQL Dump
-- version 4.8.3
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 06 Agu 2019 pada 14.43
-- Versi server: 10.1.37-MariaDB
-- Versi PHP: 7.2.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `dbtest`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `musik`
--

CREATE TABLE `musik` (
  `namalagu` text NOT NULL,
  `namaartis` text NOT NULL,
  `link` text NOT NULL,
  `durasi` int(11) NOT NULL,
  `id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `musik`
--

INSERT INTO `musik` (`namalagu`, `namaartis`, `link`, `durasi`, `id`) VALUES
('ahoy', 'harits', 'gophyuu.com', 0, 1),
('JohnsonJohnson', 'Hubert', 'gophyaa.com', 178, 2),
('jedar', 'huruharax2', 'cekcekcek.com', 150, 3);

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `musik`
--
ALTER TABLE `musik`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `musik`
--
ALTER TABLE `musik`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
