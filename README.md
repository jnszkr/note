# note ![Testing](https://github.com/jnszkr/note/actions/workflows/test.yml/badge.svg)

Simple cli that makes your thoughts captured in the moment.

## Install

```
go install github.com/jnszkr/note@latest
```

## Usage

```bash
> note This is my first note that I am creating.
> note What a rainy day.
```

The notes will be stored in your current directory to `.notes` file.

```bash
> ls -al
total 48
drwxr-xr-x   9 -  staff   288 Dec 23 14:05 .
drwxr-xr-x  27 -  staff   864 Dec 23 14:01 ..
-rw-r--r--   1 -  staff   114 Dec 23 14:11 .notes
```

To look at your notes you can do:

```bash
> note
2021-12-23T14:11:35+01:00 This is my first note that I am creating.
2021-12-23T14:11:42+01:00 What a rainy day.
```

To search in notes:

```bash
> note -s "rainy day"
 • 
	2021-12-23T14:11:42+01:00 What a rainy day.
```

Search would do a recursive search on your current dir, so if there are more
notes in any of the subdirectories, search would be performed on it.

```bash
> mkdir subfolder && cd subfolder
> note This is my first note that I am creating in the subfolder.
> cd ..
> note -s "this"
 • 
    2021-12-23T14:11:35+01:00 This is my first note that I am creating.
 • subfolder
    2021-12-23T14:15:41+01:00 This is my first note that I am creating in the subfolder.
```
