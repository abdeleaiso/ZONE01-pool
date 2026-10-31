## Solution
```bash
$ touch to-git-or-not-to-git.sh
$ echo '#!/bin/bash' > to-git-or-not-to-git.sh 
$ echo 'curl -s "https://learn.zone01oujda.ma/assets/superhero/all.json" | jq -r ".[] | select(.id == 170) | .name, .powerstats.power, .appearance.gender"' >> to-git-or-not-to-git.sh 
$ bash to-git-or-not-to-git.sh 
Chameleon
28
Male
```