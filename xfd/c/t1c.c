/*
 *
 * gcc -o md5lib -O3 -lm md5lib.c  | para gerar executavel
 * gcc -m32 -shared -o md5lib.so -fPIC md5lib.c  | gerar biblioteca 32 bits L
 * gcc -m32 -shared -o md5lib.dll -fPIC md5lib.c | gerar biblioteca 32 bits W
 *
 */
#include <stdio.h>
#include <string.h>

int retornamd5(char *s, char *retorno, long *len)
{

// #ifdef DEBUG
    printf("Chegou = %s", s);
    printf("Tamanho = %d\n", len);
// #endif

    char dest[50];
    char* str1 = "quick";
    char* str2 = "brown";
    char* str3 = "lazy";    

    // snprintf(dest, sizeof dest, "Isto %s é um %s teste %s", str1, str2, str3);
    strcpy(dest , "tua porra de saco de merda");

    strncpy(retorno, dest, sizeof dest);

    return 7;
}
