#include <stdio.h>
#include <stdlib.h>

extern void SendGoMessage(char* message);

void bluetooth_scan()
{
    printf("Hello\n");
    SendGoMessage("Hello world");
}
