CREATE TABLE topics (
    topic_id uuid NOT NULL,
    title varchar(255) NOT NULL,
    body varchar(255) NULL,
    votes int NOT NULL,
    created_date TIMESTAMP NULL,
    updated_date TIMESTAMP NULL,
    PRIMARY KEY(topic_id)
);