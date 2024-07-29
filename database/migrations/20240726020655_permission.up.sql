CREATE TABLE permissions (
  id varchar(255) PRIMARY KEY NOT NULL,
  name varchar(255) DEFAULT '' NOT NULL,
  created_at datetime NOT NULL,
  updated_at datetime NOT NULL,
  deleted_at datetime DEFAULT NULL
);