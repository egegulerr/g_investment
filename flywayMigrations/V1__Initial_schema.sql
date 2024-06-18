CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255),
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITHOUT TIME ZONE,
    deleted_at TIMESTAMP WITHOUT TIME ZONE

    );

CREATE TABLE news (
    id SERIAL PRIMARY KEY,
    url VARCHAR(255) UNIQUE NOT NULL,
    title VARCHAR(255),
    author VARCHAR(255),
    sentimental_analysis_score FLOAT,
    date TIMESTAMP WITHOUT TIME ZONE,
    summary TEXT,
    image VARCHAR(255),
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITHOUT TIME ZONE,
    deleted_at TIMESTAMP WITHOUT TIME ZONE
);


CREATE TABLE stocks (
    id SERIAL PRIMARY KEY,
    symbol VARCHAR(50) UNIQUE NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITHOUT TIME ZONE,
    deleted_at TIMESTAMP WITHOUT TIME ZONE
);

CREATE TABLE news_stocks (
    id SERIAL PRIMARY KEY,
    news_id INT REFERENCES news(id),
    stock_id INT REFERENCES stocks(id),
    relevance_score VARCHAR(255),
    stock_sentiment_score VARCHAR(255),
    stock_sentiment_label VARCHAR(255),
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITHOUT TIME ZONE,
    deleted_at TIMESTAMP WITHOUT TIME ZONE


);


CREATE TABLE user_favorite_news (
    user_id INT REFERENCES users(id),
    news_id INT REFERENCES news(id),
    PRIMARY KEY (user_id, news_id)
);

INSERT INTO users (first_name, last_name, email, password)
VALUES ('Ege', 'Güler', 'mbcguler@gmail.com', 'fbguler45');