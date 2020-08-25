#[macro_use]
extern crate anyhow;

//mod another;
use t206::another;

use anyhow::Result;
use log::LevelFilter;
use log4rs::{
    append::rolling_file::{
        policy::compound::{
            roll::fixed_window::FixedWindowRoller, trigger::size::SizeTrigger, CompoundPolicy,
        },
        RollingFileAppender,
    },
    config::{Appender, Config, Root},
    encode::json::JsonEncoder,
};

const LOG_FILE_NAME: &str = "./t206.log";
const LOG_FILE_SIZE: u64 = 1024;

fn main() -> Result<()> {
    let roller = FixedWindowRoller::builder()
        .base(0)
        .build(&format!("{}.{{}}.gz", LOG_FILE_NAME), 4)
        .map_err(|e| anyhow!("could not create the `FixedWindowRoller`: {:#?}", e))?;

    let stdout = RollingFileAppender::builder()
        .encoder(Box::new(JsonEncoder::new()))
        .build(
            LOG_FILE_NAME,
            Box::new(CompoundPolicy::new(
                Box::new(SizeTrigger::new(LOG_FILE_SIZE)),
                Box::new(roller),
            )),
        )?;

    let config = Config::builder()
        .appender(Appender::builder().build("fixed_roller", Box::new(stdout)))
        .build(
            Root::builder()
                .appender("fixed_roller")
                .build(LevelFilter::Debug),
        )?;

    let _handle = log4rs::init_config(config)?;

    log::info!("YOLO t206");

    another::module::log("I'm here too!");
    another::hidden::module::log("Hey, me too!");

    Ok(())
}
