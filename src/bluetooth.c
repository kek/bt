#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>

#include "bluetooth.h"

extern void SendGoMessage(char* message);
extern void SendGoDone();

static void adapter_on_scan_start(simpleble_adapter_t adapter, void* userdata) {
  SendGoMessage("SCAN START");
}

static void adapter_on_scan_stop(simpleble_adapter_t adapter, void* userdata) {
  SendGoDone();
}

static void adapter_on_scan_found(simpleble_adapter_t adapter, simpleble_peripheral_t peripheral, void* userdata) {
  SendGoMessage("SCAN FOUND");
}

#define assert(condition, msg) if (!condition) { printf("%s", msg); exit(255); }

void bluetooth_scan() {
  printf("Hello\n");

  simpleble_err_t err_code = SIMPLEBLE_SUCCESS;

  size_t adapter_count = simpleble_adapter_get_count();
  printf("Adapter count: %d\n", (int)adapter_count);

  simpleble_adapter_t adapter = simpleble_adapter_get_handle(0);
  assert(adapter, "No adapter was found.\n");

  simpleble_adapter_set_callback_on_scan_start(adapter, adapter_on_scan_start, NULL);
  simpleble_adapter_set_callback_on_scan_stop(adapter, adapter_on_scan_stop, NULL);
  simpleble_adapter_set_callback_on_scan_found(adapter, adapter_on_scan_found, NULL);

  simpleble_adapter_scan_for(adapter, 5000);
}
