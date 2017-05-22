CREATE TABLE users (
  id serial,
  provider varchar(100) NOT NULL,
  name varchar(100) NOT NULL,
  email varchar(150) NOT NULL,
  first_name varchar(100),
  last_name varchar(100),
  nickname varchar(100),
  description varchar(255),
  user_id varchar(150) NOT NULL,
  avatar_url varchar(255),
  location varchar(100),
  access_token text NOT NULL,
  access_token_secret text,
  refresh_token text
);

CREATE TABLE income_category (
  id serial,
  title varchar(100) UNIQUE,
  description varchar(255)
);

CREATE TABLE expense_category (
  id serial,
  title varchar(100) UNIQUE,
  description varchar(255)
);

CREATE TABLE currency (
  id serial,
  title varchar(100) UNIQUE,
  description varchar(255)
);

CREATE TABLE income (
  id serial,
  user_id integer NOT NULL,
  category_id integer NOT NULL,
  date_time timestamp NOT NULL,
  title varchar(100),
  reference varchar(150),
  value numeric NOT NULL CONSTRAINT positive_value CHECK (value > 0.0),
  currency_id integer NOT NULL,
  rate numeric DEFAULT 1.0 CONSTRAINT positive_rate CHECK (rate > 0.0)
);

CREATE TABLE expense (
  id serial,
  user_id integer NOT NULL,
  category_id integer NOT NULL,
  date_time timestamp NOT NULL,
  title varchar(100),
  reference varchar(150),
  value numeric NOT NULL CONSTRAINT positive_value CHECK (value > 0.0),
  currency_id integer NOT NULL,
  rate numeric DEFAULT 1.0 CONSTRAINT positive_rate CHECK (rate > 0.0)
);

ALTER TABLE users ADD CONSTRAINT user_pkey PRIMARY KEY (id);
ALTER TABLE users ADD CONSTRAINT email_unique UNIQUE (email);
ALTER TABLE income_category ADD CONSTRAINT income_category_pkey PRIMARY KEY (id);
ALTER TABLE expense_category ADD CONSTRAINT expense_category_pkey PRIMARY KEY (id);
ALTER TABLE currency ADD CONSTRAINT currency_pkey PRIMARY KEY (id);
ALTER TABLE income ADD CONSTRAINT income_pkey PRIMARY KEY (user_id, category_id, date_time, title);
ALTER TABLE expense ADD CONSTRAINT expense_pkey PRIMARY KEY (user_id, category_id, date_time, title);
ALTER TABLE income ADD CONSTRAINT income_users_fkey FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;
ALTER TABLE expense ADD CONSTRAINT expense_users_fkey FOREIGN KEY (user_id) REFERENCES users(id)  ON DELETE CASCADE;
ALTER TABLE income ADD CONSTRAINT income_category_fkey FOREIGN KEY (category_id) REFERENCES income_category(id);
ALTER TABLE expense ADD CONSTRAINT expense_category_fkey FOREIGN KEY (category_id) REFERENCES expense_category(id);
ALTER TABLE income ADD CONSTRAINT income_currency_fkey FOREIGN KEY (currency_id) REFERENCES currency(id);
ALTER TABLE expense ADD CONSTRAINT expense_currency_fkey FOREIGN KEY (currency_id) REFERENCES currency(id);
