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

cps_api_return_code_t CPSCreateIPv4Route(char *destNw, uint32_t prefix, char *nhIp) {
	cps_api_object_t obj = cps_api_object_create();
        cps_api_return_code_t retVal = cps_api_ret_code_OK;

	cps_api_key_from_attr_with_qual(cps_api_object_key(obj),
		   BASE_ROUTE_OBJ_OBJ,cps_api_qualifier_TARGET);


	cps_api_object_attr_add_u32(obj,BASE_ROUTE_OBJ_ENTRY_AF,AF_INET);
	cps_api_object_attr_add_u32(obj,BASE_ROUTE_OBJ_ENTRY_VRF_ID,0);
	cps_api_object_attr_add_u32(obj,BASE_ROUTE_OBJ_ENTRY_PREFIX_LEN, prefix);

	uint32_t ip;
	struct in_addr a;
	inet_aton(destNw, &a);
	ip=a.s_addr;

	cps_api_object_attr_add(obj,BASE_ROUTE_OBJ_ENTRY_ROUTE_PREFIX,&ip,sizeof(ip));
	cps_api_attr_id_t ids[3];
	const int ids_len = sizeof(ids)/sizeof(*ids);
	ids[0] = BASE_ROUTE_OBJ_ENTRY_NH_LIST;
	ids[1] = 0;
	ids[2] = BASE_ROUTE_OBJ_ENTRY_NH_LIST_NH_ADDR;

	//inet_aton("127.0.0.1", &a);
	inet_aton(nhIp, &a);
	ip=a.s_addr;
	cps_api_object_e_add(obj,ids,ids_len,cps_api_object_ATTR_T_BIN,
					&ip,sizeof(ip));

	cps_api_object_attr_add_u32(obj,BASE_ROUTE_OBJ_ENTRY_NH_COUNT,1);


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

cps_api_return_code_t CPSDeleteIPv4Route(char *destNw, uint32_t prefix, char *nhIp) {
	cps_api_object_t obj = cps_api_object_create();
        cps_api_return_code_t retVal = cps_api_ret_code_OK;

	cps_api_key_from_attr_with_qual(cps_api_object_key(obj),
		   BASE_ROUTE_OBJ_OBJ,cps_api_qualifier_TARGET);


	cps_api_object_attr_add_u32(obj,BASE_ROUTE_OBJ_ENTRY_AF,AF_INET);
	cps_api_object_attr_add_u32(obj,BASE_ROUTE_OBJ_ENTRY_VRF_ID,0);
	cps_api_object_attr_add_u32(obj,BASE_ROUTE_OBJ_ENTRY_PREFIX_LEN, prefix);

	uint32_t ip;
	struct in_addr a;
	inet_aton(destNw, &a);
	ip=a.s_addr;

	cps_api_object_attr_add(obj,BASE_ROUTE_OBJ_ENTRY_ROUTE_PREFIX,&ip,sizeof(ip));
	cps_api_attr_id_t ids[3];
	const int ids_len = sizeof(ids)/sizeof(*ids);
	ids[0] = BASE_ROUTE_OBJ_ENTRY_NH_LIST;
	ids[1] = 0;
	ids[2] = BASE_ROUTE_OBJ_ENTRY_NH_LIST_NH_ADDR;

	//inet_aton("127.0.0.1", &a);
	inet_aton(nhIp, &a);
	ip=a.s_addr;
	cps_api_object_e_add(obj,ids,ids_len,cps_api_object_ATTR_T_BIN,
					&ip,sizeof(ip));

	cps_api_object_attr_add_u32(obj,BASE_ROUTE_OBJ_ENTRY_NH_COUNT,1);


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


#if 0
void nas_route_dump_nht_object_content(cps_api_object_t obj){
	char str[INET6_ADDRSTRLEN];
	uint32_t addr_len = 0, af_data = 0;
	//uint32_t nhc = 0, nh_itr = 0;
	//char if_name[IFNAMSIZ];

	cps_api_object_attr_t af_attr = cps_api_get_key_data(obj, BASE_ROUTE_OBJ_ENTRY_AF);
	if (af_attr != CPS_API_ATTR_NULL) {
		af_data = cps_api_object_attr_data_u32(af_attr) ;
		//std::cout<<"Address Family "<<((af_data == AF_INET) ? "IPv4" : "IPv6")<<std::endl;
		addr_len = ((af_data == AF_INET) ? INET_ADDRSTRLEN : INET6_ADDRSTRLEN);
	}

	cps_api_object_attr_t dest_attr = cps_api_get_key_data(obj, BASE_ROUTE_OBJ_ENTRY_ROUTE_PREFIX);
	if (dest_attr != CPS_API_ATTR_NULL) {
        //std::cout<<"Dest Address "<<inet_ntop(af_data,cps_api_object_attr_data_bin(dest_attr),str,addr_len)<<std::endl;
		printf("Dest Address: %s\n", inet_ntop(af_data,cps_api_object_attr_data_bin(dest_attr),str,addr_len));
	}
}

void get_route(void){
	uint32_t af = AF_INET;
	int ix;
	cps_api_get_params_t gp;
	cps_api_get_request_init(&gp);

	cps_api_object_t obj = cps_api_object_list_create_obj_and_append(gp.filters);
	cps_api_key_from_attr_with_qual(cps_api_object_key(obj),BASE_ROUTE_OBJ_ENTRY,
                                        cps_api_qualifier_TARGET);

	cps_api_set_key_data (obj, BASE_ROUTE_OBJ_ENTRY_AF, cps_api_object_ATTR_T_U32,&af, sizeof(af));

	if (cps_api_get(&gp)==cps_api_ret_code_OK) {
		size_t mx = cps_api_object_list_size(gp.list);

		for ( ix = 0 ; ix < mx ; ++ix ) {
			obj = cps_api_object_list_get(gp.list,ix);
            
			nas_route_dump_nht_object_content(obj);
			//cps_api_object_print(obj);
            
		}
	}

	cps_api_get_request_close(&gp);
}
#endif
