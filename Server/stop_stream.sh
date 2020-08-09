#!/bin/bash
kill $(ps aux | grep '[r]aspivid' | awk '{print $2}')
