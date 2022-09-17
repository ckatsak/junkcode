use std::process;

use anyhow::{Context, Result};
use nix::sched::{unshare, CloneFlags};
use tokio::{fs, process::Command};

#[tokio::main]
async fn main() -> Result<()> {
    let pid = process::id();

    let my_netns = fs::read_link("/proc/self/ns/net").await?;
    println!("[{}]: {:?}", pid, my_netns);

    let mut child = unsafe {
        Command::new("readlink")
            .arg("/proc/self/ns/net")
            .pre_exec(|| {
                let old_netns = std::fs::read_link("/proc/self/ns/net")?;
                println!("[{}]: pre_exec BEFORE: {:?}", process::id(), old_netns);

                unshare(CloneFlags::CLONE_NEWNET)?;

                let new_netns = std::fs::read_link("/proc/self/ns/net")?;
                println!("[{}]: pre_exec AFTER: {:?}", process::id(), new_netns);
                Ok(())
            })
            .spawn()
            //.output()
            //.await
            .with_context(|| "failed to fork!")?
    };
    println!("[{}]: spawned {:?}", pid, child);
    let _status = child.wait().await?;
    println!("[{}]: reaped {:?}", pid, child);

    //let stdout = String::from_utf8_lossy(&output.stdout);
    //let stderr = String::from_utf8_lossy(&output.stderr);
    //println!("[{}]: child's stdout: '{}'", pid, stdout);
    //println!("[{}]: child's stderr: '{}'", pid, stderr);

    Ok(())
}
