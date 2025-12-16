CREATE TABLE subscriptions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    
    symbol TEXT NOT NULL,
    category TEXT NOT NULL,   
    
    look_back_days INTEGER NOT NULL, 
    trigger_limit REAL NOT NULL,   
    
    last_alerted_at DATETIME,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    
    UNIQUE(symbol, category),       
    CHECK(category IN ('stock', 'crypto'))
);

CREATE TABLE alert_history (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    subscription_id INTEGER,
    
    symbol TEXT NOT NULL,
    alert_price REAL NOT NULL,
    peak_price REAL NOT NULL,   
    drop_percentage REAL NOT NULL,
    triggered_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    
    FOREIGN KEY(subscription_id) REFERENCES subscriptions(id) ON DELETE SET NULL
);

