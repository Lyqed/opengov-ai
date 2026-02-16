use serde::{Deserialize, Serialize};
use chrono::{DateTime, Utc};

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct Agent {
    pub id: String,
    pub agent_type: String,
    pub version: String,
    pub host: String,
    pub owner: String,
    pub department: String,
    pub mcp_servers: Vec<MCPConnection>,
    pub risk_score: f64,
    pub first_seen: DateTime<Utc>,
    pub last_seen: DateTime<Utc>,
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct MCPConnection {
    pub server_id: String,
    pub server_name: String,
    pub transport: String,
    pub tools: Vec<String>,
    pub permissions: Vec<String>,
}
