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
#include <sys/socket.h>
#include <sys/ioctl.h>
#include <linux/if.h>
#include <netdb.h>
#include <unistd.h>

int get_vlan_if_mac_addr(char *ifName, char *macAddr) {
	struct ifreq s;
	int fd = socket(PF_INET, SOCK_DGRAM, IPPROTO_IP);
	if (fd < 0) {
		return -1;
	}

	strcpy(s.ifr_name, ifName);
	if (0 == ioctl(fd, SIOCGIFHWADDR, &s)) {
		sprintf(macAddr, "%02x:%02x:%02x:%02x:%02x:%02x",
			(unsigned char) s.ifr_addr.sa_data[0],
			(unsigned char) s.ifr_addr.sa_data[1],
			(unsigned char) s.ifr_addr.sa_data[2],
			(unsigned char) s.ifr_addr.sa_data[3],
			(unsigned char) s.ifr_addr.sa_data[4],
			(unsigned char) s.ifr_addr.sa_data[5]);
		close(fd);
		return 0;
	}
	close(fd);
	return -1;
}

cps_api_return_code_t _add_ports_to_vlan(char *vlanName, int numOfTagPorts, char **tagPortNameList, int numOfUntagPorts, char **untagPortNameList) {
	cps_api_return_code_t retVal = cps_api_ret_code_OK;
	cps_api_object_t obj = cps_api_object_create();
	int idx = 0;
	int ret = 0;

	cps_api_key_from_attr_with_qual(cps_api_object_key(obj), DELL_BASE_IF_CMN_IF_INTERFACES_INTERFACE_OBJ, cps_api_qualifier_TARGET);

	cps_api_object_attr_add(obj,IF_INTERFACES_INTERFACE_TYPE, (const char *)IF_INTERFACE_TYPE_IANAIFT_IANA_INTERFACE_TYPE_IANAIFT_L2VLAN, sizeof(IF_INTERFACE_TYPE_IANAIFT_IANA_INTERFACE_TYPE_IANAIFT_L2VLAN));
	cps_api_object_attr_add(obj, IF_INTERFACES_INTERFACE_NAME, vlanName, strlen(vlanName)+1);
	char macAddr[18] = {0};
	memset(macAddr, 0, 18);
	ret = get_vlan_if_mac_addr(vlanName, macAddr);
	if (ret < 0) {
		syslog(LOG_ERR, "Failed to get_vlan_if_macAddr ret: %d\n", ret);
		return ret;
	}
	syslog(LOG_INFO, "get_vlan_if_macAddr %s\n", macAddr);

	cps_api_object_attr_add(obj, DELL_IF_IF_INTERFACES_INTERFACE_PHYS_ADDRESS, macAddr, strlen(macAddr) + 1);
	for (idx = 0; idx < numOfTagPorts; idx++) {
		printf("Vlan %s Tag Port %s\n", vlanName, tagPortNameList[idx]);
		cps_api_object_attr_add(obj, DELL_IF_IF_INTERFACES_INTERFACE_TAGGED_PORTS, tagPortNameList[idx], strlen(tagPortNameList[idx])+1);
	}

	for (idx = 0; idx < numOfUntagPorts; idx++) {
		printf("Vlan %s Untag Port %s\n", vlanName, untagPortNameList[idx]);
		cps_api_object_attr_add(obj, DELL_IF_IF_INTERFACES_INTERFACE_UNTAGGED_PORTS, untagPortNameList[idx], strlen(untagPortNameList[idx])+1);
	}
	cps_api_object_attr_add_u32(obj,IF_INTERFACES_INTERFACE_ENABLED,true);

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


cps_api_return_code_t _create_vlan(uint32_t vlanId) {
	cps_api_object_t obj = cps_api_object_create();
	cps_api_return_code_t retVal = cps_api_ret_code_OK;

	cps_api_key_from_attr_with_qual(cps_api_object_key(obj), DELL_BASE_IF_CMN_IF_INTERFACES_INTERFACE_OBJ, cps_api_qualifier_TARGET);

	cps_api_object_attr_add(obj,IF_INTERFACES_INTERFACE_TYPE, (const char *)IF_INTERFACE_TYPE_IANAIFT_IANA_INTERFACE_TYPE_IANAIFT_L2VLAN, sizeof(IF_INTERFACE_TYPE_IANAIFT_IANA_INTERFACE_TYPE_IANAIFT_L2VLAN));
	cps_api_object_attr_add_u32(obj,BASE_IF_VLAN_IF_INTERFACES_INTERFACE_ID, vlanId);

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

cps_api_return_code_t CPSDeleteVlan(uint32_t vlanId) {
	cps_api_object_t obj = cps_api_object_create();
	cps_api_return_code_t retVal = cps_api_ret_code_OK;
	char br_name[256] = {0};

	cps_api_key_from_attr_with_qual(cps_api_object_key(obj), DELL_BASE_IF_CMN_IF_INTERFACES_INTERFACE_OBJ, cps_api_qualifier_TARGET);

	cps_api_object_attr_add(obj,IF_INTERFACES_INTERFACE_TYPE, (const char *)IF_INTERFACE_TYPE_IANAIFT_IANA_INTERFACE_TYPE_IANAIFT_L2VLAN, sizeof(IF_INTERFACE_TYPE_IANAIFT_IANA_INTERFACE_TYPE_IANAIFT_L2VLAN));
	//cps_api_object_attr_add_u32(obj,BASE_IF_VLAN_IF_INTERFACES_INTERFACE_ID, vlanId);

	sprintf(br_name, "br%u", vlanId);
	cps_api_set_key_data(obj, IF_INTERFACES_INTERFACE_NAME, cps_api_object_ATTR_T_BIN,
                             br_name, strlen(br_name) + 1 );

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

cps_api_return_code_t CPSCreateVlan(uint32_t vlanId, uint32_t numOfTagPorts, char **tagPorts, uint32_t numOfUntagPorts, char **untagPorts) {
	cps_api_return_code_t retVal = cps_api_ret_code_OK;
	char vlanName[256] = {0};
	int idx = 0;

	retVal = _create_vlan(vlanId);
	if (retVal != cps_api_ret_code_OK) {
		printf("Failed _create_vlan() retVal: %d\n", retVal);
		return retVal;
	}
	syslog(LOG_INFO, "num of tag Ports: %d\n", numOfTagPorts);
	syslog(LOG_INFO, "num of untag Ports: %d\n", numOfUntagPorts);
	for (idx = 0; idx < numOfTagPorts; idx++) {
		syslog(LOG_INFO, "Tag Port: %s\n", tagPorts[idx]);
	}
	for (idx = 0; idx < numOfUntagPorts; idx++) {
		syslog(LOG_INFO, "UnTag Port: %s\n", untagPorts[idx]);
	}
	sprintf(vlanName, "br%u", vlanId);
	retVal = _add_ports_to_vlan(vlanName, numOfTagPorts, tagPorts, numOfUntagPorts, untagPorts);
	if (retVal != cps_api_ret_code_OK) {
		printf("Failed _add_ports_to_vlan() retVal: %d\n", retVal);
		return retVal;
	}

	return retVal;
}

#if 0
void get_vlan_prop_from_obj(cps_api_object_t obj, cps_vlan_prop_t *vlanProp){
	cps_api_object_it_t it;
	cps_api_object_it_begin(obj,&it);

	for ( ; cps_api_object_it_valid(&it) ; cps_api_object_it_next(&it) ) {
		int id = (int) cps_api_object_attr_id(it.attr);
		switch(id) {
		case DELL_BASE_IF_CMN_IF_INTERFACES_INTERFACE_IF_INDEX:
		    //std::cout<<"VLAN INDEX "<<cps_api_object_attr_data_u32(it.attr)<<std::endl;
			vlanProp->ifIdx = cps_api_object_attr_data_u32(it.attr);
			printf("Vlan IfIdx: %d\n", vlanProp->ifIdx);
			break;
		case BASE_IF_VLAN_IF_INTERFACES_INTERFACE_ID:
		    //std::cout<<"VLAN ID "<<cps_api_object_attr_data_u32(it.attr)<<std::endl;
			vlanProp->vlanId = cps_api_object_attr_data_u32(it.attr);
			printf("Vlan Id: %d\n", vlanProp->vlanId);
			break;
		case IF_INTERFACES_INTERFACE_NAME:
			sprintf(vlanProp->vlanName, "%s", (char *)cps_api_object_attr_data_bin(it.attr));
			printf("Vlan Name: %s\n", vlanProp->vlanName);
		    break;
		case DELL_IF_IF_INTERFACES_INTERFACE_TAGGED_PORTS:
		    //std::cout<<"Tagged Port "<<cps_api_object_attr_data_u32(it.attr)<<std::endl;
		    break;
		case DELL_IF_IF_INTERFACES_INTERFACE_UNTAGGED_PORTS:
		    //std::cout<<"Untagged Port "<<cps_api_object_attr_data_u32(it.attr)<<std::endl;
		    break;
		case DELL_IF_IF_INTERFACES_INTERFACE_PHYS_ADDRESS:
		    //printf("MAC Address %s \n", (char *)cps_api_object_attr_data_bin(it.attr));
		    break;
		case IF_INTERFACES_INTERFACE_ENABLED:
		    //std::cout<<"Admin Status "<< cps_api_object_attr_data_u32(it.attr)<<std::endl;
			vlanProp->operStatus = cps_api_object_attr_data_u32(it.attr);
			break;
		case DELL_IF_IF_INTERFACES_INTERFACE_LEARNING_MODE:
		    //std::cout<<"Learning mode "<<cps_api_object_attr_data_u32(it.attr)<<std::endl;
		    break;

		default :
		    break;

		}
	}
}


void get_all_vlan_prop(int *count, cps_vlan_prop_t *vlanPropList) {
	cps_api_get_params_t gp;
	cps_api_get_request_init(&gp);
	cps_api_key_t *key;
	size_t mx, ix;

	cps_api_object_t obj = cps_api_object_list_create_obj_and_append(gp.filters);
	key = cps_api_object_key(obj);
	cps_api_key_from_attr_with_qual(key, DELL_BASE_IF_CMN_IF_INTERFACES_INTERFACE_OBJ, cps_api_qualifier_TARGET);

	cps_api_object_attr_add(obj,IF_INTERFACES_INTERFACE_TYPE, (const char *)IF_INTERFACE_TYPE_IANAIFT_IANA_INTERFACE_TYPE_IANAIFT_L2VLAN,sizeof(IF_INTERFACE_TYPE_IANAIFT_IANA_INTERFACE_TYPE_IANAIFT_L2VLAN));

	gp.key_count = 1;
	gp.keys = key;

	if (cps_api_get(&gp)==cps_api_ret_code_OK) {
		mx = cps_api_object_list_size(gp.list);
		for (ix = 0 ; ix < mx ; ++ix ) {
	    		obj = cps_api_object_list_get(gp.list,ix);
	    		get_vlan_prop_from_obj(obj, &(vlanPropList[*count]));
			(*count)++;
		}
	}

	cps_api_get_request_close(&gp);
}
#endif
