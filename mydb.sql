-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Mar 03, 2024 at 12:11 PM
-- Server version: 10.4.32-MariaDB
-- PHP Version: 8.2.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `mydb`
--

-- --------------------------------------------------------

--
-- Table structure for table `tbl_doctors`
--

CREATE TABLE `tbl_doctors` (
  `doctor_id` varchar(64) NOT NULL,
  `hp_id` varchar(64) NOT NULL,
  `last_name` varchar(32) NOT NULL,
  `first_name` varchar(32) NOT NULL,
  `middle_name` varchar(32) NOT NULL,
  `specialization` varchar(32) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `tbl_patients`
--

CREATE TABLE `tbl_patients` (
  `patient_id` varchar(64) NOT NULL,
  `last_name` varchar(32) NOT NULL,
  `first_name` varchar(32) NOT NULL,
  `middle_name` varchar(32) NOT NULL,
  `age` int(2) NOT NULL,
  `gender` varchar(6) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `tbl_reservation_details`
--

CREATE TABLE `tbl_reservation_details` (
  `reserve_id` varchar(64) NOT NULL,
  `patient_id_fk` varchar(64) DEFAULT NULL,
  `rd_id_fk` varchar(64) DEFAULT NULL,
  `description` varchar(256) NOT NULL,
  `date` date NOT NULL,
  `time` time NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `tbl_rooms`
--

CREATE TABLE `tbl_rooms` (
  `room_id` varchar(64) NOT NULL,
  `room_number` varchar(3) NOT NULL,
  `room_type` varchar(32) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `tbl_room_doctor`
--

CREATE TABLE `tbl_room_doctor` (
  `rd_id` varchar(64) NOT NULL,
  `doctor_id_fk` varchar(64) DEFAULT NULL,
  `room_id_fk` varchar(64) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `tbl_doctors`
--
ALTER TABLE `tbl_doctors`
  ADD PRIMARY KEY (`doctor_id`);

--
-- Indexes for table `tbl_patients`
--
ALTER TABLE `tbl_patients`
  ADD PRIMARY KEY (`patient_id`);

--
-- Indexes for table `tbl_reservation_details`
--
ALTER TABLE `tbl_reservation_details`
  ADD PRIMARY KEY (`reserve_id`),
  ADD KEY `patient_id_fk` (`patient_id_fk`),
  ADD KEY `rd_id_fk` (`rd_id_fk`);

--
-- Indexes for table `tbl_rooms`
--
ALTER TABLE `tbl_rooms`
  ADD PRIMARY KEY (`room_id`);

--
-- Indexes for table `tbl_room_doctor`
--
ALTER TABLE `tbl_room_doctor`
  ADD PRIMARY KEY (`rd_id`),
  ADD KEY `doctor_id_fk` (`doctor_id_fk`),
  ADD KEY `room_id_fk` (`room_id_fk`);

--
-- Constraints for dumped tables
--

--
-- Constraints for table `tbl_reservation_details`
--
ALTER TABLE `tbl_reservation_details`
  ADD CONSTRAINT `tbl_reservation_details_ibfk_1` FOREIGN KEY (`patient_id_fk`) REFERENCES `tbl_patients` (`patient_id`),
  ADD CONSTRAINT `tbl_reservation_details_ibfk_2` FOREIGN KEY (`rd_id_fk`) REFERENCES `tbl_room_doctor` (`rd_id`);

--
-- Constraints for table `tbl_room_doctor`
--
ALTER TABLE `tbl_room_doctor`
  ADD CONSTRAINT `tbl_room_doctor_ibfk_1` FOREIGN KEY (`doctor_id_fk`) REFERENCES `tbl_doctors` (`doctor_id`),
  ADD CONSTRAINT `tbl_room_doctor_ibfk_2` FOREIGN KEY (`room_id_fk`) REFERENCES `tbl_rooms` (`room_id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
