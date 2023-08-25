CREATE TYPE wallettype AS ENUM ('enabled', 'disabled');

CREATE TABLE wallets
(
    id          uuid NOT NULL,
    owned_by    uuid NOT NULL,
    status      wallettype NOT NULL,
    enabled_at  timestamp,
    disabled_at timestamp,
    balance     real NOT NULL,
    CONSTRAINT wallets_pk PRIMARY KEY (id),
    CONSTRAINT wallets_accounts_fk FOREIGN KEY (owned_by) REFERENCES accounts (id) ON DELETE CASCADE ON UPDATE CASCADE
);
