//
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
// _______  __       __________   ___      _______.____    __    ____  __  .___________.  ______  __    __
// |   ____||  |     |   ____\  \ /  /     /       |\   \  /  \  /   / |  | |           | /      ||  |  |  |
// |  |__   |  |     |  |__   \  V  /     |   (----` \   \/    \/   /  |  | `---|  |----`|  ,----'|  |__|  |
// |   __|  |  |     |   __|   >   <       \   \      \            /   |  |     |  |     |  |     |   __   |
// |  |     |  `----.|  |____ /  .  \  .----)   |      \    /\    /    |  |     |  |     |  `----.|  |  |  |
// |__|     |_______||_______/__/ \__\ |_______/        \__/  \__/     |__|     |__|      \______||__|  |__|
//

#include <cps.h>

cps_api_return_code_t CPSCreateIPv4Neighbor(char *nbrIp, char *intf, uint8_t macAddr[6]) {
        cps_api_object_t obj = cps_api_object_create();
        cps_api_return_code_t retVal = cps_api_ret_code_OK;

	cps_api_key_from_attr_with_qual(cps_api_object_key(obj), BASE_ROUTE_OBJ_OBJ,cps_api_qualifier_TARGET);
    	cps_api_object_attr_add_u32(obj,BASE_ROUTE_OBJ_NBR_AF,AF_INET);

    	uint32_t ip;
    	struct in_addr a;
    	inet_aton(nbrIp, &a);
    	ip=a.s_addr;

    	cps_api_object_attr_add(obj,BASE_ROUTE_OBJ_NBR_ADDRESS,&ip,sizeof(ip));
    	int port_index = if_nametoindex(intf);
    	cps_api_object_attr_add_u32(obj,BASE_ROUTE_OBJ_NBR_IFINDEX, port_index);

	hal_mac_addr_t mac_addr;
	mac_addr[0] = macAddr[0];
	mac_addr[1] = macAddr[1];
	mac_addr[2] = macAddr[2];
	mac_addr[3] = macAddr[3];
	mac_addr[4] = macAddr[4];
	mac_addr[5] = macAddr[5];

    	cps_api_object_attr_add(obj, BASE_ROUTE_OBJ_NBR_MAC_ADDR, &mac_addr, HAL_MAC_ADDR_LEN);

        cps_api_transaction_params_t tr;
        retVal = cps_api_transaction_init(&tr);
        if (retVal != cps_api_ret_code_OK) {
                printf("Failed cps_api_transaction_init() retVal: %d\n", retVal);
                return retVal;
        }

        retVal = cps_api_create(&tr, obj);
        if (retVal != cps_api_ret_code_OK) {
                cps_api_transaction_close(&tr);
                printf("Failed cps_api_create() retVal: %d\n", retVal);
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

cps_api_return_code_t CPSDeleteIPv4Neighbor(char *nbrIp, char *intf, uint8_t macAddr[6]) {
        cps_api_object_t obj = cps_api_object_create();
        cps_api_return_code_t retVal = cps_api_ret_code_OK;

	cps_api_key_from_attr_with_qual(cps_api_object_key(obj), BASE_ROUTE_OBJ_OBJ,cps_api_qualifier_TARGET);
    	cps_api_object_attr_add_u32(obj,BASE_ROUTE_OBJ_NBR_AF,AF_INET);

    	uint32_t ip;
    	struct in_addr a;
    	inet_aton(nbrIp, &a);
    	ip=a.s_addr;

    	cps_api_object_attr_add(obj,BASE_ROUTE_OBJ_NBR_ADDRESS,&ip,sizeof(ip));
    	int port_index = if_nametoindex(intf);
    	cps_api_object_attr_add_u32(obj,BASE_ROUTE_OBJ_NBR_IFINDEX, port_index);

	hal_mac_addr_t mac_addr;
	mac_addr[0] = macAddr[0];
	mac_addr[1] = macAddr[1];
	mac_addr[2] = macAddr[2];
	mac_addr[3] = macAddr[3];
	mac_addr[4] = macAddr[4];
	mac_addr[5] = macAddr[5];

    	cps_api_object_attr_add(obj, BASE_ROUTE_OBJ_NBR_MAC_ADDR, &mac_addr, HAL_MAC_ADDR_LEN);

        cps_api_transaction_params_t tr;
        retVal = cps_api_transaction_init(&tr);
        if (retVal != cps_api_ret_code_OK) {
                printf("Failed cps_api_transaction_init() retVal: %d\n", retVal);
                return retVal;
        }

        retVal = cps_api_delete(&tr, obj);
        if (retVal != cps_api_ret_code_OK) {
                cps_api_transaction_close(&tr);
                printf("Failed cps_api_delete() retVal: %d\n", retVal);
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
