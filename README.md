# Set-up

1. Using of Visual Studio Code
2. Download and install [Go](https://go.dev/doc/install). Add to Path in environment variables
3. Place the projects under ```xampp\htdocs``` as we are using mysql

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

Create table tbl_reservation_details(
    reserve_id varchar(64) primary key,
    patient_id_fk varchar(64),
    rd_id_fk varchar(64),
    description varchar(256) not null,
    date date not null,
    time time not null,
    foreign key (patient_id_fk) references tbl_patients(patient_id),
    foreign key (rd_id_fk) references tbl_room_doctor(rd_id)
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
```