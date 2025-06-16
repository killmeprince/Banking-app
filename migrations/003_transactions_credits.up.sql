CREATE TABLE IF NOT EXISTS transactions (
  id            SERIAL PRIMARY KEY,
  account_id    INT NOT NULL REFERENCES accounts(id),
  amount        NUMERIC(18,2) NOT NULL,
  type          TEXT NOT NULL,
  created_at    TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE IF NOT EXISTS credits (
  id            SERIAL PRIMARY KEY,
  account_id    INT NOT NULL REFERENCES accounts(id),
  principal     NUMERIC(18,2) NOT NULL,
  rate          NUMERIC(5,4)  NOT NULL,
  term_months   INT NOT NULL,
  margin        NUMERIC(5,4) NOT NULL,
  created_at    TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE IF NOT EXISTS payment_schedules (
  id            SERIAL PRIMARY KEY,
  credit_id     INT NOT NULL REFERENCES credits(id),
  due_date      DATE NOT NULL,
  amount        NUMERIC(18,2) NOT NULL,
  paid          BOOLEAN DEFAULT FALSE
);
