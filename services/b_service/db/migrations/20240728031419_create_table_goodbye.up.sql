CREATE TABLE goodbyes (
                          id CHAR(36) PRIMARY KEY DEFAULT (UUID()),
                          message TEXT NOT NULL,
                          user_uuid CHAR(36) NOT NULL,
                          timestamp TIMESTAMP NOT NULL,
                          created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                          updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                          deleted_at TIMESTAMP NULL
);
