dayDir := cmd/day${day}
session_cookie := $(shell cat session_cookie)

new:
	@echo "Creating day directory..."
	mkdir -p ${dayDir}
	touch ${dayDir}/main.go
	@echo "package main\n" >> ${dayDir}/main.go
	@echo "Downloading input..." 
	curl -b session=${session_cookie} https://adventofcode.com/${year}/day/${day}/input --output ${dayDir}/input.txt
