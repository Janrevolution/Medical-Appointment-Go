# Set-up

1. Using of Visual Studio Code
2. Download and install [Go](https://go.dev/doc/install). Add to Path in environment variables
3. Place the project under ```xampp\htdocs``` as we are using mysql in terminal with ```git clone (git link)```
4. Download the sql file and transfer it to ```xampp\mysql\bin```
5. In the terminal create a database named ```mydb```
6. Import the downloade database to your created database using ```mysql -u root mydb < mydb.sql```
7. Done

# [TEMP] To-do:
<!-- 1. Add the logic in printing to have where date in appointment_details -->
2. Finish doctor view
3. Finish edit function in functions that need it 

# Creation of tables [FOR DOCUMENTATION]
```mysql
Create table tbl_doctor(
    doctor_id varchar(64) primary key,
    last_name varchar(32),
    first_name varchar(32),
    middle_name varchar(32),
    specialization varchar(32)
);

Create table tbl_room_doctor(
    rd_id varchar(64) primary key,
    doctor_id_fk varchar(64),
    room_id_fk varchar(64),
    foreign key (doctor_id_fk) references tbl_doctor(doctor_id),
    foreign key (room_id_fk) references tbl_rooms (room_id)
);

Create table tbl_patients(
    patient_id varchar(64) primary key,
    last_name varchar(32) not null,
    first_name varchar(32) not null,
    middle_name varchar(32) not null,
    age int(2) not null,
    gender varchar(6) not null
);

Create table tbl_appointment_details(
    reserve_id varchar(64),
    patient_id_fk varchar(64),
    rd_id varchar(64),
    date date not null,
    time varchar(64) not null,
    secretary_id varchar(64),
    description varchar(256),
    primary key (reserve_id, date, time),
    foreign key (patient_id_fk) references tbl_patients(patient_id),
    foreign key (rd_id) references tbl_time_doctor(rd_id),
    foreign key(secretary_id) references tbl_accounts(emp_id)
);

Create table tbl_accounts(
    emp_id varchar(64) primary key,
    username varchar(32) not null,
    password varchar(32) not null,
    foreign key (emp_id) references tbl_employees(emp_id)
);

Create table tbl_secretary(
    sec_id int primary key auto increment,
    name varchar(128) not null
);

Create table tbl_time(
    time_id varchar(64) primary key,
    start_time time not null,
    end_time time not null
);

Create table tbl_status(
    status_id varchar(64) primary key,
    status_name varchar(32) not null
);
<-Connect to doctor not room since we follow doctor time not room->

Create table tbl_time_doctor(
    rd_id varchar(64), 
    time_id varchar(64),
    ad_id varchar(64) not null,
    primary key(rd_id, time_id),
    foreign key(rd_id) references tbl_room_doctor(rd_id),
    foreign key(time_id) references tbl_time(time_id),
    unique (ad_id)
);

Create table tbl_avail_doctor(
    ad_id varchar(64),
    date date,
    primary key(ad_id, date),
    foreign key (ad_id) references tbl_time_doctor(ad_id)
);

Create table tbl_patient_diagnosis(
    reserve_id varchar(64) primary key,
    diagnosis varchar(512),
    doctor_id varchar(64),
    foreign key(reserve_id) references tbl_appointment_details (reserve_id),
    foreign key(doctor_id) references tbl_accounts(emp_id)
);

# Query for tbl_time
INSERT INTO tbl_time (time_id, start_time, end_time) 
VALUES 
    (UUID(), '01:00:00', '02:00:00'),
    (UUID(), '02:00:00', '03:00:00'),
    (UUID(), '03:00:00', '04:00:00'),
    (UUID(), '04:00:00', '05:00:00'),
    (UUID(), '05:00:00', '06:00:00'),
    (UUID(), '06:00:00', '07:00:00'),
    (UUID(), '07:00:00', '08:00:00'),
    (UUID(), '08:00:00', '09:00:00'),
    (UUID(), '09:00:00', '10:00:00'),
    (UUID(), '10:00:00', '11:00:00'),
    (UUID(), '11:00:00', '12:00:00'),
    (UUID(), '12:00:00', '13:00:00'),
    (UUID(), '13:00:00', '14:00:00'),
    (UUID(), '14:00:00', '15:00:00'),
    (UUID(), '15:00:00', '16:00:00'),
    (UUID(), '16:00:00', '17:00:00'),
    (UUID(), '17:00:00', '18:00:00'),
    (UUID(), '18:00:00', '19:00:00'),
    (UUID(), '19:00:00', '20:00:00'),
    (UUID(), '20:00:00', '21:00:00'),
    (UUID(), '21:00:00', '22:00:00'),
    (UUID(), '22:00:00', '23:00:00'),
    (UUID(), '23:00:00', '24:00:00'),
    (UUID(), '00:00:00', '01:00:00');

# Query for tbl_status
INSERT INTO tbl_status(status_id, status_name)
VALUES
    (UUID(), 'Available'),
    (UUID(), 'Unavailable');
```