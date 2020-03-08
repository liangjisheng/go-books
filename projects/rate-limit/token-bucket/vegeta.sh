#!/bin/bash

vegeta attack -duration=5s -rate=100 -targets=./vegeta.conf | vegeta report > report.txt