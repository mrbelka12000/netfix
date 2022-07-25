CREATE TABLE IF NOT EXISTS Wallets(
    Id serial primary key,
    OwnerID integer not null,
    Amount float8 not null,
    FOREIGN KEY (OwnerID) REFERENCES general(id)
)