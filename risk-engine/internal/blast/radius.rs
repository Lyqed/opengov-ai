//! Blast radius calculation for AI agent security incidents.
//!
//! Maps the potential impact of a compromised agent:
//! - What data stores can it reach?
//! - What systems can it modify?
//! - What's the estimated breach cost?

use serde::{Deserialize, Serialize};

/// BlastRadius describes the potential impact of a compromised asset.
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct BlastRadius {
    pub asset_id: String,
    pub data_exposure: DataExposure,
    pub system_impact: Vec<SystemImpact>,
    pub estimated_cost_usd: f64,
    pub affected_users: u64,
    pub affected_records: u64,
    pub compliance_impact: Vec<ComplianceImpact>,
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct DataExposure {
    pub databases: Vec<DataStore>,
    pub file_systems: Vec<String>,
    pub api_endpoints: Vec<String>,
    pub total_records_accessible: u64,
    pub contains_pii: bool,
    pub contains_phi: bool,
    pub contains_financial: bool,
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct DataStore {
    pub name: String,
    pub store_type: String, // "postgres", "s3", "salesforce", etc.
    pub record_count: u64,
    pub classification: String, // "public", "internal", "confidential", "restricted"
    pub access_level: String,   // "read", "write", "admin"
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct SystemImpact {
    pub system: String,
    pub impact_type: String, // "data_exfiltration", "data_modification", "service_disruption"
    pub severity: String,
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct ComplianceImpact {
    pub framework: String, // "SOC2", "GDPR", "HIPAA", "PCI-DSS"
    pub control_id: String,
    pub description: String,
}

/// Calculator maps blast radius for assets.
pub struct Calculator {
    // TODO: Data store inventory
    // TODO: Permission mapping
    // TODO: Cost model (IBM Cost of a Breach methodology)
}

impl Calculator {
    pub fn new() -> Self {
        Calculator {}
    }

    /// Calculate the blast radius for a given asset.
    pub fn calculate(&self, _asset_id: &str) -> BlastRadius {
        // TODO: Map agent permissions to data stores
        // TODO: Calculate transitive access (agent -> MCP -> DB)
        // TODO: Estimate breach cost based on record types and count
        // TODO: Map to compliance framework impacts
        todo!("Blast radius calculation not yet implemented")
    }
}
