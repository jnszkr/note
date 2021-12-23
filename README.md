# note

Simple cli that makes your thoughts captured in the moment.

## Install

```
go install github.com/jnszkr/note
```

## Usage

```bash
> note This is my first note that I am creating.
> note What a rainy day.
```

The notes will be stored in your currend directory to `.notes` file.

```bash
> ls -al
total 48
drwxr-xr-x   9 jnszkr  staff   288 Dec 23 14:05 .
drwxr-xr-x  27 jnszkr  staff   864 Dec 23 14:01 ..
-rw-r--r--   1 jnszkr  staff   114 Dec 23 14:11 .notes
```

To look at your notes you can do:

```bash
> cat .notes
2021-12-23T14:11:35+01:00 This is my first note that I am creating.
2021-12-23T14:11:42+01:00 What a rainy day.
```

To search in notes:

```bash
> note -s rainy
/Users/jnszkr/go/src/note/.notes
        2021-12-23T14:11:42+01:00 What a rainy day.
```
