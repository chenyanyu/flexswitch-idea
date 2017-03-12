#include <cps.h>
#include "_cgo_export.h"

static bool cps_if_event_cb(cps_api_object_t obj, void *param) {
  int port_ifindex = -1;
  int port_operstate = -1;
  cps_api_object_it_t itr;
  cps_api_attr_id_t attr_id;
  cps_api_object_it_begin(obj, &itr);

  cps_api_key_t *key;
  key = cps_api_object_key(obj);

  while (cps_api_object_it_valid(&itr)) {
    attr_id = cps_api_object_attr_id(itr.attr);
    if (attr_id == IF_INTERFACES_STATE_INTERFACE_IF_INDEX) {
      port_ifindex = cps_api_object_attr_data_u32(itr.attr);
    } else if (attr_id == IF_INTERFACES_STATE_INTERFACE_OPER_STATUS) {
      port_operstate = cps_api_object_attr_data_u32(itr.attr);
    }
    if ((port_ifindex != -1) && (port_operstate != -1)) {
        //Call go exported func
        HandleLinkNotifications(port_ifindex, port_operstate);
        break;
    }
    cps_api_object_it_next(&itr);
  }
  return true;
}

cps_api_return_code_t CPSRegisterNotificationHandler()
{
  cps_api_return_code_t ret;

  //Initialize the event service
  ret = cps_api_event_service_init();
  if (ret != cps_api_ret_code_OK)   {
    return ret;
  }
  // Initialize   the   event   handling   thread
  ret = cps_api_event_thread_init();
  if (ret != cps_api_ret_code_OK) {
    return ret;
  }
  //Create   and initialize   the   key
  cps_api_key_t   key;
  cps_api_key_from_attr_with_qual(&key, DELL_BASE_IF_CMN_IF_INTERFACES_STATE_INTERFACE, cps_api_qualifier_OBSERVED);

  //Create   the   registration   object
  cps_api_event_reg_t   reg;
  memset(&reg, 0, sizeof(reg));

  reg.number_of_objects = 1;
  reg.objects = &key;

  // Register   to   receive   events   for   key created   above
  ret = cps_api_event_thread_reg(&reg, cps_if_event_cb, NULL);
  if (ret != cps_api_ret_code_OK) {
    return ret;
  }
  return cps_api_ret_code_OK;
}

void CPSUnregisterNotificationHandler()
{
  cps_api_return_code_t ret;

  cps_api_event_thread_shutdown();
  return;
}
