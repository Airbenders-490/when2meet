CREATE TABLE when2meet(
W2W_ID VARCHAR(10) primary key not null,
W2W_name VARCHAR(10)
);

create table student(
sid VARCHAR(10) primary key not null,
name VARCHAR(10),
W2W_ID VARCHAR(10),
foreign key (W2W_ID) references when2meet(W2W_ID)
);


create table possibilities( 
PID VARCHAR(10) primary	key not null,
PDay VARCHAR(10),
start_Time integer,
end_Time integer
);

-- DROP TABLE possibilities;



create table Availibilities(
AID VARCHAR(10) primary	key not null,
SID VARCHAR(10),
ADAY VARCHAR(10),
Times integer[2]
);

-- DROP TABLE Availibilities;

create table sets_possibilities( 
W2W_ID VARCHAR(10),
PID VARCHAR(10),
FOREIGN key (W2W_ID) references when2meet(W2W_ID),
FOREIGN key (PID) references Possibilities(PID),
primary key (W2W_ID, PID)
);

-- DROP TABLE sets_possibilities;

create table choose_availibilities(
SID VARCHAR(10),
AID VARCHAR(10),
FOREIGN key (SID) references Student(SID),
FOREIGN key (AID) references Availibilities(AID),
primary key (SID, AID)
);

-- DROP TABLE choose_availibilities;
