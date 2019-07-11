#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>


int main(){

int i = 10;

    char* argv[4];
    argv[0] = "./day5";
    argv[1] = NULL;
    execv(argv[0], argv);
printf("%d", i);
//system(" ./day5");


return 0;
}