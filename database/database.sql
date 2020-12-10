CREATE TABLE workbench.app_third_party (
	trp_id VARCHAR (60) PRIMARY KEY,
	trp_code VARCHAR (60) UNIQUE NOT NULL,
	trp_name VARCHAR (255) NOT NULL,
	trp_description TEXT DEFAULT NULL,
	trp_url VARCHAR ( 255 ) NOT NULL,
    trp_command VARCHAR (63) NOT NULL,
    trp_status INTEGER DEFAULT 1,
    trp_created_at DATETIME NOT NULL,
    trp_updated_at DATETIME NOT NULL
);

