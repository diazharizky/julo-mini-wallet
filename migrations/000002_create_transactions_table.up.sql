CREATE TYPE transactionstatus AS ENUM ('success', 'failed');

CREATE TYPE transactiontype AS ENUM ('deposit', 'withdrawal');

CREATE TABLE transactions
(
    id            uuid NOT NULL,
    wallet_id     uuid NOT NULL,
    status        transactionstatus NOT NULL,
    transacted_at timestamp NOT NULL,
    type          transactiontype NOT NULL,
    amount        real NOT NULL,
    reference_id  uuid NOT NULL,
    CONSTRAINT transactions_pk PRIMARY KEY (id),
    CONSTRAINT transactions_unq_ref_id UNIQUE (reference_id),
    CONSTRAINT transactions_wallets_fk FOREIGN KEY (wallet_id) REFERENCES wallets (id) ON DELETE CASCADE ON UPDATE CASCADE
);
