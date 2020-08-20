//! Change an environment variable of a running process and check whether that process, as well as
//! others, are able to see the change through `proc(5)`. (Hint: They shouldn't!)
//!
//! $ rustc t204.rs
//! $ ENV_VAR_T204=something ./t204

use std::collections::HashMap;
use std::env;
use std::error::Error;
use std::process;

const ENV_VAR: &str = "ENV_VAR_T204";

/// Read `/proc/self/eviron` and parse all environment variables into a
/// `std::collections::HashMap`.
fn parse_environ() -> Result<HashMap<String, String>, Box<dyn Error + 'static>> {
    let pid = process::id();
    let environ = std::fs::read_to_string(format!("/proc/{}/environ", pid))?;
    eprintln!("Content of '/proc/{}/environ':\n{:#?}\n\n", pid, environ);

    Ok(environ
        .split('\0')
        .map(|kv| kv.split('=').map(|s| s.to_owned()).collect::<Vec<_>>())
        .map(|kv| {
            (
                kv.get(0).unwrap().to_owned(),
                kv.get(1).unwrap_or(&"".to_owned()).to_owned(),
            )
        })
        .collect::<HashMap<_, _>>())
}

/// Report whether the given environment variable `var` exists as a key in the given `vars`
/// collection.
fn check_one(var: &str, vars: &HashMap<String, String>) -> Result<(), Box<dyn Error + 'static>> {
    if vars.contains_key(&ENV_VAR.to_owned()) {
        eprintln!("Found: {} = {}\n", var, vars.get(ENV_VAR).unwrap());
        Ok(())
    } else {
        Err(format!(
            "{} was not found! Are you sure that you ran it correctly?",
            var
        )
        .into())
    }
}

/// Pause until user input from stdin.
fn pause() {
    eprintln!("Press <ENTER> to continue...");
    let stdin = std::io::stdin();
    stdin.read_line(&mut String::with_capacity(0)).unwrap();
}

fn main() -> Result<(), Box<dyn Error + 'static>> {
    // Parse all environment variables
    let env_vars = parse_environ()?;
    eprintln!("All environment variables:\n{:#?}\n\n", env_vars);

    // Check the one of interest
    check_one(ENV_VAR, &env_vars)?;

    // Change the value of ENV_VAR
    env::set_var(ENV_VAR, "YOLOYOLOYOLO");
    // Now is the time to check it from another process too; therefore pause until user input from
    // stdin
    pause();

    // Parse all environment variables again
    let env_vars = parse_environ()?;
    eprintln!("All environment variables:\n{:#?}\n\n", env_vars);

    // Check the one of interest, again
    check_one(ENV_VAR, &env_vars)?;
    eprintln!("env::var({}) returns: {}\n\n", ENV_VAR, env::var(ENV_VAR)?);

    Ok(())
}
