CREATE TABLE sites (
    id INT AUTO_INCREMENT PRIMARY KEY,  
    guid CHAR(36) NOT NULL,             
    siteId VARCHAR(255) NOT NULL,      
    name VARCHAR(255) NOT NULL,         
    createdAt BIGINT NOT NULL,          
    updatedAt BIGINT DEFAULT NULL,      
    deletedAt BIGINT DEFAULT NULL,      
    UNIQUE KEY unique_guid (guid),      
    UNIQUE KEY unique_siteId (siteId)  
);
