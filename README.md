# files_meta_times
get files meta times 

<h2>run in docker</h2>

```
docker run --rm -it -v "$PWD":/go/src/files_meta_times -w /go/src/files_meta_times golang:latest
```

<h2>How to use</h2>

```
files_meta_times export {path_to_folder} {file_name_to_export (optional)}
```

For now, supports only one commands ``export`` and ``list``

Using ``export`` command you'll get report in .csv file in the same directory with program
Using ``list`` command you'll see report on you screen without saving to file 