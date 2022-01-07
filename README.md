# note ![Testing](https://github.com/jnszkr/note/actions/workflows/test.yml/badge.svg)

Simple cli that makes your thoughts captured in the moment.

## Install

```
go install github.com/jnszkr/note/cmd/note@latest
```

## Usage

```bash
$ cd ./docs/example-notes/running
$ note "I have decided that I am going to do running."
$ note "My goal is to finish a 5K"
$ note "But first I need a plan."
$ note "Let's start with 1K first time and see how that goes."
...
```

The notes will be stored in your current directory to `.notes` file.

```bash
$ ls -al
total 8
drwxr-xr-x  3 -  staff   96 Jan  7 21:58 .
drwxr-xr-x  3 -  staff   96 Jan  7 21:48 ..
-rw-r--r--  1 -  staff  498 Jan  7 21:58 .notes
```

To look at your notes you can do:

```bash
$ note 
2022-01-07 21:52:22 I have decided that I am going to do running.
           21:52:55 My goal is to finish a 5K
           21:53:06 But first I need a plan.
           21:53:40 Let's start with 1K first time and see how that goes.
2022-01-08 10:23:47 1K was fine, I am not tired at all.
           10:23:54 Let's see how 2K goes in two days.
2022-01-10 09:12:18 Well, it was much harder!
           09:12:31 I keep this distance for a while now.
```

To search in notes:

```bash
$ note -s "hard"
 • 
   2022-01-10 09:12:18 Well, it was much harder!
```

Recursive search would try to find all the notes in the subdirectories:

```bash
$ cd ..
$ note "This is the example folder"
$ note -r -s is
 • 
   2022-01-07 22:08:07 This is the example folder.
 • running
   2022-01-07 21:52:55 My goal is to finish a 5K
   2022-01-10 09:12:31 I keep this distance for a while now.
```
