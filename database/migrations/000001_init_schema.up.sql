CREATE TABLE IF NOT EXISTS processed_signals (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    order_id TEXT NOT NULL,
    amount NUMERIC(10, 2),
    vendor_id TEXT,
    processed_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    is_anomaly BOOLEAN DEFAULT FALSE
);