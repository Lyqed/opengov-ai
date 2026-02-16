use tracing::{info, Level};
use tracing_subscriber::FmtSubscriber;

const VERSION: &str = "0.1.0-dev";

#[tokio::main]
async fn main() -> anyhow::Result<()> {
    let subscriber = FmtSubscriber::builder()
        .with_max_level(Level::INFO)
        .json()
        .finish();
    tracing::subscriber::set_global_default(subscriber)?;

    info!(version = VERSION, "starting OpenGov AI risk engine");

    // TODO: Subscribe to NATS JetStream for incoming events
    // TODO: Initialize risk scoring models
    // TODO: Initialize blast radius calculator
    // TODO: Start HTTP API for risk queries
    // TODO: Run continuous risk assessment loop

    println!("risk-engine: scaffolding complete, implementation pending");

    tokio::signal::ctrl_c().await?;
    info!("shutting down risk engine");
    Ok(())
}
