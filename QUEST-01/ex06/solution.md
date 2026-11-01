## Solution
```bash
$ touch look
$ echo '#!/bin/bash' > look
$ echo 'find . -type f \( -name "a*" -o -name "*z" -o -name "z*a!" \)' >> look
$ bash look
```

## ToLearn
    1. find command
    2. linux regex