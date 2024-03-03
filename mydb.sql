-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Mar 03, 2024 at 09:05 PM
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
-- Table structure for table `tbl_accounts`
--

CREATE TABLE `tbl_accounts` (
  `emp_id` varchar(64) NOT NULL,
  `username` varchar(32) NOT NULL,
  `password` varchar(32) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `tbl_employees`
--

CREATE TABLE `tbl_employees` (
  `emp_id` varchar(64) NOT NULL,
  `hp_id` varchar(64) NOT NULL,
  `last_name` varchar(32) NOT NULL,
  `first_name` varchar(32) NOT NULL,
  `middle_name` varchar(32) NOT NULL,
  `profession` varchar(64) NOT NULL,
  `specialization` varchar(32) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `tbl_employees`
--

INSERT INTO `tbl_employees` (`emp_id`, `hp_id`, `last_name`, `first_name`, `middle_name`, `profession`, `specialization`) VALUES
('1c86ce4d-3732-4cc7-8c74-26c4fec0ca72', 'HPID-1709494313206', 'Villadores', 'Janrev', 'Florig', 'Doctor', 'Cardiologist'),
('57b195c4-3894-4512-baa8-d62e60a70754', 'HPID-1709494505751', 'Villadores', 'Lance', 'Florig', 'Secretary', 'N/A');

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
  `time` time NOT NULL,
  `secretary_id` varchar(64) DEFAULT NULL,
  `doctor_id` varchar(64) DEFAULT NULL
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

--
-- Dumping data for table `tbl_rooms`
--

INSERT INTO `tbl_rooms` (`room_id`, `room_number`, `room_type`) VALUES
('b5cd0284-6c4d-47bb-89a9-40c5d4bc38e2', '1', 'Basic');

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
-- Dumping data for table `tbl_room_doctor`
--

INSERT INTO `tbl_room_doctor` (`rd_id`, `doctor_id_fk`, `room_id_fk`) VALUES
('8b7cf4b3-c4fe-4ce8-93c9-116e128dae14', '1c86ce4d-3732-4cc7-8c74-26c4fec0ca72', 'b5cd0284-6c4d-47bb-89a9-40c5d4bc38e2');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `tbl_accounts`
--
ALTER TABLE `tbl_accounts`
  ADD PRIMARY KEY (`emp_id`);

--
-- Indexes for table `tbl_employees`
--
ALTER TABLE `tbl_employees`
  ADD PRIMARY KEY (`emp_id`);

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
  ADD KEY `rd_id_fk` (`rd_id_fk`),
  ADD KEY `secretary_id` (`secretary_id`),
  ADD KEY `doctor_id` (`doctor_id`);

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
-- Constraints for table `tbl_accounts`
--
ALTER TABLE `tbl_accounts`
  ADD CONSTRAINT `tbl_accounts_ibfk_1` FOREIGN KEY (`emp_id`) REFERENCES `tbl_employees` (`emp_id`);

--
-- Constraints for table `tbl_reservation_details`
--
ALTER TABLE `tbl_reservation_details`
  ADD CONSTRAINT `tbl_reservation_details_ibfk_1` FOREIGN KEY (`patient_id_fk`) REFERENCES `tbl_patients` (`patient_id`),
  ADD CONSTRAINT `tbl_reservation_details_ibfk_2` FOREIGN KEY (`rd_id_fk`) REFERENCES `tbl_room_doctor` (`rd_id`),
  ADD CONSTRAINT `tbl_reservation_details_ibfk_3` FOREIGN KEY (`secretary_id`) REFERENCES `tbl_accounts` (`emp_id`),
  ADD CONSTRAINT `tbl_reservation_details_ibfk_4` FOREIGN KEY (`doctor_id`) REFERENCES `tbl_accounts` (`emp_id`);

--
-- Constraints for table `tbl_room_doctor`
--
ALTER TABLE `tbl_room_doctor`
  ADD CONSTRAINT `tbl_room_doctor_ibfk_1` FOREIGN KEY (`doctor_id_fk`) REFERENCES `tbl_employees` (`emp_id`),
  ADD CONSTRAINT `tbl_room_doctor_ibfk_2` FOREIGN KEY (`room_id_fk`) REFERENCES `tbl_rooms` (`room_id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
