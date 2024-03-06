-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Mar 06, 2024 at 02:18 PM
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
-- Table structure for table `tbl_avail_doctor`
--

CREATE TABLE `tbl_avail_doctor` (
  `ad_id` varchar(64) NOT NULL,
  `date` date NOT NULL
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
('67efc7e0-4445-4370-a3dd-55bd2adfe8ae', 'HPID-1709663170820', 'TempLast', 'TempFirst', 'TempMiddle', 'Doctor', 'General Surgery'),
('7ce88b9f-d80f-4945-a018-e448620d3f76', 'HPID-1709657591169', 'Villadores', 'Janrev Lance', 'Florig', 'Doctor', 'Cardiologist');

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
('86244457-fe8a-4613-bcbc-35dfbf67dfed', '2', 'Basic'),
('95082190-9d88-483a-88b6-80b8723bc93d', '1', 'Basic');

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
('309369fd-3510-480e-9734-6c75429dfa1f', '7ce88b9f-d80f-4945-a018-e448620d3f76', '95082190-9d88-483a-88b6-80b8723bc93d'),
('46435dcb-8fef-488a-89d5-0aa011a15dec', '67efc7e0-4445-4370-a3dd-55bd2adfe8ae', '86244457-fe8a-4613-bcbc-35dfbf67dfed');

-- --------------------------------------------------------

--
-- Table structure for table `tbl_time`
--

CREATE TABLE `tbl_time` (
  `time_id` varchar(64) NOT NULL,
  `start_time` time NOT NULL,
  `end_time` time NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `tbl_time`
--

INSERT INTO `tbl_time` (`time_id`, `start_time`, `end_time`) VALUES
('88fa065d-db0d-11ee-9efc-902e16b789a2', '01:00:00', '02:00:00'),
('88fa1a44-db0d-11ee-9efc-902e16b789a2', '02:00:00', '03:00:00'),
('88fa1ad2-db0d-11ee-9efc-902e16b789a2', '03:00:00', '04:00:00'),
('88fa1b15-db0d-11ee-9efc-902e16b789a2', '04:00:00', '05:00:00'),
('88fa1b45-db0d-11ee-9efc-902e16b789a2', '05:00:00', '06:00:00'),
('88fa1b71-db0d-11ee-9efc-902e16b789a2', '06:00:00', '07:00:00'),
('88fa1ba0-db0d-11ee-9efc-902e16b789a2', '07:00:00', '08:00:00'),
('88fa1bc7-db0d-11ee-9efc-902e16b789a2', '08:00:00', '09:00:00'),
('88fa1bf6-db0d-11ee-9efc-902e16b789a2', '09:00:00', '10:00:00'),
('88fa1c51-db0d-11ee-9efc-902e16b789a2', '10:00:00', '11:00:00'),
('88fa1c96-db0d-11ee-9efc-902e16b789a2', '11:00:00', '12:00:00'),
('88fa1cc2-db0d-11ee-9efc-902e16b789a2', '12:00:00', '13:00:00'),
('88fa1cef-db0d-11ee-9efc-902e16b789a2', '13:00:00', '14:00:00'),
('88fa1d1b-db0d-11ee-9efc-902e16b789a2', '14:00:00', '15:00:00'),
('88fa1d41-db0d-11ee-9efc-902e16b789a2', '15:00:00', '16:00:00'),
('88fa1d6b-db0d-11ee-9efc-902e16b789a2', '16:00:00', '17:00:00'),
('88fa1d96-db0d-11ee-9efc-902e16b789a2', '17:00:00', '18:00:00'),
('88fa1dc0-db0d-11ee-9efc-902e16b789a2', '18:00:00', '19:00:00'),
('88fa1def-db0d-11ee-9efc-902e16b789a2', '19:00:00', '20:00:00'),
('88fa1e1c-db0d-11ee-9efc-902e16b789a2', '20:00:00', '21:00:00'),
('88fa1e44-db0d-11ee-9efc-902e16b789a2', '21:00:00', '22:00:00'),
('88fa1e71-db0d-11ee-9efc-902e16b789a2', '22:00:00', '23:00:00'),
('88fa1e9f-db0d-11ee-9efc-902e16b789a2', '23:00:00', '24:00:00'),
('88fa1ed7-db0d-11ee-9efc-902e16b789a2', '00:00:00', '01:00:00');

-- --------------------------------------------------------

--
-- Table structure for table `tbl_time_doctor`
--

CREATE TABLE `tbl_time_doctor` (
  `rd_id` varchar(64) NOT NULL,
  `time_id` varchar(64) NOT NULL,
  `ad_id` varchar(64) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `tbl_time_doctor`
--

INSERT INTO `tbl_time_doctor` (`rd_id`, `time_id`, `ad_id`) VALUES
('309369fd-3510-480e-9734-6c75429dfa1f', '88fa1e1c-db0d-11ee-9efc-902e16b789a2', '2e74dd53-c198-468f-8c55-e7d087989640'),
('309369fd-3510-480e-9734-6c75429dfa1f', '88fa1dc0-db0d-11ee-9efc-902e16b789a2', 'd8adaeb2-d3a4-46f0-b689-569d320daa51');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `tbl_accounts`
--
ALTER TABLE `tbl_accounts`
  ADD PRIMARY KEY (`emp_id`);

--
-- Indexes for table `tbl_avail_doctor`
--
ALTER TABLE `tbl_avail_doctor`
  ADD PRIMARY KEY (`ad_id`,`date`);

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
-- Indexes for table `tbl_time`
--
ALTER TABLE `tbl_time`
  ADD PRIMARY KEY (`time_id`);

--
-- Indexes for table `tbl_time_doctor`
--
ALTER TABLE `tbl_time_doctor`
  ADD PRIMARY KEY (`rd_id`,`time_id`),
  ADD UNIQUE KEY `ad_id` (`ad_id`),
  ADD KEY `time_id` (`time_id`);

--
-- Constraints for dumped tables
--

--
-- Constraints for table `tbl_accounts`
--
ALTER TABLE `tbl_accounts`
  ADD CONSTRAINT `tbl_accounts_ibfk_1` FOREIGN KEY (`emp_id`) REFERENCES `tbl_employees` (`emp_id`);

--
-- Constraints for table `tbl_avail_doctor`
--
ALTER TABLE `tbl_avail_doctor`
  ADD CONSTRAINT `tbl_avail_doctor_ibfk_1` FOREIGN KEY (`ad_id`) REFERENCES `tbl_time_doctor` (`ad_id`);

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

--
-- Constraints for table `tbl_time_doctor`
--
ALTER TABLE `tbl_time_doctor`
  ADD CONSTRAINT `tbl_time_doctor_ibfk_1` FOREIGN KEY (`rd_id`) REFERENCES `tbl_room_doctor` (`rd_id`),
  ADD CONSTRAINT `tbl_time_doctor_ibfk_2` FOREIGN KEY (`time_id`) REFERENCES `tbl_time` (`time_id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
