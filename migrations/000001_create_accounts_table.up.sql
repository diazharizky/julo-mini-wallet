CREATE TABLE accounts
(
    id         uuid NOT NULL,
    created_at timestamp,
    CONSTRAINT accounts_pk PRIMARY KEY (id)
);
