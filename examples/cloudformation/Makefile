COMMON = $(wildcard common/*/*.*)
PROJECT_A = $(wildcard projects/A/*.*)

project-A.template: $(COMMON) $(PROJECT_A)
	merje $^ -f json -o $@
