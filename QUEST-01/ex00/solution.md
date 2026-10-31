## Step 1 — Configure Git
config your account with global flag to tell Git to remember your password by doing this on your terminal:
```bash
$ git config --global user.name "your-name"
$ git config --global user.email "your-email"
$ git config --global credential.helper store
```

## Step 2 — Create and clone the repository
1. Create **piscine-go-oujda** on Gitea WebSite.
2. Clone it by doing this cmd on your terminal: 
```bash
$ git clone https://learn.zone01oujda.ma/git/<your-name>/piscine-go-oujda.git
$ cd piscine-go-oujda
```

## Step 3 — Create your script 
#### METHOD 1:
```bash
$ touch hello.sh # create an empty file
$ echo 'echo "Hello <your-name>!"' > hello.sh # add to it echo Hello <your-name>!
$ bash hello.sh # Execute and verify
```

#### METHOD 2:
```bash
$ vim hello.sh # enter the file and type echo "Hello <your-name>!" and exit with "esc" then ":x" then hit enter
$ bash hello.sh # compile it (excute it)
```

## Step 5 — Submit your solution
```bash
$ git add hello.sh
$ git commit -m "My very first commit"
$ git push
```