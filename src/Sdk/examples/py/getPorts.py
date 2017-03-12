import sys
import os
import simplejson as json
sys.path.append(os.path.abspath('../../py'))
from flexswitchV2 import FlexSwitch

"""This program demonstrates 
   simple read APIs supported by FlexSwitch SDK. 
   This specific example is reading all the Port states.
   getAllPortStates returns a list of ports. Each item in the
   list of ports would contain port attributes. """
   

if __name__ == '__main__':
    swtch = FlexSwitch ('10.1.10.240', 8080)  # Instantiate object to talk to flexSwitch
    ports = swtch.getAllPortStates()          # Get all the state information for all ports
    port = ports[0]      #Just print first port information
    print json.dumps(port['Object'])



