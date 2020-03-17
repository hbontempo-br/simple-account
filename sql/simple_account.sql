CREATE DATABASE IF NOT EXISTS account DEFAULT CHARACTER SET utf8 DEFAULT COLLATE utf8_general_ci;

USE account;

CREATE TABLE Account
(
    id              INT                NOT NULL UNIQUE AUTO_INCREMENT,
    document_number VARCHAR(11) UNIQUE NOT NULL,
    created_at      TIMESTAMP          NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);


CREATE TABLE OperationType
(
    id          INT                 NOT NULL UNIQUE AUTO_INCREMENT,
    enumerator  VARCHAR(100) UNIQUE NOT NULL,
    description VARCHAR(100) UNIQUE NOT NULL,
    created_at  TIMESTAMP           NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

INSERT INTO OperationType (enumerator, description)
VALUES ('single_installment_purchase', 'COMPRA A VISTA'),
       ('multi_installment_purchase', 'COMPRA PARCELADA'),
       ('withdraw', 'SAQUE'),
       ('payment', 'PAGAMENTO');

CREATE TABLE Transaction
(
    id                INT            NOT NULL UNIQUE AUTO_INCREMENT,
    account_id        INT            NOT NULL,
    operation_type_id INT            NOT NULL,
    amount            NUMERIC(19, 2) NOT NULL,
    event_date        DATETIME       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at        TIMESTAMP      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY fk_account (account_id) REFERENCES Account (id),
    FOREIGN KEY fk_operation_type (operation_type_id) REFERENCES OperationType (id)
);