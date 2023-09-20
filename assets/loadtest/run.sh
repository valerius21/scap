#!/bin/bash
# Run load test

for i in ./*.yml; do
	echo "Running $i"
	testName=$(basename $i .yml)
	artillery run -o $testName $i
done
