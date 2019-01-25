// redirect_gettime.c
#include <dlfcn.h>
#include <time.h>

static int(*real_clock_gettime)(clockid_t, struct timespec *) = NULL;

int clock_gettime(clockid_t clk_id, struct timespec *tp) {
    if (real_clock_gettime == NULL) {
        real_clock_gettime = dlsym(RTLD_NEXT, "clock_gettime");
    }
    if (clk_id == CLOCK_MONOTONIC_RAW) {
        clk_id = CLOCK_MONOTONIC;
    }
    return real_clock_gettime(clk_id, tp);
}
