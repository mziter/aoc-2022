dayDir := cmd/day${day}
session_cookie := $(shell cat session.dat)

new:
	@echo "Creating day directory..."
	mkdir -p ${dayDir}
	touch ${dayDir}/main.go
	@echo "package main\n" >> ${dayDir}/main.go
	touch ${dayDir}/partOne.go
	@echo "package main\n" >> ${dayDir}/partOne.go
	touch ${dayDir}/partTwo.go
	@echo "package main\n" >> ${dayDir}/partTwo.go
	curl -b session=${session_cookie} https://adventofcode.com/${year}/day/${day}/input --output ${dayDir}/input.txt
