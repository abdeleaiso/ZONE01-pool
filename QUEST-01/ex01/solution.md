## Step 1 — Create first directorie
```bash
$ mkdir 0
```

## Step2 — Create files and change there dates
```bash
$ touch -d "1986-11-13 00:01" 1
$ touch -d "1988-03-05 00:10" 2
$ ln -s 0 3 # symbolic link
$ touch -hd "1990-02-16 00:11" 3
$ touch -d "1986-01-05 00:00" 0
$ touch -d "1990-10-07 01:00" 4
$ touch -d "1990-11-07 01:01" 5
$ touch -d "1991-02-08 01:10" 6 
$ touch -d "1991-03-08 01:11" 7
$ touch -d "1994-05-20 10:00" 8
$ touch -d "1994-06-10 10:01" 9
$ touch -d "1995-04-10 10:10" A
```

## Step3 — Change permissions
```bash
$ chmod 401 0
$ chmod 402 1
$ chmod 604 2
$ chmod 510 4
$ chmod 460 5
$ chmod 460 6
$ chmod 510 7
$ chmod 604 8
$ chmod 402 9
$ chmod 404 A
```

## Step4 — Verification
```bash
$ TZ=utc gls -l --time-style='+%F %R' | sed 1d | awk '{print $1, $6, $7, $8, $9, $10}'
```

## Step5 — create the file ``done.tar`` to be submitted.
```bash
$ tar -cf done.tar *
$ ls
0  1  2  3  4  5  6  7  8  9  A  done.tar
$
```

## ToLearn:
#### 1.Files & Folders
#### 2.Hard & Symbolic Link
#### 3.File Permission:
read (4), write (2), and execute (1) permissions, ranging from 0 (no access) to 7 (full access).