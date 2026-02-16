// API client for the OpenGov AI backend services.

const API_BASE = import.meta.env.VITE_API_URL || 'http://localhost:8080/api/v1';

export interface Agent {
  id: string;
  type: string;
  name: string;
  host: string;
  version: string;
  owner: string;
  department: string;
  risk_score: number;
  risk_level: 'low' | 'medium' | 'high' | 'critical';
  mcp_connections: number;
  first_seen: string;
  last_seen: string;
}

export interface RiskOverview {
  total_agents: number;
  total_mcp_servers: number;
  critical_risks: number;
  high_risks: number;
  medium_risks: number;
  low_risks: number;
  total_blast_radius_usd: number;
  unmonitored_agents: number;
}

export interface BlastRadius {
  asset_id: string;
  estimated_cost_usd: number;
  affected_records: number;
  affected_users: number;
  data_stores: DataStore[];
  compliance_impact: ComplianceImpact[];
}

export interface DataStore {
  name: string;
  type: string;
  record_count: number;
  classification: string;
  access_level: string;
}

export interface ComplianceImpact {
  framework: string;
  control_id: string;
  description: string;
}

export interface Incident {
  id: string;
  timestamp: string;
  severity: string;
  title: string;
  asset_id: string;
  status: 'open' | 'investigating' | 'resolved' | 'closed';
  playbook_id: string;
  actions_taken: string[];
}

async function fetchAPI<T>(path: string): Promise<T> {
  const res = await fetch(`${API_BASE}${path}`);
  if (!res.ok) throw new Error(`API error: ${res.status}`);
  return res.json();
}

export const api = {
  getRiskOverview: () => fetchAPI<RiskOverview>('/risk/overview'),
  getAgents: () => fetchAPI<Agent[]>('/agents'),
  getAgent: (id: string) => fetchAPI<Agent>(`/agents/${id}`),
  getBlastRadius: (id: string) => fetchAPI<BlastRadius>(`/risk/blast-radius/${id}`),
  getIncidents: () => fetchAPI<Incident[]>('/incidents'),
};
