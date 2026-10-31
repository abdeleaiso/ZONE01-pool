## Solution 
```bash
$ touch who-are-you.sh
$ echo '#!/bin/bash' > who-are-you.sh
$ echo 'curl -s "https://learn.zone01oujda.ma/assets/superhero/all.json" | jq -r ".[] | select(.id == 70 | .name)"' >> who-are-you.sh
$ chmod +x who-are-you.sh
$ ./who-are-you.sh
Batman
```

## To Learn
    1. JSON fils
    2. curl command
    3. jq command
    4. shebang line 