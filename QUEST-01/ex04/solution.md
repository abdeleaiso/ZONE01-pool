## Solution
```bash
$ touch mastertheLS
$ echo '#!/bin/bash' > mastertheLS
$ echo 'ls -put | tr "\n" "," | sed "s/,$//"' >> mastertheLS
$ bash mastertheLS 
```

## ToLearn
    1. ls options
    2. tr command
    3. sed command