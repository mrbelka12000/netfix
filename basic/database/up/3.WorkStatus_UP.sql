CREATE TABLE IF NOT EXISTS  WorkStatus(
    WorkID integer not null ,
    Status boolean not null ,
    FOREIGN KEY (WorkID)REFERENCES works(id)
)