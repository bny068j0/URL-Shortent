-- Create clicks table
CREATE TABLE IF NOT EXISTS clicks (
    id          BIGSERIAL PRIMARY KEY,
    short_code  TEXT NOT NULL,
    clicked_at  TIMESTAMPTZ DEFAULT NOW(),
    user_agent  TEXT,
    ip_address  INET,
    referer     TEXT,
    country     TEXT,
    device_type TEXT
);

CREATE INDEX IF NOT EXISTS idx_clicks_short_code ON clicks(short_code);
CREATE INDEX IF NOT EXISTS idx_clicks_clicked_at ON clicks(clicked_at);
