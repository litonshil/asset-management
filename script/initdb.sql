-- Create Database
CREATE DATABASE IF NOT EXISTS asset_management;
USE asset_management;

-- Create Table: Assets
CREATE TABLE IF NOT EXISTS assets (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    type ENUM('audio', 'video', 'image', 'interactive') NOT NULL,
    user_id INT NOT NULL,
    description TEXT,
    file_url TEXT NOT NULL,
    metadata JSON,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL
);
CREATE TABLE IF NOT EXISTS asset_dependencies (
    id INT AUTO_INCREMENT PRIMARY KEY,
    asset_id INT,
    dependency_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (asset_id) REFERENCES assets(id) ON DELETE CASCADE,
    FOREIGN KEY (dependency_id) REFERENCES assets(id) ON DELETE CASCADE
);

-- Indexes for Optimization
CREATE INDEX idx_assets_type ON assets (type);
CREATE INDEX idx_assets_user_id ON assets (user_id);
