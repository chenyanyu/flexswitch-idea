//Copyright [2016] [SnapRoute Inc]
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//       Unless required by applicable law or agreed to in writing, software
//       distributed under the License is distributed on an "AS IS" BASIS,
//       WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//       See the License for the specific language governing permissions and
//       limitations under the License.
//
//   This is a auto-generated file, please do not edit!
// _______   __       __________   ___      _______.____    __    ____  __  .___________.  ______  __    __
// |   ____||  |     |   ____\  \ /  /     /       |\   \  /  \  /   / |  | |           | /      ||  |  |  |
// |  |__   |  |     |  |__   \  V  /     |   (----  \   \/    \/   /  |  |  ---|  |---- |  ,---- |  |__|  |
// |   __|  |  |     |   __|   >   <       \   \      \            /   |  |     |  |     |  |     |   __   |
// |  |     |  `----.|  |____ /  .  \  .----)   |      \    /\    /    |  |     |  |     |  `----.|  |  |  |
// |__|     |_______||_______/__/ \__\ |_______/        \__/  \__/     |__|     |__|      \______||__|  |__|
//

#include <cps.h>

cps_api_return_code_t CPSGetAllPortCfg(PortCfg_t *portCfg, uint8_t *count) {
	cps_api_return_code_t retVal = cps_api_ret_code_OK;
        cps_api_key_t *key;
        cps_api_get_params_t gp;

        cps_api_get_request_init(&gp);
        cps_api_object_t obj = cps_api_object_list_create_obj_and_append(gp.filters);

        key = cps_api_object_key(obj);
        cps_api_key_from_attr_with_qual(key, BASE_IF_PHY_FRONT_PANEL_PORT_OBJ, cps_api_qualifier_TARGET);
        gp.key_count=1;
        gp.keys = key;
	retVal = cps_api_get(&gp);
        if (retVal == cps_api_ret_code_OK) {
                size_t ix, mx = cps_api_object_list_size(gp.list);
                for (ix = 0; ix < mx; ix++) {
                        cps_api_object_t obj = cps_api_object_list_get(gp.list, ix);
                        cps_api_object_attr_t port_name = cps_api_object_attr_get(obj, BASE_IF_PHY_FRONT_PANEL_PORT_DEFAULT_NAME);
                        if (port_name == CPS_API_ATTR_NULL) {
                                fprintf(stderr, "Unable to get port name for obj at idx %d", (int)ix);
                        } else {
                                fprintf(stderr, "Iterating through obj %s\n", (char*)cps_api_object_attr_data_bin(port_name));
				snprintf(portCfg[(int)ix].PortName, strlen((char*)cps_api_object_attr_data_bin(port_name))+1, "%s", (char*)cps_api_object_attr_data_bin(port_name));
				(*count)++;
                        }
                }
        } else {
                fprintf(stderr, "Failed to get front panel port information");
		cps_api_get_request_close(&gp);
		return retVal;
        }
        //Close get request
        cps_api_get_request_close(&gp);
	fprintf(stderr, "Count = %u\n", *count);
	return retVal;

}

cps_api_return_code_t CPSSetPortAdminState(char *intfRef, uint8_t val) {
	bool state = false;
	cps_api_return_code_t retVal = cps_api_ret_code_OK;
    	cps_api_object_t obj = cps_api_object_create();

    	cps_api_key_from_attr_with_qual(cps_api_object_key(obj),
        	DELL_BASE_IF_CMN_IF_INTERFACES_INTERFACE_OBJ, cps_api_qualifier_TARGET);

    	cps_api_set_key_data(obj,IF_INTERFACES_INTERFACE_NAME,cps_api_object_ATTR_T_BIN,
               intfRef,strlen(intfRef)+1);

    	cps_api_object_attr_add_u32(obj,IF_INTERFACES_INTERFACE_ENABLED, (uint32_t)val);
	cps_api_object_attr_add_u32(obj,DELL_IF_IF_INTERFACES_INTERFACE_AUTO_NEGOTIATION,state);

        cps_api_transaction_params_t tr;
        retVal = cps_api_transaction_init(&tr);
        if (retVal != cps_api_ret_code_OK) {
                printf("Failed cps_api_transaction_init() retVal: %d\n", retVal);
                return retVal;
        }

        retVal = cps_api_set(&tr, obj);
        if (retVal != cps_api_ret_code_OK) {
                cps_api_transaction_close(&tr);
                printf("Failed cps_api_set() retVal: %d\n", retVal);
                return retVal;
        }

        retVal = cps_api_commit(&tr);
        if (retVal != cps_api_ret_code_OK) {
                cps_api_transaction_close(&tr);
                printf("Failed cps_api_commit() retVal: %d\n", retVal);
                return retVal;
        }

        cps_api_transaction_close(&tr);
	return retVal;
}

