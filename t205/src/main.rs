//use std::process;

use anyhow::Result;
//use daemonize::Daemonize;
use log::LevelFilter;
use log4rs::append::console::ConsoleAppender;
use log4rs::config::{Appender, Config, Root};
use log4rs::encode::pattern::PatternEncoder;

fn main() -> Result<()> {
    //let stdout = std::fs::File::create("/tmp/t205b.out")?;
    //let stderr = std::fs::File::create("/tmp/t205b.err")?;

    //let d = Daemonize::new()
    //    .pid_file("/tmp/t205b.pid")
    //    .chown_pid_file(true)
    //    .working_directory("/tmp")
    //    .user("christos")
    //    .group("daemon")
    //    //.stdout(stdout)
    //    //.stderr(stderr)
    //    .exit_action(|| eprintln!("[{}] before parent exits", process::id()))
    //    .privileged_action(|| eprintln!("[{}] before dropping privileges", process::id()));
    //d.start()?;

    let config = if atty::is(atty::Stream::Stdout) {
        let stdout = ConsoleAppender::builder().build();
        Config::builder()
            .appender(Appender::builder().build("is_atty", Box::new(stdout)))
            .build(Root::builder().appender("is_atty").build(LevelFilter::Info))
            .unwrap()
    } else {
        let stdout = ConsoleAppender::builder()
            .encoder(Box::new(PatternEncoder::new("[{l}] {t} - {m}{n}")))
            .build();
        Config::builder()
            .appender(Appender::builder().build("no_atty", Box::new(stdout)))
            .build(Root::builder().appender("no_atty").build(LevelFilter::Info))
            .unwrap()
    };
    log4rs::init_config(config)?;

    let mut ticks = 0u64;
    loop {
        ticks += 1;
        //println!("tick {}", ticks);
        log::info!("tick {}", ticks);
        std::thread::sleep(std::time::Duration::from_secs(10));
    }
}
