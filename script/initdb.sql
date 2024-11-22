-- Create Database
CREATE DATABASE IF NOT EXISTS asset_management;
USE asset_management;

-- Create Table: Assets
CREATE TABLE IF NOT EXISTS assets (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    type_id INT NOT NULL,
    user_id INT NOT NULL,
    description TEXT,
    file_url TEXT NOT NULL,
    metadata JSON,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (type_id) REFERENCES types(id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS asset_dependencies (
    id INT AUTO_INCREMENT PRIMARY KEY,
    asset_id INT,
    dependency_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (asset_id) REFERENCES assets(id) ON DELETE CASCADE,
    FOREIGN KEY (dependency_id) REFERENCES assets(id) ON DELETE CASCADE
);

-- Create Table: assigned_assets
CREATE TABLE IF NOT EXISTS assigned_assets (
    id INT AUTO_INCREMENT PRIMARY KEY,
    asset_id INT,
    user_id INT NOT NULL,
    assigned_by INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (asset_id) REFERENCES assets(id) ON DELETE CASCADE
);

Create TABLE IF NOT EXISTS types (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE, -- ('audio', 'video', 'image', 'interactive')
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Indexes for Optimization
CREATE INDEX idx_assets_type ON assets (type);
CREATE INDEX idx_assets_user_id ON assets (user_id);
