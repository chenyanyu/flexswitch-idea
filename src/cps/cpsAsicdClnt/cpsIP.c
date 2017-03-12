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

cps_api_return_code_t CPSCreateIPv4Intf(char *intfRef, char *ipAddr, uint32_t prefix) {
        cps_api_object_t obj = cps_api_object_create();
        cps_api_return_code_t retVal = cps_api_ret_code_OK;

        int index = if_nametoindex(intfRef);

        cps_api_key_from_attr_with_qual(cps_api_object_key(obj), BASE_IP_IPV4_ADDRESS, cps_api_qualifier_TARGET);

        cps_api_object_attr_add_u32(obj, BASE_IP_IPV4_ADDRESS_PREFIX_LENGTH, prefix);

        uint32_t ip;
        struct in_addr a;
        inet_aton(ipAddr ,&a);
        ip=a.s_addr;
        cps_api_object_attr_add(obj, BASE_IP_IPV4_ADDRESS_IP, &ip,sizeof(ip));
        cps_api_object_attr_add_u32(obj, BASE_IP_IPV4_IFINDEX, index);


        cps_api_transaction_params_t tr;
        retVal = cps_api_transaction_init(&tr);
        if (retVal != cps_api_ret_code_OK) {
                printf("Failed cps_api_transaction_init: %d\n", retVal);
                return retVal;
        }
        cps_api_create(&tr,obj);
        retVal = cps_api_commit(&tr);
        if (retVal != cps_api_ret_code_OK) {
                printf("Failed cps_api_commit : %d\n", retVal);
                return retVal;
        }
        cps_api_transaction_close(&tr);
	return retVal;
}

cps_api_return_code_t CPSDeleteIPv4Intf(char *intfRef, char *ipAddr, uint32_t prefix) {
        cps_api_object_t obj = cps_api_object_create();
        cps_api_return_code_t retVal = cps_api_ret_code_OK;

        int index = if_nametoindex(intfRef);

        cps_api_key_from_attr_with_qual(cps_api_object_key(obj), BASE_IP_IPV4_ADDRESS, cps_api_qualifier_TARGET);

        cps_api_object_attr_add_u32(obj, BASE_IP_IPV4_ADDRESS_PREFIX_LENGTH, prefix);

        uint32_t ip;
        struct in_addr a;
        inet_aton(ipAddr ,&a);
        ip=a.s_addr;
        cps_api_object_attr_add(obj, BASE_IP_IPV4_ADDRESS_IP, &ip,sizeof(ip));
        cps_api_object_attr_add_u32(obj, BASE_IP_IPV4_IFINDEX, index);


        cps_api_transaction_params_t tr;
        retVal = cps_api_transaction_init(&tr);
        if (retVal != cps_api_ret_code_OK) {
                printf("Failed cps_api_transaction_init: %d\n", retVal);
                return retVal;
        }
        cps_api_delete(&tr,obj);
        retVal = cps_api_commit(&tr);
        if (retVal != cps_api_ret_code_OK) {
                printf("Failed cps_api_commit : %d\n", retVal);
                return retVal;
        }
        cps_api_transaction_close(&tr);
	return retVal;
}

#if 0
cps_api_return_code_t update_ipv4_addr(char *intfRef, char *ipAddr, uint32_t prefix) {
        cps_api_object_t obj = cps_api_object_create();
        cps_api_return_code_t retVal = cps_api_ret_code_OK;

        int index = if_nametoindex(intfRef);

        cps_api_key_from_attr_with_qual(cps_api_object_key(obj), BASE_IP_IPV4_ADDRESS, cps_api_qualifier_TARGET);

        cps_api_object_attr_add_u32(obj, BASE_IP_IPV4_ADDRESS_PREFIX_LENGTH, prefix);

        uint32_t ip;
        struct in_addr a;
        inet_aton(ipAddr ,&a);
        ip=a.s_addr;
        cps_api_object_attr_add(obj, BASE_IP_IPV4_ADDRESS_IP, &ip,sizeof(ip));
        cps_api_object_attr_add_u32(obj, BASE_IP_IPV4_IFINDEX, index);


        cps_api_transaction_params_t tr;
        retVal = cps_api_transaction_init(&tr);
        if (retVal != cps_api_ret_code_OK) {
                printf("Failed cps_api_transaction_init: %d\n", retVal);
                return retVal;
        }
        cps_api_set(&tr,obj);
        retVal = cps_api_commit(&tr);
        if (retVal != cps_api_ret_code_OK) {
                printf("Failed cps_api_commit : %d\n", retVal);
                return retVal;
        }
        cps_api_transaction_close(&tr);
	return retVal;
}
#endif
