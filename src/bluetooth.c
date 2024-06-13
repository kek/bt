#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "bluetooth.h"

extern void GoDeviceFound(char* peripheral_identifier, char* peripheral_address);
extern void SendGoDone();

static void adapter_on_scan_start(simpleble_adapter_t adapter, void* userdata) {
}

static void adapter_on_scan_stop(simpleble_adapter_t adapter, void* userdata) {
  SendGoDone();
}

static void adapter_on_scan_found(simpleble_adapter_t adapter, simpleble_peripheral_t peripheral, void* userdata) {
  char* peripheral_identifier = simpleble_peripheral_identifier(peripheral);
  char* peripheral_address = simpleble_peripheral_address(peripheral);

  char* result = malloc(strlen(peripheral_identifier) + strlen(peripheral_address) + 3);
  if (strlen(peripheral_identifier) != 0) {
    sprintf(result, "%s: %s", peripheral_identifier, peripheral_address);
  } else {
    sprintf(result, "%s", peripheral_address);
  }
  GoDeviceFound(peripheral_identifier, peripheral_address);
  free(result);
}

#define assert(condition, msg) if (!(condition)) { printf("%s", msg); exit(255); }

void bluetooth_scan() {
  simpleble_err_t err_code = SIMPLEBLE_SUCCESS;

  size_t adapter_count = simpleble_adapter_get_count();
  assert(adapter_count != 0, "No adapter was found.\n");

  simpleble_adapter_t adapter = simpleble_adapter_get_handle(0);
  assert(adapter, "No adapter was found.\n");

  simpleble_adapter_set_callback_on_scan_start(adapter, adapter_on_scan_start, NULL);
  simpleble_adapter_set_callback_on_scan_stop(adapter, adapter_on_scan_stop, NULL);
  simpleble_adapter_set_callback_on_scan_found(adapter, adapter_on_scan_found, NULL);

  simpleble_adapter_scan_for(adapter, 5000);
}
