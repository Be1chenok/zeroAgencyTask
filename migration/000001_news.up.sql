CREATE TABLE IF NOT EXISTS news (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS news_categories (
    id BIGSERIAL PRIMARY KEY,
    news_id bigint NOT NULL,
    category_id bigint NOT NULL,
    FOREIGN KEY (news_id) REFERENCES news (id) ON DELETE CASCADE
);