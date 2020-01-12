# files_meta_times
get files meta times 

<h2>run in docker</h2>

```
docker run --rm -it -v "$PWD":/go/src/files_meta_times -w /go/src/files_meta_times golang:latest
```

<h2>How to use</h2>

```
files_meta_times times {path_to_folder} {file_name_to_export (optional)}
```

For now, supports only one command ```times```