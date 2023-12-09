.PHONY: gen run test bench

gen:
	@if [ -n "$(DAY)" ]; then \
        day=$(DAY); \
    else \
        day=$$(ls | egrep "day[0-9]+$$" | \
            sed 's/^day*//g' | sort -nr | head -1); \
        day=$$((day + 1)); \
	fi; \
	echo "Generating day $$day..."; \
	go run ./gen -day $$day

run:
	@if [ -n "$(DAY)" ]; then \
        day=$(DAY); \
    else \
        day=$$(ls | egrep "day[0-9]+$$" | \
            sed 's/^day*//g' | sort -nr | head -1); \
    fi; \
	if [ -n "$(PART)" ]; then \
		go run ./day$$day -part $(PART); \
	else \
		go run ./day$$day -part -1; \
	fi

test:
	@if [ -n "$(DAY)" ]; then \
        day=$(DAY); \
    else \
        day=$$(ls | egrep "day[0-9]+$$" | \
            sed 's/^day*//g' | sort -nr | head -1); \
    fi; \
	go test ./day$$day

bench:
	@if [ -n "$(DAY)" ]; then \
        day=$(DAY); \
    else \
        day=$$(ls | egrep "day[0-9]+$$" | \
            sed 's/^day*//g' | sort -nr | head -1); \
    fi; \
	go test -bench=. ./day$$day
