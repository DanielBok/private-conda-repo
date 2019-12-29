CREATE TABLE IF NOT EXISTS package_counts
(
    id           SERIAL PRIMARY KEY,
    channel      VARCHAR(50) NOT NULL,
    package      VARCHAR(64) NOT NULL,
    platform     VARCHAR(10) NOT NULL,
    build_string VARCHAR(64) NOT NULL,
    build_number INT         NOT NULL DEFAULT 0,
    version      VARCHAR(30) NOT NULL,
    count        INT         NOT NULL DEFAULT 0
);

CREATE INDEX idx_package_counts__channel_package on package_counts
    (channel, package);

CREATE INDEX idx_package_counts__channel_package_platform on package_counts
    (channel, package, platform);
