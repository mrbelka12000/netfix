CREATE TABLE IF NOT EXISTS Billing(
    ApplyId integer not null,
    Finished boolean,
    Amount float8,
    FOREIGN KEY (ApplyId) REFERENCES apply(id)
)