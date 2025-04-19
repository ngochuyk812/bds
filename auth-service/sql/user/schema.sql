CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,  
    guid CHAR(36) NOT NULL,             
    siteId VARCHAR(255) NOT NULL,      
    email VARCHAR(100) NOT NULL,    
    hash_password VARCHAR(255) NOT NULL,         
    salt VARCHAR(255) NOT NULL,         
    active BOOLEAN DEFAULT false,
    createdAt BIGINT NOT NULL,          
    updatedAt BIGINT DEFAULT NULL,      
    deletedAt BIGINT DEFAULT NULL,      
    UNIQUE KEY unique_guid (guid),      
    UNIQUE KEY unique_siteId_email (email, siteId)  
);
