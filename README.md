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

<h3>Export</h3>
Using ``export`` command you'll get report in .csv file in the same directory with program<br>
It requires at least one argument - path to file or directory which will be scanned<br>
Second parameter - name of report (not required)

<h3>List</h3>
Using ``list`` command you'll see report on you screen without saving to file 
<br>It requires one argument - path to file or directory which will be scanned

<h3>FsCheck</h3>
``fscheck`` command helps you to check your file system for supporting last access time. 
It will check is it enabled or not<br>
This command has no arguments