# t192

Using Apache HTTP client to upload a file (`multipart/form-data`).


## Build

```sh
$ mvn clean install -Dmaven.test.skip=true
```

## Example run

```sh
$ java -jar target/t192-0.1.0-SNAPSHOT-jar-with-dependencies.jar target/maven-status/maven-compiler-plugin/compile/default-compile/inputFiles.lst http://gold1.ckatsak:8080/yo2data/flink-schedule
```
