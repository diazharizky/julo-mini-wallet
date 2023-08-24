CREATE TYPE wallettype AS ENUM ('enabled', 'disabled');

CREATE TABLE wallets
(
    id          uuid NOT NULL,
    owned_by    uuid NOT NULL,
    status      wallettype NOT NULL,
    enabled_at  timestamp NOT NULL,
    disabled_at timestamp,
    balance     real NOT NULL,
    CONSTRAINT wallets_pk PRIMARY KEY (id)
);
