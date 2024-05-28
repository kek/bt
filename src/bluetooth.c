#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>

#include "simpleble_c/adapter.h"
#include "simpleble_c/simpleble.h"
#include "simpleble_c/types.h"

extern void SendGoMessage(char *message);

void bluetooth_scan() {
  printf("Hello\n");

  simpleble_err_t err_code = SIMPLEBLE_SUCCESS;

  size_t adapter_count = simpleble_adapter_get_count();
  printf("Adapter count: %d\n", (int)adapter_count);
  simpleble_adapter_get_handle(0);
  SendGoMessage("Hello world");
}
