CREATE TABLE user_details (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_guid CHAR(36) NOT NULL,
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    phone VARCHAR(20),
    avatar VARCHAR(255),
    address TEXT,
    createdAt BIGINT NOT NULL,
    updatedAt BIGINT DEFAULT NULL,
    deletedAt BIGINT DEFAULT NULL,
    FOREIGN KEY (user_guid) REFERENCES users(guid),
    UNIQUE KEY unique_user_guid (user_guid)
);