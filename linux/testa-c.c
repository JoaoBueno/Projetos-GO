#include <stdio.h>
#include <string.h>
#include <stdlib.h>

#include "libimgutil32.h"

int
main(int argc, char **argv) {
    char *path, *retorno;
    GoUint w, h, err;

    if (argc < 2) {
        fprintf(stderr, "missing argument\n");
        return 1;
    }

    path = strdup(argv[1]);
    err = ImgutilGetImageSize(path, &retorno, &w, &h);
    free(path);

    if (err != 200) {
        fprintf(stderr, "error: %s\n", err);
        free(err);
        return 1;
    }

    printf("%s: %llux%llu\n", argv[1], w, h);
    printf("%s: ", path);
    printf("%s: ", retorno);

    return 0;
}
