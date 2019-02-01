Weak attempt at diagnosing why we see a thread at 100%.

Found this : https://github.com/pytorch/pytorch/issues/3390

Below is the command to compile the code.  Stick the .so file in your Binaries/Linux directory and put "LD_PRELOAD=redirect_gettime.so" in front of the ./ShooterGameServer command used to start it up.

Unsure if this makes things better or worse.

gcc -D_GNU_SOURCE -fPIC -shared  -o redirect_gettime.so redirect_gettime.c -ldl
