#!/bin/sh
#set -xe
echo "Booting EdgeMcuMgr"
# For Bluetooth to work
echo "Initializing bluetooth"
# Initialize bluez
echo -e "\tbluetoothd start commanded...\n"
#/usr/libexec/bluetooth/bluetoothd --debug &
echo -e "\tbluetoothd started.\n"
# BTATTACH to /dev/ttyACM# for reel board
echo -e "\Generating HCI interface from UART over USB to ZephyrOS-based BT Dongle running zephyr/samples/bluetooth/hci_uart"
btattach -B /dev/ttyACM0 -P h4 &
echo -e "\tHCI interface generated.\n"
# Launch the API server
echo -e "Starting API server ... "
./main --conntype ble -i 1 --connstring peer_name='Zephyr' restsvc start