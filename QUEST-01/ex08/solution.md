## Solution
```bash
$ touch lookagain.sh
$ echo '#!/bin/bash' > lookagain.sh
$ echo 'find . -type f -name "*.sh" | sed 's|^\./||; s/\.sh$//' | sort -r' >> lookagain.sh
$ chmod +x lookagain.sh
$ ./lookagain.sh
```

## ToLearn
    1. sort command