#!/bin/bash

echo "Запуск Ganache"
ganache-cli -a 1 -e 1000 --gaslimit 6721975 --host 0.0.0.0 --port 8545 --cors "*"