#!/bin/bash
# Run load test

for i in ./*.yml; do
	echo "Running $i"
	testName=$(basename $i .yml)
	artillery run --output=$testName_$(date +%s).json $i
done
