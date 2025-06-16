CREATE TABLE IF NOT EXISTS accounts (
  id         SERIAL PRIMARY KEY,
  user_id    INT NOT NULL REFERENCES users(id),
  balance    NUMERIC(18,2) NOT NULL DEFAULT 0,
  created_at TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE IF NOT EXISTS cards (
  id           SERIAL PRIMARY KEY,
  account_id   INT NOT NULL REFERENCES accounts(id),
  number_enc   BYTEA  NOT NULL,
  exp_enc      BYTEA  NOT NULL,
  cvv_hash     TEXT   NOT NULL,
  hmac         TEXT   NOT NULL,
  created_at   TIMESTAMPTZ DEFAULT now()
);
