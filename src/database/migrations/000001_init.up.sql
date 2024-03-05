CREATE TABLE IF NOT EXISTS "user" (
    id varchar(255) PRIMARY KEY,
    name varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    password varchar(255) NOT NULL,
    username varchar(255) NOT NULL,
    phone varchar(255),
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS account_type (
    id varchar(255) PRIMARY KEY,
    name varchar(255) NOT NULL,
    logo varchar(255) NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS account (
    id varchar(255) PRIMARY KEY,
    user_id varchar(255) NOT NULL,
    account_type_id varchar(255) NOT NULL,
    balance decimal(10, 2) NOT NULL,
    currency varchar(255) NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES "user"(id),
    FOREIGN KEY (account_type_id) REFERENCES account_type(id)
);

CREATE TABLE IF NOT EXISTS tax_type (
    id varchar(255) PRIMARY KEY,
    name varchar(255) NOT NULL,
    rate decimal(10, 2) NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS merchant (
    id varchar(255) PRIMARY KEY,
    name varchar(255) NOT NULL,
    description TEXT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS product (
    id varchar(255) PRIMARY KEY,
    name varchar(255) NOT NULL,
    description TEXT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS product_merchant (
    id varchar(255) PRIMARY KEY,
    product_id varchar(255) NOT NULL,
    merchant_id varchar(255) NOT NULL,
    price decimal(10, 2) NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (product_id) REFERENCES product(id),
    FOREIGN KEY (merchant_id) REFERENCES merchant(id)
);

CREATE TABLE IF NOT EXISTS product_merchant_price_log (
    id varchar(255) PRIMARY KEY,
    product_merchant_id varchar(255) NOT NULL,
    price decimal(10, 2) NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (product_merchant_id) REFERENCES product_merchant(id)
);

CREATE TABLE IF NOT EXISTS expense_main_category (
    id varchar(255) PRIMARY KEY,
    name varchar(255) NOT NULL,
    description TEXT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS expense_sub_category (
    id varchar(255) PRIMARY KEY,
    expense_main_category_id varchar(255) NOT NULL,
    name varchar(255) NOT NULL,
    description TEXT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (expense_main_category_id) REFERENCES expense_main_category(id)
);

CREATE TABLE IF NOT EXISTS expense (
    id varchar(255) PRIMARY KEY,
    account_id varchar(255) NOT NULL,
    expense_main_type_id varchar(255) NOT NULL,
    expense_sub_type_id varchar(255) NULL,
    merchant_id varchar(255) NULL,
    tax_type_id varchar(255) NULL,
    transaction_name varchar(255) NOT NULL,
    transaction_date timestamp NOT NULL,
    amount decimal(10, 2) NOT NULL,
    total_tax decimal(10, 2) NOT NULL,
    total_amount decimal(10, 2) NOT NULL,
    is_summarized boolean NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (account_id) REFERENCES account(id),
    FOREIGN KEY (expense_main_type_id) REFERENCES expense_main_category(id),
    FOREIGN KEY (expense_sub_type_id) REFERENCES expense_sub_category(id),
    FOREIGN KEY (merchant_id) REFERENCES merchant(id),
    FOREIGN KEY (tax_type_id) REFERENCES tax_type(id)
);

CREATE TABLE IF NOT EXISTS expense_item (
    id varchar(255) PRIMARY KEY,
    expense_id varchar(255) NOT NULL,
    product_merchant_id varchar(255) NOT NULL,
    quantity decimal(10, 2) NOT NULL,
    price decimal(10, 2) NOT NULL,
    discount decimal(10, 2) NULL,
    total_amount decimal(10, 2) NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (expense_id) REFERENCES expense(id),
    FOREIGN KEY (product_merchant_id) REFERENCES product_merchant(id)
);

CREATE TABLE IF NOT EXISTS income_source (
    id varchar(255) PRIMARY KEY,
    name varchar(255) NOT NULL,
    email varchar(255) NULL,
    phone varchar(255) NULL,
    contact_person_name varchar(255) NOT NULL,
    description TEXT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS income_type (
    id varchar(255) PRIMARY KEY,
    name varchar(255) NOT NULL,
    description TEXT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS income (
    id varchar(255) PRIMARY KEY,
    income_type_id varchar(255) NOT NULL,
    account_id varchar(255) NOT NULL,
    income_source_id varchar(255) NOT NULL,
    transaction_name varchar(255) NOT NULL,
    transaction_date timestamp NOT NULL,
    amount decimal(10, 2) NOT NULL,
    is_summarized boolean NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (income_type_id) REFERENCES income_type(id),
    FOREIGN KEY (account_id) REFERENCES account(id),
    FOREIGN KEY (income_source_id) REFERENCES income_source(id)
);

CREATE TABLE IF NOT EXISTS person (
    id varchar(255) PRIMARY KEY,
    name varchar(255) NOT NULL,
    address varchar(255) NULL,
    email varchar(255) NULL,
    phone varchar(255) NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TYPE payment_status AS ENUM('PAID', 'PARTIALLY_PAID', 'UNPAID');

CREATE TYPE transaction_type AS ENUM('DEBT', 'LOAN');

CREATE TABLE IF NOT EXISTS debt_loan (
    id varchar(255) PRIMARY KEY,
    account_id varchar(255) NOT NULL,
    person_id varchar(255) NOT NULL,
    type transaction_type NOT NULL,
    status payment_status NOT NULL,
    transaction_name varchar(255) NOT NULL,
    transaction_date timestamp NOT NULL,
    amount decimal(10, 2) NOT NULL,
    description TEXT NULL,
    is_summarized boolean NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (account_id) REFERENCES account(id),
    FOREIGN KEY (person_id) REFERENCES person(id)
);

CREATE TABLE IF NOT EXISTS debt_loan_payment (
    id varchar(255) PRIMARY KEY,
    debt_loan_id varchar(255) NOT NULL,
    amount decimal(10, 2) NOT NULL,
    transaction_date timestamp NOT NULL,
    description TEXT NULL,
    is_calculated boolean NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (debt_loan_id) REFERENCES debt_loan(id)
);

CREATE TABLE IF NOT EXISTS income_expense_summary (
    id varchar(255) PRIMARY KEY,
    account_id varchar(255) NOT NULL,
    user_id varchar(255) NOT NULL,
    transaction_date date NOT NULL,
    income decimal(10, 2) NOT NULL,
    expense decimal(10, 2) NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (account_id) REFERENCES account(id),
    FOREIGN KEY (user_id) REFERENCES "user"(id)
);

CREATE TABLE IF NOT EXISTS debt_loan_summary (
    id varchar(255) PRIMARY KEY,
    account_id varchar(255) NOT NULL,
    user_id varchar(255) NOT NULL,
    transaction_date date NOT NULL,
    debt decimal(10, 2) NOT NULL,
    loan decimal(10, 2) NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (account_id) REFERENCES account(id),
    FOREIGN KEY (user_id) REFERENCES "user"(id)
);