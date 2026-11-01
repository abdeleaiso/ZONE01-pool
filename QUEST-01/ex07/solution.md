## Solution
```bash
$ touch myfamily.sh
$ echo '#!/bin/bash' > myfamily.sh
$ echo 'curl -s "https://learn.zone01oujda.ma/assets/superhero/all.json" | jq -r ".[] | select(.id == $HERO_ID) | .connections.relatives"
' >> myfamily.sh
$ chmod +x myfamily.sh
$ export HERO_ID=1
$ ./myfamily.sh
Marlo Chandler-Jones (wife); Polly (aunt); Mrs. Chandler (mother-in-law); Keith Chandler, Ray Chandler, three unidentified others (brothers-in-law); unidentified father (deceased); Jackie Shorr (alleged mother; unconfirmed) 
```