-- 事前作業：ユーザ／データベースの作成、データベース切替
-- スーパユーザでログイン (> psql -U postgres)
-- CREATE ROLE denki_user WITH LOGIN PASSWORD 'denki_user';
-- CREATE DATABASE denki;
-- GRANT ALL PRIVILEGES ON DATABASE denki TO denki_user;
-- \c denki;

-- テーブルの作成
-- \i [[filepath]]

CREATE TABLE CLASS (
	class_id BIGINT PRIMARY KEY,
	class_name VARCHAR(100) NOT NULL
);

CREATE TABLE MAKER (
	maker_id BIGINT PRIMARY KEY,
	maker_name VARCHAR(100) NOT NULL
);

CREATE TABLE GOODS (
	goods_id BIGINT PRIMARY KEY,
	goods_name VARCHAR(100) NOT NULL,
	class_id INTEGER NOT NULL REFERENCES CLASS(CLASS_ID),
	maker_id INTEGER NOT NULL REFERENCES MAKER(MAKER_ID),
	indicated_price INTEGER,
	purchase_price NUMERIC,
	stock INTEGER,
	update_super_visor_id VARCHAR(30) NOT NULL,
	update_date TIMESTAMP NOT NULL,
	update_version_id BIGINT NOT NULL
);