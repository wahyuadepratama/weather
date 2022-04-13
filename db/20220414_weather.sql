-- phpMyAdmin SQL Dump
-- version 5.0.4
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Apr 14, 2022 at 12:50 AM
-- Server version: 10.4.16-MariaDB
-- PHP Version: 7.4.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `weather`
--

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` varchar(20) NOT NULL,
  `email` varchar(190) NOT NULL,
  `name` varchar(190) NOT NULL,
  `password` varchar(190) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `email`, `name`, `password`, `created_at`, `updated_at`) VALUES
('U1649868242', 'wahyu.mailist@gmail.com', 'Wahyu Ade Pratama', '$2a$08$yxLDlk/m5SLfcj8H/g3meOX3yXWg249ZY4bH9E7XFLI2StAG5gAHi', '2022-04-13 23:44:02', '2022-04-13 23:44:02'),
('U1649868391', 'wahyu.mailist1@gmail.com', 'Wahyu Ade Pratama', '$2a$08$cSLe7vlm7bhPo8zEnVh4f.NiS6wT6FBAuZfR9CVHgMPVYpQ7mMeey', '2022-04-13 23:46:31', '2022-04-13 23:46:31');

-- --------------------------------------------------------

--
-- Table structure for table `weather`
--

CREATE TABLE `weather` (
  `id` int(11) NOT NULL,
  `lat` double(9,4) NOT NULL,
  `lon` double(9,4) NOT NULL,
  `timezone` varchar(190) NOT NULL,
  `pressure` int(5) NOT NULL,
  `humidity` int(5) NOT NULL,
  `wind_speed` double(9,2) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `weather`
--

INSERT INTO `weather` (`id`, `lat`, `lon`, `timezone`, `pressure`, `humidity`, `wind_speed`, `created_at`) VALUES
(6, -6.2953, 106.6383, 'Asia/Jakarta', 1008, 88, 3.09, '2022-04-14 04:52:13'),
(7, -6.2953, 106.6383, 'Asia/Jakarta', 1008, 88, 1.54, '2022-04-14 05:06:50');

-- --------------------------------------------------------

--
-- Table structure for table `weather_detail`
--

CREATE TABLE `weather_detail` (
  `id` int(11) NOT NULL,
  `weather_id` int(11) NOT NULL,
  `main` varchar(50) NOT NULL,
  `description` varchar(190) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `weather_detail`
--

INSERT INTO `weather_detail` (`id`, `weather_id`, `main`, `description`) VALUES
(721, 6, 'Haze', 'haze'),
(721, 7, 'Haze', 'haze');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `email` (`email`);

--
-- Indexes for table `weather`
--
ALTER TABLE `weather`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `weather_detail`
--
ALTER TABLE `weather_detail`
  ADD PRIMARY KEY (`id`,`weather_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `weather`
--
ALTER TABLE `weather`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
