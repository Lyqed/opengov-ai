use serde::{Deserialize, Serialize};
use chrono::{DateTime, Utc};

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct NormalizedEvent {
    pub id: String,
    pub timestamp: DateTime<Utc>,
    pub source: String,
    pub agent_id: String,
    pub user_id: String,
    pub action: String,
    pub mcp_server: String,
    pub tool: String,
    pub status: String,
    pub risk_signals: Vec<RiskSignal>,
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct RiskSignal {
    pub signal_type: String,
    pub severity: String,
    pub detail: String,
    pub score: f64,
}
