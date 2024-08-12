CREATE TABLE consumer (
     id BIGINT NOT NULL AUTO_INCREMENT,
     nik VARCHAR(255) NOT NULL,
     nik_index VARCHAR(255) NOT NULL,
     password VARCHAR(1024) NOT NULL,
     pin VARCHAR(1024) NOT NULL,
     fullname VARCHAR(255) NOT NULL,
     legal_name VARCHAR(255) NOT NULL,
     date_of_birth VARCHAR(255) NOT NULL,
     place_of_birth VARCHAR(255) NOT NULL,
     salary BIGINT NOT NULL,
     identity_photo VARCHAR(255) NOT NULL,
     photo VARCHAR(255) NOT NULL,
     created_at TIMESTAMP NOT NULL default NOW(),
     created_by VARCHAR(255) NOT NULL,
     updated_at TIMESTAMP NOT NULL default NOW(),
     updated_by VARCHAR(255) NOT NULL,
     deleted_at TIMESTAMP,
     deleted_by VARCHAR(255),
     PRIMARY KEY (id)
);

CREATE TABLE credit_limit (
     id BIGINT NOT NULL AUTO_INCREMENT,
     consumer_id BIGINT NOT NULL,
     duration INT NOT NULL,
     amount BIGINT NOT NULL,
     created_at TIMESTAMP NOT NULL default NOW(),
     created_by VARCHAR(255) NOT NULL,
     updated_at TIMESTAMP NOT NULL default NOW(),
     updated_by VARCHAR(255) NOT NULL,
     deleted_at TIMESTAMP,
     deleted_by VARCHAR(255),
     PRIMARY KEY (id),
     INDEX limit_consumer_index (consumer_id),
     FOREIGN KEY (consumer_id) REFERENCES consumer(id) ON DELETE CASCADE
);

CREATE TABLE partner (
     id BIGINT NOT NULL AUTO_INCREMENT,
     name VARCHAR(255) NOT NULL,
     client_id VARCHAR(42) NOT NULL,
     api_key VARCHAR(255) NOT NULL,
     webhook VARCHAR(255),
     created_at TIMESTAMP NOT NULL default NOW(),
     created_by VARCHAR(255) NOT NULL,
     updated_at TIMESTAMP NOT NULL default NOW(),
     updated_by VARCHAR(255) NOT NULL,
     deleted_at TIMESTAMP,
     deleted_by VARCHAR(255),
     PRIMARY KEY (id)
);

CREATE TABLE credit_transaction (
     id BIGINT NOT NULL AUTO_INCREMENT,
     contract_number VARCHAR(255) NOT NULL,
     asset_name VARCHAR(255) NOT NULL,
     consumer_id BIGINT NOT NULL,
     partner_id BIGINT NOT NULL,
     otr BIGINT NOT NULL,
     admin_fee BIGINT NOT NULL,
     total_installment BIGINT NOT NULL,
     interest BIGINT NOT NULL,
     created_at TIMESTAMP NOT NULL default NOW(),
     created_by VARCHAR(255) NOT NULL,
     updated_at TIMESTAMP NOT NULL default NOW(),
     updated_by VARCHAR(255) NOT NULL,
     deleted_at TIMESTAMP,
     deleted_by VARCHAR(255),
     PRIMARY KEY (id),
     INDEX transaction_consumer_index (consumer_id),
     FOREIGN KEY (consumer_id) REFERENCES consumer(id) ON DELETE CASCADE,
     FOREIGN KEY (partner_id) REFERENCES partner(id) ON DELETE CASCADE
);

CREATE TABLE credit_payment (
     id BIGINT NOT NULL AUTO_INCREMENT,
     transaction_id BIGINT NOT NULL,
     consumer_id BIGINT NOT NULL,
     total_amount BIGINT NOT NULL,
     monthly_amount BIGINT NOT NULL,
     duration INT NOT NULL,
     interest BIGINT NOT NULL,
     status SMALLINT NOT NULL, 
     created_at TIMESTAMP NOT NULL default NOW(),
     created_by VARCHAR(255) NOT NULL,
     updated_at TIMESTAMP NOT NULL default NOW(),
     updated_by VARCHAR(255) NOT NULL,
     deleted_at TIMESTAMP,
     deleted_by VARCHAR(255),
     PRIMARY KEY (id),
     INDEX payment_consumer_index (consumer_id),
     FOREIGN KEY (consumer_id) REFERENCES consumer(id) ON DELETE CASCADE,
     FOREIGN KEY (transaction_id) REFERENCES credit_transaction(id) ON DELETE CASCADE
);

CREATE TABLE credit_payment_installment (
     id BIGINT NOT NULL AUTO_INCREMENT,
     payment_id BIGINT NOT NULL,
     amount BIGINT NOT NULL,
     status SMALLINT NOT NULL, 
     created_at TIMESTAMP NOT NULL default NOW(),
     created_by VARCHAR(255) NOT NULL,
     updated_at TIMESTAMP NOT NULL default NOW(),
     updated_by VARCHAR(255) NOT NULL,
     deleted_at TIMESTAMP,
     deleted_by VARCHAR(255),
     PRIMARY KEY (id),
     FOREIGN KEY (payment_id) REFERENCES credit_payment(id) ON DELETE CASCADE
);
